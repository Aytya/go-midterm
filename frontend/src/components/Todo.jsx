import React from 'react';

export const Todo = ({ task, toggleComplete, deleteTodo, updateTodo, isCompleted }) => {

    const formatDate = (dateString) => {
        const options = { year: 'numeric', month: 'long', day: 'numeric' };
        return new Date(dateString).toLocaleDateString(undefined, options);
    };

    const formatTime = (timeString) => {
        if (timeString.startsWith("0000-01-01T")) {
            timeString = timeString.replace("0000-01-01T", "");
        }
        if (timeString.endsWith(":00Z")) {
            timeString = timeString.replace(":00Z", "");
        }

        const [hours, minutes] = timeString.split(':');
        const date = new Date();
        date.setHours(hours, minutes);
        const options = { hour: '2-digit', minute: '2-digit', hour12: true };
        return date.toLocaleTimeString(undefined, options);
    };

    return (
        <div className={`Todo ${isCompleted ? 'completed-task' : ''}`}>
            <p
                onClick={() => toggleComplete(task.id)}
                className={task.status ? 'completed' : 'incompleted'}>
                {task.title}
            </p>
            <div>
                <span>{formatDate(task.date)}</span>
                <span>{formatTime(task.time)}</span>
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
