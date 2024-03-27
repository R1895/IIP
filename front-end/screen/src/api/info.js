import http from '../utils/request.js'


const info = {
    getEnvInfoAnalysis(start, end, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
            start_date: start,
            end_date: end
        }
        return http.post(`/screen/environment`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },
    getNoiseInfoAnalysis(start, end, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
            start_date: start,
            end_date: end
        }
        return http.post(`/screen/noise`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },
    getSmellInfoAnalysis(start, end, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
            start_date: start,
            end_date: end
        }
        return http.post(`/screen/smell`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },
    getSafeInfoAnalysis(start, end, id) {
        if (!id) id = 1
        const params = {
            workshop_id: id,
            start_date: start,
            end_date: end
        }
        return http.post(`/screen/safe`, params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
        })
    },
}

export default info;
