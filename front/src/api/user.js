import request from '@/utils/request'

export function login(data) {
  return request({
    url: 'login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: 'userInfo',
    method: 'get',
    params: { token }
  })
}

export function changePass(token,data) {
  return request({
    url: 'changePass',
    method: 'post',
    params: { token },
    data:data
  })
}

export function logout() {
  return request({
    url: 'logout',
    method: 'post'
  })
}

export function getUserList() {
  return request({
    url: '/user/list',
    method: 'post'
  })
}

export function addUser(data) {
  return request({
    url: '/user/add',
    method: 'post',
    data:data
  })
}

export function editUser(data) {
  return request({
    url: '/user/edit',
    method: 'post',
    data:data
  })
}

export function delUser(data) {
  return request({
    url: '/user/del',
    method: 'post',
    data:data
  })
}
