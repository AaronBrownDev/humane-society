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
                console.log("Initializing auth context...");
                const token = localStorage.getItem('accessToken');

                if (token) {
                    console.log("Found token in storage, attempting to validate");

                    try {
                        // Validate token by refreshing it
                        const response = await authService.refreshToken();
                        console.log("Token refresh successful", response);

                        // Store new token
                        localStorage.setItem('accessToken', response.accessToken);

                        setAuth({
                            accessToken: response.accessToken,
                            userId: response.userId,
                            isAuthenticated: true
                        });
                    } catch (refreshError) {
                        console.error("Token refresh failed:", refreshError);
                        // Invalid or expired token, clear it
                        localStorage.removeItem('accessToken');
                        setAuth({});
                    }
                } else {
                    console.log("No token found in storage");
                    setAuth({});
                }
            } catch (error) {
                console.error("Auth initialization error:", error);
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
            console.log("Login attempt for:", email);

            const response = await authService.login({ email, password });
            console.log("Login successful, response data:", response);

            // Store token in localStorage
            localStorage.setItem('accessToken', response.accessToken);

            // Update auth state with proper case and ensure isAuthenticated is set
            setAuth({
                accessToken: response.accessToken,
                userId: response.userId,
                isAuthenticated: true  // Make sure this is set
            });

            return { success: true };
        } catch (error) {
            console.error("Login error:", error);
            return {
                success: false,
                error: typeof error === 'string'
                    ? error
                    : error.message || 'Login failed'
            };
        }
    };

    // Register function
    const register = async (userData) => {
        try {
            console.log("Registration attempt for:", userData.email);

            const response = await authService.register(userData);
            console.log("Registration successful:", response);

            return { success: true, userId: response.userId };
        } catch (error) {
            console.error("Registration error:", error);
            return {
                success: false,
                error: typeof error === 'string'
                    ? error
                    : error.message || 'Registration failed'
            };
        }
    };

    // Logout function
    const logout = async () => {
        try {
            console.log("Logging out...");
            await authService.logout();
            console.log("Logout successful");
        } catch (error) {
            console.error("Logout error:", error);
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