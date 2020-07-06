import request from '@/utils/request'

export function checkInit(query) {
  return request({
    url: '/check_init',
    method: 'get',
    params: query
  })
}
