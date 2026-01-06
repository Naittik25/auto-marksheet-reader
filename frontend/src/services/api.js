import axios from 'axios';

// Point to your Go Backend
const API_URL = 'http://localhost:8080/api';

export const uploadFile = async (file) => {
    const formData = new FormData();
    formData.append('file', file);

    try {
        const response = await axios.post(`${API_URL}/upload`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });
        return response.data;
    } catch (error) {
        console.error("Upload failed", error);
        throw error;
    }
};