import request from '@/utils/request'
import {http} from '@/utils/request'

export function getNetworkList(params) {
  return http({
    url: '/network/networks',
    method: 'get',
    params,
  })
}
export function getList(params) {
  return request({
    url: '/vue-admin-template/table/list',
    method: 'get',
    params,
  })
}
