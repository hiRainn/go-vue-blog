<template>
	<div>
		<ArticleItem :list="list"/>
	</div>
</template>

<script>
	import ArticleItem from '@/components/ArticleItem'
	export default {
		name :'init',
		components: {
			ArticleItem,
		},
		data () {
			return {
				list:[]
			}
		},
		methods:{
			getArticles() {
				alert(1)
			},
			dian() {
				alert(1)
			}
		},
		mounted() {
		}
		
	}
</script>

<style>
</style>
