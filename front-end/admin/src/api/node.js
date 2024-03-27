import { http } from '@/utils/request'

export function getNodeList(params) {
  return http({
    url: '/node/nodes',
    method: 'get',
    params,
  })
}
