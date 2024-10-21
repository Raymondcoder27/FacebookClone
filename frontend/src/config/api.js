import axios from 'axios'

const api = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_URL
})

api.interceptors.request.use((config) =>{
    return config
})

// api.interceptors.response.use(
//     (response) => response,
//     (error) => {
//         if (error.response != undefined && error.response.status === 401){

//         }
//         return Promise.reject(error)
//     }
// )

// Add a request interceptor to include the token in headers
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  }, (error) => {
    return Promise.reject(error);
  });
  
//   export default api;

export default api