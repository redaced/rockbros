import axios from 'axios';

// Create an axios instance
const instance = axios.create({
    baseURL: 'http://localhost:3000' // replace with your server's address and port
});

// Add a request interceptor
instance.interceptors.request.use(function (config) {
    // Do something before the request is sent
    // For example, you can add an authorization header
    // config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});

// Add a response interceptor
instance.interceptors.response.use(function (response) {
    // Do something with response data
    return response;
}, function (error) {
    // Do something with response error
    return Promise.reject(error);
});

export default instance;