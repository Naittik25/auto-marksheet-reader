import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/globals.css'; 

const Navbar = () => {
    return (
        <nav className="navbar">
            <div className="nav-brand">
                <h1>📊 Marksheet Reader</h1>
            </div>
            <div className="nav-links">
                <Link to="/" className="nav-link">Home</Link>
                <Link to="/scan" className="nav-link">Scan & Upload</Link>
                <Link to="/results" className="nav-link">View Results</Link>
            </div>
        </nav>
    );
};

export default Navbar;