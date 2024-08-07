import React from 'react';

export const Todo = ({task, toggleComplete, deleteTodo, updateTodo}) => {
    return (
        <div className="Todo">
            <p onClick={() => toggleComplete(task.id)} className={`${task.completed ? 'completed' : ""}`}>{task.title}</p>
            <div>
                <img className="fa-pen" src="https://static-00.iconduck.com/assets.00/edit-icon-1022x1024-kes437mc.png" alt="Pen Icon"
                    onClick={() => updateTodo(task.id)}/>
                <img className="fa-trash" src="https://img.icons8.com/ios_filled/200/FFFFFF/trash.png" alt="Trash Icon"
                     onClick={() => deleteTodo(task.id)} />
            </div>
        </div>
    )
}