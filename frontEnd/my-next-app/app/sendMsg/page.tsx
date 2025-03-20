'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function SendMessage() {
  const router = useRouter();
  const [message, setMessage] = useState('');
  const [useAI, setUseAI] = useState(false);
  const [prompt, setPrompt] = useState('');
  const [isSending, setIsSending] = useState(false);

  const handleSend = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSending(true);
    try {
      // TODO: 实现发送消息逻辑
      console.log('发送信息：', {
        message,
        useAI,
        prompt: useAI ? prompt : undefined
      });
      // 发送成功后清空表单
      setMessage('');
      setPrompt('');
    } catch (error) {
      console.error('发送失败：', error);
    } finally {
      setIsSending(false);
    }
  };

  const handleOpenSettings = () => {
    router.push('/settings');
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-2xl mx-auto bg-white rounded-lg shadow-md p-6">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold">发送消息</h1>
          <button
            onClick={handleOpenSettings}
            className="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg
                     transition-all duration-300 ease-in-out transform
                     hover:bg-gray-200 hover:scale-105 active:scale-95
                     flex items-center space-x-2"
          >
            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
            <span>设置</span>
          </button>
        </div>
        <form onSubmit={handleSend} className="space-y-4">
          {/* 消息输入框 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              消息内容
            </label>
            <textarea
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              className="w-full h-32 p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              placeholder="请输入要发送的消息..."
              required
            />
          </div>

          {/* AI 开关 - 修复点击效果 */}
          <div className="flex items-center space-x-3">
            <div className="relative w-14 h-7 flex items-center">
              <input
                type="checkbox"
                id="useAI"
                checked={useAI}
                onChange={(e) => setUseAI(e.target.checked)}
                className="sr-only"
              />
              <div
                className={`w-14 h-7 rounded-full transition-all duration-300 ease-in-out cursor-pointer
                          ${useAI ? 'bg-blue-500' : 'bg-gray-200'}`}
                onClick={() => setUseAI(!useAI)}
              >
                <div
                  className={`absolute w-6 h-6 rounded-full bg-white shadow-md transform transition-all duration-300 ease-in-out
                            ${useAI ? 'translate-x-7' : 'translate-x-1'} top-0.5`}
                ></div>
              </div>
            </div>
            <label 
              htmlFor="useAI" 
              className="text-sm font-medium cursor-pointer select-none
                        transition-all duration-300 ease-in-out
                        hover:text-blue-600"
            >
              使用 AI 优化消息
            </label>
          </div>

          {/* AI 提示词输入框 */}
          <div 
            className={`overflow-hidden transition-all duration-500 ease-in-out
                      ${useAI ? 'max-h-[200px] opacity-100' : 'max-h-0 opacity-0'}`}
          >
            <div 
              className={`transform transition-all duration-500 ease-in-out
                        ${useAI ? 'translate-y-0' : '-translate-y-4'}`}
            >
              <label className="block text-sm font-medium text-gray-700 mb-2">
                AI 提示词
              </label>
              <textarea
                value={prompt}
                onChange={(e) => setPrompt(e.target.value)}
                className="w-full h-24 p-3 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                          transition-all duration-300 ease-in-out"
                placeholder="请输入 AI 提示词，例如：'使用正式的语气'"
              />
            </div>
          </div>

          {/* 发送按钮 */}
          <div className="flex justify-end">
            <button
              type="submit"
              disabled={isSending}
              className={`px-6 py-2 bg-blue-500 text-white rounded-lg 
                        transition-all duration-300 ease-in-out transform hover:scale-105
                        hover:bg-blue-600 hover:shadow-lg active:scale-95
                        ${isSending ? 'opacity-50 cursor-not-allowed scale-100' : ''}`}
            >
              {isSending ? '发送中...' : '发送'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
} 