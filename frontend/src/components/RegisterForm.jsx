import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Register } from "../../wailsjs/go/handler/App";


export const RegisterForm = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLoginClick = (e) => {
        e.preventDefault();
        navigate('/');
    };

    const handleRegister = async (event) => {
        event.preventDefault();
        try {
            const message = await Register(username, password);
            alert(message);
            navigate('/');
        } catch (error) {
            console.error('Registration failed:', error.message);
        }
    };

    return (
        <form onSubmit={handleRegister} className="auth-form">
            <h2>Register</h2>
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
            <button type="submit">Register</button>
            <button className="login-button" onClick={handleLoginClick}>Already have an account? Log in</button>
        </form>
    );
};
