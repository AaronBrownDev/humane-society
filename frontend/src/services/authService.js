import api from './api';

// Base API URL
const AUTH_URL = '/api/auth';

export const login = async (credentials) => {
    try {
        const response = await api.post(`${AUTH_URL}/login`, credentials);
        return response.data;
    } catch (error) {
        throw error.response?.data?.error || 'Login failed';
    }
};

export const register = async (userData) => {
    try {
        console.log('Sending registration data:', userData);

        const response = await api.post(`${AUTH_URL}/register`, userData);
        return response.data;
    } catch (error) {
        console.error('Registration error:', error);
        throw error.response?.data?.error || 'Registration failed';
    }
};

export const logout = async () => {
    try {
        const response = await api.post(`${AUTH_URL}/logout`);
        return response.data;
    } catch (error) {
        console.error('Logout error:', error);
        // Still clear local state even if server logout fails
        return { message: 'Logged out locally' };
    }
};

export const refreshToken = async () => {
    try {
        const response = await api.post(`${AUTH_URL}/refresh`);
        return response.data;
    } catch (error) {
        throw error.response?.data?.error || 'Token refresh failed';
    }
};

// Export as an object for easier imports
const authService = {
    login,
    register,
    logout,
    refreshToken
};

export default authService;