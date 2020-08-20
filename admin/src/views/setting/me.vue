<template>
	<div class="continer">
		<el-form ref="form" :model="form" :rules="rules" label-width="80px" label-position="left">
			<el-form-item :label="$t('user.avatar')">
				<el-avatar shape="circle" :size="100" fit="fill" :src="form.avatar"></el-avatar>
				<el-upload
				  :on-success="(resopnse,file,fileList) => {avatatSucc(resopnse,file,fileList)}"
				  :headers="{'X-Token':token}"
				  name="image"
				  action="http://localhost:8080/bac/user/avatar"
				  class="upload-demo"
				  :limit="1"
				  >
				  <el-button size="small" type="primary">{{$t('user.upload')}}</el-button>
				</el-upload>
			</el-form-item>
			<el-form-item :label="$t('user.username')" prop="username">
				<el-input v-model="form.username"></el-input>
			</el-form-item>

			<el-form-item :label="$t('user.nickname')" prop="nickname">
				<el-input v-model="form.nickname"></el-input>
			</el-form-item>

			<el-form-item :label="$t('user.sign')" prop="sign">
				<el-input v-model="form.sign"></el-input>
			</el-form-item>

			<el-form-item :label="$t('user.city')" prop="city">
				<el-input v-model="form.city"></el-input>
			</el-form-item>

			<el-form-item :label="$t('user.gender')" prop="gender">
				<el-select v-model="form.gender" filterable :placeholder="$t('user.gender')" style="width: 100%;">
					<el-option class="option" v-for="item in gender_list" :key="item.id" :label="item.name" :value="item.id">
					</el-option>
				</el-select>
			</el-form-item>

			<el-form-item :label="$t('user.birthday')">
				<el-date-picker v-model="form.birthday" type="datetime" format="yyyy-MM-dd" placeholder="yyyy-mm-dd" value-format="yyyy-MM-dd">
				</el-date-picker>
			</el-form-item>

			<el-form-item label="github" prop="github">
				<el-input v-model="form.github"></el-input>
			</el-form-item>
			<el-form-item label="twitter" prop="twitter">
				<el-input v-model="form.twitter"></el-input>
			</el-form-item>
			<el-form-item label="facebook" prop="facebook">
				<el-input v-model="form.facebook"></el-input>
			</el-form-item>
			<el-form-item label="weibo" prop="weibo">
				<el-input v-model="form.weibo"></el-input>
			</el-form-item>

			<el-form-item :label="$t('user.intro')">
				<mavon-editor ref='md' @imgAdd="imgAdd" :language="language" codeStyle="atelier-lakeside-dark" v-model="form.intro"
				 style="height: 600px" :placeholder="$t('user.intro')"></mavon-editor>
			</el-form-item>

			<el-form-item>
				<el-button type="primary" @click="setUserSetting()">{{$t('article.save')}}</el-button>
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
		getUserSetting,
		setUserSetting
	} from '@/api/setting.js'

	import {
		uploadAvatar,
		uploadArticleImage
	} from '@/api/upload.js'
	export default {
		name: 'add',
		components: {
			mavonEditor
			// or 'mavon-editor': mavonEditor
		},
		data() {
			return {
				token:'',
				addDsiabled: false,
				gender_list: [{
						id: 0,
						name: this.$i18n.t('os.secret')
					},
					{
						id: 1,
						name: this.$i18n.t('os.male')
					},
					{
						id: 2,
						name: this.$i18n.t('os.female')
					}
				],
				language: 'en',
				form: {
					id: 0,
					username: '',
					nickname: '',
					birthday: '',
					email: '',
					intro: '',
					sign: '',
					twitter: '',
					facebook: '',
					weibo: '',
					github: '',
					city: '',
					gender: 0,
					avatar:''
				},
				rules: {
					username: [{
						required: true,
						message: this.$i18n.t('article.title_rule'),
						trigger: 'blur'
					}, ],
					nickname: [{
						required: true,
						message: this.$i18n.t('article.cate_rule'),
						trigger: 'blur'
					}, ],
				}
			}
		},
		methods: {
			avatatSucc(r,file,list) {
				this.form.avatar ='http://sorahei.com/static/upload/' + r.data.filename
			},
			getUserSetting() {
				getUserSetting().then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						var data = new Date(parseInt(response.data.birthday) * 1000)
						var YY = data.getFullYear() + '-';
						var MM = (data.getMonth() + 1 < 10 ? '0' + (data.getMonth() + 1) : data.getMonth() + 1) + '-';
						var DD = (data.getDate() < 10 ? '0' + (data.getDate()) : data.getDate());
						response.data.birthday = YY + MM + DD
						
						if (response.data.avatar.indexOf('http') < 0 ) {
							response.data.avatar = 'http://' + response.data.avatar
						}
						this.form = response.data
					}
				}).catch(err => {
					console.log(err)
				})
			},
			setUserSetting() {
				setUserSetting(this.form).then(response => {
					this.$message({
						message: this.$i18n.t('os.success'),
						type: 'success'
					});
				}).catch(err => {
					console.log(err)
					this.addDsiabled = false
				})

				console.log(this.form)
			},

			imgAdd(pos, $file) {
				// 第一步.将图片上传到服务器.
				var formdata = new FormData();
				formdata.append('image', $file);
				var md = this.$refs['md']
				uploadArticleImage(formdata).then(response => {
					if (response.code) {
						this.$alert(response.msg)
					} else {
						var data = response.data
						md.$img2Url(pos, 'http://sorahei.com/static/upload/'+data.filename);

					}

				}).catch(err => {
					console.log(err)
				})
			}
		},

		mounted() {
			this.getUserSetting()
			var language = localStorage.getItem('locate') ? localStorage.getItem('locate') : 'en'
			if (language == 'zh') {
				language = 'zh-CN'
			}
			this.language = language
			this.token = this.$store.getters.token;
		}

	}
</script>

<style scoped="scoped">

</style>
