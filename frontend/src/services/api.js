import axios from 'axios';

// Create an axios instance with defaults
const api = axios.create({
    baseURL: 'http://localhost:8080', // Change to your API URL in production
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: true // Important for cookies (refresh token)
});

// Request interceptor for adding auth token
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('accessToken');
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

// Response interceptor for handling common errors
api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        // If error is 401 and we haven't retried yet
        if (error.response?.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;

            try {
                // Try to refresh the token
                const response = await axios.post(
                    'http://localhost:8080/api/auth/refresh',
                    {},
                    { withCredentials: true }
                );

                const { accessToken } = response.data;
                localStorage.setItem('accessToken', accessToken);

                // Retry the original request with new token
                originalRequest.headers['Authorization'] = `Bearer ${accessToken}`;
                return api(originalRequest);
            } catch (refreshError) {
                // If refresh fails, redirect to login
                window.location.href = '/LoginPage';
                return Promise.reject(refreshError);
            }
        }

        return Promise.reject(error);
    }
);

export default api;