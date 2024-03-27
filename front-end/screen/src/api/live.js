import http from '../utils/request.js'


const live = {

    getLiveInformation(id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
        }
        return http.post(`/screen/productionStatus`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },
}

export default live;
