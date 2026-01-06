import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
        <div className="home-container">
            <div className="hero-section">
                <h1>Automate Your Marksheet Entry</h1>
                <p>Upload images or PDFs of marksheets and instantly extract student data into Excel.</p>
                <Link to="/scan">
                    <button className="cta-button">Start Scanning Now</button>
                </Link>
            </div>
            
            <div className="features-grid">
                <div className="feature-card">
                    <h3>🚀 Fast OCR</h3>
                    <p>Extracts text in seconds using Tesseract AI.</p>
                </div>
                <div className="feature-card">
                    <h3>📂 Bulk Upload</h3>
                    <p>Process multiple marksheets at once.</p>
                </div>
                <div className="feature-card">
                    <h3>📊 Excel Export</h3>
                    <p>Download clean data directly to Excel.</p>
                </div>
            </div>
        </div>
    );
};

export default Home;