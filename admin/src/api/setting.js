import request from '@/utils/request'


export function getBlogSetting() {
  return request({
    url: '/blog',
    method: 'get',
  })
}

export function setBlogSetting(data) {
  return request({
    url: '/blog',
    method: 'post',
	data : data
  })
}

export function getUserSetting() {
  return request({
    url: '/user',
    method: 'get',
  })
}

export function setUserSetting(data) {
  return request({
    url: '/user',
    method: 'post',
	data : data
  })
}


export function changPass(data) {
  return request({
    url: '/change_pass',
    method: 'post',
	data : data
  })
}

export function getFriendsLink() {
	return request({
	  url: 'friends',
	  method: 'get',
	})
}

export function addFriend(data) {
  return request({
    url: '/friend',
    method: 'post',
	data : data
  })
}

export function delFriend(data) {
  return request({
    url: '/friend/' + data.id,
    method: 'delete',
    data: data
  })
}



