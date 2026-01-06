import React, { useRef, useState, useCallback } from 'react';
import Webcam from "react-webcam";
import { uploadFile } from '../services/api';
import '../styles/globals.css';

const CameraScanner = ({ onScanSuccess }) => {
    const webcamRef = useRef(null);
    const [image, setImage] = useState(null);
    const [loading, setLoading] = useState(false);

    // 1. Capture Image logic
    const capture = useCallback(() => {
        const imageSrc = webcamRef.current.getScreenshot();
        setImage(imageSrc);
    }, [webcamRef]);

    // 2. Convert Base64 to File and Upload
    const handleUpload = async () => {
        if (!image) return;

        setLoading(true);
        try {
            // Convert base64 string to a real file object
            const res = await fetch(image);
            const blob = await res.blob();
            const file = new File([blob], "camera_capture.jpg", { type: "image/jpeg" });

            // Send to Backend
            const data = await uploadFile(file);
            onScanSuccess(data); // Pass result back to parent
        } catch (error) {
            alert("Upload failed!");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="camera-container">
            <h3>📷 Scan via Camera</h3>
            
            {!image ? (
                <>
                    <Webcam
                        audio={false}
                        ref={webcamRef}
                        screenshotFormat="image/jpeg"
                        className="webcam-view"
                    />
                    <button onClick={capture} className="btn-capture">Capture Photo</button>
                </>
            ) : (
                <div className="preview-box">
                    <img src={image} alt="Captured" className="img-preview" />
                    <div className="btn-group">
                        <button onClick={() => setImage(null)} className="btn-retry">Retake</button>
                        <button onClick={handleUpload} disabled={loading} className="btn-upload">
                            {loading ? "Processing..." : "Upload Photo"}
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default CameraScanner;