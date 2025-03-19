import { useState, useEffect } from 'react';

interface CreateTaskDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onSubmit: (taskData: any) => void;
  onMethodSelect: (method: string) => void;
}

type SendMethod = 'email' | 'sms' | 'wechat' | 'dingtalk' | 'feishu';

export default function CreateTaskDialog({ isOpen, onClose, onSubmit, onMethodSelect }: CreateTaskDialogProps) {
  const [content, setContent] = useState('');
  const [sendMethod, setSendMethod] = useState<SendMethod | ''>('');
  const [isVisible, setIsVisible] = useState(false);

  useEffect(() => {
    if (isOpen) {
      setIsVisible(true);
    } else {
      setIsVisible(false);
    }
  }, [isOpen]);

  const handleMethodChange = (method: SendMethod) => {
    setSendMethod(method);
    onMethodSelect(method);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({ content, sendMethod });
    onClose();
  };

  const handleClose = () => {
    setIsVisible(false);
    setTimeout(onClose, 300); // 等待动画完成后再关闭
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
                  onChange={(e) => handleMethodChange(e.target.value as SendMethod)}
                  className="w-full px-4 py-2 border border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-gray-200 focus:border-transparent transition-all duration-200"
                >
                  <option value="">请选择发送方式</option>
                  <option value="email">邮件</option>
                  <option value="sms">短信</option>
                  <option value="wechat">微信公众号</option>
                  <option value="dingtalk">钉钉API</option>
                  <option value="feishu">飞书API</option>
                </select>
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
                />
              </div>

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
                  className="px-4 py-2 text-sm font-medium text-white bg-black rounded-lg hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-200 focus:ring-offset-2 transition-all duration-200"
                >
                  下一步
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
} 