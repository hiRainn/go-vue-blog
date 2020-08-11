<template>
	<div>
		<a-row>
			<h2 class="title">{{article.title}}</h2>
		</a-row>
		<a-row class="subinfo">
			<span class="info" v-if="!show_modify_time">{{$t('article.publish')}} {{article.post_at}}</span>
			<span class="info" v-if="show_modify_time">{{$t('article.modify')}} {{article.modify_at}}</span>
			<span class="info">
				<router-link style="color: #2C3E50;" :to="'/article/cate_id/' + article.cate_id">{{article.cate_name}}</router-link>
			</span>
			<a-icon type="eye" />{{article.views}}
			<a-icon type="message" />{{ comment_number }}</i>
		</a-row>
		<a-row style="z-index: 0;">
			<mavon-editor ref='md' :language="language" :editable="false" codeStyle="atelier-lakeside-dark" v-model="article.content"
			 :toolbarsFlag="false" defaultOpen="preview" :subfield="false"></mavon-editor>
		</a-row>

		<a-row class="comment-list">
			
		</a-row>
		<a-row>
			<comment 
				@submit="comment" 
				:reportText="'举报'" 
				:showName="true" 
				:showEmail="true"
				:replayText="'回复'" 
				:list="getCommentList" 
				:content="form.content" 
				ref="comment" />
		</a-row>
		<a-row style="height: 200px;">

		</a-row>
	</div>
</template>

<script>
	import {
		getArticleByid,
		getComments,
		commentArticle
	} from '@/api/article.js'
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
				
				language: "en",
				show_modify_time: true,
				save: false,
				comment_list: [],
				comment_active: [1, 2],
				comment_number: 0,
				article: {

				},
				form: {
					pid:0,
					name: '',
					email: '',
					content: '',
					article_id: '',
					submit: false  //this to tag whether the submit is maybe-sensitive
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
					this.article = r.data
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
			comment(value) {
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
						this.$confirm(this.$i18n.t('os.sensitive_word_notice') + sensitive, this.$i18n.t('os.tips'), {
							confirmButtonText: this.$i18n.t('comment.continue'),
							cancelButtonText: this.$i18n.t('comment.edit'),
							type: 'warning'
						}).then(() => {
							this.form.submit = true
							this.comment(this.form.content)
							return
						}).catch(() => {
							return
						});
					} else {
						this.$message({
							type: 'success',
							message: this.$i18n.t('comment.success')
						});
						var data = {
							id:r.data.id,
							name: this.form.name,
							content: this.htmlEscape(this.form.content),
							
						}
						
						if(r.data.status == 0) {
							this.comment_number++
						}
						this.form.content = ''
						this.$refs.comment.cleanContent()
					}

				}).catch(e => {
					console.log(e)
				})
			},
			
			htmlEscape(text){
			  return text.replace(/[<>"&]/g, function(match, pos, originalText){
			    switch(match){
			    case "<": return "&lt;";
			    case ">":return "&gt;";
			    case "&":return "&amp;";
			    case "\"":return "&quot;";
			  }
			  });
			}
		},
		computed:{
			getCommentList(){
				return this.comment_list
			},
		},
		mounted() {
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

	.comment-list {
		text-align: left;
		margin-top: 30px;
		/* background-color:#fbfbfb ; */
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
