import React, { useState } from 'react';

export const EditTodoForm = ({ updateTodo, task }) => {
    const [title, setTitle] = useState(task.title);
    const [date, setDate] = useState(task.date ? task.date.split('T')[0] : '');
    const [time, setTime] = useState(task.time ? task.time.split('T')[1].substring(0, 5) : '');

    const handleSubmit = (e) => {
        e.preventDefault();
        updateTodo(task.id, title, date, time);
        setTitle('');
        setDate('');
        setTime('');
    };

    return (
        <form className='todo-form' onSubmit={handleSubmit}>
            <input
                type='text'
                className='todo-input'
                value={title}
                placeholder='Update Task'
                onChange={(e) => setTitle(e.target.value)}
            />
            <input
                type='date'
                className='todo-input-data'
                value={date}
                onChange={(e) => setDate(e.target.value)}
            />
            <input
                type='time'
                className='todo-input-time'
                value={time}
                onChange={(e) => setTime(e.target.value)}
            />
            <button type='submit' className='todo-btn'>Update Task</button>
        </form>
    );
};
