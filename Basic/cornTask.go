// cron 库的封装
package basic

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	cronInstance *cron.Cron                          // 全局 cron 实例
	cronTasks    = make(map[cron.EntryID]([]string)) // 通过 entryID 来获取 cron 表达式
	cronMutex    sync.RWMutex                        // 使用读写锁
	cronOnce     sync.Once                           // 保证 cronInstance 只被初始化一次
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	cronOnce.Do(func() {
		cronInstance = cron.New(cron.WithSeconds())
		cronInstance.Start()
	})
	return cronInstance
}

// 设置定时任务
func SetCronTask(cronName, cronExpr string, task func()) error {
	// 创建新的 cron 实例
	c := newWithSeconds()

	// 添加任务
	entryID, err := c.AddFunc(cronExpr, task)
	if err != nil {
		return err
	}
	// 更新任务列表
	cronMutex.Lock()
	cronTasks[entryID] = append(cronTasks[entryID], cronName)
	cronTasks[entryID] = append(cronTasks[entryID], cronExpr)
	cronMutex.Unlock()

	return nil
}

// 更新定时任务
func UpdateCronTask(entryID cron.EntryID, entryName, cronExpr string, task func()) bool {
	c := newWithSeconds()
	// 获取任务表达式
	cronMutex.RLock()
	_, exists := cronTasks[entryID]
	cronMutex.RUnlock()
	// 如果任务不存在，返回 false
	if !exists {
		return false
	}

	// 更新任务
	c.Remove(entryID)
	newEntryID, err := c.AddFunc(cronExpr, task)
	if err != nil {
		return false
	}

	// 更新任务列表
	cronMutex.Lock()
	cronTasks[newEntryID] = append(cronTasks[newEntryID], entryName)
	cronTasks[newEntryID] = append(cronTasks[newEntryID], cronExpr)
	delete(cronTasks, entryID)
	cronMutex.Unlock()

	return true
}

// 删除定时任务
func DeleteCronTask(entryID cron.EntryID) bool {
	c := newWithSeconds()

	// 获取任务表达式
	cronMutex.RLock()
	_, exists := cronTasks[entryID]
	cronMutex.RUnlock()

	// 如果任务不存在，返回 false
	if !exists {
		return false
	}

	// 删除任务
	c.Remove(entryID)

	cronMutex.Lock()
	delete(cronTasks, entryID)
	cronMutex.Unlock()

	return true
}

// 获取所有定时任务
func ListCronTasks() []struct {
	CronName string
	ID       int
	EntryID  cron.EntryID
	Expr     string
} {
	// c := newWithSeconds()
	var tasks []struct {
		CronName string
		ID       int
		EntryID  cron.EntryID
		Expr     string
	}

	// 获取所有任务
	cronMutex.RLock()
	for entryID := range cronTasks {
		tasks = append(tasks, struct {
			CronName string
			ID       int
			EntryID  cron.EntryID
			Expr     string
		}{
			CronName: cronTasks[entryID][0],
			ID:       int(entryID),
			EntryID:  entryID,
			Expr:     cronTasks[entryID][1],
		})
	}
	cronMutex.RUnlock()

	return tasks
}
