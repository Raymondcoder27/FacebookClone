import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_URL
})

api.interceptors.request.use((config) => {
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
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      console.error("Unauthorized request:", error.response.data); // Log unauthorized access
      // You might want to trigger a logout here if necessary
    }
    return Promise.reject(error);
  }
);


// Add a request interceptor to include the token in headers
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `${token}`;
    console.log("Adding token to headers:", config.headers.Authorization); // Log the header
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

//   export default api;

export default api