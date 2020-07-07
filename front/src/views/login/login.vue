<template>
	<div class="login-container">
		<div class="login">
			<el-form :model="loginForm" ref="loginForm" label-width="100px" class="demo-loginForm">
				<el-form-item label="username" prop="region">
					<el-input v-model="loginForm.username" />
				</el-form-item>
				<el-form-item label="password" prop="delivery">
					<el-input type="password" v-model="loginForm.password" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="submitForm('loginForm')">log in</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script>
	import {
		checkInit
	} from '@/api/init.js'
	import {
		login
	} from '@/api/user.js'
	export default {
		name: 'Login',
		data() {
			return {
				loginForm: {
					username: '',
					password: ''
				},
				init: false
			}
		},
		watch: {
			$route: {
				handler: function(route) {
					const query = route.query
					if (query) {
						this.redirect = query.redirect
						this.otherQuery = this.getOtherQuery(query)
					}
				},
				immediate: true
			}
		},
		created() {
			// window.addEventListener('storage', this.afterQRScan)
		},

		destroyed() {
			// window.removeEventListener('storage', this.afterQRScan)
		},
		methods: {
			checkInit() {
				checkInit().then(response => {
					if (!response.code) {
						this.init = true
						this.$router.push({
							path: '/init'
						})
					} else {
						return;
					}
				})
			},
			checkCapslock({
				shiftKey,
				key
			} = {}) {
				if (key && key.length === 1) {
					if (shiftKey && (key >= 'a' && key <= 'z') || !shiftKey && (key >= 'A' && key <= 'Z')) {
						this.capsTooltip = true
					} else {
						this.capsTooltip = false
					}
				}
				if (key === 'CapsLock' && this.capsTooltip === true) {
					this.capsTooltip = false
				}
			},
			showPwd() {
				if (this.passwordType === 'password') {
					this.passwordType = ''
				} else {
					this.passwordType = 'password'
				}
				this.$nextTick(() => {
					this.$refs.password.focus()
				})
			},
			handleLogin() {
				this.$store.dispatch('user/login', this.loginForm)
					.then(() => {
						this.$router.push({
							path: 'dashboard'
						})
						this.loading = false
					})
					.catch(() => {
						this.loading = false
					})
			},
			getOtherQuery(query) {
				return Object.keys(query).reduce((acc, cur) => {
					if (cur !== 'redirect') {
						acc[cur] = query[cur]
					}
					return acc
				}, {})
			}
			// afterQRScan() {
			//   if (e.key === 'x-admin-oauth-code') {
			//     const code = getQueryObject(e.newValue)
			//     const codeMap = {
			//       wechat: 'code',
			//       tencent: 'code'
			//     }
			//     const type = codeMap[this.auth_type]
			//     const codeName = code[type]
			//     if (codeName) {
			//       this.$store.dispatch('LoginByThirdparty', codeName).then(() => {
			//         this.$router.push({ path: this.redirect || '/' })
			//       })
			//     } else {
			//       alert('第三方登录失败')
			//     }
			//   }
			// }
		},
		mounted() {
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
