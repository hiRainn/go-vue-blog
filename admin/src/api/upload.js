import request from '@/utils/request'

//upload article image
export function uploadArticleImage(data) {
  return request({
    url: '/article/upload',
    method: 'post',
    data: data,
    headers: {
    	'Content-Type': 'multipart/form-data'
    },
  })
}

export function uploadAvatar(data) {
  return request({
    url: 'user/avatar',
    method: 'post',
    data: data,
    headers: {
    	'Content-Type': 'multipart/form-data'
    },
  })
}







