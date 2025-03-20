export const checkAuth = () => {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('apiKey');
  }
  return null;
};

export const getAuthHeader = () => {
  const apiKey = checkAuth();
  return apiKey ? `?api_key=${apiKey}` : '';
}; 