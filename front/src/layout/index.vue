<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<div class="header">
			<el-row>
			  <el-col :span="24" class="hidden-md-and-down">
				  <header></header>
			  </el-col>
			</el-row>
		</div>   
		<div class="menu">
			<!-- menu for phone -->  
			<div class="hidden-sm-and-up">
				手机菜单
			</div>
			<!-- menu for pad -->
			<div class="hidden-xs-only hidden-lg-and-up">
				平板菜单
			</div>
			<div class="hidden-md-only hidden-md-and-down">
				电脑菜单
			</div>
		</div>
		<el-row>
			<el-col :offset="3" :span="18">
				<menu>
					<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
					  <el-menu-item index="1">处理中心</el-menu-item>
					  <el-submenu index="2">
					    <template slot="title">我的工作台</template>
					    <el-menu-item index="2-1">选项1</el-menu-item>
					    <el-menu-item index="2-2">选项2</el-menu-item>
					    <el-menu-item index="2-3">选项3</el-menu-item>
					    <el-submenu index="2-4">
					      <template slot="title">选项4</template>
					      <el-menu-item index="2-4-1">选项1</el-menu-item>
					      <el-menu-item index="2-4-2">选项2</el-menu-item>
					      <el-menu-item index="2-4-3">选项3</el-menu-item>
					    </el-submenu>
					  </el-submenu>
					  <el-menu-item index="3" disabled>消息中心</el-menu-item>
					  <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
					</el-menu>
				</menu>
			</el-col>
			<el-clo :span="3"></el-clo>
		</el-row>
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
	import 'element-ui/lib/theme-chalk/display.css'
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
	header{
		height: 100px;
		border: 1px solid black;
	}
	
	a {
	  text-decoration: none;
	}
	 
	.router-link-active {
	  text-decoration: none;
	}
</style>
