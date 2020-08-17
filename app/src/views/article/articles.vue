<template>
	<div class="container">
		<a-row v-if="articleHide">
			<a-row class="item">
				<a-skeleton active />
			</a-row>
			<a-row class="item">
				<a-skeleton active />
			</a-row>
			<a-row class="item">
				<a-skeleton active />
			</a-row>
			
		</a-row>
		<a-row v-if="!articleHide">
			<ArticleItem :list="list" />
		</a-row>

		<a-row>
			<a-button v-if="!noMore && !articleHide" :loading="loading" @click="nextPage()">{{C_load_text}}</a-button>
			<a-button disabled v-if="noMore  && !articleHide">{{nomore_text}}</a-button>
		</a-row>

	</div>
</template>

<script>
	import ArticleItem from '@/components/ArticleItem'
	import {
		getArticles
	} from '@/api/article.js'
	export default {
		components: {
			ArticleItem,
		},
		data() {
			return {
				articleHide:true,
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
						this.articleHide = false
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
	.item{
		background: #fff;
		border: 1px solid #eee;
		text-align: left;
		border-radius: 10px;
		padding: 10px;
		margin-bottom: 20px;
	}
</style>
