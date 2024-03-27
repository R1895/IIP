import http from '../utils/request.js'


const order = {
    getOrderInfoAnalysis(limit, start, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
            limit: limit,
            start: start
        }
        return http.post(`/screen/order`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },

}

export default order;
