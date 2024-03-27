import { http } from '@/utils/request'

export function getNetworkList(params) {
  return http({
    url: '/network/networks',
    method: 'get',
    params
  })
}

export function CreateNetwork(params) {
  return http({
    url: '/network',
    method: 'post',
    params
  })
}

export function DeleteNetwork(params) {
  return http({
    url: `/network/${params.id}`,
    method: 'delete',
    params
  })
}

export function UpdateNetwork(params) {
  return http({
    url: `/network/${params.id}`,
    method: 'put',
    params
  })
}
export function GetNetwork(params) {
  return http({
    url: `/network/${params.id}`,
    method: 'get',
    params
  })
}
