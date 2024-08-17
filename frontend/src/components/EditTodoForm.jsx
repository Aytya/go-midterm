import React, { useEffect, useState } from 'react';

export const EditTodoForm = ({ updateTodo, task }) => {
    const [title, setTitle] = useState('');
    const [dateTime, setDateTime] = useState('');
    const [priority, setPriority] = useState('');

    useEffect(() => {
        if (task) {
            setTitle(task.title || '');
            const localDateTime = new Date(task.datetime).toLocaleString('sv-SE').replace(' ', 'T');
            setDateTime(localDateTime || '');
            setPriority(task.priority || '');
        }
    }, [task]);

    const handleDateChange = (e) => {
        const newDate = e.target.value;
        const time = dateTime.split('T')[1] || '00:00:00';
        setDateTime(`${newDate}T${time}`);
    };

    const handleTimeChange = (e) => {
        const newTime = `${e.target.value}:00`;
        const date = dateTime.split('T')[0] || '';
        setDateTime(`${date}T${newTime}`);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const isoDateTime = new Date(dateTime).toISOString();
        updateTodo(task.id, title, isoDateTime, priority);
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
                value={dateTime.split('T')[0]}
                onChange={handleDateChange}
            />
            <input
                type='time'
                className='todo-input-time'
                value={dateTime.split('T')[1]?.substring(0, 5) || ''}
                onChange={handleTimeChange}
            />
            <select
                className="todo-input-priority"
                value={priority}
                onChange={(e) => setPriority(e.target.value)}>
                <option value="">Select Priority</option>
                <option value="High">High</option>
                <option value="Medium">Medium</option>
                <option value="Low">Low</option>
            </select>
            <button type='submit' className='todo-btn'>
                Update Task
            </button>
        </form>
    );
};
