<template>
	<div class="comment-wrap">
		<el-input type="textarea" id="textpanel" rows="3" v-model="content"></el-input>
		<!-- <el-input type="textarea" class="comment-input" placeholder="" id="textpanel" v-model="content"></el-input> -->
		<div class="opration">
			<div class="emoji-panel-btn" @click="showEmojiPanel">
				<img src="./assets/img/face_logo.png" />
			</div>
			<el-button @click="saveComment" round size="small" style="float: right;" :disabled="content.trim() == ''">{{buttonText}}</el-button>
		</div>
		<el-row>
			<emoji-panel @emojiClick="appendEmoji" v-if="isShowEmojiPanel"></emoji-panel>
		</el-row>
	</div>
</template>
<script>
	import EmojiPanel from "./children/EmojiPanel.vue";
	export default {
		props: {
			buttonText: {
				type: String,
				default: '提交'
			},
		},
		data() {
			return {
				content: "",
				isShowEmojiPanel: false,
			}
		},
		components: {
			EmojiPanel
		},
		methods: {
			saveComment() {
				// this.comments.push(this.content.replace(/:.*?:/g, this.emoji)); // 替换":"符号包含的字符串,通过emoji方法生成表情<span></span>
				this.isShowEmojiPanel = false;
				this.$emit('submit', this.content)
			},
			showEmojiPanel() {
				this.isShowEmojiPanel = !this.isShowEmojiPanel;
			},
			appendEmoji(text) {
				const el = document.getElementById("textpanel");
				this.isShowEmojiPanel = false;
				this.content = el.value + ":" + text + ":";
			}
		}
	};
</script>
<style lang="scss">
	// 注意 这里因为v-html的原因 不能使用scoped 不然样式不能失效
	.comment-wrap {
		.emoji-item-common {
			background: url("./assets/img/emoji_sprite.png");
			display: inline-block;

			&:hover {
				cursor: pointer;
			}
		}

		.emoji-size-small {
			// 表情大小
			zoom: 0.4;
		}

		.emoji-size-large {
			zoom: 0.4; // emojipanel表情大小
			margin: 2px;
		}

		.comments-list {
			margin-top: 20px;

			.comments-list-item {
				margin-bottom: 20px;

				.comments-list-item-heading {
					display: inline-block;

					img {
						height: 32px;
						width: 32px;
						border-radius: 50%;
						vertical-align: middle;
					}

					.comments-list-item-username {
						margin-left: 5px;
						font-weight: bold;
					}
				}

				.comments-list-item-content {
					margin: 10px 0px;
					border-bottom: 1px solid #cccccc;

					&:last-child {
						border-bottom: 0;
					}

					span {
						vertical-align: top;
					}
				}
			}
		}

		.comment-input {
			height: 100px;
			width: 500px;
			border: 1px solid #cccccc;
			border-radius: 5px;
			padding: 10px;
			resize: none;

			&:focus {
				outline: none;
			}
		}

		.opration {
			display: flex;
			justify-content: space-between;
			position: relative;

			.emoji-panel-btn {
				display: inline-block;

				&:hover {
					cursor: pointer;
					opacity: 0.8;
				}

				img {
					height: 24px;
					width: 24px;
				}
			}

			.comment-send-btn {
				width: 80px;
				height: 30px;
				line-height: 30px;
				text-align: center;
				border: 1px solid #1da1f2;
				border-radius: 100px;
				box-sizing: border-box;
				font-weight: bold;
				font-size: 13px;
				color: #ffffff;
				background-color: #4ab3f4;

				&:hover {
					cursor: pointer;
				}
			}

			.comment-send-btn-bounce {
				position: relative;
			}

			.comment-send-btn-bounce:focus {
				outline: none;
			}

			.comment-send-btn-bounce:after {
				content: "";
				display: block;
				position: absolute;
				top: 0px;
				left: 0px;
				right: 0px;
				bottom: 0px;
				border-radius: 100px;
				border: 0px solid #1da1f2;
				opacity: 0.7;
				transition: all 0.1s;
			}

			.comment-send-btn-bounce:active:after {
				//.bounce active时 伪元素:after的样式
				opacity: 1;
				top: -5px;
				left: -5px;
				right: -5px;
				bottom: -5px;
				border-radius: 100px;
				border: 2px solid #1da1f2;
				transition: all 0.2s;
			}
		}
	}

	@import "./assets/css/emoji.css"; // 导入精灵图样式
</style>
