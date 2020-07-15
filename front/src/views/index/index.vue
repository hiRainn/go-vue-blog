<template>
	<div id="page">
		<span>this is heihei's home</span>
		<div style="width: 100%" class="flex-container column">
			<video id="video" src="../../static/audios/yhys.mp3" muted="true" onpagehide="auto" controls="controls" autoplay="autoplay"
			 loop="loop">

			</video>

			<img id="img" class="play" style="transform: translate(-1584.4%, 121.5%) scale(1)" :src="sound" @click="play()">

		</div>
	</div>
</template>


<script>
	export default {

		data() {
			return {
				sound: require('../../static/img/33.jpg'),
				video: ''
			}
		},
		methods: {
			play() {
				let vo = document.getElementById("video")

			},
			mulo($ccc) {
				this.video[$ccc] = !this.video[$ccc]
			}
		},
		mounted() {
			this.video = document.getElementById("video")
			this.video.addEventListener('click', this.mulo);
			this.mulo('autolpay');
			this.mulo('muted');
			if (this.video.paused) { //如果已暂停则播放
				this.video.play(); //播放控制
			} else { // 已播放点击则暂停
				this.video.pause(); //暂停控制
			}

		}
	}
</script>
