import http from '../utils/request.js'


const rank = {
    // 返回所有achievements
    getRankInfoAnalysis(start, end, id) {
        return http.get(`/workerAttendance/statworkerattendance?start_date=${start}&end_date=${end}&workshop_id=${id}`)
    },

}

export default rank;
