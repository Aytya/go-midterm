import React from 'react';

export const Todo = ({ task, toggleComplete, deleteTodo, updateTodo, isCompleted }) => {
    const formatDate = (dateTimeString) => {
        const date = new Date(dateTimeString);
        if (isNaN(date)) {
            return "Invalid Date";
        }
        return date.toLocaleDateString(undefined, {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
        });
    };

    const formatTime = (dateTimeString) => {
        const date = new Date(dateTimeString);
        if (isNaN(date)) {
            return "Invalid Time";
        }
        return date.toLocaleTimeString(undefined, {
            hour: 'numeric',
            minute: 'numeric',
            hour12: true,
        });
    };

    return (
        <div className={`Todo ${isCompleted ? 'completed-task' : '' } priority-${task.priority.toLowerCase()}`}>
            <p
                onClick={() => toggleComplete(task.id)}
                className={task.status ? 'completed' : 'incompleted'}>
                {task.title}
            </p>
            <div>
                <span>{formatDate(task.datetime)}</span>
                <span>{formatTime(task.datetime)}</span>
                <img
                    className="fa-pen"
                    src="https://static-00.iconduck.com/assets.00/edit-icon-1022x1024-kes437mc.png"
                    alt="Pen Icon"
                    onClick={() => updateTodo(task.id)}/>
                <img
                    className="fa-trash"
                    src="https://img.icons8.com/ios_filled/200/FFFFFF/trash.png"
                    alt="Trash Icon"
                    onClick={() => deleteTodo(task.id)}/>
            </div>
        </div>
    );
};
