<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<el-row>
			<el-col :span="3">
				<div class="sideBar">
					<el-menu :default-active="active" :select="handleClickOutside" class="el-menu-vertical-demo" background-color="#d9f6ff">
						<router-link to='/'>
							<el-menu-item index="/">
								<i class="el-icon-menu"></i>
								<span slot="title" class="menu-span">{{$t('menu.index')}}</span>
							</el-menu-item>
						</router-link>
						<el-submenu index="">
							<template slot="title">
								<i class="el-icon-location"></i>
								<span class="menu-span">{{$t('menu.article_manage')}}</span>
							</template>
							<el-menu-item-group>
								<router-link to='/article'>
									<el-menu-item index="article" class="menu-span">{{$t('menu.list')}}</el-menu-item>
								</router-link>
								<router-link to='/article/add'>
									<el-menu-item index="article/add" class="menu-span">{{$t('menu.drafts')}}</el-menu-item>
								</router-link>
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
			}
		},
		mounted() {
			// alert(this.style)
			this.menu_list = this.$router.options.routes;

			this.active = this.$route.path.trimLeft('/')
		}
	}
</script>

<style lang="scss" scoped>
	.sideBar {
		height: 100vh;
		background-color: #d9f6ff;
	}
	
	.menu-span{
		text-decoration: none !important;
		font-size: 18px !important;
	}
	el-menu-item{
		padding-left: 20px;
	}
	
	a {
	  text-decoration: none;
	}
	 
	.router-link-active {
	  text-decoration: none;
	}
</style>
