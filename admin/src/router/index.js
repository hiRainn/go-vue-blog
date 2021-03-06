import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/layout'

Vue.use(VueRouter)

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push (location) {

return originalPush.call(this, location).catch(err => err)

}

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
		path: '/article',
		component: Layout,
		hidden: false,
		children: [
			{
			path: 'add',
			component: () => import('@/views/article/add'),
			hidden: false
			}, 
			{
			path: 'drafts',
			component: () => import('@/views/article/drafts'),
			hidden: false,
			},
			{
			path: '/',
			component: () => import('@/views/article/articles'),
			hidden: false
			},
		]
	},
	{
		path: '/setting',
		component: Layout,
		hidden: false,
		children: [
			{
			path: 'me',
			component: () => import('@/views/setting/me'),
			hidden: false
			}, 
			{
			path: 'changepass',
			component: () => import('@/views/setting/changepass'),
			hidden: false
			},
			{
			path: 'friends',
			component: () => import('@/views/setting/friends'),
			hidden: false
			},
			{
			path: '/',
			component: () => import('@/views/setting/blog'),
			hidden: false
			},
		]
	},
	{
		path: '/',
		component: Layout,
		hidden: false,
	},
	
]


export default new VueRouter({
	mode: 'history',
	routes: Routes
})
