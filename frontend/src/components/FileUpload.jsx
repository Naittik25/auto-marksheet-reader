import React, { useState } from 'react';
import { uploadFile } from '../services/api';
// Make sure to import the CSS file if you haven't imported it in main.jsx
import '../styles/globals.css'; 

const FileUpload = () => {
    const [selectedFile, setSelectedFile] = useState(null);
    const [result, setResult] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
        setError(""); // Clear old errors
        setResult(null); // Clear old results
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            setError("Please select a file first!");
            return;
        }

        setLoading(true);
        setError("");
        
        try {
            const data = await uploadFile(selectedFile);
            setResult(data);
        } catch (err) {
            console.error(err);
            setError("Failed to upload file. Check if Backend is running.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="container">
            <h2>📄 Auto Marksheet Reader</h2>
            
            {/* Upload Section */}
            <div className="upload-box">
                <input type="file" onChange={handleFileChange} accept="image/*, application/pdf" />
                <br />
                <button onClick={handleUpload} disabled={loading} className="upload-btn">
                    {loading ? "Processing..." : "Upload & Scan"}
                </button>
            </div>

            {/* Error Message */}
            {error && <div className="error-msg">{error}</div>}

            {/* Results Section */}
            {result && (
                <div className="result-box">
                    <h3>✅ Success!</h3>
                    <p><strong>File Path:</strong> {result.path}</p>
                    <div className="text-box">
                        <strong>Extracted Text:</strong>
                        <pre>{result.text}</pre>
                    </div>
                </div>
            )}
        </div>
    );
};

export default FileUpload;