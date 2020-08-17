<template>
	<div>
		<a-row>
			<a-collapse>
				<a-collapse-panel key="1" header="This is panel header 1">
					<p>{{ text }}</p>
				</a-collapse-panel>
				<a-collapse-panel key="2" header="This is panel header 2" :disabled="false">
					<p>{{ text }}</p>
				</a-collapse-panel>
				<a-collapse-panel key="3" header="This is panel header 3" disabled>
					<p>{{ text }}</p>
				</a-collapse-panel>
			</a-collapse>
		</a-row>
	</div>
</template>


<script>
	import {getAllArticles} from '@/api/article.js'
	export default {

		data() {
			return {
				show: false,
				counter: 0,
				sound: require('../../static/img/33.jpg'),
				video: ''
			}
		},
		methods: {
			play() {
				let vo = document.getElementById("video")

			},
			sbsb() {
				this.show = !this.show
			},
			plus() {
				this.counter++
			},
			mulo($ccc) {
				this.video[$ccc] = !this.video[$ccc]
			}
		},
		mounted() {

		}
	}
</script>
