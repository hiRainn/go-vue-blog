<template>
	<div id="page">
		<a-row style="margin-top: 20px;">
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
		</a-row>



	</div>
</template>


<script>
	import {
		indexArticle
	} from '@/api/article.js'
	import ArticleItem from '@/components/ArticleItem'
	export default {
		components: {
			ArticleItem,
		},
		data() {
			return {
				articleHide: true,
				list: []
			}
		},
		methods: {
			indexArticle() {
				indexArticle().then(r => {
					if (r.code) {
						this.$message.error(r.msg)
						this.articleHide = false
						return
					}
					this.articleHide = false
					this.list = r.data
				}).catch(e => {
					this.$message.error(e)
					console.log(e)
					this.articleHide = false
				})
			}
		},
		mounted() {
			this.indexArticle()
		}
	}
</script>

<style scoped>
	/* For demo */
	.item {
		background: #fff;
		border: 1px solid #eee;
		text-align: left;
		border-radius: 10px;
		padding: 10px;
		margin-bottom: 20px;
	}

</style>
