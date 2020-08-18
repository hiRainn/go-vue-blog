import request from '@/utils/request'


export function init() {
  return request({
    url: 'info',
    method: 'get',
  })
}

export function getCateArticle() {
	return request({
	  url: 'cate_article',
	  method: 'get',
	})
}

export function getFriendsLink() {
	return request({
	  url: 'friends',
	  method: 'get',
	})
}

export function getLatestComment() {
	return request({
	  url: 'latest_comments',
	  method: 'get',
	})
}

export function getTags() {
	return request({
	  url: '/tags',
	  method: 'get',
	})
}

export function getClickMost() {
	return request({
	  url: 'most',
	  method: 'get',
	})
}

export function getStat() {
	return request({
	  url: '/stat',
	  method: 'get',
	})
}

export function About() {
	return request({
	  url: '/about',
	  method: 'get',
	})
}


