import axios from 'axios'
// import router from '@/router/router'

// Add a request interceptor
axios.interceptors.request.use(function (config) {
    // Do something before request is sent
    // For example, add your authorization token to the header
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});

// Add a response interceptor
axios.interceptors.response.use(function (response) {
    // Do something with response data
    return response;
}, function (error) {
    // Do something with response error
    // For example, handle global error here
    // if (error.response.status === 401) {
    //     console.log('Unauthorized, logging out ...');
    //     localStorage.removeItem('token');
    //     router.push('/login');
    // }
    return Promise.reject(error);
});