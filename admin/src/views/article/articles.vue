<template>
	<div class="app-container">
	  <div class="filter-container">
	    <el-input :placeholder="$t('article.title')" v-model="search.title" style="width: 200px;" class="filter-item" />
		<el-select v-model="search.cate_id" filterable clearable class="filter-item" loading-text="loading" :placeholder="$t('article.cate')">
		  <el-option class="option" v-for="item in cate_list" :key="item.id" :label="item.name" :value="item.id">
		  </el-option>
		</el-select>
		<el-select v-model="search.tag_id" filterable clearable class="filter-item" loading-text="loading" :placeholder="$t('article.tags')">
		  <el-option class="option" v-for="item in tags_list" :key="item.id" :label="item.name" :value="item.id">
		  </el-option>
		</el-select>
	    <el-button class="filter-item" type="primary" icon="el-icon-search" @click="searchList(0)">{{$t('os.search')}}</el-button>
	  </div>
	  <br>
	  
	  <router-link to="/article/add"><el-button type="success" style="float: right;margin-bottom: 10px;">{{$t('article.post')}}</el-button></router-link>
	  <br>
	  <el-table :data="list" border >
	    <el-table-column prop="id" label="ID" width="80px">
	    </el-table-column>
	    <el-table-column :label="$t('article.title')" >
			<template slot-scope="scope">
				<span>{{scope.row.title}}</span>
			</template>
	    </el-table-column>
		<el-table-column prop="view" :label="$t('article.view')" width="150px">
		</el-table-column>
		<el-table-column prop="comment" :label="$t('article.comment')" width="150px">
		</el-table-column>
	    
	  
	    
	    
	    <el-table-column :label="$t('os.operate')" width="250px">
	      <template slot-scope="scope">
	        <el-button v-if="!scope.row.edit" type="primary" size="small" icon="el-icon-edit" @click="scope.row.edit = true">编辑</el-button>
	        <el-button v-if="scope.row.edit" type="success" :disabled="scope.row.disabled" size="small" icon="el-icon-circle-check-outline"
	          @click="confirmEdit(scope.row)">确定</el-button>
	        <el-button v-if="scope.row.edit" type="danger" size="small" icon="el-icon-circle-check-outline" @click="scope.row.edit = false">取消</el-button>
	        <el-button v-if="!scope.row.edit" type="danger" size="small" icon="el-icon-delete" @click="delcus(scope.row.id)">删除</el-button>
	      </template>
	    </el-table-column>
	  </el-table>
	  <el-pagination background @current-change="searchList(page.p)" :current-page.sync="page.p" :page-size="page.max" layout="total, prev, pager, next"
	    :total="page.total">
	  </el-pagination>
	
	</div>
</template>

<script>
	import {getSelectCateList} from '@/api/cate.js'
	import {getSelectTagsList} from '@/api/tags.js'
	import {getArticles} from '@/api/article.js'
	export default {
		name :'init',
		data () {
			return {
				page:{
					p:1,
					max:0,
					total:0
				},
				list:[],
				cate_list:[],
				tags_list:[],
				search:{
					cate_id:'',
					tag_id:'',
					title:'',
					p:''
				},
				search_data:{
					cate_id:'',
					tag_id:'',
					title:'',
					p:''
				},
				initForm:{
					
				}
			}
		},
		methods:{
			getArticles(data) {
				getArticles().then( response => {
					if(!response.code) {
						var data = response.data
						this.list = data.list
						this.page.p = data.p
						this.page.total = data.total
						this.page.max = data.page_num
					}
				}).catch(err => {
					console.log(err)
				}) 
			},
			getSelectCateList() {
				getSelectCateList().then( response => {
					if(!response.code) {
						this.cate_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				}) 
			},
			getSelectTagsList() {
				getSelectTagsList().then( response => {
					if(!response.code) {
						this.tags_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				}) 
			},
			searchList(p) {
				if(p != 0) {
					this.search.p = p
				}
				getArticles(this.search)
			}
		},
		mounted() {
			this.getArticles({})
			this.getSelectCateList()
			this.getSelectTagsList()
		}
		
	}
</script>

<style>
</style>
