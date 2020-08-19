import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Cookies from 'js-cookie'
import i18n from './i18n/i18n'


import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';

Vue.use(Antd);

var language = localStorage.getItem('locate')?localStorage.getItem('locate'):'zh'
Vue.config.productionTip = false

router.beforeEach((to, from, next) => {
  window.document.title = to.meta.title == undefined?'heihei is smiling':to.meta.title,
  next()
});
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
//select language

i18n.locale = language

//弹出框禁止滑动
Vue.prototype.noScroll = function () {
  var mo = function (e) { e.preventDefault() }
  document.body.style.overflow = 'hidden'
  document.addEventListener('touchmove', mo, false)// 禁止页面滑动
}
 
//弹出框可以滑动
Vue.prototype.canScroll = function () {
  var mo = function (e) {
    e.preventDefault()
  }
  document.body.style.overflow = ''// 出现滚动条
  document.removeEventListener('touchmove', mo, false)
}

new Vue({
	i18n,
    router,
    render: h => h(App),
}).$mount('#app')
