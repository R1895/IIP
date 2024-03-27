import http from '../utils/request.js'


const status = {

    // getProductionStatus(limit, start, id) {
    //     if (!id) return http.get(`/task/tasks?limit=${limit}&start=${start}`)
    //     return http.get(`/task/tasks?limit=${limit}&start=${start}&workshop_id=${id}`)
    // }
    getProductionStatus(limit, start, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id
        }
        return http.post(`/screen/productionBoard`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    }
}

export default status;
