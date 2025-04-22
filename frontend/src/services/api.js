import axios from 'axios';

// Create an axios instance with defaults
const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080', // Use environment variable if available
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: true // Important for cookies (refresh token)
});

// Request interceptor for adding auth token
api.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('accessToken');

        // Log requests in development
        if (import.meta.env.DEV) {
            console.log(`API Request: ${config.method.toUpperCase()} ${config.url}`);
        }

        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

// Response interceptor for handling common errors
api.interceptors.response.use(
    (response) => {
        // Log responses in development
        if (import.meta.env.DEV) {
            console.log(`API Response: ${response.status} ${response.config.url}`);
        }
        return response;
    },
    async (error) => {
        const originalRequest = error.config;

        // Add useful debugging in development
        if (import.meta.env.DEV) {
            console.error(`API Error: ${error.response?.status} ${originalRequest?.url}`);
            console.error(error.response?.data || error.message);
        }

        // If error is 401 (Unauthorized) and we haven't retried yet
        if (error.response?.status === 401 && !originalRequest._retry) {
            // Mark this request to prevent infinite retry loops
            originalRequest._retry = true;

            try {
                console.log("Attempting to refresh token");

                // Try to refresh the token
                const response = await axios.post(
                    `${import.meta.env.VITE_API_URL || 'http://localhost:8080'}/api/auth/refresh`,
                    {},
                    { withCredentials: true }
                );

                const { accessToken } = response.data;

                if (accessToken) {
                    console.log("Token refresh successful");
                    localStorage.setItem('accessToken', accessToken);

                    // Retry the original request with new token
                    originalRequest.headers['Authorization'] = `Bearer ${accessToken}`;
                    return api(originalRequest);
                } else {
                    console.error("Token refresh response missing accessToken");
                    // Clear invalid token
                    localStorage.removeItem('accessToken');
                    throw new Error("Token refresh failed");
                }
            } catch (refreshError) {
                console.error("Token refresh error:", refreshError);

                // Clear invalid token
                localStorage.removeItem('accessToken');

                // Let consuming code handle redirection to login
                return Promise.reject({
                    ...error,
                    isAuthError: true,
                    message: "Authentication expired"
                });
            }
        }

        // Format error message for easier consumption
        if (error.response?.data?.error) {
            error.message = error.response.data.error;
        }

        return Promise.reject(error);
    }
);

export default api;