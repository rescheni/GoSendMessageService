import { useState, useEffect } from 'react';

interface TaskConfigPanelProps {
  isOpen: boolean;
  onClose: () => void;
  sendMethod: string;
  onSubmit: (config: any) => void;
}

export default function TaskConfigPanel({ isOpen, onClose, sendMethod, onSubmit }: TaskConfigPanelProps) {
  const [recipient, setRecipient] = useState('');
  const [datetime, setDatetime] = useState('');
  const [cronExpression, setCronExpression] = useState('');
  const [isRepeat, setIsRepeat] = useState(false);
  const [isVisible, setIsVisible] = useState(false);

  useEffect(() => {
    if (isOpen) {
      setIsVisible(true);
    } else {
      setIsVisible(false);
    }
  }, [isOpen]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      recipient,
      datetime,
      cronExpression,
      isRepeat
    });
    handleClose();
  };

  const handleClose = () => {
    setIsVisible(false);
    setTimeout(onClose, 300); // 等待动画完成后再关闭
  };

  return (
    <div 
      className={`fixed inset-y-0 right-0 w-96 bg-white shadow-xl transform transition-all duration-300 ease-out ${
        isVisible ? 'translate-x-0 opacity-100' : 'translate-x-full opacity-0'
      } ${!isOpen && !isVisible ? 'pointer-events-none' : ''}`}
    >
      <div className="h-full flex flex-col">
        <div className="p-6 border-b border-gray-200">
          <div className="flex justify-between items-center">
            <h3 className="text-lg font-medium text-gray-900">发送配置</h3>
            <button
              onClick={handleClose}
              className="text-gray-400 hover:text-gray-500 transition-colors duration-200"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <div className="flex-1 overflow-y-auto p-6">
          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                发送方式
              </label>
              <div className="px-4 py-2 bg-gray-50 border border-gray-200 rounded-lg text-sm text-gray-700">
                {sendMethod === 'email' && '邮件'}
                {sendMethod === 'sms' && '短信'}
                {sendMethod === 'wechat' && '微信公众号'}
                {sendMethod === 'dingtalk' && '钉钉API'}
                {sendMethod === 'feishu' && '飞书API'}
              </div>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                收件人
              </label>
              <input
                type="text"
                value={recipient}
                onChange={(e) => setRecipient(e.target.value)}
                className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                placeholder="请输入收件人信息"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                发送时间
              </label>
              <div className="space-y-2">
                <input
                  type="datetime-local"
                  value={datetime}
                  onChange={(e) => setDatetime(e.target.value)}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                />
                <div className="text-xs text-gray-500">
                  或使用 cron 表达式：
                  <input
                    type="text"
                    value={cronExpression}
                    onChange={(e) => setCronExpression(e.target.value)}
                    className="w-full mt-1 px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                    placeholder="例如：0 0 12 * * ?"
                  />
                </div>
              </div>
            </div>

            <div className="flex items-center">
              <input
                type="checkbox"
                id="repeat"
                checked={isRepeat}
                onChange={(e) => setIsRepeat(e.target.checked)}
                className="h-4 w-4 text-gray-900 focus:ring-gray-200 border-gray-300 rounded"
              />
              <label htmlFor="repeat" className="ml-2 text-sm text-gray-700">
                重复发送
              </label>
            </div>

            <div className="pt-4">
              <button
                type="submit"
                className="w-full px-4 py-2 bg-black text-white rounded-lg text-sm font-medium hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-200 focus:ring-offset-2 transition-all duration-200"
              >
                确认配置
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
} 