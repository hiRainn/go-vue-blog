<template>
	<div>
		<button @click="dian"></button>
	</div>
</template>

<script>
	import {checkInit} from '@/api/init.js'
	export default {
		name :'init',
		data () {
			return {
				initForm:{
					
				}
			}
		},
		methods:{
			checkInit() {
				checkInit().then( response => {
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
