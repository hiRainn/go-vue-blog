import request from '@/utils/request'

export function getArticles(query) {
  return request({
    url: '/article',
    method: 'get',
    params: query
  })
}

export function getArticleInfo(id) {
  return request({
    url: '/article/'+ id ,
    method: 'get',
  })
}

export function postArticle(data) {
  return request({
    url: '/article',
    method: 'post',
    data: data
  })
}

export function updateArticle(data) {
  return request({
    url: '/article',
    method: 'put',
    data: data,
  })
}

export function postSaveArticle(data) {
  return request({
    url: '/save_article' ,
    method: 'put',
    data: data,
  })
}



export function saveArticle(data) {
  return request({
    url: '/save_article',
    method: 'post',
    data: data
  })
}

export function delArticle(data) {
  return request({
    url: '/article/' + data.id,
    method: 'delete',
    data: data
  })
}


