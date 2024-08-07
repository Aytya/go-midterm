import React, {useEffect, useState} from 'react';
import {TodoForm} from "./TodoForm.jsx";
import {v4 as uuidv4} from 'uuid';
import {Todo} from "./Todo";
import {EditTodoForm} from "./EditTodoForm";
import { GetAllTodos, DeleteTodo, CreateTodo} from "../../wailsjs/go/main/App";
uuidv4();

export const TodoWrapper = () => {
    const [todos, setTodos] = useState([]);

    const toggleComplete = id => {
        setTodos(todos.map(todo => todo.id === id ? {...todo, completed :!todo.completed} : todo));
    }

    const updateTodo = id => {
        setTodos(todos.map(todo => todo.id === id ? {...todo, isEditing: !todo.isEditing} : todo));
    }

    const editTask = (task,id) => {
        setTodos(todos.map(todo => todo.id === id ? {...todo, task, isEditing: !todo.isEditing} : todo));
    }

    return (
        <div className='TodoWrapper'>
            <h1>TODO LIST</h1>
            <TodoForm addTodo={CreateTodo}/>
            {todos.map((todo, index) => (
                todo.isEditing ? (
                    <EditTodoForm updateTodo={editTask} task={todo}/>
                ) : (
                    <Todo task={todo} key={index}
                          toggleComplete={toggleComplete}
                          deleteTodo={deleteTodo}
                          updateTodo={updateTodo}
                    />
                )
            ))}
        </div>
    )
}
