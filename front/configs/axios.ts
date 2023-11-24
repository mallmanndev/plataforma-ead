import axiosLib from "axios";
import { getSession } from "next-auth/react";

const defaultOptions = {
  baseURL: process.env.NEXT_PUBLIC_API_URL,
};

const axios = axiosLib.create(defaultOptions);
axios.interceptors.request.use(async (config) => {
  const session = await getSession();
  const token = session?.token;
  config.headers.Authorization = token ? `Bearer ${token}` : "";
  return config;
});

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.log(`error`, error);
    throw new Error(error.response.data.message);
  }
);

export default axios;
