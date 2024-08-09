import React, {useEffect, useState} from 'react';

export const EditTodoForm = ({ updateTodo, task }) => {
    const [title, setTitle] = useState(task?.title || '');
    const [date, setDate] = useState(task?.date ? task.date.split('T')[0] : '');
    const [time, setTime] = useState(() => {
        const timePart = task?.time ? task.time.substring(0, 5) : '';
        return timePart;
    });
    const [priority, setPriority] = useState(task?.priority);

    useEffect(() => {
        setTitle(task?.title || '');
        setDate(task?.date ? task.date.split('T')[0] : '');
        setTime(task?.time && task.time.includes('T') ? task.time.split('T')[1].substring(0, 5) : '');
        setPriority(task?.priority);
    }, [task]);

    const handleSubmit = (e) => {
        e.preventDefault();
        updateTodo(task.id, title, date, time, priority);


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
            <select
                className="todo-input-priority"
                value={priority}
                onChange={e => setPriority(e.target.value)}>
                <option value="">Select Priority</option>
                <option value="High">High</option>
                <option value="Medium">Medium</option>
                <option value="Low">Low</option>
            </select>
            <button type='submit' className='todo-btn'>Update Task</button>
        </form>
    );
};
