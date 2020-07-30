<template>
	<div class="app-container">
	 
	  <el-table :data="list" border >
	    <el-table-column prop="id" label="ID" width="80px">
	    </el-table-column>
	    <el-table-column :label="$t('article.title')" >
			<template slot-scope="scope">
				<span>{{scope.row.title}}</span>
			</template>
	    </el-table-column>
		<el-table-column prop="cate_name" :label="$t('article.cate_name')" width="150px">
		</el-table-column>
		<el-table-column prop="tags" :label="$t('article.tags')" width="150px">
		</el-table-column>
		<el-table-column prop="views" :label="$t('article.view')" width="150px">
		</el-table-column>
		<el-table-column prop="comments" :label="$t('article.comment')" width="150px">
		</el-table-column>
		<el-table-column prop="post_at" :label="$t('article.create_at')" width="150px">
		</el-table-column>
	    
	  
	    
	    
	    <el-table-column :label="$t('os.operate')" width="250px">
	      <template slot-scope="scope">
			  <el-button v-if="!scope.row.edit" type="primary" size="small" icon="el-icon-edit">
			  	<router-link :to="{path:'/article/add',query:{id:scope.row.id}}">{{$t('os.edit')}}</router-link>
			  </el-button>
	        <el-button v-if="!scope.row.edit" type="danger" size="small" icon="el-icon-delete" @click="del(scope.$index,scope.row.id)">{{$t('os.delete')}}</el-button>
	      </template>
	    </el-table-column>
	  </el-table>
	  <el-pagination background @current-change="searchList(page.p)" :current-page.sync="page.p" :page-size="page.max" layout="total, prev, pager, next"
	    :total="page.total">
	  </el-pagination>
	
	</div>
</template>

<script>

	import {getArticles,delArticle} from '@/api/article.js'
	export default {
		name :'init',
		data () {
			return {
				page:{
					p:1,
					page_size:20,
					total:0
				},
				list:[],
			}
		},
		methods:{
			getArticles(data) {
				getArticles(data).then( response => {
					if(!response.code) {
						var data = response.data
						this.list = data.list
						this.page.p = data.p
						this.page.total = data.total
						this.page.page_size = data.page_size
					}
				}).catch(err => {
					console.log(err)
				}) 
			},
			changPage(p) {
				var data = {
					p:p,
					page_size : this.page.page_size,
					status : 1
				}
				this.getArticles(data)
			},
			del(index,id) {
				this.$confirm(this.$i18n.t('article.confirm_delete'), this.$i18n.t('os.tip'), {
					confirmButtonText: this.$i18n.t('os.confirm'),
					cancelButtonText: this.$i18n.t('os.cancle'),
					type: 'warning'
				}).then(() => {
					delArticle({id: id}).then(response => {
						this.$message({
							message: this.$i18n.t('os.success'),
							type: 'success'
						});
						this.list.splice(index,1)
						this.page.total = this.page.total - 1
					}).catch(error => {
						this.$alert(error, this.$i18n.t('os.tip'))
					})
				}).catch(() => {
					this.$message({
						type: 'info',
						message: this.$i18n.t('os.cancle')
					});
				});
			}
			
		},
		mounted() {
			var data = {
				status:1
			}
			this.getArticles(data)
		}
		
	}
</script>

<style scoped="scoped">
	a {
	    text-decoration: none;
		color: #FFFFFF;
	}
</style>
