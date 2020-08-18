<template>
	<div class="item">
		<div v-if="contentHide">
			<a-skeleton :paragraph="{ rows: 50 }" />
		</div>
		<div v-if="!contentHide">
			
			<a-row style="z-index: 0;">
				<mavon-editor ref='md' :editable="false" codeStyle="atelier-lakeside-dark" v-model="about"
				 :toolbarsFlag="false" defaultOpen="preview" :subfield="false"></mavon-editor>
			</a-row>

		</div>
	</div>
</template>

<script>
	import {
		About,
	} from '@/api/init.js'

	import {
		mavonEditor
	} from 'mavon-editor'

	export default {
		components: {
			mavonEditor,
		},
		data() {
			return {
				contentHide: true,
				about: "",
			}
		},
		methods: {
			//get
			About() {
				About().then(r => {
					if (r.code) {
						this.$alert(r.msg)
						return;
					} else {
						this.about = r.data
						this.contentHide = false
					}
				}).catch(e => {
					console.log(e)
				})
			},
		},
		created() {
			this.About()
		}
	}
</script>

<style scoped="scoped">
	.item {
		background: #fff;
		border: 1px solid #eee;
		text-align: left;
		border-radius: 10px;
		margin-bottom: 20px;
	}

	.title {
		margin: 0px;
		float: left;
	}

	a {
		text-decoration: none;
	}

	.comment {
		text-align: left;
		margin-top: 30px;
		padding: 15px;
		background-color: #fbfbfb;
	}

	.like-article {
		text-align: center;
		margin: 30px 0px;

		button {
			display: inline-block;
			text-align: center;
			padding: 20px 60px;
		}

		i {
			margin: 0px;
		}
	}

	.title {
		font-size: 24px;
	}

	.router-link-active {
		text-decoration: none;
	}

	.subinfo {
		text-align: left;
	}

	.el-icon-arrow-right>i {
		display: none !important;
	}

	.info {
		color: #606266;
		margin-right: 15px;
		font-size: 14px;
	}

	i {
		margin-right: 10px;
		font-size: 14px;
	}
</style>
