import http from '@/utils/request'
import request from '@/utils/request'
import axios from 'axios'
// 用户管理，权限管理，角色管理


// 1.用户管理
// 查询
export function getUserManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/user/users`
    } else {
        url = `/user/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
        params
    })
}
// 添加
export function addUserManagerCol(params) {
    return request({
        url: `/user`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateUserManagerCol(params) {
    return request({
        url: `/user/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteUserManagerCol(params) {
    return request({
        url: `/user/${params.id}`,
        method: 'delete'
    })
}

// 2.角色管理
// 查询
export function getRoleManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/role/roles?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/role/${params.id}`
    }
    return request({
        url: url,
        method: 'get'
    })
}
// 添加
export function addRoleManagerCol(params) {
    return request({
        url: `/role`,
        method: 'post',
        data:params
    })
}
// 更新
export function updateRoleManagerCol(params) {
    return request({
        url: `/role/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deleteRoleManagerCol(params) {
    return request({
        url: `/role/${params.id}`,
        method: 'delete'
    })
}

// 3.权限管理
// 查询
export function getPermissionManagerCols(params) {
    let url
    if(params.id==""||params.id==void 0){
        url = `/permission/permissions?limit=${params.limit}&start=${params.start}`
    } else {
        url = `/permission/${params.id}`
    }
    return request({
        url: url,
        method: 'get',
    })
}
// 添加
export function addPermissionManagerCol(params) {
    return request({
        url: `/permission`,
        method: 'post',
        data:params
    })
}
// 更新
export function updatePermissionManagerCol(params) {
    return request({
        url: `/permission/${params.id}`,
        method: 'put',
        data:params
    })
}
// 删除
export function deletePermissionManagerCol(params) {
    return request({
        url: `/permission/${params.id}`,
        method: 'delete'
    })
}

  
