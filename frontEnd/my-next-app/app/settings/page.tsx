'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';

interface ConfigSection {
  title: string;
  fields: {
    key: string;
    label: string;
    type: string;
    placeholder?: string;
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
  // 使用正确的类型定义初始状态
  const [config, setConfig] = useState<ConfigData>({});
  const [isSaving, setIsSaving] = useState(false);

  const configSections: Record<string, ConfigSection> = {
    email: {
      title: '邮件服务配置',
      fields: [
        { key: 'email_address', label: '邮箱地址', type: 'email', placeholder: 'example@126.com' },
        { key: 'username', label: '用户名', type: 'text' },
        { key: 'smtp_server', label: 'SMTP服务器', type: 'text', placeholder: 'smtp.126.com' },
        { key: 'smtp_port', label: 'SMTP端口', type: 'number', placeholder: '587' },
        { key: 'auth_code', label: '授权码', type: 'password' },
        { key: 'encryption', label: '加密方式', type: 'select' },
        { key: 'sender_name', label: '发件人昵称', type: 'text' },
      ]
    },
    dingding: {
      title: '钉钉配置',
      fields: [
        { key: 'access_token', label: 'Access Token', type: 'text' },
      ]
    },
    feishu: {
      title: '飞书配置',
      fields: [
        { key: 'feishu_app_id', label: 'App ID', type: 'text' },
        { key: 'feishu_app_secret', label: 'App Secret', type: 'password' },
        { key: 'feishu_user_id', label: '用户 ID', type: 'text' },
      ]
    },
    serverJiang: {
      title: 'Server酱配置',
      fields: [
        { key: 'server_jiang_key', label: 'Server酱 Key', type: 'text' },
        { key: 'server_jiang_desp', label: '描述', type: 'text' },
      ]
    },
    wxPush: {
      title: '微信推送配置',
      fields: [
        { key: 'wx_push_key', label: 'WxPush Key', type: 'text' },
        { key: 'default_uid', label: '默认 UID', type: 'text' },
      ]
    },
    napcat: {
      title: 'Napcat配置',
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
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-4xl mx-auto bg-white rounded-lg shadow-md p-6">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold">消息发送配置</h1>
          <button
            onClick={() => router.back()}
            className="px-4 py-2 text-gray-600 hover:text-gray-800"
          >
            返回
          </button>
        </div>

        <div className="flex space-x-4">
          {/* 左侧标签栏 */}
          <div className="w-1/4 space-y-2">
            {Object.entries(configSections).map(([key, section]) => (
              <button
                key={key}
                onClick={() => setActiveTab(key)}
                className={`w-full text-left px-4 py-2 rounded-lg transition-all duration-200
                          ${activeTab === key 
                            ? 'bg-blue-500 text-white' 
                            : 'hover:bg-gray-100'}`}
              >
                {section.title}
              </button>
            ))}
          </div>

          {/* 右侧配置表单 */}
          <div className="w-3/4 bg-gray-50 rounded-lg p-6">
            <h2 className="text-xl font-semibold mb-4">
              {configSections[activeTab].title}
            </h2>
            <div className="space-y-4">
              {configSections[activeTab].fields.map(field => (
                <div key={field.key}>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    {field.label}
                  </label>
                  {field.type === 'select' ? (
                    <select
                      className="w-full p-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
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
                      className="w-full p-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
                      value={config[activeTab]?.[field.key] || ''}
                      onChange={(e) => handleInputChange(activeTab, field.key, e.target.value)}
                    />
                  )}
                </div>
              ))}
            </div>

            <div className="mt-6 flex justify-end">
              <button
                onClick={handleSave}
                disabled={isSaving}
                className={`px-6 py-2 bg-blue-500 text-white rounded-lg
                          transition-all duration-300 ease-in-out transform
                          hover:bg-blue-600 hover:scale-105 active:scale-95
                          ${isSaving ? 'opacity-50 cursor-not-allowed' : ''}`}
              >
                {isSaving ? '保存中...' : '保存配置'}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 