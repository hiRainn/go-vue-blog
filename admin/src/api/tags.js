import request from '@/utils/request'

//get all tags for select 
export function getSelectTagsList() {
  return request({
    url: '/tags_for_select',
    method: 'get',
  })
}