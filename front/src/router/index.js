import Vue from 'vue'
import VueRouter from 'vue-router'
// import Layout from '@/layout'

Vue.use(VueRouter)

const Routes = [
	{
		path: '/',
		component: () => import('@/views/index/index'),
		hidden: false,
	},
	
	
	
]



export default new VueRouter({
	mode:'history',
	routes:Routes
})


