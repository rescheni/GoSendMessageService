'use client';

import { useState, useEffect } from 'react';
import { getAuthHeader } from '@/src/utils/auth';

interface CreateTaskDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (taskData: any) => void;
  onMethodSelect: (method: string) => void;
}

// 定义发送方式和对应的接口路径
const SEND_METHODS = {
  'email': '/send/email',
  'dingding': '/send/dingding',
  'server_jiang': '/send/server_jiang',
  'feishu': '/send/feishu',
  'wxpusher': '/send/wxpusher',
  'napcat_qq': '/send/napcat_qq'
};

export default function CreateTaskDialog({ isOpen, onClose, onSubmit }: CreateTaskDialogProps) {
  const [content, setContent] = useState('');
  const [title, setTitle] = useState('');
  const [sendMethod, setSendMethod] = useState('');
  const [isVisible, setIsVisible] = useState(false);
  const [isSending, setIsSending] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    if (isOpen) {
      setIsVisible(true);
    } else {
      setIsVisible(false);
    }
  }, [isOpen]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSending(true);
    setError('');

    try {
      const endpoint = SEND_METHODS[sendMethod as keyof typeof SEND_METHODS];
      if (!endpoint) {
        throw new Error('请选择发送方式');
      }

      const response = await fetch(`/api${endpoint}${getAuthHeader()}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          api_key: localStorage.getItem('apiKey'),
          message: content,
          title
        })
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || '发送失败');
      }

      console.log('发送成功');
      handleClose();
    } catch (error: any) {
      console.error('发送错误:', error);
      setError(error.message || '发送失败');
    } finally {
      setIsSending(false);
    }
  };

  const handleClose = () => {
    setIsVisible(false);
    setTimeout(onClose, 300);
  };

  if (!isOpen) return null;

  return (
    <div 
      className={`fixed inset-0 bg-black/30 backdrop-blur-sm transition-opacity duration-300 ease-out ${
        isVisible ? 'opacity-100' : 'opacity-0'
      }`}
    >
      <div className="fixed inset-0 flex items-center justify-center">
        <div 
          className={`bg-white rounded-xl shadow-xl w-full max-w-2xl transform transition-all duration-300 ease-out ${
            isVisible ? 'translate-y-0 opacity-100' : 'translate-y-4 opacity-0'
          }`}
        >
          <div className="p-6">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-xl font-medium text-gray-900">新建任务</h2>
              <button
                onClick={handleClose}
                className="text-gray-400 hover:text-gray-500 transition-colors duration-200"
              >
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                </svg>
              </button>
            </div>

            <form onSubmit={handleSubmit} className="space-y-6">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  发送方式
                </label>
                <select
                  value={sendMethod}
                  onChange={(e) => setSendMethod(e.target.value)}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                  required
                >
                  <option value="">请选择发送方式</option>
                  <option value="email">邮件</option>
                  <option value="dingding">钉钉</option>
                  <option value="server_jiang">Server酱</option>
                  <option value="feishu">飞书</option>
                  <option value="wxpusher">微信推送</option>
                  <option value="napcat_qq">QQ</option>
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  内容标题
                </label>
                <input
                  type="text"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                  placeholder="请输入消息标题"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  消息内容
                </label>
                <textarea
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  rows={6}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200 resize-none"
                  placeholder="请输入要发送的消息内容..."
                  required
                />
              </div>

              {error && (
                <div className="text-sm text-red-500 bg-red-50 border border-red-100 rounded-lg px-4 py-2">
                  {error}
                </div>
              )}

              <div className="flex justify-end space-x-3">
                <button
                  type="button"
                  onClick={handleClose}
                  className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-gray-200 focus:ring-offset-2 transition-all duration-200"
                >
                  取消
                </button>
                <button
                  type="submit"
                  disabled={isSending}
                  className="px-4 py-2 text-sm font-medium text-white bg-black rounded-lg hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-200 focus:ring-offset-2 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {isSending ? (
                    <div className="flex items-center justify-center">
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
      </div>
    </div>
  );
} 