import request from '@/utils/request'

export function getArticles(query) {
  return request({
    url: '/article',
    method: 'get',
    params: query
  })
}

export function postArticle(data) {
  return request({
    url: '/article',
    method: 'post',
    data: data
  })
}


