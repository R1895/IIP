import http from '../utils/request.js'


const stock = {
    getStockInfoAnalysis(limit, start, id) {
        let url = `/stock/stocks?limit=${limit}&start=${start}`
        if (id) url += `&workshop_id=${id}`
        return http.get(url)
    },

}

export default stock;
