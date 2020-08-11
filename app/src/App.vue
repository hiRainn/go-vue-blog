<template>
	<a-config-provider :locale="locale">
		<div id="app">
			<router-view />
		</div>
	</a-config-provider>
</template>

<script>
	import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
	import enUS from 'ant-design-vue/lib/locale-provider/en_US'
	import ConfigProvider from 'ant-design-vue';
	export default {
		name: 'App',
		provide() { //父组件中通过provide来提供变量，在子组件中通过inject来注入变量。                                             
			return {
				reload: this.reload
			}
		},
		components:{
			ConfigProvider
		},
		data() {
			return {
				locale:enUS,
				isRouterAlive: true //控制视图是否显示的变量
			}
		},
		methods: {
			reload() {
				this.isRouterAlive = false; //先关闭，
				this.$nextTick(function() {
					this.isRouterAlive = true; //再打开
				})
			},
		},
		mounted(){
			var language = localStorage.getItem('locate')?localStorage.getItem('locate'):'en_US';
			if( language == 'en_US') {
				this.locale = enUS;
			} else if (language == 'zh_CN') {
				this.locale = zhCN;
			}
		}
	}
</script>

<style>
	#app {
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: center;
		color: #2c3e50;
	}
	body{
		margin: 0px;
		padding: 0px;
	}
</style>
