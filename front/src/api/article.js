import request from '@/utils/request'

export function getArticles(query) {
  return request({
    url: '/article',
    method: 'get',
    params: query
  })
}

export function getArticleByid(id) {
	return request({
		url : '/article/' + id,
		method:'get'
	})
}

export function getComments(id){
	return request({
		url : '/comment/' + id,
		method:'get'
	})
}

export function commentArticle(data){
	return request({
		url : '/comment',
		method:'post',
		data:data
	})
}
