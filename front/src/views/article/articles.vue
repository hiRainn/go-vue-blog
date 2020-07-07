<template>
	<div>
		<button @click="dian"></button>
	</div>
</template>

<script>
	import {getArticles} from '@/api/article.js'
	export default {
		name :'init',
		data () {
			return {
				initForm:{
					
				}
			}
		},
		methods:{
			getArticles() {
				getArticles().then( response => {
					console.log(response)
				}).catch(err => {
					console.log(err)
				}) 
			},
			dian() {
				alert(1)
			}
		},
		mounted() {
			this.getArticles()
		}
		
	}
</script>

<style>
</style>
