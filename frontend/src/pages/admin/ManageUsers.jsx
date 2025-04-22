import React, { useState, useEffect } from 'react';
import userService from '../../services/userService';
import '../../styles/manageUsers.css';

export default function ManageUsers() {
    // State for the form
    const [formData, setFormData] = useState({
        FirstName: '',
        LastName: '',
        Email: '',
        Phone: '',
        Password: '',
        Confirmation: '',
    });

    // State for the user list
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    // Fetch users on component mount
    useEffect(() => {
        const fetchUsers = async () => {
            try {
                setLoading(true);
                const data = await userService.getAllUsers();
                setUsers(data);
                setError(null);
            } catch (err) {
                setError('Failed to load users: ' + (err.message || 'Unknown error'));
                console.error('Error loading users:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchUsers();
    }, []);

    // Handle form input changes
    function handleChange(event) {
        const { name, value } = event.target;
        setFormData(prevState => ({ ...prevState, [name]: value }));
    }

    // API call for form submission
    function handleSubmit(formData) {
        const data = Object.fromEntries(formData);
        console.log(data);
        // Add API call here to create new user
    }

    // Format date for display
    const formatDate = (dateString) => {
        if (!dateString) return 'N/A';
        return new Date(dateString).toLocaleString();
    };

    return (
        <div className="users-container">
            <div className="top">
                <h1>Manage Users</h1>
            </div>

            {/* Display users table */}
            <div className="users-table-container">
                <h2>User Accounts</h2>

                {loading && <p>Loading users...</p>}

                {error && <p className="error-message">{error}</p>}

                {!loading && !error && (
                    <table className="users-table">
                        <thead>
                        <tr>
                            <th>User ID</th>
                            <th>Last Login</th>
                            <th>Status</th>
                            <th>Created At</th>
                            <th>Failed Logins</th>
                            <th>Is Locked</th>
                            <th>Lockout Ends</th>
                        </tr>
                        </thead>
                        <tbody>
                        {users.length === 0 ? (
                            <tr>
                                <td colSpan="7">No users found</td>
                            </tr>
                        ) : (
                            users.map(user => (
                                <tr key={user.userId}>
                                    <td>{user.userId}</td>
                                    <td>{formatDate(user.lastLogin)}</td>
                                    <td>{user.isActive ? 'Active' : 'Inactive'}</td>
                                    <td>{formatDate(user.createdAt)}</td>
                                    <td>{user.failedLoginAttempts}</td>
                                    <td>{user.isLocked ? 'Yes' : 'No'}</td>
                                    <td>{formatDate(user.lockoutEnd)}</td>
                                </tr>
                            ))
                        )}
                        </tbody>
                    </table>
                )}
            </div>

            {/* Existing form for adding users */}
            <div className="bottom">
                <div className="left">
                    <img src="../../assets/placeholder-1024x1024.png" alt="Placeholder" />
                </div>
                <div className="right">
                    <h2>Add New User</h2>
                    <form action={handleSubmit}>
                        <label>First Name</label>
                        <input
                            type="text"
                            name="FirstName"
                            value={formData.FirstName}
                            onChange={handleChange}
                            placeholder="First Name"
                        />
                        <label>Last Name</label>
                        <input
                            type="text"
                            name="LastName"
                            value={formData.LastName}
                            onChange={handleChange}
                            placeholder="Last Name"
                        />
                        <label>Email</label>
                        <input
                            type="email"
                            name="Email"
                            value={formData.Email}
                            onChange={handleChange}
                            placeholder="email@gmail.com"
                        />
                        <label>Phone Number</label>
                        <input
                            type="tel"
                            name="Phone"
                            value={formData.Phone}
                            onChange={handleChange}
                            placeholder="123-123-1234"
                        />
                        <label>Password</label>
                        <input
                            type="password"
                            name="Password"
                            value={formData.Password}
                            onChange={handleChange}
                        />
                        <label>Confirm Password</label>
                        <input
                            type="password"
                            name="Confirmation"
                            value={formData.Confirmation}
                            onChange={handleChange}
                        />
                        <button type="submit">Add User</button>
                    </form>
                </div>
            </div>
        </div>
    );
}