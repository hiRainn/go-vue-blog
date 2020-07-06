import request from '@/utils/request'

export function checkInit(query) {
  return request({
    url: '/bac/check_init',
    method: 'get',
    params: query
  })
}
