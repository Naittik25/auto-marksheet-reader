import React from 'react';

const ResultTable = ({ data }) => {
    // Dummy data structure for UI testing
    const defaultData = [
        { subject: "Mathematics", marks: 85, grade: "A" },
        { subject: "Science", marks: 78, grade: "B+" },
        { subject: "English", marks: 92, grade: "O" },
    ];

    const displayData = data || defaultData;

    return (
        <div className="table-wrapper">
            <table className="styled-table">
                <thead>
                    <tr>
                        <th>Subject</th>
                        <th>Marks Obtained</th>
                        <th>Grade</th>
                    </tr>
                </thead>
                <tbody>
                    {displayData.map((row, index) => (
                        <tr key={index}>
                            <td>{row.subject}</td>
                            <td>{row.marks}</td>
                            <td>{row.grade}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            <button className="export-btn">Download Excel</button>
        </div>
    );
};

export default ResultTable;