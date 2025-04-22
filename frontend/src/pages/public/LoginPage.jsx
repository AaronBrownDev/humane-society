import React, { useState, useRef, useEffect } from "react";
import { useNavigate, useLocation, NavLink } from "react-router-dom";
import useAuth from "../../hooks/useAuth";
import "../../styles/Login.css";

export default function LoginPage() {
    const { auth, login } = useAuth();
    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.from?.pathname || "/Adopt"; // Default to Adopt if no redirect location

    const userRef = useRef();
    const errRef = useRef();

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [errMsg, setErrMsg] = useState('');
    const [successMessage, setSuccessMessage] = useState('');

    // Check if already logged in
    useEffect(() => {
        if (auth?.isAuthenticated) {
            console.log("Already authenticated, redirecting to:", from);
            navigate(from, { replace: true });
        }
    }, [auth, navigate, from]);

    useEffect(() => {
        userRef.current?.focus();
    }, []);

    useEffect(() => {
        setErrMsg('');
    }, [email, password]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log("Form submitted");

        // Reset messages
        setErrMsg('');
        setSuccessMessage('');

        // Validate inputs
        if (!email || !password) {
            setErrMsg('Please enter both email and password.');
            return;
        }

        setIsLoading(true);

        try {
            console.log("Calling login function with email:", email);
            const result = await login(email, password);
            console.log("Login result:", result);

            if (result.success) {
                console.log("Login successful, preparing to redirect to:", from);
                setSuccessMessage('Login successful!');

                // Small delay for better UX before redirect
                setTimeout(() => {
                    console.log("Navigating to:", from);
                    navigate(from, { replace: true });
                    console.log("Navigation function called");
                }, 1000);
            } else {
                console.error("Login failed with error:", result.error);
                setErrMsg(result.error);
                errRef.current?.focus();
            }
        } catch (error) {
            console.error("Exception during login:", error);
            setErrMsg('Login failed. Please try again.');
            errRef.current?.focus();
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className='loginContainer'>
            <section className='loginForm'>
                <p ref={errRef} className={errMsg ? "errmsg" : "offscreen"} aria-live="assertive">
                    {errMsg}
                </p>
                <h1>Log In</h1>

                {successMessage && (
                    <div className="success-message">
                        {successMessage}
                    </div>
                )}

                <form onSubmit={handleSubmit}>
                    <label htmlFor="email">Email</label>
                    <input
                        type="email"
                        id="email"
                        ref={userRef}
                        autoComplete="off"
                        onChange={(e) => setEmail(e.target.value)}
                        value={email}
                        required
                    />

                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        id="password"
                        onChange={(e) => setPassword(e.target.value)}
                        value={password}
                        required
                    />

                    <button
                        type="submit"
                        className="login-button"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Signing in...' : 'Sign In'}
                    </button>
                </form>

                <p>
                    Need an Account?<br/>
                    <span className='line'>
                        <NavLink to="/Registration">Register Me</NavLink>
                    </span>
                </p>
            </section>
        </div>
    );
}