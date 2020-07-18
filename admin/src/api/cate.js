import request from '@/utils/request'

//get all tags for select 
export function getSelectCateList() {
  return request({
    url: '/cate_for_select',
    method: 'get',
  })
}

export function addTag(data) {
  return request({
    url: '/cate',
    method: 'post',
	data : data
  })
}




