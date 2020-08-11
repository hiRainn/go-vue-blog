import request from '@/utils/request'


export function init() {
  return request({
    url: '/info',
    method: 'get',
  })
}

export function getCateArticle() {
	return request({
	  url: '/cate_article',
	  method: 'get',
	})
}
