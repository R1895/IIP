import http from '../utils/request.js'


const oee = {

    getOeeAnalysis() {
        return http.get(`/api/v1/screen/oee`)
    }

}

export default oee;
