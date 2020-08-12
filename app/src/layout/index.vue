<template>
	<div class="" style="margin-top: 0px;padding-top: 0px;">
		<!-- header hidden under phone and pad -->
		<a-row>
			<a-col class="header" :xs="0" :md="0" :xl="24">
				<header class="pc-header" style="background-image: url('/static/img/37.jpg');">
					<a-col :offset="3" :span="14">
						<p class="pc-title" id="mytitle">{{blog.title}}</p>
					</a-col>
				</header>
			</a-col>
		</a-row>
		<!-- menu -->
		<a-row>
			<!-- menu for phone -->
			<a-row>
				<a-col :xs="24" :md="0">
					<a-affix :offset-top="0" :offsetBottom="0">
						<a-row>
							<header class="phone-header" style="display: block;">
								<a-button @click="drawer = true" class="phone-title-botton">
									<a-icon type="menu-fold" v-if="drawer" />
									<a-icon type="menu-unfold" v-if="!drawer" />
								</a-button>
								<span class="phone-title">{{ blog.title}}</span>
				
								<a-drawer :maskClosable="true" :bodyStyle="{padding:'24px 0px'}" width="200" placement="left" :closable="false"
								 :visible.async="drawer" @close="ondrawerClose">
									<a-menu mode="vertical" @click="menuClick" :selectable="false">
										<a-menu-item key="/" :disabled="($route.path == '' || $route.path == '/')">
											<a-icon type="home" />
											{{$t('menu.index')}}
										</a-menu-item>
										<a-sub-menu key="article">
											<span slot="title">
												<a-icon type="book" />
												<span>{{$t('menu.article')}}</span>
											</span>
											<a-menu-item v-for="item in menu_list"  :key="item.key" :disabled="$route.path == item.key">
												{{item.name}}
											</a-menu-item>
										</a-sub-menu>
									</a-menu>
								</a-drawer>
							</header>
						</a-row>
					</a-affix>
					<a-row v-if="this.$route.path == '/'">
						<a-col :offset="5" :span="14">
							<a-avatar :size="100" :src="avatar"  style="margin-top: 20px;"></a-avatar>
							<div class="phone-author-name"><b>{{author.nickname}}</b></div>
						</a-col>
						<a-col :span="5"></a-col>
					</a-row>
				</a-col>
			</a-row>
			

			<!-- menu for pad && pc -->
			<a-row>
				<a-affix :offset-top="0" :offsetBottom="0">
					<a-col class="pp-menu" :xs="0" :md="24">
						<a-col :md="{offset:1,span:22}" :lg="{offset:3,span:14}" :xs="0">
							<a-menu @click="menuClick" :selectable="false"
							 class="pp-menu" mode="horizontal">
								<a-menu-item key="/" :disabled="($route.path == '' || $route.path == '/')">
									<a-icon type="home" />
									{{$t('menu.index')}}
								</a-menu-item>
								<a-sub-menu key="article">
									<span slot="title">
										<a-icon type="book" />
										<span>{{$t('menu.article')}}</span>
									</span>
									<a-menu-item v-for="item in menu_list"  :key="item.key" :disabled="$route.path == item.key">
										{{item.name}}
									</a-menu-item>
								</a-sub-menu>
							</a-menu>
						</a-col>
						<a-col :md="1" :lg="0" :xs="0"></a-col>
						<a-col :md="0" :lg="{offset:2,span:4}" :xs="0">
							<video id="video" src="../static/audios/yhys.mp3" muted="true" onpagehide="auto" controls="controls" autoplay="autoplay"
							 loop="loop" style="display: inline-block;height: 46px;width: 100%;margin:2px;line-height: 50px;">
							</video>
						</a-col>
					</a-col>
				</a-affix>
			</a-row>
			
		</a-row>

		
			<a-row>
				<a-col :lg="{span:16,offset:4}" :xs="{span:22,offset:1}">
				<a-row>
					<!-- content -->
					<a-col :md="17" :xs="24">
						<a-row :hidden="!Cshow_breadcumb" class="breadcrumb">
							<a-breadcrumb separator="/" style="float: left;margin: 5px;font-size: 14px ;">
								<a-breadcrumb-item >
									<template>
										<router-link to="/" >
											<span @click="show_breadcumb = false">
												<a-icon type="home" />{{$t('menu.index')}}
											</span>
										</router-link>
									</template>
								</a-breadcrumb-item>
							</a-breadcrumb>
						</a-row>
						<a-row>
							<app-main />
						</a-row>
					</a-col>
					<a-col :md="7" :xs="0">
						<div>
							<a-row style="margin-top: 20px;">
								<a-col :span="8" :offset="1">
									<a-avatar :size="100" :src="avatar"  style="margin-top: 20px;"></a-avatar>
								</a-col>
								<a-col style="text-align: left;" :span="15">
									<div><b>{{author.nickname}}</b></div>
									<div style="font-size: 13px;">{{author.intro}}</div>
								</a-col>
							</a-row>
						</div>
					</a-col>
				</a-row>
			</a-col>
			<a-col :md="4" :xs="1" :sm="1">
				<a-back-top />
			</a-col>
		</a-row>
	

	</div>
</template>

<script>
	import {
		init,
		getCateArticle
	} from '@/api/init.js'
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
				avatar: "http://localhost:8080/static/images/avatar.jpg",
				drawer: false,
				active: '',
				menu_list: [],
				style: 'top',
				author: {

				},
				blog: {

				},
				show_breadcumb: true
			}
		},
		methods: {
			menuClick(item) {
				if (item.key == '/') {
					this.show_breadcumb = false
				} else {
					this.show_breadcumb = true
				}
				if(this.active != item.key) {
					this.active = item.key
					this.$router.push({
						path: item.key
					})
				} 
				
				
			},
			ondrawerClose() {
				this.drawer = false
			},
			init() {
				init().then(r => {
					if (r.code != 0) {
						this.$alert(r.msg)
						return;
					}
					this.author = r.data.author
					this.blog = r.data.blog_info
				}).catch(e => {
					console.log(e)
				})
			},
			getCateArticle() {
				getCateArticle().then(r => {
					if (r.code != 0) {
						console.log(r)
					} else {
						var data = r.data
						var all_menu = {
							key:'/article',
							name:''
						}
						var total = 0
						for (var p in data) {
							this.menu_list.push({
								key: '/article/cate_id/' + data[p]['id'],
								name: data[p]['cate_name'] + '(' + data[p]['num'] + ')'
							})
							total = total + parseInt(data[p]['num'])
						}
						all_menu.name = this.$i18n.t('menu.all_article') + '(' + total + ')'
						this.menu_list.unshift(all_menu)
					
					}
				}).catch(e => {
					console.log(e)
				})
			},
			show() {
				if (this.$route.path == "" || this.$route.path == '/') {
					this.show_breadcumb = false
				} else {
					this.show_breadcumb = true
				}
			},
		},
		computed: {
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
		color: #FFF8E6;
		line-height: 40px;
	}

	.phone-title-botton {
		background-color: rgba(256, 256, 256, 0);
		color: #FFFFFF;
		border: none;
		display: inline-block;
		float: left;
		line-height: 40px;
	}


	.phone-author-name {
		font-size: 20px;
	}

	.pp-menu {
		background: #F0C78A;
		height: 50px;
		color: #000000;
		line-height: 50px;
		font-size: 17px;
		display: inline-block;
		float: left;
		i{
			font-size: 17px;
		}
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
		-webkit-animation: mymove 5s infinite;
		/* Chrome, Safari, Opera */
		animation: mymove 5s infinite;
	}

	/* Chrome, Safari, Opera */
	@-webkit-keyframes mymove {
		50% {
			text-color: black;
		}
	}

	/* Standard syntax */
	@keyframes mymove {
		50% {
			color: blue;
		}
	}
</style>
