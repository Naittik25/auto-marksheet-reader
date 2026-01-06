import React from 'react';
import FileUpload from '../components/FileUpload';

const Scan = () => {
    return (
        <div className="page-container">
            <h2>Upload Marksheet</h2>
            <p className="page-subtitle">Select an image or PDF to extract data.</p>
            <FileUpload />
        </div>
    );
};

export default Scan;