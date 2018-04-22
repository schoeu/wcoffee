import axios from 'axios'

//axios.defaults.withCredentials = true;
// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    if (response.data.status === 0) {
        return response.data.data;
    }
    return Promise.reject(response);
    
}, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
});

export const getTags = function (param) {
    return axios.get('http://localhost:8910/api/tags')
}

export const getList = function (param) {
    return axios.get('http://localhost:8910/api/list')
}
