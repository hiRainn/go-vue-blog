import Vue from 'vue'
import VueRouter from 'vue-router'
// import Layout from '@/layout'

Vue.use(VueRouter)

const Routes = [
	{
		path: '/login',
		component: () => import('@/views/login/login'),
		hidden: true,
	},
	{
		path: '/init',
		component: () => import('@/views/login/init'),
		hidden: false,
	},
	{
		path: '/index',
		component: () => import('@/views/index/index'),
		hidden: false,
	},
	{
		path: '/article',
		component: () => import('@/views/article/articles'),
		hidden: false,
	},
	
	
]



export default new VueRouter({
	mode:'hash',
	routes:Routes
})


