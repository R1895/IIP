import http from '../utils/request.js'


const light = {
    // 返回所有achievements
    getTest() {
        return http.get(`user/users?limit=10&start=1`)
    },
    getLightEfficiency() {
        return http.get(`/api/v1/screen/light`)
    }

}

export default light;
