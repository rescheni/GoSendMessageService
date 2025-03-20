'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

interface ConfigSection {
  title: string;
  description: string;
  icon: React.ReactNode;
  fields: {
    key: string;
    label: string;
    type: string;
    placeholder?: string;
    description?: string;
  }[];
}

// 定义配置对象的类型
interface ConfigData {
  [section: string]: {
    [key: string]: string;
  };
}

export default function Settings() {
  const router = useRouter();
  const [activeTab, setActiveTab] = useState('email');
  const [config, setConfig] = useState<ConfigData>({});
  const [isSaving, setIsSaving] = useState(false);

  const configSections: Record<string, ConfigSection> = {
    email: {
      title: '邮件服务',
      description: '配置SMTP邮件服务，用于发送邮件通知',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
        </svg>
      ),
      fields: [
        { key: 'email_address', label: '邮箱地址', type: 'email', placeholder: 'example@126.com', description: '用于发送邮件的邮箱地址' },
        { key: 'username', label: '用户名', type: 'text', description: '邮箱账号用户名' },
        { key: 'smtp_server', label: 'SMTP服务器', type: 'text', placeholder: 'smtp.126.com', description: '邮件服务器地址' },
        { key: 'smtp_port', label: 'SMTP端口', type: 'number', placeholder: '587', description: 'SMTP服务器端口，一般为587(TLS)或465(SSL)' },
        { key: 'auth_code', label: '授权码', type: 'password', description: '邮箱的应用专用密码或授权码' },
        { key: 'encryption', label: '加密方式', type: 'select', description: '连接加密方式' },
        { key: 'sender_name', label: '发件人昵称', type: 'text', description: '发件人显示的名称' },
      ]
    },
    dingding: {
      title: '钉钉',
      description: '配置钉钉机器人，通过webhook发送消息',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
        </svg>
      ),
      fields: [
        { key: 'access_token', label: 'Access Token', type: 'text', description: '钉钉机器人的访问令牌' },
      ]
    },
    feishu: {
      title: '飞书',
      description: '配置飞书机器人，用于发送飞书消息',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
        </svg>
      ),
      fields: [
        { key: 'feishu_app_id', label: 'App ID', type: 'text' },
        { key: 'feishu_app_secret', label: 'App Secret', type: 'password' },
        { key: 'feishu_user_id', label: '用户 ID', type: 'text' },
      ]
    },
    serverJiang: {
      title: 'Server酱',
      description: '配置Server酱，推送消息到微信',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
        </svg>
      ),
      fields: [
        { key: 'server_jiang_key', label: 'Server酱 Key', type: 'text' },
        { key: 'server_jiang_desp', label: '描述', type: 'text' },
      ]
    },
    wxPush: {
      title: '微信推送',
      description: '配置微信推送服务',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z" />
        </svg>
      ),
      fields: [
        { key: 'wx_push_key', label: 'WxPush Key', type: 'text' },
        { key: 'default_uid', label: '默认 UID', type: 'text' },
      ]
    },
    napcat: {
      title: 'Napcat',
      description: '配置Napcat，推送消息到QQ',
      icon: (
        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
        </svg>
      ),
      fields: [
        { key: 'napcat_url', label: 'Napcat URL', type: 'text' },
        { key: 'napcat_token', label: 'Token', type: 'password' },
        { key: 'napcat_qq', label: 'QQ号码', type: 'text' },
      ]
    },
  };

  const handleSave = async () => {
    setIsSaving(true);
    try {
      // TODO: 实现保存配置的API调用
      console.log('保存配置：', config);
      // 模拟API调用延迟
      await new Promise(resolve => setTimeout(resolve, 1000));
    } catch (error) {
      console.error('保存失败：', error);
    } finally {
      setIsSaving(false);
    }
  };

  const handleInputChange = (section: string, key: string, value: string) => {
    setConfig((prev: ConfigData) => ({
      ...prev,
      [section]: {
        ...(prev[section] || {}),
        [key]: value
      }
    }));
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* 顶部导航 */}
        <div className="flex justify-between items-center mb-8">
          <div className="flex items-center space-x-3">
            <button
              onClick={() => router.back()}
              className="p-2 hover:bg-gray-100 rounded-full transition-colors"
            >
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 19l-7-7m0 0l7-7m-7 7h18" />
              </svg>
            </button>
            <h1 className="text-xl font-medium">消息通道配置</h1>
          </div>
          <button
            onClick={handleSave}
            disabled={isSaving}
            className={`px-4 py-2 bg-black text-white rounded-lg text-sm font-medium
                      transition-all duration-200 flex items-center space-x-2
                      ${isSaving ? 'opacity-50 cursor-not-allowed' : 'hover:bg-gray-800'}`}
          >
            {isSaving ? (
              <>
                <svg className="animate-spin h-4 w-4" viewBox="0 0 24 24">
                  <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none" />
                  <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                </svg>
                <span>保存中...</span>
              </>
            ) : (
              <>
                <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                </svg>
                <span>保存配置</span>
              </>
            )}
          </button>
        </div>

        <div className="flex gap-6">
          {/* 左侧导航栏 */}
          <div className="w-56">
            <div className="space-y-1">
              {Object.entries(configSections).map(([key, section]) => (
                <button
                  key={key}
                  onClick={() => setActiveTab(key)}
                  className={`w-full flex items-center space-x-3 px-3 py-2 rounded-lg transition-all duration-200
                            ${activeTab === key 
                              ? 'bg-black text-white' 
                              : 'text-gray-600 hover:bg-gray-100'}`}
                >
                  <span className={activeTab === key ? 'text-white' : 'text-gray-400'}>
                    {section.icon}
                  </span>
                  <div className="text-left">
                    <div className="font-medium text-sm">{section.title}</div>
                    <div className={`text-xs ${activeTab === key ? 'text-gray-300' : 'text-gray-400'}`}>
                      {section.description}
                    </div>
                  </div>
                </button>
              ))}
            </div>
          </div>

          {/* 右侧配置表单 */}
          <div className="flex-1 bg-white rounded-xl p-6 border border-gray-200">
            <div className="mb-6">
              <h2 className="text-lg font-medium mb-1">{configSections[activeTab].title}</h2>
              <p className="text-sm text-gray-500">{configSections[activeTab].description}</p>
            </div>
            
            <div className="space-y-5">
              {configSections[activeTab].fields.map(field => (
                <div key={field.key} className="group">
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    {field.label}
                  </label>
                  {field.type === 'select' ? (
                    <select
                      className="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg
                               focus:ring-2 focus:ring-black focus:ring-offset-1 focus:outline-none
                               transition-colors duration-200"
                      value={config[activeTab]?.[field.key] || ''}
                      onChange={(e) => handleInputChange(activeTab, field.key, e.target.value)}
                    >
                      <option value="TLS">TLS</option>
                      <option value="SSL">SSL</option>
                    </select>
                  ) : (
                    <input
                      type={field.type}
                      placeholder={field.placeholder}
                      className="w-full px-3 py-2 bg-gray-50 border border-gray-200 rounded-lg
                               focus:ring-2 focus:ring-black focus:ring-offset-1 focus:outline-none
                               transition-colors duration-200"
                      value={config[activeTab]?.[field.key] || ''}
                      onChange={(e) => handleInputChange(activeTab, field.key, e.target.value)}
                    />
                  )}
                  {field.description && (
                    <p className="mt-1 text-xs text-gray-500">{field.description}</p>
                  )}
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 