import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import useAuth from "../../hooks/useAuth";
import "../../styles/Registration.css";

export default function Registration() {
    const { register } = useAuth();
    const navigate = useNavigate();

    const [formData, setFormData] = useState({
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
        physicalAddress: '',
        mailingAddress: ''
    });

    // UI state
    const [isLoading, setIsLoading] = useState(false);
    const [errors, setErrors] = useState({});
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
        if (errors[name]) {
            setErrors({
                ...errors,
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

        // Validate addresses
        if (!formData.physicalAddress.trim()) {
            newErrors.physicalAddress = 'Physical address is required';
        }

        if (!formData.mailingAddress.trim()) {
            newErrors.mailingAddress = 'Mailing address is required';
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
            setErrors(formErrors);
            return;
        }

        // Start loading
        setIsLoading(true);

        try {
            // Format data for API
            const userData = {
                firstName: formData.firstName,
                lastName: formData.lastName,
                emailAddress: formData.email,
                password: formData.password,
                physicalAddress: formData.physicalAddress,
                mailingAddress: formData.mailingAddress,
                // Set optional fields
                birthDate: null,
                phoneNumber: ''
            };

            const result = await register(userData);

            if (result.success) {
                // Show success message
                setSuccessMessage('Registration successful! You can now log in.');

                // Clear form
                setFormData({
                    firstName: '',
                    lastName: '',
                    email: '',
                    password: '',
                    confirmPassword: '',
                    physicalAddress: '',
                    mailingAddress: ''
                });

                // Redirect to login page after delay
                setTimeout(() => {
                    navigate('/LoginPage');
                }, 2000);
            } else {
                setApiError(result.error);
            }
        } catch (error) {
            setApiError(error.message || 'An error occurred during registration');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="registrationContainer">
            <div className="formContainer">
                <h2>Registration</h2>

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

                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="firstName">First Name</label>
                        <input
                            type="text"
                            id="firstName"
                            name="firstName"
                            value={formData.firstName}
                            onChange={handleChange}
                            className={errors.firstName ? 'error' : ''}
                        />
                        {errors.firstName && <div className="error-text">{errors.firstName}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="lastName">Last Name</label>
                        <input
                            type="text"
                            id="lastName"
                            name="lastName"
                            value={formData.lastName}
                            onChange={handleChange}
                            className={errors.lastName ? 'error' : ''}
                        />
                        {errors.lastName && <div className="error-text">{errors.lastName}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="email">Email</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            value={formData.email}
                            onChange={handleChange}
                            className={errors.email ? 'error' : ''}
                        />
                        {errors.email && <div className="error-text">{errors.email}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="physicalAddress">Physical Address</label>
                        <input
                            type="text"
                            id="physicalAddress"
                            name="physicalAddress"
                            value={formData.physicalAddress}
                            onChange={handleChange}
                            className={errors.physicalAddress ? 'error' : ''}
                        />
                        {errors.physicalAddress && <div className="error-text">{errors.physicalAddress}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="mailingAddress">Mailing Address</label>
                        <input
                            type="text"
                            id="mailingAddress"
                            name="mailingAddress"
                            value={formData.mailingAddress}
                            onChange={handleChange}
                            className={errors.mailingAddress ? 'error' : ''}
                        />
                        {errors.mailingAddress && <div className="error-text">{errors.mailingAddress}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            value={formData.password}
                            onChange={handleChange}
                            className={errors.password ? 'error' : ''}
                        />
                        {errors.password && <div className="error-text">{errors.password}</div>}
                    </div>

                    <div className="form-group">
                        <label htmlFor="confirmPassword">Confirm Password</label>
                        <input
                            type="password"
                            id="confirmPassword"
                            name="confirmPassword"
                            value={formData.confirmPassword}
                            onChange={handleChange}
                            className={errors.confirmPassword ? 'error' : ''}
                        />
                        {errors.confirmPassword && <div className="error-text">{errors.confirmPassword}</div>}
                    </div>

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
    );
}