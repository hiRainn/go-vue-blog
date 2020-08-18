<template>
	<div id="page">
		<span>this is heihei's home</span>
		<div style="width: 100%" class="flex-container column">

		</div>
	</div>
</template>


<script>
	export default {

		data() {
			return {
				show:false,
				counter:0,
				sound: require('../../static/img/33.jpg'),
				video: ''
			}
		},
		methods: {
			play() {
				let vo = document.getElementById("video")

			},
			sbsb(){
				this.show = !this.show
			},
			plus(){
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
