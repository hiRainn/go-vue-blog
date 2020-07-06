<template>
	<div class="app-container">
	  <div class="filter-container">
	    <el-select v-model="search1.user_id" filterable clearable class="filter-item" loading-text="正在加载数据" placeholder="选择后台人员">
	      <el-option class="option" v-for="item in user_list" :key="item.id" :label="item.name" :value="item.id">
	      </el-option>
	    </el-select>
	    <el-date-picker v-model="search1.start_at" type="date" placeholder="开始日期" format="yyyy-MM-dd" value-format="yyyy-MM-dd">
	    </el-date-picker>
	    <el-date-picker v-model="search1.end_at" type="date" placeholder="优惠结束日期" format="yyyy-MM-dd" value-format="yyyy-MM-dd">
	    </el-date-picker>
	    <el-button class="filter-item" type="primary" icon="el-icon-search" @click="searchList">搜索</el-button>
	  </div>
	  <br>
	  <el-table :data="list" border v-loading="loading">
	    <el-table-column prop="id" label="编号" width="50px">
	    </el-table-column>
	    <el-table-column label="标题">
	      <template slot-scope="scope">
	        <span>{{ scope.row.title }}</span>
	      </template>
	    </el-table-column>
	    <el-table-column label="操作内容">
	      <template slot-scope="scope">
	        <span>{{ scope.row.content }}</span>
	      </template>
	    </el-table-column>
	    <el-table-column label="操作时间">
	      <template slot-scope="scope">
	        <span>{{ scope.row.created_at }}</span>
	      </template>
	    </el-table-column>
	    <el-table-column label="操作前内容">
	      <template slot-scope="scope">
	        <span>{{ scope.row.old_data }}</span>
	      </template>
	    </el-table-column>
	    <el-table-column label="操作后内容">
	      <template slot-scope="scope">
	        <span>{{ scope.row.new_data }}</span>
	      </template>
	    </el-table-column>
	
	
	  </el-table>
	  <el-pagination background @current-change="changePage(p)" :current-page.sync="p" :page-size="max" layout="total, prev, pager, next"
	    :total="count">
	  </el-pagination>
	
	</div>
</template>

<script>
	import {checkInit} from '@/api/article.js'
	export default {
		name :'article',
		data () {
			return {
				list:[],
				p: 0,
				max: 50, //每页显示数量
				count: 0,
			}
		},
		methods:{
			getArticles() {
				getArticles().then( response => {
					if(response.code) {
						alert(response.msg)
					}
					else {
						this.list = response.data
						this.p = response.data.list.current_page;
						this.count = response.data.count;
					}
				}).catch(err => {
					console.log(err)
				}) 
			},
			dian() {
				alert(1)
			}
		},
		mounted() {
			this.getArticles()
		}
		
	}
</script>

<style>
</style>
