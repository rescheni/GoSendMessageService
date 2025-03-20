'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

type SendMethod = 'email' | 'dingding' | 'server_jiang' | 'feishu' | 'wxpusher' | 'napcat_qq';

interface SendConfig {
  title: string;
  endpoint: string;
  fields: {
    key: string;
    label: string;
    type: string;
    required?: boolean;
  }[];
}

export default function SendMessage() {
  const router = useRouter();
  const [message, setMessage] = useState('');
  const [selectedMethod, setSelectedMethod] = useState<SendMethod>('email');
  const [recipientData, setRecipientData] = useState<Record<string, string>>({});
  const [isSending, setIsSending] = useState(false);
  const [error, setError] = useState('');

  const sendConfigs: Record<SendMethod, SendConfig> = {
    email: {
      title: '邮件发送',
      endpoint: '/api/send/email',
      fields: [
        { key: 'to', label: '收件人', type: 'email', required: true },
        { key: 'subject', label: '主题', type: 'text', required: true }
      ]
    },
    dingding: {
      title: '钉钉发送',
      endpoint: '/api/send/dingding',
      fields: []  // 钉钉使用配置文件中的 webhook
    },
    server_jiang: {
      title: 'Server酱发送',
      endpoint: '/api/send/server_jiang',
      fields: [
        { key: 'title', label: '标题', type: 'text', required: true }
      ]
    },
    feishu: {
      title: '飞书发送',
      endpoint: '/api/send/feishu',
      fields: []  // 飞书使用配置文件中的用户ID
    },
    wxpusher: {
      title: '微信发送',
      endpoint: '/api/send/wxpusher',
      fields: [
        { key: 'uids', label: '用户ID', type: 'text', required: true }
      ]
    },
    napcat_qq: {
      title: 'QQ发送',
      endpoint: '/api/send/napcat_qq',
      fields: []  // QQ使用配置文件中的QQ号
    }
  };

  const handleSend = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSending(true);
    setError('');

    try {
      const apiKey = localStorage.getItem('apiKey');
      if (!apiKey) {
        router.push('/login');
        return;
      }

      console.log('发送请求:', {
        method: selectedMethod,
        message,
        recipientData
      });

      const config = sendConfigs[selectedMethod];
      const response = await fetch(`${config.endpoint}?api_key=${apiKey}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          message,
          ...recipientData
        })
      });

      const data = await response.json();
      console.log('服务器响应:', data);

      if (!response.ok) {
        throw new Error(data.error || `发送失败: ${response.status}`);
      }

      alert('发送成功！');
      setMessage('');
      setRecipientData({});
    } catch (error: any) {
      console.error('发送失败：', error);
      setError(error.message || '发送失败');
    } finally {
      setIsSending(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 p-6">
      <div className="max-w-2xl mx-auto bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-xl font-medium">发送消息</h1>
          <button
            onClick={() => router.push('/settings')}
            className="px-4 py-2 text-gray-600 hover:text-gray-800 flex items-center space-x-2"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"/>
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
            </svg>
            <span>设置</span>
          </button>
        </div>

        <form onSubmit={handleSend} className="space-y-6">
          {/* 发送方式选择 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              发送方式
            </label>
            <div className="grid grid-cols-2 md:grid-cols-3 gap-2">
              {(Object.entries(sendConfigs) as [SendMethod, SendConfig][]).map(([method, config]) => (
                <button
                  key={method}
                  type="button"
                  onClick={() => {
                    setSelectedMethod(method);
                    setRecipientData({}); // 切换方式时清空接收方数据
                  }}
                  className={`px-4 py-2 rounded-lg text-sm font-medium transition-colors
                            ${selectedMethod === method 
                              ? 'bg-black text-white' 
                              : 'bg-gray-100 text-gray-600 hover:bg-gray-200'}`}
                >
                  {config.title}
                </button>
              ))}
            </div>
          </div>

          {/* 接收方信息 */}
          {sendConfigs[selectedMethod].fields.map(field => (
            <div key={field.key}>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                {field.label}
              </label>
              <input
                type={field.type}
                value={recipientData[field.key] || ''}
                onChange={(e) => setRecipientData(prev => ({
                  ...prev,
                  [field.key]: e.target.value
                }))}
                className="w-full px-4 py-2 border border-gray-200 rounded-lg
                         focus:ring-2 focus:ring-black focus:ring-offset-1 focus:outline-none"
                required={field.required}
              />
            </div>
          ))}

          {/* 消息内容 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              消息内容
            </label>
            <textarea
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              className="w-full h-32 px-4 py-2 border border-gray-200 rounded-lg
                       focus:ring-2 focus:ring-black focus:ring-offset-1 focus:outline-none"
              placeholder="请输入要发送的消息内容"
              required
            />
          </div>

          {error && (
            <div className="text-sm text-red-500 bg-red-50 border border-red-100 rounded-lg px-4 py-2">
              {error}
            </div>
          )}

          <div className="flex justify-end">
            <button
              type="submit"
              disabled={isSending}
              className="px-6 py-2 bg-black text-white rounded-lg text-sm font-medium
                       hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-200
                       focus:ring-offset-2 transition-colors disabled:opacity-50
                       disabled:cursor-not-allowed"
            >
              {isSending ? (
                <div className="flex items-center">
                  <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></div>
                  发送中...
                </div>
              ) : (
                '发送'
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
} 