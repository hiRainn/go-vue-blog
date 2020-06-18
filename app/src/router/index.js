import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

/* Layout */
import Layout from '@/layout'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/',
      component: Layout,
      redirect: 'noredirect',
      children: [
        {
          path: 'index',
          component: () => import('@/views/index'),
          name: '首页',
          meta: { title: '首页', icon: 'chart', affix: true }
        },
        {
          path: 'life',
          component: () => import('@/views/life'),
          name: '生活',
          meta: { title: '生活', icon: 'chart', affix: true }
        },

      ]
    },
  ]
})
