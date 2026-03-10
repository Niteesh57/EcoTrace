import axios from 'axios';

const API_URL = 'http://your-blockchain-api-url';

export const submitTransaction = async (qrData) => {
  const response = await axios.post(`${API_URL}/submitTransaction`, {
    shipmentData: qrData
  });
  return response.data;
};