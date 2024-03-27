// 导入axios
import axios from 'axios'

// 通过axios.create方法创建一个axios实例，用request接收
const http = axios.create({
    // 指定请求的根路径
    baseURL: 'http://159.138.20.163:10086/api/v1/'
    // baseURL: 'http://localhost:10086/api/v1/'
    // baseURL: 'http://127.0.0.1:4523/m2/3695455-0-default/130512736'
})

export default http
