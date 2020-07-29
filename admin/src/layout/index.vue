<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<el-row>
			<el-col :span="3">
				<div class="sideBar">
					<el-menu :default-active="active" text-color="#fff" background-color="#1a1918" :select="handleClickOutside" class="el-menu-vertical-demo">
						<el-menu-item index="/" @click="jump('/')">
							<i class="el-icon-menu"></i>
							<span slot="title" class="menu-span" >{{$t('menu.index')}}</span>
						</el-menu-item>
						<el-submenu index="1">
							<template slot="title">
								<i class="el-icon-location"></i>
								<span class="menu-span">{{$t('menu.article_manage')}}</span>
							</template>
							<el-menu-item-group>
								<el-menu-item index="/article" class="menu-span sub-menu-item" @click="jump('/article')">{{$t('menu.list')}}</el-menu-item>
								<el-menu-item index="/article/add" hidden ></el-menu-item>
								<el-menu-item index="/article/drafts" class="menu-span sub-menu-item" @click="jump('/article/drafts')">{{$t('menu.drafts')}}</el-menu-item>
							</el-menu-item-group>
						</el-submenu>


					</el-menu>
				</div>

			</el-col>
			<el-col :offset="1" :span="19" style="padding-top: 20px;">
				<app-main />
			</el-col>
		</el-row>

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
				style: 'top'
			}
		},
		methods: {
			handleClickOutside() {
				console.log(1)
			},
			jump(path) {
				this.$router.push({path:path})
			}
		},
		mounted() {
			// alert(this.style)
			this.menu_list = this.$router.options.routes;

			this.active = this.$route.path
			console.log(this.active)
			// alert(this.$route.path.trimLeft('/'))
		}
	}
</script>

<style lang="scss" scoped>
	.sideBar {
		height: 105vh;
		background-color: #1a1918;
	}
	
	.menu-span{
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
