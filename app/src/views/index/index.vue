<template>
	<div id="page">
		<span>this is heihei's home</span>
		<div style="width: 100%" class="flex-container column">
			<span>{{ counter }}</span>
			<a-button @click="plus">+++</a-button>

			<img id="img" class="play" style="transform: translate(-1584.4%, 121.5%) scale(1)" :src="sound" @click="play()">

		</div>
	</div>
</template>


<script>
	export default {

		data() {
			return {
				counter:0,
				sound: require('../../static/img/33.jpg'),
				video: ''
			}
		},
		methods: {
			play() {
				let vo = document.getElementById("video")

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
