import request from '@/utils/request'


export function getMessage(){
	return request({
		url : 'message',
		method:'get'
	})
}

export function postMessage(data){
	return request({
		url : 'message',
		method:'post',
		data:data
	})
}

