import { createContext, useState, useEffect } from "react";
import authService from "../services/authService";

const AuthContext = createContext({});

export const AuthProvider = ({ children }) => {
    const [auth, setAuth] = useState({});
    const [loading, setLoading] = useState(true);

    // Check if user is logged in (on page load)
    useEffect(() => {
        const initAuth = async () => {
            try {
                const token = localStorage.getItem('accessToken');
                if (token) {
                    // Validate token by refreshing it
                    const response = await authService.refreshToken();

                    // Store new token
                    localStorage.setItem('accessToken', response.accessToken);

                    setAuth({
                        accessToken: response.accessToken,
                        userId: response.userId,
                        isAuthenticated: true
                    });
                }
            } catch (error) {
                // Invalid token, clear it
                localStorage.removeItem('accessToken');
                setAuth({});
            } finally {
                setLoading(false);
            }
        };

        initAuth();
    }, []);

    // Login function
    const login = async (email, password) => {
        try {
            const response = await authService.login({ email, password });

            // Store token in localStorage
            localStorage.setItem('accessToken', response.accessToken);

            // Update auth state
            setAuth({
                accessToken: response.accessToken,
                userId: response.userId,
                isAuthenticated: true
            });

            return { success: true };
        } catch (error) {
            return { success: false, error: error.message || 'Login failed' };
        }
    };

    // Register function
    const register = async (userData) => {
        try {
            const response = await authService.register(userData);
            return { success: true, userId: response.userId };
        } catch (error) {
            return { success: false, error: error.message || 'Registration failed' };
        }
    };

    // Logout function
    const logout = async () => {
        try {
            await authService.logout();
        } finally {
            // Always clear local state even if server logout fails
            localStorage.removeItem('accessToken');
            setAuth({});
        }
    };

    const value = {
        auth,
        setAuth,
        login,
        logout,
        register,
        loading
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};

export default AuthContext;