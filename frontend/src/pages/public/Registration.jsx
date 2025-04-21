import React, {useState} from 'react'
import "../../styles/Registration.css"
import {NavLink} from "react-router-dom";

export default function Registration() {
    const [formData, setFormData] = useState({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: ''
    });

    // UI state
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState({});
    const [apiError, setApiError] = useState('');
    const [successMessage, setSuccessMessage] = useState('');

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({
            ...formData,
            [name]: value
        });

        // Clear specific field error when user types
        if (error[name]) {
            setError({
                ...error,
                [name]: ''
            });
        }
    };

    // Validate form data
    const validateForm = () => {
        const newErrors = {};

        // Validate first name
        if (!formData.firstName.trim()) {
            newErrors.firstName = 'First name is required';
        }

        // Validate last name
        if (!formData.lastName.trim()) {
            newErrors.lastName = 'Last name is required';
        }

        // Validate email
        if (!formData.email.trim()) {
            newErrors.email = 'Email is required';
        } else if (!/\S+@\S+\.\S+/.test(formData.email)) {
            newErrors.email = 'Email is invalid';
        }

        // Validate password
        if (!formData.password) {
            newErrors.password = 'Password is required';
        } else if (formData.password.length < 8) {
            newErrors.password = 'Password must be at least 8 characters';
        }

        // Validate password confirmation
        if (!formData.confirmPassword) {
            newErrors.confirmPassword = 'Please confirm your password';
        } else if (formData.password !== formData.confirmPassword) {
            newErrors.confirmPassword = 'Passwords do not match';
        }

        return newErrors;
    };

    // Handle form submission
    const handleSubmit = async (e) => {
        e.preventDefault();

        // Reset messages
        setApiError('');
        setSuccessMessage('');

        // Validate form
        const formErrors = validateForm();
        if (Object.keys(formErrors).length > 0) {
            setError(formErrors);
            return;
        }

        // Start loading
        setIsLoading(true);

        try {
            // Make API request to registration endpoint
            const response = await fetch('https://api.example.com/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    firstName: formData.firstName,
                    lastName: formData.lastName,
                    email: formData.email,
                    password: formData.password
                }),
            });

            const data = await response.json();

            if (!response.ok) {
                // Handle error response from server
                throw new Error(data.message || 'Registration failed');
            }

            // Show success message
            setSuccessMessage('Registration successful! You can now log in.');

            // Clear form
            setFormData({
                firstName: '',
                lastName: '',
                email: '',
                password: '',
                confirmPassword: ''
            });

            // Optional: Redirect to login page after delay
            setTimeout(() => {
                window.location.href = '/login';
            }, 3000);

        } catch (error) {
            setApiError(error.message || 'An error occurred during registration');
        } finally {
            setIsLoading(false);
        }
    };



    return (
        <div className="registrationContainer">
            <div className="formconcontainer">
                {apiError && (
                    <div className="error-message">
                        {apiError}
                    </div>
                )}

                {successMessage && (
                    <div className="success-message">
                        {successMessage}
                    </div>
                )}
                <form action={handleSubmit}>
                    <label htmlFor="first_name">First Name</label>
                    <input
                        type="text"
                        id="first_name"
                        name="first_name"
                        placeholder="First Name"
                        onChange={handleChange}
                    />
                    <label htmlFor="last_name">Last Name</label>
                    <input
                        type="text"
                        id="last_name"
                        name="last_name"
                        placeholder="Last Name"
                        onChange={handleChange}
                    />
                    <label htmlFor="email">Email</label>
                    <input
                        type = "email"
                        id='email'
                        name = "email"
                        placeholder="Email"
                        onChange={handleChange}
                    />
                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="Password"
                        onChange={handleChange}
                        />
                    <label htmlFor="confirmPassword">Confirm Password</label>
                    <input
                        type="password"
                        id="confirmPassword"
                        name="confirmPassword"
                        placeholder="Confirm Password"
                        onChange={handleChange}
                    />

                    <button
                        type="submit"
                        className="register-button"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Creating Account...' : 'Create Account'}
                    </button>


                </form>
            </div>

        </div>

    )
}