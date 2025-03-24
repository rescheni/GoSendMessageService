'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { getAuthHeader } from '@/src/utils/auth';

interface TaskForm {
  title: string;
  content: string;
  sendMethod: string;
  scheduledTime?: string;
  cronExpression?: string;
  scheduleType?: 'datetime' | 'cron';
  config: Record<string, any>;
  cronName: string;
  toUser?: string;
}

export default function CreateTaskPage() {
  const router = useRouter();
  const [formData, setFormData] = useState<TaskForm>({
    title: '',
    content: '',
    sendMethod: '',
    config: {},
    cronName: '',
    cronExpression: '',
    toUser: ''
  });
  const [step, setStep] = useState(1); // 1: 基本信息, 2: 配置信息

  const sendMethods = [
    { value: 'email', label: '邮件' },
    { value: 'dingding', label: '钉钉' },
    { value: 'server_jiang', label: 'Server酱' },
    { value: 'feishu', label: '飞书' },
    { value: 'wx_push', label: '微信推送' },
    { value: 'napcat_qq', label: 'QQ' }
  ];

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (!formData.sendMethod || !formData.content || !formData.title || 
          !formData.cronExpression || !formData.cronName) {
        alert('请填写所有必要信息');
        return;
      }

      // 如果是邮件方式，验证收件人
      if (formData.sendMethod === 'email' && !formData.toUser) {
        alert('请填写收件人邮箱');
        return;
      }

      const apiKey = localStorage.getItem('apiKey');
      const requestBody = {
        api_key: apiKey,
        cron_expr: `0 ${formData.cronExpression}`,
        cron_name: formData.cronName,
        message: formData.content,
        title: formData.title,
        task_type: formData.sendMethod,
        to_user: formData.toUser || undefined,  // 仅在有值时添加
        is_open: true
      };

      const response = await fetch(`/api/cron/set?api_key=${apiKey}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(requestBody)
      });

      if (!response.ok) {
        const text = await response.text();
        throw new Error(text);
      }

      alert('定时任务创建成功');
      router.push('/');
    } catch (error: any) {
      console.error('创建任务失败:', error);
      alert('创建任务失败: ' + (error.message || '未知错误'));
    }
  };

  const handleDirectSend = async () => {
    try {
      if (!formData.sendMethod || !formData.content) {
        alert('请填写完整信息');
        return;
      }

      const response = await fetch(`/api/send/${formData.sendMethod}${getAuthHeader()}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          api_key: localStorage.getItem('apiKey'),
          title: formData.title,
          message: formData.content
        })
      });

      if (!response.ok) {
        throw new Error('发送失败');
      }

      alert('发送成功');
      router.push('/'); // 发送成功后返回主页
    } catch (error) {
      console.error('发送失败:', error);
      alert('发送失败，请重试');
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 py-12">
      <div className="max-w-3xl mx-auto px-4">
        <div className="mb-8">
          <button
            onClick={() => router.back()}
            className="flex items-center text-gray-600 hover:text-gray-900"
          >
            <svg className="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
            返回
          </button>
        </div>

        <div className="bg-white rounded-xl shadow-sm p-8">
          <h1 className="text-2xl font-medium text-gray-900 mb-8">创建新任务</h1>
          
          <div className="mb-8">
            <div className="flex items-center justify-center">
              <div className={`w-10 h-10 rounded-full flex items-center justify-center ${
                step >= 1 ? 'bg-black text-white' : 'bg-gray-200 text-gray-600'
              }`}>1</div>
              <div className={`h-1 w-20 ${step >= 2 ? 'bg-black' : 'bg-gray-200'}`}></div>
              <div className={`w-10 h-10 rounded-full flex items-center justify-center ${
                step >= 2 ? 'bg-black text-white' : 'bg-gray-200 text-gray-600'
              }`}>2</div>
            </div>
          </div>

          <form onSubmit={handleSubmit}>
            {step === 1 ? (
              <div className="space-y-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    发送方式
                  </label>
                  <select
                    value={formData.sendMethod}
                    onChange={(e) => setFormData({ ...formData, sendMethod: e.target.value })}
                    className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                  >
                    <option value="">请选择发送方式</option>
                    {sendMethods.map((method) => (
                      <option key={method.value} value={method.value}>
                        {method.label}
                      </option>
                    ))}
                  </select>
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    消息标题
                  </label>
                  <input
                    type="text"
                    value={formData.title}
                    onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                    className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                    placeholder="请输入消息标题"
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    消息内容
                  </label>
                  <textarea
                    value={formData.content}
                    onChange={(e) => setFormData({ ...formData, content: e.target.value })}
                    rows={4}
                    className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                    placeholder="请输入消息内容"
                  />
                </div>

                <div className="flex justify-end space-x-4">
                  <button
                    type="button"
                    onClick={handleDirectSend}
                    className="px-6 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors"
                  >
                    立即发送
                  </button>
                  <button
                    type="button"
                    onClick={() => setStep(2)}
                    className="px-6 py-2 bg-black text-white rounded-lg hover:bg-gray-800 transition-colors"
                  >
                    下一步
                  </button>
                </div>
              </div>
            ) : (
              <div className="space-y-6">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    任务名称
                  </label>
                  <input
                    type="text"
                    value={formData.cronName}
                    onChange={(e) => setFormData({ ...formData, cronName: e.target.value })}
                    className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                    placeholder="请输入任务名称"
                    required
                  />
                </div>

                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-2">
                    Cron表达式
                  </label>
                  <input
                    type="text"
                    value={formData.cronExpression}
                    onChange={(e) => setFormData({ ...formData, cronExpression: e.target.value })}
                    className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                    placeholder="*/5 * * * *"
                    required
                  />
                  <p className="mt-1 text-xs text-gray-500">
                    示例: "*/5 * * * *" 表示每5分钟执行一次
                  </p>
                </div>

                {formData.sendMethod === 'email' && (
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      收件人邮箱
                    </label>
                    <input
                      type="email"
                      value={formData.toUser}
                      onChange={(e) => setFormData({ ...formData, toUser: e.target.value })}
                      className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200"
                      placeholder="请输入收件人邮箱"
                      required={formData.sendMethod === 'email'}
                    />
                  </div>
                )}

                <div className="flex justify-between">
                  <button
                    type="button"
                    onClick={() => setStep(1)}
                    className="px-6 py-2 border border-gray-200 rounded-lg hover:bg-gray-50 transition-colors"
                  >
                    上一步
                  </button>
                  <button
                    type="submit"
                    className="px-6 py-2 bg-black text-white rounded-lg hover:bg-gray-800 transition-colors"
                  >
                    创建任务
                  </button>
                </div>
              </div>
            )}
          </form>
        </div>
      </div>
    </div>
  );
}