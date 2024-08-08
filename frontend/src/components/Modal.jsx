import React from 'react';

export const Modal = ({ message, onClose }) => {
    return (
        <div className="modal">
            <div className="modal-content">
                <p>{message}</p>
                <button className="close-btn" onClick={onClose}>Close</button>
            </div>
        </div>
    );
};

