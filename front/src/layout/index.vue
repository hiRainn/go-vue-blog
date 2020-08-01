<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<div class="header">
			<el-col :span="24" class="hidden-md-and-down">
				<header class="pc-header"></header>
			</el-col>
		</div>
		<div class="menu">
			<!-- menu for phone -->
			<div class="hidden-sm-and-up">
				<el-row>
					<header class="phone-header">
						<el-button @click="drawer = true" type="small" style="float: left;background-color: #001528;border: none;">
						  <i class="el-icon-s-unfold"></i>
						</el-button>
						<el-drawer
						    size="50%"
							type="primary"
							:with-header="false"
						    :visible.sync="drawer"
						    direction="ltr">
						    <el-menu default-active="1-4-1" class="el-menu-vertical-demo"  :collapse="isCollapse">
						      <el-submenu index="1">
						        <template slot="title">
						          <i class="el-icon-location"></i>
						          <span slot="title">导航一</span>
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
						      <el-menu-item index="2">
						        <i class="el-icon-menu"></i>
						        <span slot="title">导航二</span>
						      </el-menu-item>
						      <el-menu-item index="3" disabled>
						        <i class="el-icon-document"></i>
						        <span slot="title">导航三</span>
						      </el-menu-item>
						      <el-menu-item index="4">
						        <i class="el-icon-setting"></i>
						        <span slot="title">导航四</span>
						      </el-menu-item>
						    </el-menu>
						</el-drawer>
					</header>
				</el-row>
				<el-row>
					<el-col :offset="5" :span="14">
						<el-avatar :size="100" :src="avatar" fit="contain" style="margin-top: 20px;" icon="el-icon-user-solid"></el-avatar>
						<div>1234</div>
					</el-col>
					
					<el-clo :span="5"></el-clo>
				</el-row>
			</div>
			<!-- menu for pad -->
			<div class="hidden-xs-only hidden-lg-and-up">
				<el-row>
					<el-col :offset="1" :span="22">
						<el-row>
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
						</el-row>
					</el-col>
					<el-col :span="1"></el-col>
				</el-row>
			</div>
			<!-- menu for pc -->
			<div class="hidden-md-only hidden-md-and-down">
				<el-row>
					<el-col :offset="4" :span="16">
						
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

					</el-col>
					<el-clo :span="4"></el-clo>
				</el-row>
			</div>
		</div>

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
				isCollapse:true,
				avatar:"http://localhost:8080/static/images/avatar.jpg",
				drawer: false,
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
				this.$router.push({
					path: path
				})
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
	.pc-header {
		height: 100px;
		border: 1px solid black;
	}
	.phone-header {
		height: 30px;
		border: 1px solid black;
		background-color: #001528;
	}

	a {
		text-decoration: none;
	}

	.router-link-active {
		text-decoration: none;
	}
</style>
