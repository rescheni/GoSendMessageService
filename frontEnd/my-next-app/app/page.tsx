'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { checkAuth } from '@/src/utils/auth';

interface Task {
  EntryID: number;
  Expr: string;
  CronName: string;
  ID: number;
}

export default function Home() {
  const router = useRouter();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!checkAuth()) {
      router.push('/login');
      return;
    }

    const fetchTasks = async () => {
      try {
        const apiKey = localStorage.getItem('apiKey');
        const response = await fetch(`/api/cron/list?api_key=${apiKey}`);
        
        if (!response.ok) {
          throw new Error('获取任务列表失败');
        }

        const data = await response.json();
        setTasks(data.tasks || []);
      } catch (error) {
        console.error('获取任务列表失败:', error);
        setError('获取任务列表失败');
      } finally {
        setLoading(false);
      }
    };

    fetchTasks();
  }, [router]);

  const handleDelete = async (entryId: number) => {
    if (!confirm('确定要删除这个任务吗？')) {
      return;
    }

    try {
      const apiKey = localStorage.getItem('apiKey');
      const response = await fetch(`/api/cron/delete?api_key=${apiKey}&id=${entryId}`, {
        method: 'GET'
      });

      if (!response.ok) {
        throw new Error('删除失败');
      }

      // 删除成功后，更新任务列表
      setTasks(tasks.filter(task => task.EntryID !== entryId));
      alert('删除成功');
    } catch (error) {
      console.error('删除任务失败:', error);
      alert('删除任务失败，请重试');
    }
  };

  if (loading) {
    return <div className="text-center py-8">加载中...</div>;
  }

  return (
    <main className="min-h-screen p-8">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-2xl font-bold">定时任务列表</h1>
          <button
            onClick={() => router.push('/create-task')}
            className="px-4 py-2 bg-black text-white rounded-lg hover:bg-gray-800 transition-colors"
          >
            新建任务
          </button>
        </div>

        {error ? (
          <div className="text-red-500 bg-red-50 border border-red-100 rounded-lg p-4">
            {error}
          </div>
        ) : tasks.length === 0 ? (
          <div className="text-center py-8 text-gray-500">暂无任务</div>
        ) : (
          <div className="grid gap-4">
            {tasks.map((task) => (
              <div
                key={task.EntryID}
                className="bg-white rounded-lg shadow p-4"
              >
                <div className="flex justify-between items-start">
                  <div>
                    <div className="text-lg font-medium mb-2">{task.CronName}</div>
                    <div className="text-sm text-gray-500">任务ID: {task.EntryID}</div>
                    <div className="mt-2 font-mono text-sm">Cron表达式: {task.Expr}</div>
                  </div>
                  <button
                    onClick={() => handleDelete(task.EntryID)}
                    className="text-red-500 hover:text-red-700 p-2 rounded-lg hover:bg-red-50 transition-colors"
                  >
                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                    </svg>
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </main>
  );
}
