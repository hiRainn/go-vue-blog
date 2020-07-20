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

			<el-form-item :label="$t('article.is_top')">
				<el-col :span="2">
					<el-switch v-model="form.is_top" active-color="#13ce66" :active-text="$t('os.yes')">
					</el-switch>
				</el-col>
				<el-col :span="8" :hidden="!form.is_top">
					<span style="margin-right: 5px;">{{ $t('article.sort') }}:</span>
					<el-input v-model.number="form.sort" @input="checkSort" :placeholder="$t('article.sort_content')" style="width:80%"></el-input>
				</el-col>
			</el-form-item>

			<el-form-item>
				<mavon-editor v-model="form.content" style="height: 600px" :placeholder="$t('article.edit')"></mavon-editor>
			</el-form-item>
			<el-form-item :label="$t('article.tags')">
				<el-select style="width: 100%;" v-model="form.tags" :multiple-limit="5" multiple filterable allow-create
				 default-first-option :placeholder="$t('article.create_tags')">
					<el-option v-for="item in tags_list" :key="item.id" :label="item.name" :value="item.id">
					</el-option>
				</el-select>
			</el-form-item>
			<el-form-item :label="$t('article.create_at')">
				<el-date-picker 
				v-model="form.create_at" 
				type="date" 
				format="yyyy-MM-dd"
				placeholder="yyyy-mm-dd"
				value-format="yyyy-MM-dd">
				</el-date-picker>
				<span>*{{$t('article.create_desc')}}*</span>
			</el-form-item>
			<el-form-item>
				<el-button type="primary" @click="postArticle()">{{$t('article.post')}}</el-button>
				<el-button :hidden="form.is_edit" :disabled="form.is_edit" @click="saveArticle()">{{$t('article.save')}}</el-button>
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
		postArticle,saveArticle,updateArticle,postSaveArticle
	} from '@/api/article.js'
	export default {
		name: 'add',
		components: {
			mavonEditor
			// or 'mavon-editor': mavonEditor
		},
		data() {
			return {
				is_new:true,
				dialogFormVisible: false,
				addDsiabled: false,
				cate_list: [],
				tags_list: [],
				cate_name: '',
				sort: 0,
				form: {
					id:0,
					title: '',
					cate_id: '',
					content: '',
					tags: [],
					is_top: false,
					sort: '',
					create_at:'',
					is_edit:false
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
					this.form["sort"] = this['sort']
					return false;
				}
				this['sort'] = value
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
				if(this.form['sort'] == false) {
					this.form['sort'] = 0
				} else {
					this.form['sort'] = parseInt(this.form['sorm'])
				}
				
				//edit article
				if(!this.is_new) {
					if(this.form.is_edit) {
						updateArticle(this.form).then(response => {
							console.log(response)
						}).catch(err => {
							console.log(err)
						})
					} else {
						postSaveArticle(this.form).then(response => {
							console.log(response)
						}).catch(err => {
							console.log(err)
						})
					}
				} else  {
					//add article
					postArticle(this.form).then(response => {
						if (response.code) {
							this.$alert(response.msg)
							this.is_new = true
						} else {
							this.form.id = response.data.id
							this.is_new = false
							this.form.is_edit = true;
						}
					}).catch(err => {
						console.log(err)
						this.is_new = true
					})
				}
				
				console.log(this.form)
			},
			saveArticle() {
				if(this.form['sort'] == false) {
					this.form['sort'] = 0
				}else {
					this.form['sort'] = parseInt(this.form['sorm'])
				}
				saveArticle(this.form).then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						this.form.id = response.data.id
					}
				}).catch(err => {
					console.log(err)
				})
			},
			getInfo(id) {
				
			}
		},

		mounted() {
			this.getSelectCateList()
			this.getSelectTagsList()
		}

	}
</script>

<style scoped="scoped">

</style>
