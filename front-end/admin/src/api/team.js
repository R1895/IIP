import http from '@/utils/request'
import request from '@/utils/request'



export function getTaskManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/task/tasks?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/task/${params.id}`
    }
    console.log(url)
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addTaskManagerCol(params) {
    return request({
        url: `/task`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateTaskManagerCol(params) {
    return request({
        url: `/task/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteTaskManagerCol(params) {
    return request({
        url: `/task/${params.id}`,
        method: 'delete'
    })
}

// 查询
export function getWorkerManagerCols(params) {
    let url

    if(params.id==""||params.id==void 0){
        url = `/worker/workers`
    } else {
        url = `/worker/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
        params
    })
}
// 添加
export function addWorkerManagerCol(params) {
    return request({
        url: `/worker`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateWorkerManagerCol(params) {
    return request({
        url: `/worker/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteWorkerManagerCol(params) {
    return request({
        url: `/worker/${params.id}`,
        method: 'delete'
    })
}



// 查询
export function getWorkerAttendanceManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/workerAttendance/workerAttendances?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/workerAttendance/${params.id}`
    }
    console.log(url)
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addWorkerAttendanceManagerCol(params) {
    return request({
        url: `/workerAttendance`,
        method: 'post',
        params
    })
}
// 更新
export function updateWorkerAttendanceManagerCol(params) {
    return request({
        url: `/workerAttendance/${params.id}`,
        method: 'put',
        params
    })
}
// 删除
export function deleteWorkerAttendanceManagerCol(params) {
    return request({
        url: `/workerAttendance/${params.id}`,
        method: 'delete'
    })
}


  
