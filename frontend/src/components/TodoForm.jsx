import React, {useState} from 'react';

export const TodoForm = ({addTodo}) => {
    const [title, setTitle] = useState('');
    const [date, setDate] = useState('');
    const [time, setTime] = useState('');
    const [priority, setPriority] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();

        const dateTime = new Date(`${date}T${time}:00`).toISOString();
        const newTodo = {
            title: title.trim(),
            datetime: dateTime,
            priority: priority
        };
        console.log(dateTime);
        addTodo(newTodo);
        setTitle('');
        setDate('');
        setTime('');
        setPriority('');
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
            <select
                className="todo-input-priority"
                value={priority}
                onChange={e => setPriority(e.target.value)}>
                <option value="">Select Priority</option>
                <option className="priority-high" value="High">High</option>
                <option className="priority-medium" value="Medium">Medium</option>
                <option className="priority-low" value="Low">Low</option>
            </select>
            <button type='submit' className='todo-btn'>
                Add Task
            </button>
        </form>
    );
}
