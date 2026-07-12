import axios from 'axios';

export const API_BASE = import.meta.env.VITE_API_BASE || '/api/v1';

axios.defaults.baseURL = API_BASE;

const storedToken = localStorage.getItem('accessToken');
if (storedToken) {
  axios.defaults.headers.common['Authorization'] = `Bearer ${storedToken}`;
}

export function getCurrentUserId() {
  return localStorage.getItem('userId') || '';
}

export function getCurrentUserFullName() {
  return localStorage.getItem('userFullName') || '';
}

export function storeSession({ access_token, refresh_token, user }) {
  localStorage.setItem('accessToken', access_token);
  localStorage.setItem('refreshToken', refresh_token);
  if (user) {
    localStorage.setItem('userId', user.id);
    localStorage.setItem('userFullName', user.full_name);
  }
  axios.defaults.headers.common['Authorization'] = `Bearer ${access_token}`;
}

export function clearSession() {
  localStorage.removeItem('accessToken');
  localStorage.removeItem('refreshToken');
  localStorage.removeItem('userId');
  localStorage.removeItem('userFullName');
  delete axios.defaults.headers.common['Authorization'];
}

// Registered once at module load (not per React render) so a refresh is only ever attempted once per failed request.
axios.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response?.status === 401 && !originalRequest._retried) {
      originalRequest._retried = true;
      const refreshToken = localStorage.getItem('refreshToken');
      if (refreshToken) {
        try {
          const res = await axios.post('/refresh', { refresh_token: refreshToken }, {
            headers: { Authorization: undefined },
          });
          localStorage.setItem('accessToken', res.data.access_token);
          if (res.data.refresh_token) {
            localStorage.setItem('refreshToken', res.data.refresh_token);
          }
          axios.defaults.headers.common['Authorization'] = `Bearer ${res.data.access_token}`;
          originalRequest.headers['Authorization'] = `Bearer ${res.data.access_token}`;
          return axios(originalRequest);
        } catch (err) {
          clearSession();
          window.dispatchEvent(new Event('auth:unauthenticated'));
        }
      } else {
        clearSession();
        window.dispatchEvent(new Event('auth:unauthenticated'));
      }
    }
    return Promise.reject(error);
  }
);

export default axios;
