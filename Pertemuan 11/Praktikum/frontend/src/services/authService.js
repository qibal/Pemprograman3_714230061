import axios from "axios";

const API_URL = "http://127.0.0.1:8088";

export const login = async (username, password) => {
  const response = await axios.post(`${API_URL}/login`, { username, password });
  return response.data;
};

export const register = async (username, password, role) => {
  const response = await axios.post(`${API_URL}/register`, { username, password, role });
  return response.data;
};