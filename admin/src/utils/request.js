import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import i18n from '../i18n/i18n.js'

axios.defaults.baseURL = "http://localhost:8080/bac"

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

    if (store.getters.token) {
      
      config.headers['X-Token'] = getToken()
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
    if (res.code != 0 && res.code != 30101) {
		
		Message({
		  message: res.msg || 'Error',
		  type: 'error',
		  duration: 5 * 1000
		})
		
     

      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (res.code == 300001 || res.code == 50012 || res.code == 50014) {
        // to re-login
		
	
        MessageBox.confirm(i18n.t('login.login_timeout'),i18n.t('login.tips') , {
          confirmButtonText: i18n.t('login.login'),
          cancelButtonText: i18n.t('login.cancle'),
          type: 'warning'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
			localStorage.removeItem("user")
            location.reload()
          })
        })
      }
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
