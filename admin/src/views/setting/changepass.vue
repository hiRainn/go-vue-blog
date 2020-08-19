<template>
	<div class="continer">

		<el-form ref="form" :model="form" :rules="rules" label-width="80px" label-position="left">
			
			<el-form-item :label="$t('login.password')" prop="password">
				<el-input v-model="form.password"></el-input>
			</el-form-item>
			
			<el-form-item :label="$t('login.new_pass')" prop="newpass1">
				<el-input v-model="form.newpass1"></el-input>
			</el-form-item>
			
			<el-form-item :label="$t('login.repeat')" prop="newpass2">
				<el-input v-model="form.newpass2"></el-input>
			</el-form-item>

			<el-form-item>
				<el-button type="primary" @click="changPass()">{{$t('article.save')}}</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>

	import 'mavon-editor/dist/css/index.css'
	import {
		changPass,
	} from '@/api/setting.js'

	export default {
		name: 'add',
		data() {
			return {
				form:{
					password:'',
					newpass1:'',
					newpass2:'',
				},
				
				rules: {
					password: [{
						required: true,
						message: this.$i18n.t('login.password_rule'),
						trigger: 'blur'
					}, ],
					newpass1: [{
						required: true,
						message: this.$i18n.t('login.password_again'),
						trigger: 'blur'
					}, ],
					
				}
			}
		},
		methods: {
			
			changPass() {
				changPass(this.form).then(response => {
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
		},
		mounted() {
		}

	}
</script>

<style scoped="scoped">

</style>
