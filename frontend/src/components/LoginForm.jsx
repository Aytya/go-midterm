import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Login } from "../../wailsjs/go/handler/App";


export const LoginForm = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLoginClick = (e) => {
        e.preventDefault();
        navigate('/register');
    };

    const handleLogin = async (e) => {
        e.preventDefault();

        try {
            const response = await Login(username, password);

            if (response && response.token && response.role) {
                localStorage.setItem('token', response.token);
                localStorage.setItem('role', response.role);
                navigate('/todos');
            } else {
                console.error('Invalid response format:', response);
                alert('Login failed: Invalid response from server');
            }
        } catch (error) {
            console.error('Login failed:', error.message || error);
            alert('Login failed: ' + (error.response?.data || error.message));
        }
    };

    return (
        <form onSubmit={handleLogin} className="auth-form">
            <h2>Login</h2>
            <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <button type="submit">Login</button>
            <button className="login-button" onClick={handleLoginClick}>Don't have an account? Register</button>
        </form>
    );
};
