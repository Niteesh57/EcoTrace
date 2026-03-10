import axios from 'axios';

const API_URL = 'http://localhost:3000';

export const submitTransaction = async (qrData) => {
  try {
    const response = await axios.post(`${API_URL}/submit`, { qrData });
    return response.data;
  } catch (error) {
    throw error;
  }
};