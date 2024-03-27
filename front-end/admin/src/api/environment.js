import request from '@/utils/request'

// 查询
export function getSaferateManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/safe/safes?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/safe/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addSaferateManagerCol(params) {
    return request({
        url: `/safe`,
        method: 'post',
        params
    })
}
// 更新
export function updateSaferateManagerCol(params) {
    return request({
        url: `/safe/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteSaferateManagerCol(params) {
    return request({
        url: `/safe/${params.id}`,
        method: 'delete'
    })
}


// 查询
export function getNoiseManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/noise/noises?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/noise/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addNoiseManagerCol(params) {
    return request({
        url: `/noise`,
        method: 'post',
        params
    })
}
// 更新
export function updateNoiseManagerCol(params) {
    return request({
        url: `/noise/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteNoiseManagerCol(params) {
    return request({
        url: `/noise/${params.id}`,
        method: 'delete',
        params
    })
}


// 查询
export function getSmellManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/smell/smells?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/smell/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addSmellManagerCol(params) {
    return request({
        url: `/smell`,
        method: 'post',
        params
    })
}
// 更新
export function updateSmellManagerCol(params) {
    return request({
        url: `/smell/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteSmellManagerCol(params) {
    return request({
        url: `/smell/${params.id}`,
        method: 'delete'
    })
}

// 查询
export function getEnvironmentManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/environment/environments?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/environment/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addEnvironmentManagerCol(params) {
    return request({
        url: `/environment`,
        method: 'post',
        params
    })
}
// 更新
export function updateEnvironmentManagerCol(params) {
    return request({
        url: `/environment/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteEnvironmentManagerCol(params) {
    return request({
        url: `/environment/${params.id}`,
        method: 'delete'
    })
}


  
