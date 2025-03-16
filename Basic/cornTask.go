package basic

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	cronInstance *cron.Cron
	cronTasks    = make(map[cron.EntryID]string)
	cronMutex    sync.RWMutex
	cronOnce     sync.Once
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
func SetCronTask(cronExpr string, task func()) error {
	// 创建新的 cron 实例
	c := newWithSeconds()

	// 添加任务
	entryID, err := c.AddFunc(cronExpr, task)
	if err != nil {
		return err
	}

	// 更新任务列表
	cronMutex.Lock()
	cronTasks[entryID] = cronExpr
	cronMutex.Unlock()

	return nil
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
	EntryID cron.EntryID
	Expr    string
} {
	// c := newWithSeconds()
	var tasks []struct {
		EntryID cron.EntryID
		Expr    string
	}

	// 获取所有任务
	cronMutex.RLock()
	for entryID := range cronTasks {
		tasks = append(tasks, struct {
			EntryID cron.EntryID
			Expr    string
		}{
			EntryID: entryID,
			Expr:    cronTasks[entryID],
		})
	}
	cronMutex.RUnlock()

	return tasks
}
