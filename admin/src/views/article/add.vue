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
			
				<el-select v-model="form.cate_id" placeholder="" style="width: 95%;">
					<el-option class="option" v-for="(name,id) in cate_list" :key="id" :label="name" :value="id">
					</el-option>
				</el-select>
				<el-button @click="dialogFormVisible = true">+</el-button>
			
			</el-form-item>
			<el-form-item>
				<mavon-editor style="height: 600px" :placeholder="$t('article.edit')"></mavon-editor>
			</el-form-item>
			<el-form-item :label="$t('article.tags')" >
				<el-input v-model="form.tags"></el-input>
			</el-form-item>
			<el-form-item>
				<el-button type="primary">立即创建</el-button>
				<el-button>取消</el-button>
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
		getArticles
	} from '@/api/article.js'
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
				cate_name: '',
				form: {
					title: '',
					cate_id: '',
					content:'',
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
			getArticles() {
				getArticles().then(response => {
					console.log(response)
				}).catch(err => {
					console.log(err)
				})
			},
			addcate() {
				alert(1)
			}

		},
		mounted() {
			this.getArticles()
		}

	}
</script>

<style>

</style>
