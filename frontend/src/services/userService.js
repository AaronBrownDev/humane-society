import api from './api';

/**
 * Retrieves all user accounts from the system
 * @returns {Promise<Array>} List of user accounts
 */
export const getAllUsers = async () => {
    try {
        const response = await api.get('/api/users');
        return response.data;
    } catch (error) {
        console.error('Error fetching users:', error);
        throw error.response?.data?.error || error.message || 'Failed to fetch users';
    }
};

// Export as both named export and default object
const userService = {
    getAllUsers
};

export default userService;