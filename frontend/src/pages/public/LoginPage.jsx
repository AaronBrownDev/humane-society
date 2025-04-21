import React from "react";
import useAuth from "../../hooks/useAuth.js";
import  "../../styles/Login.css";

import { NavLink,useNavigate, useLocation } from "react-router-dom";



export default function LoginPage() {

    const {setAuth} = useAuth;

    const navigate = useNavigate();
    const location = useLocation();
    const from = location.state?.from?.pathname || "/";

    const userRef = React.useRef();
    const errRef = React.useRef();

    const [email, setEmail]= React.useState('');
    const [password,setPassword] =React.useState('')

    const [isLoading, setIsLoading]= React.useState('');
    const [errMsg, setErrMsg] =React.useState('');
    const [successMessage, setSuccessMessage]= React.useState('');

    React.useEffect(() => {
        userRef.current.focus();
        },[])

    React.useEffect(() =>{
        setErrMsg('');
        },[email,password])

    const handleSubmit = async () => {

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
            // Change to user endpoint
            const response = await fetch('https://api.example.com/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
            });

            const accessToken = response?.data?.accessToken;
            const roles = response?.data.roles;

            setAuth({email, password,roles,accessToken});
            setEmail('');
            setPassword('');
            navigate(from, {replace: true});

            {/**setTimeout(() => {
                window.location.href = '/dashboard';
            }, 1500);*/}

         } catch (error) {
            if (!error?.response) {
                setErrMsg('No Server Response');
            } else if (error.response?.status === 400) {
                setErrMsg('Missing Username or Password');
            } else if (error.response?.status === 401) {
                setErrMsg('Unauthorized');
            } else {
                setErrMsg('Login Failed');
            }
            errRef.current.focus();
        }
    }
        return (
            <div className='loginContainer'>
                <section className='loginForm'>
                    <p ref={errRef} className={errMsg ? "errmsg" :
                        "offscreen"} aria-live="assertive">{errMsg}</p>
                    <h1> Log In</h1>
                    {errMsg && (
                        <div className="error-message">
                            {errMsg}
                        </div>
                    )}

                    {successMessage && (
                        <div className="success-message">
                            {successMessage}
                        </div>
                    )}
                    <form action={handleSubmit}>
                        <label htmlFor="email"> Email</label>
                        <input
                            type="email"
                            id="username"
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
                            autoComplete="off"
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
                        <NavLink to ="/Registration"> Register Me </NavLink>
                    </span>

                    </p>

                </section>
            </div>

        )
    }