import axios from "axios";


const mock = {

    getProductionStatus() {
        return axios.get(`http://127.0.0.1:4523/m2/3695455-0-default/130512736`)
    },
    getOeeAnalysis() {
        return axios.get(`http://127.0.0.1:4523/m1/3695455-0-default/api/v1/screen/oee`)
    },
    getProductionLive() {
        return axios.get(`http://127.0.0.1:4523/m1/3695455-0-default/api/v1/screen/live`)
    },
    getLightEfficiency() {
        return axios.get(`http://127.0.0.1:4523/m1/3695455-0-default/api/v1/screen/light`)
    },
    getWorkingHoursRank() {

    }
}

export default mock;
