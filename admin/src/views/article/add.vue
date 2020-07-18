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
			<el-form-item>
				<el-button type="primary" @click="postArticle()">{{$t('article.post')}}</el-button>
				<el-button>{{$t('article.save')}}</el-button>
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
		getSelectCateList,addTag
	} from '@/api/cate.js'
	import {getSelectTagsList} from '@/api/tags.js'
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
				form: {
					title: '',
					cate_id: '',
					content: '',
					tags: [],
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
					if(response.code) {
						this.$alert(response.msg)
					} else {
						this.cate_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			getSelectTagsList() {
				getSelectTagsList().then( response => {
					if(response.code) {
						this.$alert(response.msg)
					} else {
						this.tags_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			addcate() {
				if(this.cate_name == false) {
					this.$alert(this.$i18n.t('article.cate_rule'));
					return false;
				}
				var data = {
					name : this.cate_name
				}
				this.addDsiabled = true
				addTag(data).then(response => {
					if(response.code) {
						this.$alert(response.msg)
					} else {
						var add = {
							id : response.data.id,
							name: this.cate_name
						}
						this.cate_list.unshift(add)
						this.$message({
							message:this.$i18n.t('os.success'),
							type:'success'
						})
					}
					this.dialogFormVisible = false
					this.addDsiabled = false
				}).catch( err => {
					this.$alert(err)
					this.addDsiabled = false
				})
			},
			postArticle() {
				console.log(this.form)
			},

		},
		
		mounted() {
			this.getSelectCateList()
			this.getSelectTagsList()
		}

	}
</script>

<style scoped="scoped">

</style>
