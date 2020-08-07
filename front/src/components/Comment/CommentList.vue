<template>
	<div class="comment-wrap">
		<div class="comments-list">
			<div class="comments-list-item" v-for="(item,index) in comments" v-bind:key="index">
				<div class="comments-list-item-heading">
					<img src="./assets/img/heading.jpg" />
					<span class="comments-list-item-username">{{item.name}}</span>
				</div>
				<div class="comments-list-item-content" v-html="item.content"></div>
			</div>
		</div>
		
		<hr>
		
	</div>
</template>
<script>
	export default {
		name:'CommentsList',
		props: {
			comments: {
				type: Array,
				default: []
			}
		},
		data() {
			return {
				
			}
		},
		methods: {
			emoji(word) {
				// 生成html
				const type = word.substring(1, word.length - 1);
				return `<span class="emoji-item-common emoji-${type} emoji-size-small" ></span>`;
			},
		},
		updated(){
			for(let p in this.comments) {
				this.comments[p]['name'] = this.comments[p]['name']
				this.comments[p]['content'] = this.comments[p]['content'].replace(/:.*?:/g, this.emoji);
			}
		},
	};
</script>
<style lang="scss">


</style>
