import http from '@/utils/request'
import request from '@/utils/request'
// 用户管理，权限管理，角色管理



export function getMessageManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/message/messages?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/message/${params.id}`
    }
    console.log(url)
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addMessageManagerCol(params) {
    return request({
        url: `/message`,
        method: 'post',
        params
    })
}
// 更新
export function updateMessageManagerCol(params) {
    return request({
        url: `/message/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteMessageManagerCol(params) {
    return request({
        url: `/message/${params.id}`,
        method: 'delete'
    })
}

// 2.车间管理
// 查询
export function getDictionaryItemManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/dictionaryItem/dictionaryItems?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/dictionaryItem/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addDictionaryItemManagerCol(params) {
    return request({
        url: `/dictionaryItem`,
        method: 'post',
        params
    })
}
// 更新
export function updateDictionaryItemManagerCol(params) {
    return request({
        url: `/dictionaryItem/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteDictionaryItemManagerCol(params) {
    return request({
        url: `/dictionaryItem/${params.id}`,
        method: 'delete',
        params
    })
}


// 查询
export function getLogManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/log/logs?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/log/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addLogManagerCol(params) {
    return request({
        url: `/log`,
        method: 'post',
        params
    })
}
// 更新
export function updateLogManagerCol(params) {
    return request({
        url: `/log/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteLogManagerCol(params) {
    return request({
        url: `/log/${params.id}`,
        method: 'delete'
    })
}


  
