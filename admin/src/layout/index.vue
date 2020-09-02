<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<el-col :span="3">
			<div :style="compute_style">
				<el-menu :style="style" :default-active="active" :router="true" text-color="#fff" background-color="#1a1918" :select="handleClickOutside"
				 class="el-menu-vertical-demo">
					<el-menu-item index="/">
						<i class="el-icon-menu"></i>
						<span slot="title" class="menu-span">{{$t('menu.index')}}</span>
					</el-menu-item>
					<el-submenu index="/article">
						<template slot="title">
							<i class="el-icon-location"></i>
							<span class="menu-span">{{$t('menu.article_manage')}}</span>
						</template>
						<el-menu-item-group>
							<el-menu-item index="/article" class="menu-span sub-menu-item">{{$t('menu.list')}}</el-menu-item>
							<el-menu-item index="/article/add" hidden></el-menu-item>
							<el-menu-item index="/article/drafts" class="menu-span sub-menu-item">{{$t('menu.drafts')}}</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
					<el-submenu index="/setting">
						<template slot="title">
							<i class="el-icon-setting"></i>
							<span class="menu-span">{{$t('menu.setting')}}</span>
						</template>
						<el-menu-item-group>
							<el-menu-item index="/setting" class="menu-span sub-menu-item">{{$t('menu.blog')}}</el-menu-item>
							<el-menu-item index="/setting/me" class="menu-span sub-menu-item">{{$t('menu.me')}}</el-menu-item>
							<el-menu-item index="/setting/changepass" class="menu-span sub-menu-item">{{$t('menu.changepass')}}</el-menu-item>
							<el-menu-item index="/setting/friends" class="menu-span sub-menu-item">{{$t('menu.friends')}}</el-menu-item>
						</el-menu-item-group>
					</el-submenu>
				</el-menu>
			</div>
			
		</el-col>
		<el-col :offset="1" :span="19" style="padding-top: 20px;">
			<app-main />
		</el-col>


	</div>
</template>

<script>
	import {
		AppMain
	} from './components'
	export default {
		name: 'Layout',
		components: {
			AppMain,
		},
		data() {
			return {
				active: '',
				menu_list: [],
				style: {
					height: "100vh",
					position:'fixed',
					width:'200px',
					top:'0px'
				},
				compute_style:{
					display:'inline-block',
					width:'200px',
					height:window.screen.availHeight + 'px'
				}
			}
		},
		methods: {
			handleClickOutside() {
				console.log(1)
			},
			jump(path) {
				this.$router.push({
					path: path
				})
			}
		},
		
		mounted() {
			// alert(this.style)
			this.menu_list = this.$router.options.routes;

			this.active = this.$route.path
			console.log(document.body.clientHeight)
			// alert(this.$route.path.trimLeft('/'))
			this.style.height = document.body.scrollHeight + 'px'
			this.compute_style.width = document.getElementsByClassName("el-menu")[0].clientWidth + 'px'
			
		},
		created() {
			
		}
	}
</script>

<style lang="scss" scoped>
	.sideBar {
		background-color: #1a1918;
	}
	
	.menu-span {
		text-decoration: none !important;
		font-size: 18px !important;
	}

	.sub-menu-item {
		padding-left: 60px !important;
	}

	a {
		text-decoration: none;
	}

	.router-link-active {
		text-decoration: none;
	}
</style>
