import api from './api';

// Base API URL
const AUTH_URL = '/api/auth';

/**
 * Attempts to log in a user with the provided credentials
 * @param {Object} credentials - User login credentials
 * @param {string} credentials.email - User's email address
 * @param {string} credentials.password - User's password
 * @returns {Promise<Object>} Login response containing tokens and user info
 */
export const login = async (credentials) => {
    try {
        console.log('Login attempt with:', credentials.email);

        const response = await api.post(`${AUTH_URL}/login`, credentials);

        console.log('Login successful, received tokens');
        return response.data;
    } catch (error) {
        console.error('Login failed:', error.message);
        throw error.response?.data?.error || error.message || 'Login failed';
    }
};

/**
 * Registers a new user with the provided data
 * @param {Object} userData - New user registration data
 * @returns {Promise<Object>} Registration response
 */
export const register = async (userData) => {
    try {
        console.log('Sending registration data:', {
            ...userData,
            password: '[REDACTED]' // Don't log passwords
        });

        const response = await api.post(`${AUTH_URL}/register`, userData);

        console.log('Registration successful');
        return response.data;
    } catch (error) {
        console.error('Registration error:', error);

        // Extract proper error message
        const errorMessage = error.response?.data?.error ||
            error.message ||
            'Registration failed';

        throw errorMessage;
    }
};

/**
 * Logs out the current user by invalidating their tokens
 * @returns {Promise<Object>} Logout response
 */
export const logout = async () => {
    try {
        console.log('Attempting logout');
        const response = await api.post(`${AUTH_URL}/logout`);

        console.log('Logout successful');
        return response.data;
    } catch (error) {
        console.error('Logout error:', error);

        // Still clear local state even if server logout fails
        return { message: 'Logged out locally' };
    }
};

/**
 * Refreshes the access token using the refresh token cookie
 * @returns {Promise<Object>} New tokens
 */
export const refreshToken = async () => {
    try {
        console.log('Attempting to refresh token');

        const response = await api.post(`${AUTH_URL}/refresh`);

        console.log('Token refresh successful');
        return response.data;
    } catch (error) {
        console.error('Token refresh failed:', error);
        throw error.response?.data?.error || error.message || 'Token refresh failed';
    }
};

/**
 * Retrieves the current user's profile information
 * @returns {Promise<Object>} User profile data
 */
export const getUserProfile = async () => {
    try {
        const response = await api.get('/api/users/profile');
        return response.data;
    } catch (error) {
        throw error.response?.data?.error || error.message || 'Failed to get user profile';
    }
};

// Export as both named exports and default object
const authService = {
    login,
    register,
    logout,
    refreshToken,
    getUserProfile
};

export default authService;