<template>
	<div id="container" style="margin-top: 0px;padding-top: 50px;background: #fefbf1;">
		<!-- menu -->
		<a-row class="fix">
			<!-- menu for phone -->
			<a-row>
				<a-col :xs="24" :md="0">
					<!-- <a-affix :offsetTop="0" :style="{ position: 'absolute', top: 0, left: 0}"> -->

					<header class="phone-header" style="display: block;">
						<a-button @click="drawer = true" class="phone-title-botton">
							<a-icon type="menu-fold" v-if="drawer" />
							<a-icon type="menu-unfold" v-if="!drawer" />
						</a-button>
						<span class="phone-title">{{ blog.title}}</span>

						<a-drawer :maskClosable="true" :bodyStyle="{padding:'24px 0px'}" width="200" placement="left" :closable="false"
						 :visible.async="drawer" @close="ondrawerClose">
							<a-menu mode="vertical" @click="menuClick" :selectable="false">
								<a-menu-item key="/home" :disabled="$route.path == '/home'">
									<a-icon type="home" />
									{{$t('menu.index')}}
								</a-menu-item>
								<a-sub-menu key="article">
									<span slot="title">
										<a-icon type="book" />
										<span>{{$t('menu.article')}}</span>
									</span>
									<a-menu-item v-for="item in menu_list" :key="item.key" :disabled="$route.path == item.key">
										{{item.name}}
									</a-menu-item>
								</a-sub-menu>
								<a-menu-item key="/filing" :disabled="$route.path == '/filing'">
									<a-icon type="database" />
									{{$t('menu.filing')}}
								</a-menu-item>
								<a-menu-item key="/message" :disabled="$route.path == '/message'">
									<a-icon type="form" />
									{{$t('menu.msg')}}
								</a-menu-item>
								<a-menu-item key="/about" :disabled="$route.path == '/about'">
									<a-icon type="smile" />
									{{$t('menu.about')}}
								</a-menu-item>
							</a-menu>
						</a-drawer>
					</header>

					<!-- </a-affix> -->
				</a-col>
			</a-row>


			<!-- menu for pad && pc -->
			<a-row>
				<!-- <a-affix :offsetTop="0" > -->
				<a-col class="pp-menu" :xs="0" :md="24">
					<a-col :md="{offset:1,span:22}" :lg="{offset:3,span:14}" :xs="0">
						<a-menu @click="menuClick" :selectable="false" class="pp-menu" mode="horizontal">
							<a-menu-item key="/home" :disabled=" $route.path == '/home'">
								<a-icon type="home" />
								{{$t('menu.index')}}
							</a-menu-item>
							<a-sub-menu key="article">
								<span slot="title">
									<a-icon type="book" />
									<span>{{$t('menu.article')}}</span>
								</span>
								<a-menu-item v-for="item in menu_list" :key="item.key" :disabled="$route.path == item.key">
									{{item.name}}
								</a-menu-item>
							</a-sub-menu>
							<a-menu-item key="/filing" :disabled="$route.path == '/filing'">
								<a-icon type="database" />
								{{$t('menu.filing')}}
							</a-menu-item>
							<a-menu-item key="/message" :disabled="$route.path == '/message'">
								<a-icon type="form" />
								{{$t('menu.msg')}}
							</a-menu-item>
							<a-menu-item key="/about" :disabled="$route.path == '/about'">
								<a-icon type="smile" />
								{{$t('menu.about')}}
							</a-menu-item>
						</a-menu>
					</a-col>
					<a-col :md="1" :lg="0" :xs="0"></a-col>
					<a-col :md="0" :lg="{offset:2,span:4}" :xs="0">
					<!-- 	<video id="video" src="../static/audios/yhys.mp3" muted="true" onpagehide="auto" controls="controls" autoplay="autoplay"
						 loop="loop" style="display: inline-block;height: 46px;width: 100%;margin:2px;line-height: 50px;">
						</video> -->
					</a-col>
				</a-col>
				<!-- </a-affix> -->
			</a-row>
		</a-row>

		<!-- banner -->
		<!-- <a-row>
			<a-col :xs="0" :xl="24"> 
				<div id="banner" :style="{width:getClientWidth}">
				</div>
			</a-col>
		</a-row> -->

		<!-- phone header -->
		<a-row v-if="this.$route.path == '/home'">
			<a-col :xs="24" :md="0">
				<a-row>
					<a-col :offset="5" :span="14">
						<a-avatar :size="100" :src="author.avatar" style="margin-top: 20px;"></a-avatar>
						<div class="phone-author-name"><b>{{author.nickname}}</b></div>
					</a-col>
					<a-col :span="5"></a-col>
				</a-row>
			</a-col>
		</a-row>


		<!-- body -->
		<a-row style="margin-top:20px">
			<a-col :xl="{span:16,offset:4}" :xs="{span:22,offset:1}">
				<a-row>
					<!-- content -->
					<a-col :md="18" :xs="24">
						<!-- <a-row :hidden="!Cshow_breadcumb" class="breadcrumb">
							<a-breadcrumb separator="/" style="float: left;margin: 5px;font-size: 14px ;">
								<a-breadcrumb-item>
									<template>
										<router-link to="/">
											<span @click="show_breadcumb = false">
												<a-icon type="home" />{{$t('menu.index')}}
											</span>
										</router-link>
									</template>
								</a-breadcrumb-item>
							</a-breadcrumb>
						</a-row> -->
						<a-row>
							<app-main />
						</a-row>
					</a-col>
					<!-- right area -->
					<a-col :md="6" :xs="0">
						<!-- info -->
						<a-row class="side-info">
							<a-row type="flex" class="side-info-item" justify="center">
								<a-avatar :size="100" :src="author.avatar" style="margin-top: 20px;"></a-avatar>
							</a-row>
							<a-row type="flex" justify="center" align="top" class="side-info-item">
								<a-col><b style="font-size: 20px;">{{author.nickname}}</b></a-col>
								<a-col v-if="author.gender == 1">
									<a-icon type="man" :style="{ fontSize: '16px', color: '#4b97f3','margin':'4px' }" />
								</a-col>
								<a-col v-if="author.gender == 2">
									<a-icon type="woman" :style="{ fontSize: '16px', color: 'pink','margin':'4px'  }" />
								</a-col>

							</a-row type="flex" justify="center" class="side-info-item">
							<a-row class="side-info-item" type="flex" justify="start">
								<span> {{author.sign}}</span>
							</a-row>

							<a-row v-if="author.city" type="flex" justify="start" class="">
								<a-icon type="environment" :style="{ 'margin': '3px', color:'#25292e'  }" />{{author.city}}
							</a-row>
							<a-row v-if="author.email" type="flex" justify="start" class="side-info-item">
								<a :href="'mailto:' + author.email" style="color: #25292e;">
									<a-icon type="mail" :style="{ 'margin': '3px', color:'#25292e'  }" />{{author.email}}
								</a>
							</a-row>
							<a-row class="side-info-item" type="flex" justify="start">
								<a-col class='icon' :span="3" v-if="author.github">
									<a :href="getUrl(author.github)" target="_blank">
										<icon-font type="icon-github" :style="{ fontSize: '24px', color:'#25292e'  }" />
									</a>
								</a-col>
								<a-col class='icon' :span="3" v-if="author.facebook">
									<a :href="getUrl(author.facebook)" target="_blank">
										<icon-font type="icon-facebook" :style="{ fontSize: '24px', color:'#317aea'   }" />
									</a>
								</a-col>
								<a-col class='icon' :span="3" v-if="author.twitter">
									<a :href="getUrl(author.twitter)" target="_blank">
										<icon-font type="icon-twitter" :style="{ fontSize: '24px', color:'#48a1eb'  }" />
									</a>
								</a-col>
								<a-col class='icon' :span="3" v-if="author.weibo">
									<a :href="getUrl(author.weibo)" target="_blank">
										<icon-font type="icon-weibo" :style="{ fontSize: '24px', color:'#d43437'  }" />
									</a>
								</a-col>
							</a-row>
						</a-row>

						<a-row class="side-info-center">
							<template v-if="statHide">
								<a-skeleton active />
							</template>
							<a-row type="flex" v-if="!statHide">
								<a-col :span="6">
									<a-statistic :valueStyle="{'font-size':'14px'}" :title="$t('stat.article')" :value="stat.article_num" />
								</a-col>
								<a-col :span="6">
									<a-statistic :valueStyle="{'font-size':'14px'}" :title="$t('stat.like')" :value="stat.like_number" />
								</a-col>
								<a-col :span="6">
									<a-statistic :valueStyle="{'font-size':'14px'}" :title="$t('stat.comments')" :value="stat.comments_number" />
								</a-col>
								<a-col :span="6">
									<a-statistic :valueStyle="{'font-size':'14px'}" :title="$t('stat.message')" :value="stat.msg_number" />
								</a-col>
							</a-row>
							<a-row v-if="!statHide">
								<span style="font-size: 14px;">
									{{getDay}}
								</span>

							</a-row>

						</a-row>

						<a-row class="side-info" v-if="article_list.length > 0">
							<a-row>
								<b style="font-size: 19px;">{{$t('menu.click_most')}}</b>
							</a-row>
							<a-row v-for="(item,index) in article_list" v-if="item.num > 0">
								<router-link style="font-size: 15px;color: #2C3E50;" :to="'/article/' + item.id">{{item.title}}({{item.num}})</router-link>
							</a-row>

						</a-row>

						<a-row class="side-info" v-if="comments.length > 0">
							<a-row>
								<b style="font-size: 19px;">{{$t('menu.last_comments')}}</b>
							</a-row>
							<a-row v-for="(item,index) in comments">
								<a-col>
									<a-popover placement="left">
										<template slot="content">
											{{item.content}}
										</template>
										<a style="font-size: 15px;" v-if="$route.path.indexOf('/article/') > -1" :href="'/article/' + item.article_id + '#PcAnchor' + item.id">{{item.name}}</a>
										<router-link v-else style="font-size: 15px;" :to="'/article/' + item.article_id + '#PcAnchor' + item.id">{{item.name}}</router-link>
									</a-popover>
									<span>
										- <router-link style="font-size: 15px;color: #2C3E50;" :to="'/article/' + item.article_id">{{item.article_title}}</router-link>
									</span>
								</a-col>
							</a-row>
						</a-row>

						<a-row class="side-info" v-if="tags.length > 0">
							<a-row>
								<b style="font-size: 19px;">{{$t('menu.tags')}}</b>
							</a-row>
							<a-row>
								<router-link :to="'/article/tag/' + item.name" v-for="(item,index) in tags">
									<a-tag @click="clickTag" :color="colorArray[Math.round(Math.random()*10)]" v-if="item.num > 0" style="margin-top: 8px;">
										{{item.name}}({{item.num}})
									</a-tag>
								</router-link>

							</a-row>
						</a-row>


						<a-row class="side-info" v-if="links.length > 0">
							<a-row>
								<b style="font-size: 19px;">{{$t('menu.friends')}}</b>
							</a-row>
							<a-row>
								<span style="font-size: 15px;margin-right: 20px;" class="friend-link" v-for="(item,index) in links">
									<a style="font-size: 15px;color: #2C3E50;" target="_blank" :href="item.link">{{item.name}}</a>
								</span>

							</a-row>
						</a-row>

						<a-row class="side-info">
							<template v-if="statHide">
								<a-skeleton active />
							</template>
							<a-row>
								<b style="font-size: 19px;">{{$t('os.site')}}</b>
							</a-row>
							<a-row class="stat" v-if="!statHide">
								<span>{{$t('stat.today')}}:</span><span>{{stat.today}}</span>
							</a-row>
							<a-row class="stat" v-if="!statHide">
								<span>{{$t('stat.week')}}:</span><span>{{stat.week}}</span>
							</a-row>
							<a-row class="stat" v-if="!statHide">
								<span>{{$t('stat.month')}}:</span><span>{{stat.month}}</span>
							</a-row>
							<a-row class="stat" v-if="!statHide">
								<span>{{$t('stat.year')}}:</span><span>{{stat.year}}</span>
							</a-row>
							<a-row class="stat" v-if="!statHide">
								<span>{{$t('stat.total')}}:</span><span>{{stat.total}}</span>
							</a-row>
						</a-row>
					</a-col>
				</a-row>
			</a-col>
			<a-col :md="4" :xs="1" :sm="1">
				<a-back-top />
			</a-col>
		</a-row>
		<!-- footer -->
		<a-row>
			<footer style="margin-top: 20px;">
				<a-row>
					<span>
						Copyright © 2020 <a href="http://sorahei.com">嘿嘿</a>
					</span>
					<span style="margin-left: 3px;">
						<a href="https://github.com/hirainn" target="_blank">
							<icon-font type="icon-github" :style="{color:'#25292e'  }" />
						</a>
					</span>
					<span style="margin-left: 3px;">
						<a href="https://twitter.com/HhheyW">
							<icon-font type="icon-twitter" :style="{color:'#25292e'  }" />
						</a>
					</span>
				</a-row>
				<a-row>
					<span>
						<a href="/licence">Licence</a>
					</span>
					<span style="margin-left: 3px;">
						www.sorahei.com All rights reserved.
					</span>

				</a-row>
			</footer>
		</a-row>


	</div>
</template>

<script>
	import $ from 'jquery'
	import {
		Icon
	} from 'ant-design-vue';

	const IconFont = Icon.createFromIconfontCN({
		scriptUrl: '//at.alicdn.com/t/font_8d5l8fzk5b87iudi.js',
	});
	import {
		init,
		getCateArticle,
		getFriendsLink,
		getLatestComment,
		getTags,
		getStat,
		getClickMost
	} from '@/api/init.js'
	import {
		AppMain
	} from './components'
	export default {
		name: 'Layout',
		components: {
			AppMain,
			IconFont
		},
		data() {
			return {
				statHide: true,
				tags: [],
				avatar: "http://localhost:8080/static/images/avatar.jpg",
				drawer: false,
				active: '',
				links: [],
				comments: [],
				menu_list: [],
				article_list: [],
				stat: [],
				style: 'top',
				author: {

				},
				blog: {

				},
				stat: {

				},
				show_breadcumb: true,
				colorArray: ['pink', 'red', 'orange', 'green', 'cyan', 'blue', 'purple', '#f50', '#2db7f5', '#87d068', '#108ee9']
			}
		},
		methods: {
			clickTag() {
				var target_offset = $("#container").offset();
				var target_top = target_offset.top;
				//goto that anchor by setting the body scroll top to anchor top
				$('html, body').animate({
					scrollTop: target_top
				}, 500);
			},
			menuClick(item) {
				if (item.key == '/home') {
					this.show_breadcumb = false
				} else {
					this.show_breadcumb = true
				}
				var path = this.$router.path
				if (path != this.active) {
					this.active = item.key
					this.$router.push({
						path: item.key
					})
					this.drawer = false
					var target_offset = $("#container").offset();
					var target_top = target_offset.top;
					//goto that anchor by setting the body scroll top to anchor top
					$('html, body').animate({
						scrollTop: target_top
					}, 500);
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
					if (r.data.author.avatar.indexOf('http') < 0 ) {
						r.data.author.avatar = 'http://' + r.data.author.avatar
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
							key: '/article',
							name: ''
						}
						var total = 0
						for (var p in data) {
							if(data[p]['num'] != 0) {
								this.menu_list.push({
									key: '/article/cate_id/' + data[p]['id'],
									name: data[p]['cate_name'] + '(' + data[p]['num'] + ')'
								})
								total = total + parseInt(data[p]['num'])
							} else {
								delete(data[p])
							}
						}
						all_menu.name = this.$i18n.t('menu.all_article') + '(' + total + ')'
						this.menu_list.unshift(all_menu)

					}
				}).catch(e => {
					console.log(e)
				})
			},
			getLatestComment() {
				getLatestComment().then(r => {
					if (r.code) {
						console.log(r)
						return
					}
					for (let p in r.data) {
						if (r.data[p].name == '') {
							r.data[p].name = this.$i18n.t('comment.anonymous')
						}
					}
					this.comments = r.data
				}).catch(e => {
					console.log(e)
				})
			},
			getFriendsLink() {
				getFriendsLink().then(r => {
					if (r.code) {
						console.log(r)
						return
					}
					for (let p in r.data) {
						if (typeof(r.data[p]['link']) == 'string') {
							if (r.data[p]['link'].indexOf('http') < 0) {
								r.data[p]['link'] = 'http://' + r.data[p]['link']
							}
						}
					}
					this.links = r.data
				}).catch(e => {
					console.log(e)
				})
			},
			getClickMost() {
				getClickMost().then(r => {
					if (r.code) {
						console.log(r)
						return
					}
					this.article_list = r.data
				}).catch(e => {
					console.log(e)
				})
			},
			getTags() {
				getTags().then(r => {
					if (r.code) {
						console.log(r)
						return
					}
					this.tags = r.data
				}).catch(e => {
					console.log(e)
				})
			},
			getStat() {
				getStat().then(r => {
					if (r.code) {
						console.log(r)
						return
					}
					this.stat = r.data
					this.statHide = false
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
			},
			getUrl() {
				return function(url) {
					if (typeof(url) == 'string') {
						if (url.indexOf('http') < 0) {
							return 'https://' + url
						}
					}
				}
			},
			getDay() {
				var template = this.$i18n.t('stat.first_day')
				var day = this.stat.first_day
				var time = parseInt(new Date().getTime() / 1000)
				var data = Math.ceil((time - day) / 86400)
				return template.replace('{Number}', data)
			},
			getClientWidth() {
				return document.body.clientWidth + 'px'

			}
		},
		created() {
			this.show()
			this.init()
			this.getCateArticle()
			this.getFriendsLink()
			this.getLatestComment()
			this.getTags()
			this.getClickMost()
			this.getStat()
		},
		mounted() {
			document.body.style.background = '#fefbf1'
		}
	}
</script>

<style lang="scss" scoped>
	.fix {
		position: fixed;
		left: 0;
		right: 0;
		top: 0;
		z-index: 999;
	}

	#banner {
		background: url('/static/img/banner.jpg') no-repeat;
		height: 365px;
		margin-top: -50px;
		display: inline-block;
		background-position: center;
	}


	.stat>span {
		font-size: 16px;
	}



	.phone-header {
		height: 50px;
		border: 1px solid black;
		background-color: #001528;
		text-align: left;
		width: 100%;
		display: block;
	}

	.phone-title {
		color: #FFF8E6;
		line-height: 50px;
		font-size: 18px;
	}

	.phone-title-botton {
		background-color: #001528;
		background-color: rgba(256, 256, 256, 0);
		color: #FFFFFF;
		border: none;
		display: inline-block;
		float: left;
		line-height: 50px;
	}


	.phone-author-name {
		font-size: 20px;
	}

	.pp-menu {
		background: #F0C78A;
		// background: rgba(0,0,0,0.5);
		height: 50px;
		color: #000000;
		line-height: 50px;
		font-size: 17px;
		display: inline-block;
		float: left;

		i {
			font-size: 17px;
		}
	}

	a {
		text-decoration: none;
	}

	.side-info {
		background: #fff;
		margin: 20px 10px;
		border: 1px solid #eee;
		border-radius: 5px;
		padding: 10px;
		text-align: left;
		// vertical-align:middle;
		// align-items: center;
	}

	.side-info-center {
		background: #fff;
		margin: 20px 10px;
		border: 1px solid #eee;
		border-radius: 5px;
		padding: 10px;
		font-size: 18px;
	}

	.side-info-item {
		margin-bottom: 10px;
		word-wrap: break-word;

		span {
			word-wrap: break-word;
		}

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
