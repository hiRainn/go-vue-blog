<template>
	<div class="item">
		<div v-if="contentHide">
			<a-skeleton :paragraph="{ rows: 50 }" />
		</div>
		<div v-if="!contentHide">
			<a-row>
				<h2 class="title">{{article.title}}</h2>
			</a-row>
			<a-row class="subinfo">
				<span class="info" v-if="!show_modify_time">{{$t('article.publish')}} {{article.post_at}}</span>
				<span class="info" v-if="show_modify_time">{{$t('article.modify')}} {{article.modify_at}}</span>
				<span class="info">
					<router-link style="color: #2C3E50;" :to="'/article/cate_id/' + article.cate_id">{{article.cate_name}}</router-link>
				</span>
				<a-icon type="eye" :style="{margin:'0px 2px 0px 5px'}"/>{{article.views}}
				<a-icon type="message" :style="{margin:'0px 2px 0px 5px'}"/>{{ comment_number }}</i>
				<a-icon type="heart" :style="{margin:'0px 2px 0px 5px'}"/>{{ article.like_number }}</i>
			</a-row>
			<a-row style="z-index: 0;">
				<mavon-editor ref='md' :language="language" :editable="false" codeStyle="atelier-lakeside-dark" v-model="article.content"
				 :toolbarsFlag="false" defaultOpen="preview" :subfield="false"></mavon-editor>
			</a-row>

			<a-row class="like-article">
				<a-button style="display: inline-block;height: 60px;border:none" @click="likeArticle">
					<a-icon type="heart" :theme="article.like?'filled':'outlined'" :style="{color:'red',fontSize:'60px',margin:'0px'}" />
				</a-button>
				<a-button style="display: inline-block;height: 60px;border:none">
					<a-icon type="pay-circle" theme="filled" :style="{color:'red',fontSize:'60px',margin:'0px'}" />
				</a-button>


			</a-row>
			<a-row>
				<comment @submit="comment" @clickUnlike="clickUnlike" @clickLike="clickLike" @clickReport="clickReport"
				 @cancleReport="cancleReport" @cancleLike="cancleLike" @cancleUnlike="cancleUnlike" :hideNumber="1" repeatType="cancle"
				 :title="comment_number?('当前共 ' + comment_number + ' 条评论'):'当前还没有评论，快来评论第一条吧~'" :allowComment="Boolean(article.allow_comment)"
				 :showReport="true" :reportText="'举报'" :showName="true" :showEmail="true" :replayText="'回复'" :list="comment_list"
				 :content="form.content" ref="comment" />
			</a-row>
		</div>
	</div>
</template>

<script>
	import {
		getArticleByid,
		getComments,
		commentArticle,
		likeArticle
	} from '@/api/article.js'
	import {
		like,
		dislike,
		report,
		cancleLike,
		cancleDislike,
		cancleReport
	} from '@/api/like.js'
	import {
		mavonEditor
	} from 'mavon-editor'

	import Comment from '@/components/Comment'
	export default {
		components: {
			mavonEditor,
			Comment,

			// or 'mavon-editor': mavonEditor
		},
		data() {
			return {
				contentHide: true,
				language: "en",
				show_modify_time: true,
				save: false,
				comment_list: [],
				comment_number: 0,
				article: {

				},
				form: {
					pid: 0,
					name: '',
					email: '',
					content: '',
					article_id: '',
					submit: false //this to tag whether the submit is maybe-sensitive
				}
			}
		},
		methods: {
			//get article
			getContent(id) {
				getArticleByid(id).then(r => {
					if (r.code != 0) {
						return;
					}
					if (r.data.modify_at == '' || r.data.modify_at.indexOf('1970') > -1) {
						this.show_modify_time = false
					}
					var likeArticle = localStorage.getItem('likeArticle')
					if( likeArticle != undefined && likeArticle != '') {
						var a = likeArticle.split(',')
						for(let p in a) {
							if(parseInt(a[p]) == parseInt(this.$route.params.id)) {
								r.data.like = 1
								break
							}
						}
					}
					this.article = r.data
					this.contentHide = false;
				}).catch(e => {
					console.log(e)
				})
			},
			//get comment
			getComments(id) {
				getComments(id).then(r => {
					if (r.code) {
						this.$alert(r.msg)
						return;
					} else {
						this.comment_list = r.data.list
						this.comment_number = r.data.total
					}
				}).catch(e => {
					console.log(e)
				})
			},
			//post comment
			comment(value, car) {
				this.form.pid = value.pid
				this.form.name = value.name
				this.form.email = value.email
				this.form.content = value.content
				this.form.article_id = parseInt(this.$route.params.id)
				commentArticle(this.form).then(r => {
					this.form.submit = false;
					if (r.code) {
						console.log(r)
						return;
					}
					if (r.data.status && !r.data.submit) {
						var sensitive = ''
						for (let p in r.data.msg) {
							sensitive = sensitive + p + ' '
						}
						sensitive = sensitive + '. ' + this.$i18n.t('os.sensitive_word_check')
						this.$confirm({
							content: this.$i18n.t('os.sensitive_word_notice') + sensitive,
							title: this.$i18n.t('os.tips'),
							okText: this.$i18n.t('comment.continue'),
							cancelText: this.$i18n.t('comment.edit'),
							onOk: () => {
								this.form.submit = true
								this.comment(this.form, car)
							},
							onCancel() {}
						});
					} else {
						this.$message.success(this.$i18n.t('comment.success'));

						var data = {
							id: r.data.id,
							name: this.form.name,
							content: this.htmlEscape(this.form.content),
						}

						if (r.data.status == 0) {
							this.comment_number++
						}
						this.form.content = ''
						car(data)
					}

				}).catch(e => {
					console.log(e)
				})
			},

			htmlEscape(text) {
				return text.replace(/[<>"&]/g, function(match, pos, originalText) {
					switch (match) {
						case "<":
							return "&lt;";
						case ">":
							return "&gt;";
						case "&":
							return "&amp;";
						case "\"":
							return "&quot;";
					}
				});
			},
			likeArticle() {
				if(this.article.like) {
					this.$message.error(this.$i18n.t('comment.cancle_like'))
					return false
				}
				var data = {
					article_id: parseInt(this.$route.params.id),
				}
				likeArticle(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						return
					}
					var likeArticle = localStorage.getItem('likeArticle')
					if(typeof(likeArticle) == 'undefined' || likeArticle == undefined || likeArticle == '') {
						likeArticle = this.$route.params.id
						localStorage.setItem('likeArticle',likeArticle)
					} else {
						likeArticle = likeArticle + ',' + this.$route.params.id
						localStorage.setItem('likeArticle',likeArticle)
					}
					this.$set(this.article,'like',1)
					this.article.like_number++
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
				})
			},
			clickUnlike(row, car) {
				// item.data.unlike_number++
				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				dislike(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			clickLike(row, car) {

				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				like(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			clickReport(row, car) {
				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				report(data).then(r => {
					if (r.code) {
						console.log(r)
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleReport(row, car) {
				this.$message.warning(this.$i18n.t('comment.cancke_report'));
				car(false)
				return;
				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				cancleReport(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleLike(row, car) {
				this.$message.warning(this.$i18n.t('comment.cancle_like'));
				car(false)
				return;
				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				cancleLike(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleUnlike(row, car) {
				this.$message.warning(this.$i18n.t('comment.cancle_dislike'));
				car(false)
				return;
				var data = {
					article_id: parseInt(this.$route.params.id),
					comment_id: row.id
				}
				cancleDislike(data).then(r => {
					if (r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch(e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			}
		},
		created() {
			var id = this.$route.params.id
			this.getContent(id)
			this.getComments(id)
			var language = localStorage.getItem('locate') ? localStorage.getItem('locate') : 'en'
			if (language == 'zh') {
				language = 'zh-CN'
			}
			this.language = language

			


		}
	}
</script>

<style scoped="scoped">
	.item {
		background: #fff;
		border: 1px solid #eee;
		text-align: left;
		border-radius: 10px;
		padding: 10px;
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
