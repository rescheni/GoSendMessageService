'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Login() {
  const router = useRouter();
  const [token, setToken] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      // 模拟登录验证
      if (token === 'test123') {
        localStorage.setItem('userToken', token);
        router.push('/');
      } else {
        setError('无效的令牌');
      }
    } catch (error: any) {
      setError('登录失败: ' + (error.message || '未知错误'));
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center">
      <div className="w-full max-w-md">
        <div className="bg-white border border-gray-200 rounded-xl p-8 shadow-sm">
          <div className="text-center mb-8">
            <h1 className="text-2xl font-light text-gray-900 mb-2">欢迎登录</h1>
            <p className="text-sm text-gray-500">请输入您的API令牌以继续</p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="token" className="block text-sm font-medium text-gray-700 mb-2">
                API令牌
              </label>
              <div className="relative">
                <input
                  type="text"
                  id="token"
                  value={token}
                  onChange={(e) => setToken(e.target.value)}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-colors"
                  placeholder="请输入您的API令牌"
                  required
                />
                <div className="absolute inset-y-0 right-0 flex items-center pr-3">
                  <svg className="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
                  </svg>
                </div>
              </div>
            </div>

            {error && (
              <div className="text-sm text-red-500 bg-red-50 border border-red-100 rounded-lg px-4 py-2">
                {error}
              </div>
            )}

            <button
              type="submit"
              disabled={isLoading}
              className="w-full px-4 py-2 bg-black text-white rounded-lg text-sm font-medium hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-200 focus:ring-offset-2 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {isLoading ? (
                <div className="flex items-center justify-center">
                  <div className="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin mr-2"></div>
                  登录中...
                </div>
              ) : (
                '登录'
              )}
            </button>
          </form>

          <div className="mt-6 text-center">
            <p className="text-xs text-gray-500">
              测试令牌: <span className="font-mono bg-gray-50 border border-gray-200 px-2 py-1 rounded">test123</span>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
} 