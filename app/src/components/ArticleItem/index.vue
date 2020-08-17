<template>
  <div>
	  <div class="item" v-for="item in list">
		  <a-row>
			  <span class="title">
					<router-link style="font-size: 22px;color: #303133;" :to="'/article/' + item.id" >{{item.title}}</router-link>
			  </span>
		  </a-row>
		  <a-row>
			  <span class="info">{{$t('article.publish')}} {{item.post_at}}</span>
			  <span class="info">
				  <router-link style="color: #2C3E50;" :to="'/article/cate_id/' + item.cate_id" >{{item.cate_name}}</router-link>
			  </span>
			  <a-icon type="eye" />{{item.views}}
			  <a-icon type="message" />{{item.comments}}
		  </a-row>
		  <a-row class="content">
			  {{item.content.replace(/[`~!@#$%^&*()_\-+=<>?:"{}|,.\/;'\\[\]·~！@#￥%……&*（）——\-+={}|《》？：“”【】、；‘’，。、]+$/im,'')}}...<router-link style="color: #398aec;" :to="'/article/' + item.id ">[{{$t('article.more')}}]</router-link>
		  </a-row>
		  <a-row class="tag">
			  <a-tag type="info" class="tag" color="#606266" v-if="item.tag_name != false" v-for="(v,k) in item.tag_name.split(',')">
				  <router-link style="font-size: 12px;margin-right: 5px;padding: 0px;margin: 0px ;" :to="'/article/tag/' + v" >{{v}}</router-link>
			  </a-tag>
		  </a-row>
	  </div>
	  
  </div>
</template>

<script>

export default {
  name: 'ArticleItem',
  props: {
    list: {
      type: Array,
      required: true
    },
  },
  methods:{
	  
  },
  computed: {
   
	
  }
}
</script>

<style scoped>
	.item{
		background: #fff;
		border: 1px solid #eee;
		text-align: left;
		border-radius: 10px;
		padding: 10px;
		margin-bottom: 20px;
	}
	a {
		text-decoration: none;
	}
	.title{
		font-size: 24px;
	}
	.router-link-active {
		text-decoration: none;
	}
	.info{
		color:#606266;
		margin-right: 15px;
		font-size: 14px;
	}
	i{
		margin-right: 10px;
		font-size: 14px;
	}
	.content{
		margin-top: 15px;
	}
	.tag{
		margin-top: 10px;
		margin-right:5px;
		display: inline-block;
	}
	.el-divider--horizontal{
		margin:15px 0px
	}
</style>
