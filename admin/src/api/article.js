import request from '@/utils/request'

export function getArticles() {
  return request({
    url: '/articles/',
    method: 'get',
    params: query
  })
}
