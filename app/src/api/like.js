import request from '@/utils/request'


export function like(data) {
  return request({
    url: 'like',
    method: 'post',
	data:data
  })
}

export function dislike(data) {
  return request({
    url: 'dislike',
    method: 'post',
	data:data
  })
}

export function report(data) {
  return request({
    url: 'report',
    method: 'post',
	data:data
  })
}

export function cancleLike(data) {
  return request({
    url: 'like',
    method: 'delete',
	data:data
  })
}

export function cancleDislike(data) {
  return request({
    url: 'dislike',
    method: 'delete',
	data:data
  })
}

export function cancleReport(data) {
  return request({
    url: 'report',
    method: 'delete',
	data:data
  })
}