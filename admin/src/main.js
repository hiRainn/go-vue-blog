import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Cookies from 'js-cookie'
import 'element-ui/lib/theme-chalk/index.css'
import i18n from './i18n/i18n'
import locale1 from 'element-ui/lib/locale/lang/en'
import locale2 from 'element-ui/lib/locale/lang/zh-CN'
import store from './store'

var language = localStorage.getItem('locate')?localStorage.getItem('locate'):'en'
var locale;
if(language == 'en') {
	locale = locale1
} else {
	locale = locale2
}



import Element from 'element-ui'

Vue.config.productionTip = false

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' ,// set element-ui default size,
  locale
})
//determine the page is init or not login
router.beforeEach(async(to, from, next) => {
  if(to.path !== '/init') {
	  if (to.path === '/logout') {
	    localStorage.removeItem('user');
	  }
	  let user = localStorage.getItem('user');
	  console.log(user)
	  //if user is not exists
	  if(!user) {
	  	  if(to.path !== '/login') {
	  		  next({ path:'/login' })
	  	  } else {
	  		  next()
	  	  }
	  } else {
	  	  if(to.path === '/login') {
	  		  next( {path :"/index"})
	  	  } else {
	  		  next()
	  	  }
	  }
  }
  next()
  
})
//select language

i18n.locale = language

new Vue({
	i18n,
    router,
	store,
    render: h => h(App),
}).$mount('#app')
