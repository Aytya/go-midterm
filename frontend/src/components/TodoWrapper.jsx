import React, { useEffect, useState } from 'react';
import { TodoForm } from "./TodoForm.jsx";
import { Todo } from "./Todo";
import { EditTodoForm } from "./EditTodoForm";
import { domain } from "../../wailsjs/go/models.ts";
import { Modal } from './Modal';
import { CreateTodo, CheckTodo, UpdateTodo, GetAllTodos, DeleteTodo } from "../../wailsjs/go/handler/App";
import {PriorityInfo} from "./PriorityInfo";

export const TodoWrapper = () => {
    const [todos, setTodos] = useState([]);
    const [modalVisible, setModalVisible] = useState(false);
    const [modalMessage, setModalMessage] = useState("");

    const fetchTodos = async () => {
        console.log("GetAllTodos handler:", GetAllTodos);
        try {
            const fetchedTodos = await GetAllTodos();
            setTodos(fetchedTodos.map(todo => {
                return domain.Todo.createFrom(todo);
            }));
        } catch (error) {
            console.error('Error fetching todos:', error);
            setModalMessage(`Error fetching todos: ${error.message}`);
            setModalVisible(true);
        }
    };

    useEffect(() => {
        fetchTodos();
    }, []);

    const addTodo = async ({ title, priority, datetime }) => {
        if (!title.trim() || !datetime || !['High', 'Medium', 'Low'].includes(priority)) {
            setModalMessage("Validation failed: fields cannot be empty!");
            setModalVisible(true);
            return;
        }

        try {
            const createdTodo = await CreateTodo(title.trim(), priority, datetime);
            setTodos(prevTodos => [...prevTodos, createdTodo]);
        } catch (error) {
            console.error('Error creating todo:', error);
            setModalMessage(`Error creating todo: ${error.message}`);
            setModalVisible(true);
        }
    };

    const toggleComplete = async id => {
        try {
            await CheckTodo(id);
            setTodos(todos.map(todo => todo.id === id ? { ...todo, status: !todo.status } : todo));
        } catch (error) {
            console.error('Error toggling complete status:', error);
            setModalMessage(`Error toggling complete status: ${error.message}`);
            setModalVisible(true);
        }
    };

    const deleteTodo = async id => {
        try {
            await DeleteTodo(id);
            setTodos(todos.filter(todo => todo.id !== id));
            setModalMessage("Well done! You're one step closer to becoming better 🎉");
            setModalVisible(true);
        } catch (error) {
            console.error('Error deleting todo:', error);
            setModalMessage(`Error deleting todo: ${error.message}`);
            setModalVisible(true);
        }
    };

    const updateTodo = async (id, title, datetime, priority) => {
        if (!title.trim() || !datetime || !['High', 'Medium', 'Low'].includes(priority)) {
            setModalMessage("Validation failed: fields cannot be empty!");
            setModalVisible(true);
            return;
        }

        try {
            console.log("Updating Todo with data:", { id, title, priority, datetime });
            await UpdateTodo(id, title, priority, datetime);
            const updatedTodo = domain.Todo.createFrom({
                id,
                title: title.trim(),
                priority,
                datetime,
                status: false,
            });
            setTodos(todos.map(todo => todo.id === id ? updatedTodo : todo));
        } catch (error) {
            console.error('Error updating todo:', error);
            setModalMessage(`Error updating todo: ${error.message}`);
            setModalVisible(true);
        }
    };

    const sortTodosByPriority = todos => {
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
            {activeTodos && activeTodos.map(todo => (
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
            {completedTodos && completedTodos.map(todo => (
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
            <PriorityInfo />
        </div>
    );
};
