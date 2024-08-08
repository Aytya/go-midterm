import React, {useState} from 'react';

export const TodoForm = ({addTodo}) => {
    const [title, setTitle] = useState('');
    const [date, setDate] = useState('');
    const [time, setTime] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        addTodo({
            title: title.trim(),
            date,
            time,
        });
        setTitle('');
        setDate('');
        setTime('');
    };

    return (
        <form className='TodoForm' onSubmit={handleSubmit}>
            <input
                type='text'
                className='todo-input'
                value={title}
                placeholder='Add a new task'
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
            <button type='submit' className='todo-btn'>
                Add Task
            </button>
        </form>
    );
}