import './App.css';
import {TodoWrapper} from "./components/TodoWrapper.jsx";
import {PriorityInfo} from "./components/PriorityInfo";
import {Routes, Route, useNavigate } from 'react-router-dom';
import {LoginForm} from "./components/LoginForm";
import {useEffect} from "react";
import {RegisterForm} from "./components/RegisterForm";


function App() {
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (!token) {
            navigate('/');
        }
    }, [navigate]);

    return (
        <Routes>
            <Route path="/register" element={<RegisterForm />} />
            <Route path="/" element={<LoginForm />} />
            <Route path="/todos" element={<TodoWrapper />} />
        </Routes>
    );
}
export default App