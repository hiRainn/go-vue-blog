<template>
	<div>
		
		 <mavon-editor style="height: 100%"></mavon-editor>
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
				initForm: {

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
			dian() {
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
