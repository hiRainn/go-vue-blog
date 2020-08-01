<template>
	<div class="continer">
		<el-dialog :title="$t('article.quick_add_cate')" :visible.sync="dialogFormVisible" :close-on-click-modal="false">
			<el-form :model="form" ref="form" label-width="15%">
				<el-form-item :label="$t('article.cate_name')" required>
					<el-input v-model="cate_name" autocomplete="off" style="width: 100%;"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="dialogFormVisible = false">{{$t('os.cancle')}}</el-button>
				<el-button type="primary" :disabled="addDsiabled" @click="addcate()">{{$t('os.confirm')}}</el-button>
			</div>
		</el-dialog>
		<el-form ref="form" :model="form" :rules="rules" label-width="80px" label-position="left">
			<el-form-item :label="$t('article.title')" prop="title">
				<el-input v-model="form.title"></el-input>
			</el-form-item>
			<el-form-item :label="$t('article.cate')" prop="cate_id">
				<el-select v-model="form.cate_id" filterable :placeholder="$t('article.select_cate')" style="width: 95%;">
					<el-option class="option" v-for="item in cate_list" :key="item.id" :label="item.name" :value="item.id">
					</el-option>
				</el-select>
				<el-button @click="dialogFormVisible = true">+</el-button>
			</el-form-item>
			
			<el-form-item :label="$t('article.is_origin')">
				<el-col :span="2">
					<el-switch v-model="form.is_original" active-color="#13ce66" :active-text="$t('os.yes')">
					</el-switch>
				</el-col>
				<el-col :span="8" :hidden="form.is_original">
					<span style="margin-right: 5px;">{{ $t('article.reprint_from') }}:</span>
					<el-input v-model.number="form.reprint_from" @input="checkSort" :placeholder="$t('article.reprint_msg')" style="width:80%"></el-input>
				</el-col>
			</el-form-item>

			<el-form-item :label="$t('article.is_top')">
				<el-col :span="2">
					<el-switch v-model="form.is_top" active-color="#13ce66" :active-text="$t('os.yes')">
					</el-switch>
				</el-col>
				<el-col :span="8" :hidden="!form.is_top">
					<span style="margin-right: 5px;">{{ $t('article.sort') }}:</span>
					<el-input v-model.number="form.sort_num" @input="checkSort" :placeholder="$t('article.sort_content')" style="width:80%"></el-input>
				</el-col>
			</el-form-item>
			
			<el-form-item :label="$t('article.is_self')">
				<el-col :span="2">
					<el-switch v-model="form.is_self" active-color="#13ce66" :active-text="$t('os.yes')">
					</el-switch>
				</el-col>
			</el-form-item>
			
			<el-form-item :label="$t('article.no_comment')">
				<el-col :span="2">
					<el-switch v-model="form.allow_comment" active-color="#13ce66" :active-text="$t('os.yes')">
					</el-switch>
				</el-col>
			</el-form-item>

			<el-form-item>
				<mavon-editor ref='md' @imgAdd="imgAdd" :language="language" codeStyle="atelier-lakeside-dark" v-model="form.content" style="height: 600px"
				 :placeholder="$t('article.edit')"></mavon-editor>
			</el-form-item>
			<el-form-item :label="$t('article.tags')">
				<el-select style="width: 100%;" v-model="form.tags" :multiple-limit="5" multiple filterable allow-create
				 default-first-option :placeholder="$t('article.create_tags')">
					<el-option v-for="item in tags_list" :key="item.id" :label="item.name" :value="item.id">
					</el-option>
				</el-select>
			</el-form-item>
			<el-form-item :label="$t('article.create_at')">
				<el-date-picker v-model="form.create_at" type="datetime" format="yyyy-MM-dd HH:mm" placeholder="yyyy-mm-dd HH:mm"
				 value-format="yyyy-MM-dd HH:mm">
				</el-date-picker>
				<span>*{{$t('article.create_desc')}}*</span>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="postArticle()">{{$t('article.post')}}</el-button>
				<el-button :style="{ display: (is_draft || form.id == 0)?'inline-block':'none' }" :disabled="addDsiabled" @click="saveArticle()">{{$t('article.save')}}</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
	import {
		mavonEditor
	} from 'mavon-editor'
	import 'mavon-editor/dist/css/index.css'
	import {
		getSelectCateList,
		addTag
	} from '@/api/cate.js'
	import {
		getSelectTagsList
	} from '@/api/tags.js'
	import {
		postArticle,
		saveArticle,
		updateArticle,
		postSaveArticle,
		getArticleInfo
	} from '@/api/article.js'
	import {
		Format
	} from '@/utils/index.js'
	import {uploadArticleImage} from '@/api/upload.js'
	export default {
		name: 'add',
		components: {
			mavonEditor
			// or 'mavon-editor': mavonEditor
		},
		data() {
			return {
				dialogFormVisible: false,
				addDsiabled: false,
				cate_list: [],
				tags_list: [],
				cate_name: '',
				sort: 0,
				is_draft: false,
				language: 'en',
				form: {
					id: 0,
					title: '',
					cate_id: '',
					content: '',
					tags: [],
					is_top: false,
					is_self: false,
					is_original:true,
					allow_comment:true,
					reprint_from:'',
					sort: '',
					sort_num:'',
					create_at: '',
				},
				rules: {
					title: [{
						required: true,
						message: this.$i18n.t('article.title_rule'),
						trigger: 'blur'
					}, ],
					cate_id: [{
						required: true,
						message: this.$i18n.t('article.cate_rule'),
						trigger: 'blur'
					}, ],
					content: [{
						required: true,
						message: this.$i18n.t('article.content_rule'),
						trigger: 'blur'
					}, ],
				}
			}
		},
		methods: {
			getSelectCateList() {
				getSelectCateList().then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						this.cate_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			getSelectTagsList() {
				getSelectTagsList().then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						this.tags_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			checkSort(value) {
				if (value > 255 || value < 0) {
					this.form["sort_num"] = this['sort_num']
					return false;
				}
				this['sort_num'] = value
			},
			getArticleInfo(id) {
				getArticleInfo(id).then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						var data = response.data
						if (data.is_top == 1) {
							data.is_top = true
						} else {
							data.is_top = false
						}
						if (data.is_self == 1) {
							data.is_self = true
						} else {
							data.is_self = false
						}
						if (data.is_original == 1) {
							data.is_original = true
						} else {
							data.is_original = false
						}
						if (data.allow_comment == 1) {
							data.allow_comment = true
						} else {
							data.allow_comment = false
						}
						data.sort_num = data['sort']
						if(data.tags_ids != false) {
							var tags = data.tags_ids.split(',')
							data.tags = []
							for (var p in tags) {
								data.tags.push(parseInt(tags[p]))
							}
						}
						data.create_at = Format(parseInt(data.created_at) * 1000, 'yyyy-MM-dd hh:mm')
						if (data.status == 1) {
							this.is_draft = true
						}
						this.form = response.data
					}
				}).catch(err => {
					console.log(err)
				})
			},
			addcate() {
				if (this.cate_name == false) {
					this.$alert(this.$i18n.t('article.cate_rule'));
					return false;
				}
				var data = {
					name: this.cate_name
				}
				this.addDsiabled = true
				addTag(data).then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						var add = {
							id: response.data.id,
							name: this.cate_name
						}
						this.cate_list.unshift(add)
						this.$message({
							message: this.$i18n.t('os.success'),
							type: 'success'
						})
					}
					this.dialogFormVisible = false
					this.addDsiabled = false
				}).catch(err => {
					this.$alert(err)
					this.addDsiabled = false
				})
			},
			postArticle() {
				if (this.form['sort_num'] == false) {
					this.form['sort'] = 0
				} else {
					this.form['sort'] = parseInt(this.form['sort_num'])
				}
				this.addDsiabled = true
				console.log(this.form.id)
				//edit article
				if (this.form.id != 0) {
					if (!this.is_draft) {
						updateArticle(this.form).then(response => {
							this.$message({
								message: this.$i18n.t('os.success'),
								type: 'success'
							});
							this.$router.push({
								path: '/article'
							})
						}).catch(err => {
							console.log(err)
							this.addDsiabled = false
						})
					} else {
						postSaveArticle(this.form).then(response => {
							this.$message({
								message: this.$i18n.t('os.success'),
								type: 'success'
							});
							this.$router.push({
								path: '/article'
							})
						}).catch(err => {
							console.log(err)
							this.addDsiabled = false
						})
					}
				} else {
					//add article
					postArticle(this.form).then(response => {
						if (response.code) {
							this.$alert(response.msg)
							this.addDsiabled = false
						} else {
							this.form.id = response.data.id
							this.is_draft = true;
							this.addDsiabled = false;
							this.$message({
								message: this.$i18n.t('os.success'),
								type: 'success'
							});
							this.$router.push({
								path: '/article'
							})
						}
					}).catch(err => {
						console.log(err)
						this.is_new = true
					})
				}

				console.log(this.form)
			},
			saveArticle() {
				if (this.form['sort_num'] == false) {
					this.form['sort'] = 0
				} else {
					this.form['sort'] = parseInt(this.form['sort_num'])
				}
				this.addDsiabled = true
				saveArticle(this.form).then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						this.$message({
							message: this.$i18n.t('os.success'),
							type: 'success'
						});
						this.form.id = response.data.id
						this.is_draft = true
					}
					this.addDsiabled = false
				}).catch(err => {
					console.log(err)
					this.addDsiabled = false
				})
			},
			// 绑定@imgAdd event
			imgAdd(pos, $file) {
				// 第一步.将图片上传到服务器.
				var formdata = new FormData();
				formdata.append('image', $file);
				var md = this.$refs['md']
				uploadArticleImage(formdata).then(response => {
					if(response.code) {
						this.$alert(response.msg)
					} else {
						var data = response.data
						md.$img2Url(pos,data.url );
						
					}
					
				}).catch(err => {
					console.log(err)
				})
			}
		},

		mounted() {
			this.getSelectCateList()
			this.getSelectTagsList()
			var id = this.$route.query.id
			if (id !== undefined) {
				id = parseInt(id)
				if (id != 0) {
					this.form.id = id
					this.getArticleInfo(this.form.id)
				}
			}
			var language = localStorage.getItem('locate') ? localStorage.getItem('locate') : 'en'
			if (language == 'zh') {
				language = 'zh-CN'
			}
			this.language = language
			
			
		}

	}
</script>

<style scoped="scoped">

</style>
