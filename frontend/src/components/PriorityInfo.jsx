import React from 'react';

export const PriorityInfo = () => {
    return (
        <div className="priority-info">
            <h3>Priority Levels</h3>
            <div className="priority-item high">
                <span className="priority-label">High</span>
                <span className="priority-color high"></span>
            </div>
            <div className="priority-item medium">
                <span className="priority-label">Medium</span>
                <span className="priority-color medium"></span>
            </div>
            <div className="priority-item low">
                <span className="priority-label">Low</span>
                <span className="priority-color low"></span>
            </div>
        </div>
    );
};
