<template>
  <div>
	  <div class="item" v-for="item in itemList">
		  <span class="title"><router-link :to="'/article/' + item.id" style="font-size: 24px;">{{item.title}}</router-link></span>
		  <el-divider></el-divider>
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
  computed: {
    itemList() {
      return this.list
    },
  }
}
</script>

<style scoped>
	.item{
		text-align: left;
	}
	a {
		text-decoration: none;
	}
	
	.router-link-active {
		text-decoration: none;
	}
</style>
