<template>
	<div class="login-container">
		<div class="login">
			<el-form :model="loginForm" ref="loginForm" label-width="100px" class="demo-loginForm">
				<el-form-item :label="$t('login.username')" prop="region">
					<el-input v-model="loginForm.username" />
				</el-form-item>
				<el-form-item :label="$t('login.password')" prop="delivery">
					<el-input type="password" v-model="loginForm.password" />
				</el-form-item>
				<el-form-item>
					<el-button type="primary" @click="handleLogin('loginForm')">{{ $t('login.login') }}</el-button>
				</el-form-item>
			</el-form>
		</div>
	</div>
</template>

<script>
	import { MessageBox, Message } from 'element-ui'
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
			
			handleLogin() {
				this.$store.dispatch('user/login', this.loginForm)
					.then(() => {
						localStorage.setItem("user",true)
						this.$router.push({
							path: 'index'
						})
						this.loading = false
					})
					.catch(() => {
						this.loading = false
					})
			},
		
		},
		mounted() {
			this.checkInit()
		},
		created() {
			// window.addEventListener('storage', this.afterQRScan)
		},
		
		destroyed() {
			// window.removeEventListener('storage', this.afterQRScan)
		},
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
