import React, { useEffect, useState } from 'react';
import { TodoForm } from "./TodoForm.jsx";
import { v4 as uuidv4 } from 'uuid';
import { Todo } from "./Todo";
import { EditTodoForm } from "./EditTodoForm";
import { domain } from "../../wailsjs/go/models.ts";
import {Modal} from './Modal';

export const TodoWrapper = () => {
    const [todos, setTodos] = useState([]);
    const [modalVisible, setModalVisible] = useState(false);
    const [modalMessage, setModalMessage] = useState("");

    const fetchTodos = async () => {
        try {
            const response = await fetch('http://localhost:8080');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const fetchedTodos = await response.json();
            console.log('Fetched todos:', fetchedTodos);

            if (!fetchedTodos) {
                throw new Error('Fetched todos is undefined');
            }

            if (!Array.isArray(fetchedTodos)) {
                throw new Error('Fetched todos is not an array');
            }

            const todoObjects = fetchedTodos.map(todo => domain.Todo.createFrom(todo));
            console.log('Fetched todos (processed):', todoObjects);
            setTodos(todoObjects);
        } catch (error) {
            console.error('Error fetching todos:', error);
        }
    };

    useEffect(() => {
        fetchTodos();
    }, []);

    const addTodo = async ({ title, date, time, priority }) => {
        if (!title.trim() || !date.trim() || !time.trim() || !['High', 'Medium', 'Low'].includes(priority)) {
            setModalMessage("Validation failed: fields cannot be empty !");
            setModalVisible(true);
        }

        const newTodo = {
            id: uuidv4(),
            title: title.trim(),
            date,
            time,
            active_at: new Date().toISOString(),
            status: false,
            priority
        };

        try {
            const response = await fetch('http://localhost:8080/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(newTodo),
            });
            const createdTodo = await response.json();
            setTodos([...todos, domain.Todo.createFrom(createdTodo)]);
        } catch (error) {
            console.error('Error creating todo:', error);
        }
    };


    const toggleComplete = async id => {
        console.log(`Toggling complete status for todo with ID: ${id}`);
        try {
            await fetch(`http://localhost:8080/check/${id}`, {
                method: 'PUT',
            });
            setTodos(todos.map(todo => todo.id === id ? { ...todo, status: !todo.status } : todo));
            console.log(`Todo with ID ${id} status toggled successfully.`);
        } catch (error) {
            console.error('Error toggling complete status:', error);
        }
    };

    const deleteTodo = async (id) => {
        try {
            const response = await fetch(`http://localhost:8080/${id}`, {
                method: 'DELETE',
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            setTodos(todos.filter((todo) => todo.id !== id));
            setModalMessage("Well done! You're one step closer to becoming better 🎉");
            setModalVisible(true);
        } catch (error) {
            console.error('Error deleting todo:', error);
        }
    };

    const updateTodo = async (id, title, date, time, priority) => {
        if (!title.trim() || !date.trim() || !time.trim() || !['High', 'Medium', 'Low'].includes(priority)) {
            setModalMessage("Validation failed: fields cannot be empty !");
            setModalVisible(true);
        }

        try {
            const response = await fetch(`http://localhost:8080/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ title, date, time, priority }),
            });
            const updatedTodo = await response.json();
            setTodos(todos.map(todo => todo.id === id ? updatedTodo : todo));
            console.log(`Todo with ID ${id} updated successfully.`);
        } catch (error) {
            console.error('Error updating todo:', error);
        }
    };

    const sortTodosByPriority = (todos) => {
        const priorityOrder = { 'High': 1, 'Medium': 2, 'Low': 3 };
        return todos.sort((a, b) => priorityOrder[a.priority] - priorityOrder[b.priority]);
    };

    const activeTodos = sortTodosByPriority(todos.filter(todo => !todo.status));
    const completedTodos = sortTodosByPriority(todos.filter(todo => todo.status));

    return (
        <div className='TodoWrapper'>
            <h1>TODO LIST</h1>
            <TodoForm addTodo={addTodo} />
            <h2 className="active-task">Active Tasks</h2>
            {activeTodos && activeTodos.map((todo) => (
                todo.isEditing ? (
                    <EditTodoForm updateTodo={updateTodo} key={todo.id} task={todo} />
                ) : (
                    <Todo
                        task={todo}
                        key={todo.id}
                        toggleComplete={() => toggleComplete(todo.id)}
                        deleteTodo={() => deleteTodo(todo.id)}
                        updateTodo={() => setTodos(todos.map(t => t.id === todo.id ? { ...t, isEditing: true } : t))}
                    />
                )
            ))}
            <h2 className="complete-task">Completed Tasks</h2>
            {completedTodos && completedTodos.map((todo) => (
                todo.isEditing ? (
                    <EditTodoForm updateTodo={updateTodo} key={todo.id} task={todo} />
                ) : (
                    <Todo
                        task={todo}
                        key={todo.id}
                        toggleComplete={() => toggleComplete(todo.id)}
                        deleteTodo={() => deleteTodo(todo.id)}
                        updateTodo={() => setTodos(todos.map(t => t.id === todo.id ? { ...t, isEditing: true } : t))}
                        isCompleted={true}
                    />
                )
            ))}
            {modalVisible &&
                <Modal
                    message={modalMessage}
                    onClose={() => setModalVisible(false)}
                />
            }
        </div>
    );
};
