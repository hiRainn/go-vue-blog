<template>
	<div class="container">

		<el-row>
			<ArticleItem :list="list" />
		</el-row>

		<el-row>
			<el-button v-if="!noMore" :loading="loading" @click="nextPage()">{{C_load_text}}</el-button>
			<el-button disabled v-if="noMore">{{nomore_text}}</el-button>
		</el-row>

	</div>
</template>

<script>
	import ArticleItem from '@/components/ArticleItem'
	import {
		getArticles
	} from '@/api/article.js'
	export default {
		name: 'init',
		components: {
			ArticleItem,
		},
		data() {
			return {
				loading_text: this.$i18n.t('article.load_text'),
				nomore_text: this.$i18n.t('article.nomore_text'),
				count: 10,
				loading: false,
				page: {
					p: 1,
					page_size: 8,
					total: 0
				},
				list: []
			}
		},
		methods: {
			getArticles(params) {
				var query = {}
				if (params.name == 'cate_id') {
					query.cate_id = params.id
				}
				if (params.name == 'tag') {
					query.tag = params.id
				}
				query.page_size = this.page.page_size
				query.p = this.page.p
				getArticles(query).then(r => {
					this.loading = false;
					this.loading_text = this.$i18n.t('article.load_text')
					if (r.code != 0) {
						console.log(r)
					} else {
						var data = r.data
						for (var p in data.list) {
							this.list.push(data.list[p])
						}
						this.page.p = data.p
						this.page.total = data.total
					}
				}).catch(e => {
					console.log(e)
					this.loading = false;
				})
			},
			nextPage() {
				this.page.p++
				var params = this.$route.params
				this.loading_text = this.$i18n.t('article.loading_text')
				this.loading = true
				this.getArticles(params)
			}
		},
		computed: {
			noMore() {
				return this.page.p * this.page.page_size >= this.page.total
			},
			C_load_text() {
				return this.loading_text
			}
		},
		mounted() {
			var data = this.$route.params
			this.getArticles(data)
			
		},
	}
</script>

<style>
</style>
