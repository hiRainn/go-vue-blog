<template>
	<div class="login-container">
		<div class="login">
			<el-form :model="initForm" ref="Form" :rules="rules" label-width="100px" class="demo-loginForm">
				<el-form-item :label="$t('login.language')" prop="region">
					<el-select v-model="initForm.language" @change="changeL">
						<el-option v-for="item in language" :key="item.type" :label="item.value" :value="item.type">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item :label="$t('login.title')" prop="title">
					<el-input v-model="initForm.title" />
				</el-form-item>
				<el-form-item :label="$t('login.subtitle')">
					<el-input v-model="initForm.subtitle" />
				</el-form-item>
				<el-form-item :label="$t('login.username')" prop="username">
					<el-input v-model="initForm.username" />
				</el-form-item>
				<el-form-item :label="$t('login.password')" prop="password">
					<el-input type="password" v-model="initForm.password" />
				</el-form-item>
				<el-form-item :label="$t('login.repeat')" prop="repeat">
					<el-input type="password" v-model="initForm.repeat" />
				</el-form-item>
				<el-form-item :label="$t('login.nickname')" prop="region">
					<el-input v-model="initForm.nickname" />
				</el-form-item>
				<el-form-item :label="$t('login.birthday')" prop="region">
					<el-date-picker v-model="initForm.birthday" type="date" format="yyyy-MM-dd" value-format="yyyy-MM-dd">
					</el-date-picker>
				</el-form-item>
				<el-form-item :label="$t('login.intro')" prop="region">
					<el-input type="textarea" v-model="initForm.intro" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="submitForm('Form')">{{ $t('login.init') }}</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script>
	import { MessageBox, Message } from 'element-ui'
	import {
		checkInit,
		init
	} from '@/api/init.js'
	export default {
		name: 'init',
		data() {
			var validatePass = (rule, value, callback) => {
				if (value === '') {
					callback(new Error(this.$i18n.t('login.password_again')));
				} else if (value !== this.initForm.password) {
					console.log(value, this.initForm)
					callback(new Error(this.$i18n.t('login.repeat_rule')));
				} else {
					callback();
				}
			};
			return {

				language: [{
						type: 'en',
						value: 'english'
					},
					{
						type: 'zh',
						value: '简体中文'
					},
				],
				initForm: {
					title: '',
					subtitle:'',
					nickname: '',
					intro: '',
					birthday: '',
					username: '',
					password: '',
					repeat: '',
					language: ''
				},
				rules: {
					title: [{
						required: true,
						message: this.$i18n.t('login.title_rule'),
						trigger: 'blur'
					}, ],
					username: [{
						required: true,
						message: this.$i18n.t('login.username_rule'),
						trigger: 'blur'
					}, ],
					password: [{
						required: true,
						message: this.$i18n.t('login.password_rule'),
						trigger: 'blur'
					}],
					repeat: [{
						validator: validatePass,
						trigger: 'blur'
					}]
				}
			}
		},
		methods: {
			checkInit() {
				checkInit().then(response => {
					if(response.code) {
						this.$router.push({
							path: 'login'
						})
					} else {
						return false;
					}
				}).catch(err => {
					console.log(err)
				})
			},
			submitForm(formName) {
		
				this.$refs[formName].validate((valid) => {
					if (valid) {
						init(this.initForm).then( response => {
							if(response.code) {
								Message({
								  message: response.msg || 'Error',
								  type: 'error',
								  duration: 5 * 1000
								})
							} else {
								MessageBox.confirm(this.$i18n.t('login.init_ok'), 'success', {
								  confirmButtonText: 'ok',
								  type: 'success'
								}).then(() => {
								  this.$router.push({
								  	path: 'login'
								  })
								})
							}
						}).catch( err => {
							console.log(err)
						})
					} else {
						console.log('error submit!!');
						return false;
					}
				});
			},
			changeL(value) {
				localStorage.setItem('locate', value)
				this.$router.go(0)
			}
		},
		mounted() {
			var locate = localStorage.getItem('locate')
			if (locate != false) {
				this.initForm.language = locate
			}
			this.checkInit()
		}

	}
</script>

<style lang="scss" scoped>
	.login {
		margin: 0 auto;
		width: 400px;
		border: 1px solid #cef;
		border-radius: 5px;
		margin-top: 200px;
		padding: 50px 80px;
	}
</style>
