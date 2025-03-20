'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { checkAuth } from '@/src/utils/auth';
import CreateTaskDialog from '../src/components/CreateTaskDialog';
import TaskConfigPanel from '../src/components/TaskConfigPanel';

interface Task {
  id: number;
  title: string;
  content: string;
  recipient: string;
  scheduledTime?: string;
  sentTime?: string;
  status: 'pending' | 'sent' | 'failed';
}

export default function Home() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  const [scheduledTasks, setScheduledTasks] = useState<Task[]>([]);
  const [sentTasks, setSentTasks] = useState<Task[]>([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isConfigOpen, setIsConfigOpen] = useState(false);
  const [selectedMethod, setSelectedMethod] = useState('');

  useEffect(() => {
    if (!checkAuth()) {
      router.push('/login');
      return;
    }

    // 模拟获取任务数据
    const fetchTasks = async () => {
      try {
        setIsLoading(true);
        // 模拟数据
        const mockScheduledTasks: Task[] = [
          {
            id: 1,
            title: '测试任务1',
            content: '这是一个测试任务',
            recipient: 'test@example.com',
            scheduledTime: new Date(Date.now() + 3600000).toISOString(),
            status: 'pending'
          },
          {
            id: 2,
            title: '测试任务2',
            content: '这是另一个测试任务',
            recipient: 'test2@example.com',
            scheduledTime: new Date(Date.now() + 7200000).toISOString(),
            status: 'pending'
          }
        ];

        const mockSentTasks: Task[] = [
          {
            id: 3,
            title: '已发送任务1',
            content: '这是一个已发送的任务',
            recipient: 'sent@example.com',
            sentTime: new Date(Date.now() - 3600000).toISOString(),
            status: 'sent'
          },
          {
            id: 4,
            title: '已发送任务2',
            content: '这是另一个已发送的任务',
            recipient: 'sent2@example.com',
            sentTime: new Date(Date.now() - 7200000).toISOString(),
            status: 'failed'
          }
        ];

        setScheduledTasks(mockScheduledTasks);
        setSentTasks(mockSentTasks);
      } catch (error: any) {
        setError('获取任务失败: ' + (error.message || '未知错误'));
        console.error('获取任务失败:', error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchTasks();
  }, [router]);

  const handleCreateTask = (taskData: any) => {
    setIsDialogOpen(false);
    
    console.log('发送任务:', taskData);
  };

  const handleConfigSubmit = (config: any) => {
    setIsConfigOpen(false);
  };

  if (isLoading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-red-500 text-center">
          <div className="text-xl mb-2">⚠️</div>
          <div>{error}</div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="flex justify-between items-center mb-12">
          <h1 className="text-2xl font-light text-gray-900">消息任务管理</h1>
          <div className="flex space-x-4">
            <button
              onClick={() => setIsDialogOpen(true)}
              className="px-4 py-2 bg-black text-white rounded-full text-sm hover:bg-gray-800 transition-colors"
            >
              发送新消息
            </button>
          </div>
        </div>
        
        {/* 定时任务部分 */}
        <section className="mb-16">
          <div className="flex items-center mb-6">
            <h2 className="text-lg font-medium text-gray-900">定时任务</h2>
            <span className="ml-2 px-2 py-1 text-xs bg-white text-gray-600 rounded-full border border-gray-200">
              {scheduledTasks.length}
            </span>
          </div>
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {scheduledTasks.map((task) => (
              <div key={task.id} className="group bg-white border border-gray-200 rounded-xl p-6 hover:shadow-md hover:border-gray-300 transition-all duration-200">
                <div className="flex justify-between items-start mb-4">
                  <h3 className="text-base font-medium text-gray-900">{task.title}</h3>
                  <div className="flex space-x-2 opacity-0 group-hover:opacity-100 transition-opacity">
                    <button className="p-1 text-gray-400 hover:text-gray-600">
                      <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                      </svg>
                    </button>
                    <button className="p-1 text-gray-400 hover:text-red-500">
                      <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="space-y-3 text-sm text-gray-500">
                  <p className="flex items-center">
                    <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                    {task.recipient}
                  </p>
                  <p className="flex items-center">
                    <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    {new Date(task.scheduledTime!).toLocaleString()}
                  </p>
                  <p className="text-gray-400">{task.content}</p>
                </div>
              </div>
            ))}
            {scheduledTasks.length === 0 && (
              <div className="col-span-full text-center py-12">
                <div className="text-gray-400 mb-2">
                  <svg className="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                  </svg>
                </div>
                <p className="text-gray-500">暂无定时任务</p>
              </div>
            )}
          </div>
        </section>

        {/* 已发送任务部分 */}
        <section>
          <div className="flex items-center mb-6">
            <h2 className="text-lg font-medium text-gray-900">已发送任务</h2>
            <span className="ml-2 px-2 py-1 text-xs bg-white text-gray-600 rounded-full border border-gray-200">
              {sentTasks.length}
            </span>
          </div>
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {sentTasks.map((task) => (
              <div key={task.id} className="group bg-white border border-gray-200 rounded-xl p-6 hover:shadow-md hover:border-gray-300 transition-all duration-200">
                <div className="flex justify-between items-start mb-4">
                  <h3 className="text-base font-medium text-gray-900">{task.title}</h3>
                  <span className={`px-2 py-1 text-xs rounded-full
                    ${task.status === 'sent' ? 'bg-green-50 text-green-700 border border-green-100' :
                      task.status === 'failed' ? 'bg-red-50 text-red-700 border border-red-100' :
                      'bg-yellow-50 text-yellow-700 border border-yellow-100'}`}>
                    {task.status === 'sent' ? '发送成功' :
                     task.status === 'failed' ? '发送失败' : '发送中'}
                  </span>
                </div>
                <div className="space-y-3 text-sm text-gray-500">
                  <p className="flex items-center">
                    <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                    {task.recipient}
                  </p>
                  <p className="flex items-center">
                    <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    {new Date(task.sentTime!).toLocaleString()}
                  </p>
                  <p className="text-gray-400">{task.content}</p>
                </div>
              </div>
            ))}
            {sentTasks.length === 0 && (
              <div className="col-span-full text-center py-12">
                <div className="text-gray-400 mb-2">
                  <svg className="w-12 h-12 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
                  </svg>
                </div>
                <p className="text-gray-500">暂无已发送任务</p>
              </div>
            )}
          </div>
        </section>
      </div>

      <CreateTaskDialog
        isOpen={isDialogOpen}
        onClose={() => setIsDialogOpen(false)}
        onSubmit={handleCreateTask}
        onMethodSelect={setSelectedMethod}
      />

      <TaskConfigPanel
        isOpen={false}
        onClose={() => setIsConfigOpen(false)}
        sendMethod={selectedMethod}
        onSubmit={handleConfigSubmit}
      />
    </div>
  );
}
