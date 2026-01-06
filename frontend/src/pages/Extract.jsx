import React from 'react';
import { useLocation } from 'react-router-dom';

const Extract = () => {
    // We can access data passed from the Scan page here
    const location = useLocation();
    const rawText = location.state?.text || "No text data found. Please scan a file first.";

    return (
        <div className="page-container">
            <h2>📝 Raw Text Extraction</h2>
            <p>Review the raw text read by the OCR engine.</p>
            
            <div className="text-editor-box">
                <textarea 
                    className="raw-textarea" 
                    defaultValue={rawText} 
                    rows="15"
                />
            </div>
            <button className="btn-primary">Proceed to Table View</button>
        </div>
    );
};

export default Extract;