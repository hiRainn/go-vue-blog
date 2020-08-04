<template>
	<div>
		<ArticleItem :list="list"/>
	</div>
</template>

<script>
	import ArticleItem from '@/components/ArticleItem'
	import { getArticles } from '@/api/article.js'
	export default {
		name :'init',
		components: {
			ArticleItem,
		},
		data () {
			return {
				page: {
					p: 1,
					page_size: 15,
					total: 0
				},
				list:[]
			}
		},
		methods:{
			getArticles(params) {
				var query = {}
				if(params.name == 'cate_id') {
					query.cate_id = params.id
				}
				if(params.name == 'tag_id') {
					query.tag_id = params.id
				}
				getArticles(query).then( r => {
					if(r.code != 0) {
						console.log(r)
					} else {
						var data = r.data
						this.list = data.list
						this.page.p = data.p
						this.page.total = data.total
					}
				}).catch( e => {
					console.log(e)
				})
			},
			dian() {
				alert(1)
			}
		},
		mounted() {
			var data = this.$route.params
			this.getArticles(data)
		}
		
	}
</script>

<style>
</style>
