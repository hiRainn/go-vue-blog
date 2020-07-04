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
	
]

export default new VueRouter({
	
	routes:Routes
})



