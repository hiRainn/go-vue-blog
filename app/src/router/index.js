import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/layout'

Vue.use(VueRouter)

const Routes = [
	// {
	// 	path: '/article',
	// 	component: Layout,
	// 	hidden: false,
	// 	children: [
	// 		{
	// 		path: 'add',
	// 		component: () => import('@/views/article/add'),
	// 		hidden: false
	// 		}, 
	// 		{
	// 		path: 'drafts',
	// 		component: () => import('@/views/article/drafts'),
	// 		hidden: false,
	// 		},
	// 		{
	// 		path: '/',
	// 		component: () => import('@/views/article/articles'),
	// 		hidden: false
	// 		},
	// 	]
	// },
	{
		path: '/article',
		component: Layout,
		hidden: false,
		children: [{
			path: ':name/:id',
			meta:{title:"article"},
			component: () => import('@/views/article/articles'),
			hidden: false
		},
		{
			path: ':id',
			meta:{},
			component: () => import('@/views/article/content'),
			hidden: false
		},
			,{
			path: '',
			meta:{title:"article"},
			component: () => import('@/views/article/articles'),
			hidden: false
		}, ]
	},
	{
		path: '/',
		component: Layout,
		hidden: false,
		children: [{
			path: '',
			meta:{title:"heihei's home"},
			component: () => import('@/views/index/index'),
			hidden: false
		}, ]
	},

]


export default new VueRouter({
	mode: 'history',
	routes: Routes
})