<template>
	<div class="login-container">
		<div class="login">
			<span> {{ $t('login.title') }}</span>
		</div>
	</div>
</template>

<script>
	import {
		checkInit
	} from '@/api/init.js'
	export default {
		name: 'init',
		data() {
			return {
				initForm: {
					username: '',
					password: '',
					password1: '',
				}
			}
		},
		methods: {
			checkInit() {
				checkInit().then(response => {
					console.log(response)
				}).catch(err => {
					console.log(err)
				})
			},
			dian() {
				alert(1)
			}
		},
		mounted() {
			this.checkInit()
		}

	}
</script>

<style>
</style>
