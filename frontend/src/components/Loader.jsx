import React from 'react';
import '../styles/globals.css';

const Loader = () => {
    return (
        <div className="loader-overlay">
            <div className="spinner"></div>
            <p>Scanning Marksheet...</p>
        </div>
    );
};

export default Loader;