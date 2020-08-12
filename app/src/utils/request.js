import axios from 'axios'
import {
	getToken
} from '@/utils/auth'
import md5 from 'js-md5';

axios.defaults.baseURL = "http://localhost:8080"

// create an axios instance
const service = axios.create({
	baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
	// withCredentials: true, // send cookies when cross-domain requests
	timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
	config => {
		// do something before request is sent
		var name = localStorage.getItem("comment_name")
		var email = localStorage.getItem("comment_email")
		if (name != null && email != null) {
			config.headers['X-Token'] = md5(name + email)
		}
		return config
	},
	error => {
		// do something with request error
		console.log(error) // for debug
		return Promise.reject(error)
	}
)

// response interceptor
service.interceptors.response.use(
	/**
	 * If you want to get http information such as headers or status
	 * Please return  response => response
	 */

	/**
	 * Determine the request status by custom code
	 * Here is just an example
	 * You can also judge the status by HTTP Status Code
	 */
	response => {
		const res = response.data

		// code 30101 is checkinit
		if (res.code != 0) {
			
			this.$message.error(res.msg || 'Error');
			return Promise.reject(new Error(res.msg || 'Error'))
		} else {
			return res
		}
	},
	error => {
		console.log('err' + error) // for debug
		// Message({
		//   message: error.message,
		//   type: 'error',
		//   duration: 5 * 1000
		// })
		return Promise.reject(error)
	}
)

export default service
