import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Cookies from 'js-cookie'
import 'element-ui/lib/theme-chalk/index.css'

import Element from 'element-ui'

Vue.config.productionTip = false

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})
//determine the page is init or not login
// router.beforeEach((to, from, next) => {
//   if (to.path == '/logout') {
//     sessionStorage.removeItem('user');
//   }
//   let user = sessionStorage.getItem('user');
//   if (!user && to.path != '/login') {
//     next({ path: '/login' })
//   } else {
//     next()
//   }
// })

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
