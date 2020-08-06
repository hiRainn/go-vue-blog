<template>
	<div>
		<el-row>
			<h2 class="title">{{article.title}}</h2>
		</el-row>
		<el-row class="subinfo">
			<span class="info" v-if="!show_modify_time">{{$t('article.publish')}} {{article.post_at}}</span>
			<span class="info" v-if="show_modify_time">{{$t('article.modify')}} {{article.modify_at}}</span>
			<span class="info">
				<router-link style="color: #2C3E50;" :to="'/article/cate_id/' + article.cate_id">{{article.cate_name}}</router-link>
			</span>
			<i class="el-icon-view">{{article.views}}</i>
			<i class="el-icon-chat-dot-square">{{article.comments}}</i>
		</el-row>
		<el-row>
			<mavon-editor ref='md' :language="language" :editable="false" codeStyle="atelier-lakeside-dark" v-model="article.content"
			 :toolbarsFlag="false" defaultOpen="preview" :subfield="false"></mavon-editor>
		</el-row>
		
		<el-row class="comment-list">
			<comment-list :comments="comment_list"></comment-list>
		</el-row>
		<el-row class="comment">
			<span>{{$t('comment.email_msg')}}</span>
			<el-form :label-position="label_position" label-width="80px">
				<el-form-item :label="$t('comment.name')" style="padd">
					<el-input v-model="form.name"></el-input>
				</el-form-item>
				<el-form-item :label="$t('comment.email')">
					<el-input v-model="form.email"></el-input>
				</el-form-item>

				<el-form-item :label="$t('comment.content')">
					<comment @submit="comment"></comment>
				</el-form-item>
				<el-form-item>
					<el-checkbox v-model="save">{{$t('comment.save')}}</el-checkbox>
				</el-form-item>
			</el-form>

		</el-row>
		<el-row style="height: 200px;">

		</el-row>
	</div>
</template>

<script>
	import {
		getArticleByid,getComments,commentArticle
	} from '@/api/article.js'
	import {
		mavonEditor
	} from 'mavon-editor'

	import Comment from '@/components/Comment/Comment.vue'
	import CommentList from '@/components/Comment/CommentList.vue'
	export default {
		components: {
			mavonEditor,
			Comment,
			CommentList
			// or 'mavon-editor': mavonEditor
		},
		data() {
			return {
				label_position:'top',
				language: "en",
				show_modify_time: true,
				save:false,
				comment_list:[],
				comment_active:[1,2],
				comment_number:0,
				article: {

				},
				form: {
					name: '',
					email:'',
					content:'',
					id:''
				}
			}
		},
		methods: {
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
			getComments(id) {
				getComments(id).then( r => {
					if(r.code) {
						this.$alert(r.msg)
						return ;
					} else {
						this.comment_list = r.data.list
						this.comment_number = r.data.total
					}
				}).catch( e => {
					console.log(e)
				})
			},
			comment(value) {
				
				this.form.content = value
				this.form.article_id = parseInt(this.$route.params.id)
				commentArticle(this.form).then( r => {
					console.log(r)
				}).catch( e => {
					console.log(e)
				})
				
				
			}
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
	.comment{
		text-align: left;
		margin-top: 30px;
		padding: 15px;
		background-color:#fbfbfb ;
	}
	.comment-list{
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
	.el-icon-arrow-right > i{
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
