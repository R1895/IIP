import { http } from '@/utils/request'

export function getShardingList(params) {
  return http({
    url: '/sharding/microchains',
    method: 'get',
    params,
  })
}
