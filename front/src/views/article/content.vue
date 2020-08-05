<template>
	<div>
		123
	</div>
</template>

<script>
	import {GetArticleByid} from '@/api/article.js'
	export default {
		data() {
			return {
				
			}
		},
		methods:{
			getContent(id) {
				GetArticleByid(id).then( r => {
					
				}).catch( e => {
					
				})
			}
		},
		mounted(){
			var id = this.$route.params.id
			this.getContent(id)
		}
	}
</script>

<style>
</style>
