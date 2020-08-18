<template>
	<div class="item">
		<div v-if="msgHide">
			 <a-skeleton  :paragraph="{ rows: 30 }" />
		</div>
		<a-row v-if="!msgHide">
			<comment 
				@submit="comment"
				@clickUnlike="clickUnlike" 
				@clickLike="clickLike" 
				@clickReport="clickReport"
				@cancleReport="cancleReport"
				@cancleLike="cancleLike"
				@cancleUnlike="cancleUnlike"
				:showReplay="false"
				:hideNumber="1"
				repeatType="cancle"
				:title="comment_number?('当前共 ' + comment_number + ' 条评论'):'当前还没有评论，快来评论第一条吧~'"
				:showReport="true"
				:reportText="'举报'" 
				:showName="true" 
				:showEmail="true"
				:replayText="'回复'" 
				:list="comment_list" 
				:content="form.content" 
				ref="comment" />
		</a-row>
		</div>
	</div>
</template>

<script>
	import {
		getMessage,
		postMessage
	} from '@/api/message.js'
	import {
		like,
		dislike,
		report,
		cancleLike,
		cancleDislike,
		cancleReport
	} from '@/api/like.js'
	import Comment from '@/components/Comment'
	export default {
		components: {
			Comment,
		},
		data() {
			return {
				msgHide:true,
				language: "en",
				show_modify_time: true,
				save: false,
				comment_list: [],
				comment_number: 0,
				form: {
					pid:0,
					name: '',
					email: '',
					content: '',
					submit: false  //this to tag whether the submit is maybe-sensitive
				}
			}
		},
		methods: {
			//get message
			getMessage() {
				getMessage().then(r => {
					if (r.code) {
						return;
					} else {
						this.comment_list = r.data.list
						this.comment_number = r.data.total
						this.msgHide = false
					}
				}).catch(e => {
					console.log(e)
				})
			},
			//post comment
			comment(value,car) {
				this.form.pid = value.pid
				this.form.name = value.name
				this.form.email = value.email
				this.form.content = value.content
				postMessage(this.form).then(r => {
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
							content:this.$i18n.t('os.sensitive_word_notice') + sensitive, 
							title:this.$i18n.t('os.tips'),
							okText: this.$i18n.t('comment.continue'),
							cancelText: this.$i18n.t('comment.edit'),
							onOk: () => {
								this.form.submit = true
								this.comment(this.form,car)
							},
							onCancel(){}
						});
					} else {
						 this.$message.success(this.$i18n.t('comment.success'));
						
						var data = {
							id:r.data.id,
							name: this.form.name,
							content: this.htmlEscape(this.form.content),
						}
						
						if(r.data.status == 0) {
							this.comment_number++
						}
						this.form.content = ''
						car(data)
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
			},
			clickUnlike(row,car) {
				// item.data.unlike_number++
				var data = {
					comment_id:row.id
				}
				dislike(data).then(r => {
					if(r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			clickLike(row,car) {
				var data = {
					comment_id:row.id
				}
				like(data).then(r => {
					if(r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			clickReport(row,car) {
				var data = {
					comment_id:row.id
				}
				report(data).then(r => {
					if(r.code) {
						console.log(r)
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleReport(row,car) {
				this.$message.warning(this.$i18n.t('comment.cancke_report'));
				car(false)
				return ;
				var data = {
					article_id:parseInt(this.$route.params.id),
					comment_id:row.id
				}
				cancleReport(data).then(r => {
					if(r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleLike(row,car) {
				this.$message.warning(this.$i18n.t('comment.cancle_like'));
				car(false)
				return ;
				var data = {
					article_id:parseInt(this.$route.params.id),
					comment_id:row.id
				}
				cancleLike(data).then(r => {
					if(r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			},
			cancleUnlike(row,car) {
				this.$message.warning(this.$i18n.t('comment.cancle_dislike'));
				car(false)
				return ;
				var data = {
					article_id:parseInt(this.$route.params.id),
					comment_id:row.id
				}
				cancleDislike(data).then(r => {
					if(r.code) {
						this.$message.error(r.msg);
						car(false)
						return
					}
					car(true)
				}).catch( e => {
					this.$message.error(e);
					console.log(e)
					car(false)
				})
			}
		},
		created() {
			this.getMessage()
			var language = localStorage.getItem('locate') ? localStorage.getItem('locate') : 'en'
			if (language == 'zh') {
				language = 'zh-CN'
			}
			this.language = language
		}
	}
</script>

<style scoped="scoped">
	.item{
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
