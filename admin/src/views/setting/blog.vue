<template>
	<div class="continer">

		<el-form ref="form" :model="form" :rules="rules" label-width="80px" label-position="left">
			
			<el-form-item :label="$t('blog.title')" prop="title">
				<el-input v-model="form.title"></el-input>
			</el-form-item>
			
			<el-form-item :label="$t('blog.subtitle')" prop="subtitle">
				<el-input v-model="form.subtitle"></el-input>
			</el-form-item>
			
			<el-form-item :label="$t('blog.language')" prop="language">
				<el-select v-model="form.language" filterable :placeholder="$t('language')" style="width: 100%;">
					<el-option class="option" v-for="item in language_list" :key="item.id" :label="item.name" :value="item.id">
					</el-option>
				</el-select>
			</el-form-item>



			<el-form-item>
				<el-button type="primary" @click="setBlogSetting()">{{$t('article.save')}}</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>

	import 'mavon-editor/dist/css/index.css'
	import {
		getBlogSetting,
		setBlogSetting
	} from '@/api/setting.js'

	import {
		uploadAvatar,
		uploadArticleImage
	} from '@/api/upload.js'
	export default {
		name: 'add',
		data() {
			return {
				addDsiabled: false,
				language_list: [{
						id: 'en',
						name: 'en'
					},
					{
						id: 'zh',
						name: 'zh'
					},
					
				],
				form:{
					title:'',
					subtitle:'',
					language:'',
				},
				language: 'en',
				rules: {
					title: [{
						required: true,
						message: this.$i18n.t('article.title_rule'),
						trigger: 'blur'
					}, ],
					
				}
			}
		},
		methods: {
			getBlogSetting() {
				getBlogSetting().then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						var data = response.data
						for (let p in data) {
							this.form[data[p]['key']] = data[p]['value']
						}
					}
				}).catch(err => {
					console.log(err)
				})
			},
			setBlogSetting() {
				setBlogSetting(this.form).then(response => {
					this.$message({
						message: this.$i18n.t('os.success'),
						type: 'success'
					});
					localStorage.setItem('locate',this.form.language)
				}).catch(err => {
					console.log(err)
					this.addDsiabled = false
				})

				console.log(this.form)
			},
		},

		mounted() {
			this.getBlogSetting()
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
