<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
			<!-- header hidden under phone and pad -->
			<div class="hidden-sm-and-down">
				<el-row>
					<div class="header">
						<header class="pc-header" style="background-image: url('/static/img/37.jpg');">
							<el-col :offset="3" :span="14">
								<p class="pc-title" id="mytitle">{{blog.title}}</p>
							</el-col>
						</header>
					</div>
				</el-row>
			</div>
			<!-- menu for phone -->
			<div class="hidden-sm-and-up">
				<el-row>
					<header class="phone-header" style="display: block;">
						<el-button @click="drawer = true" type="" style="background-color: #001528;border: none;display: inline-block;">
						  <i class="el-icon-s-unfold"></i>
						</el-button>
						<span class="phone-title">{{ blog.title}}</span>
						<el-drawer
						    size="50%"
							type="primary"
							:with-header="false"
						    :visible.sync="drawer"
						    direction="ltr">
						    <el-menu default-active="1-4-1" :router="true" class="el-menu-vertical-demo"  :collapse="isCollapse">
						      <el-submenu index="">
						        <template slot="title">
						          <span slot="title">首页</span>
						        </template>
						        <el-menu-item-group>
						          <span slot="title">分组一</span>
						          <el-menu-item index="1-1">选项1</el-menu-item>
						          <el-menu-item index="1-2">选项2</el-menu-item>
						        </el-menu-item-group>
						        <el-menu-item-group title="分组2">
						          <el-menu-item index="1-3">选项3</el-menu-item>
						        </el-menu-item-group>
						        <el-submenu index="1-4">
						          <span slot="title">选项4</span>
						          <el-menu-item index="1-4-1">选项1</el-menu-item>
						        </el-submenu>
						      </el-submenu>
						    </el-menu>
						</el-drawer>
					</header>
				</el-row>
				<el-row v-if="this.$route.path == '/'">
					<el-col :offset="5" :span="14">
						<el-avatar :size="100" :src="avatar" fit="contain" style="margin-top: 20px;" icon="el-icon-user-solid"></el-avatar>
						<div class="phone-author-name"><b>{{author.nickname}}</b></div>
					</el-col>
					<el-col :span="5"></el-col>
				</el-row>
			</div>
			<!-- menu for pad -->
			<div class="hidden-xs-only hidden-lg-and-up">
				<el-row class="pad-menu">
					<el-col :offset="1" :span="18">
						<el-menu  :default-active="active" :router="true" text-color="#303133" active-text-color="#409EFF" @select="select_menu"  class="el-menu-demo pc-menu" mode="horizontal">
							<el-menu-item index="/">{{$t('menu.index')}}</el-menu-item>
							<el-submenu index="/article">
								<template slot="title"><span @click="goArticle">{{$t('menu.article')}}</span></template>
								<el-menu-item v-for="item in menu_list" :index="item.index">{{item.name}}</el-menu-item>
							</el-submenu>
							<el-menu-item index="3" disabled>消息中心</el-menu-item>
							<el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
						</el-menu>
					</el-col>
				</el-row>
				
			</div>
			<!-- menu for pc -->
			<div class="hidden-md-only hidden-md-and-down">
				<el-row class="pc-menu">
					<el-col :offset="3" :span="14">
						<el-menu  :default-active="active" :router="true" @select="select_menu" text-color="#303133" active-text-color="#409EFF" class="el-menu-demo pc-menu" mode="horizontal">
							<el-menu-item index="/">{{$t('menu.index')}}</el-menu-item>
							<el-submenu index="/article">
								<template slot="title"><span @click="goArticle">{{$t('menu.article')}}</span></template>
								<el-menu-item v-for="item in menu_list" :index="item.index">{{item.name}}</el-menu-item>
							</el-submenu>
							<el-menu-item index="3" disabled>消息中心</el-menu-item>
							<el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
						</el-menu>
					</el-col>
					<el-col :offset="2" :span="4">
						<video id="video" src="../static/audios/yhys.mp3" muted="true" onpagehide="auto" controls="controls" autoplay="autoplay"
						 loop="loop" style="display: inline-block;height: 57px;width: 280px;line-height: 60px;">
						</video>
					</el-col>
				</el-row>
				
			</div>
			
			<div class="content">
				<el-row>
					<el-col :md="{span:16,offset:4}" :xs="{span:22,offset:1}" :sm="{span:22,offset:1}">
						<el-row>
							<!-- content -->
							<el-col :sm="17" :xs="24">
								<div>
									<el-row :hidden="!Cshow_breadcumb">
										<div class="breadcumb">
											<el-breadcrumb separator-class="el-icon-arrow-right">
											  <el-breadcrumb-item to="/"><i class="el-icon-s-home" @click="show_breadcumb=false">{{$t('menu.index')}}</i></el-breadcrumb-item>
											  <el-breadcrumb-item>活动管理</el-breadcrumb-item>
											  <el-breadcrumb-item>活动列表</el-breadcrumb-item>
											  <el-breadcrumb-item>活动详情</el-breadcrumb-item>
											</el-breadcrumb>
										</div>
									</el-row>
									<el-row>
										<app-main />
									</el-row>
								</div>
							</el-col>
							<el-col :sm="7" :xs="0">
								<div>
									<el-row style="margin-top: 20px;">
										<el-col :span="8" :offset="1">
											<el-avatar :size="70" :src="avatar"></el-avatar>
										</el-col>
										<el-col style="text-align: left;" :span="15"> 
											<div><b>{{author.nickname}}</b></div>
											<div style="font-size: 13px;">{{author.intro}}</div>
										</el-col>
									</el-row>
								</div>
							</el-col>
						</el-row>
					</el-col>
					<el-col :md="4" :xs="1" :sm="1" >
						<el-backtop></el-backtop>
					</el-col>
				</el-row>
			</div>
			
	</div>
</template>

<script>
	import 'element-ui/lib/theme-chalk/display.css'
	import {init,getCateArticle} from '@/api/init.js'
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
				isCollapse:true,
				avatar:"http://localhost:8080/static/images/avatar.jpg",
				drawer: false,
				active: '',
				menu_list: [],
				style: 'top',
				author:{
					
				},
				blog:{
					
				},
				show_breadcumb:true
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
			},
			init() {
				init().then(r => {
					if(r.code != 0) {
						this.$alert(r.msg)
						return ;
					}
					this.author = r.data.author
					this.blog = r.data.blog_info
				}).catch(e => {
					console.log(e)
				})
			},
			getCateArticle() {
				getCateArticle().then( r => {
					if (r.code != 0) {
						console.log(r)
					} else {
						var data = r.data
						for(var p in data) {
							this.menu_list.push({
								index:'/article/cate_id/' + data[p]['id'],
								name : data[p]['cate_name'] + '(' + data[p]['num'] + ')'
							})
						}
						console.log(this.menu_list)
					}
				}).catch( e => {
					console.log(e)
				})
			},
			show() {
				if(this.$route.path == "" ||  this.$route.path == '/') {
					this.show_breadcumb = false
				} else {
					this.show_breadcumb = true
				}
			},
			goArticle() {
				this.$router.push({
					path: '/article'
				})
				this.show_breadcumb = true
				this.active = '/article'
			},
			select_menu(index) {
				if(index == '/') {
					this.show_breadcumb = false
				} else {
					this.show_breadcumb = true
				}
				this.active = index
			}
		},
		computed:{
			Cshow_breadcumb() {
				return this.show_breadcumb;
			}
		},
		mounted() {
			this.show()
			this.init()
			this.getCateArticle()
		}
	}
	


	
</script>

<style lang="scss" scoped>
	
	.pc-header {
		height: 140px;
		// border: 1px solid black;
	}
	.phone-header {
		height: 40px;
		border: 1px solid black;
		background-color: #001528;
		text-align: left;
		width: 100%;
		display: block;
	}
	.phone-title {
		color:#FFF8E6;
		line-height:40px;
	}
	.pc-title{
		float: left;
	}
	.phone-author-name {
		font-size: 20px;
	}
	.pc-menu {
		background: #F0C78A;
		height: 60px;
		color:#000000;
	}
	.pad-menu {
		background: #F0C78A;
	}
	a {
		text-decoration: none;
	}
	.breadcumb {
		height: 30px;
		line-height: 30px;
		margin-top: 10px;
	}
	.router-link-active {
		text-decoration: none;
	}
	p {
	  margin: 5px;
	  font-size: 2em;
	  font-weight: 600;
	}
	#mytitle {
	    -webkit-animation: mymove 5s infinite; /* Chrome, Safari, Opera */
	    animation: mymove 5s infinite;
	}
	
	/* Chrome, Safari, Opera */
	@-webkit-keyframes mymove {
	    50% {text-color: black;}
	}
	
	/* Standard syntax */
	@keyframes mymove {
	    50% {color: blue;}
	}
</style>
