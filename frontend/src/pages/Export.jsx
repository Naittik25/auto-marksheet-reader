import React from 'react';

const Export = () => {
    // These functions will eventually call your Backend export endpoints
    const handleDownloadExcel = () => {
        alert("Downloading Excel... (Backend integration needed)");
        // window.open('http://localhost:8080/api/export/excel', '_blank');
    };

    const handleDownloadPDF = () => {
        alert("Downloading PDF... (Backend integration needed)");
        // window.open('http://localhost:8080/api/export/pdf', '_blank');
    };

    return (
        <div className="page-container">
            <h2>📤 Export Data</h2>
            <p className="page-subtitle">Download your extracted data in your preferred format.</p>

            <div className="export-options">
                <div className="export-card excel">
                    <h3>📊 Excel Report</h3>
                    <p>Best for analysis and sorting data.</p>
                    <button onClick={handleDownloadExcel} className="btn-excel">Download .XLSX</button>
                </div>

                <div className="export-card pdf">
                    <h3>📄 PDF Document</h3>
                    <p>Best for printing and archiving.</p>
                    <button onClick={handleDownloadPDF} className="btn-pdf">Download .PDF</button>
                </div>
            </div>
        </div>
    );
};

export default Export;