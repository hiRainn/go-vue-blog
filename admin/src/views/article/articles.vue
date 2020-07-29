<template>
	<div class="app-container">
		<div class="filter-container">
			<el-input :placeholder="$t('article.title')" v-model="search.title" style="width: 200px;" class="filter-item" />
			<el-select v-model="search.cate_id" filterable clearable class="filter-item" loading-text="loading" :placeholder="$t('article.cate')">
				<el-option class="option" v-for="item in cate_list" :key="item.id" :label="item.name" :value="item.id">
				</el-option>
			</el-select>
			<el-select v-model="search.tag_id" multiple filterable clearable class="filter-item" loading-text="loading"
			 :placeholder="$t('article.tags')">
				<el-option class="option" v-for="item in tags_list" :key="item.id" :label="item.name" :value="item.id">
				</el-option>
			</el-select>
			<el-button class="filter-item" type="primary" icon="el-icon-search" @click="searchList()">{{$t('os.search')}}</el-button>
		</div>
		<br>

		<router-link to="/article/add">
			<el-button type="success" style="float: right;margin-bottom: 10px;">{{$t('article.post')}}</el-button>
		</router-link>
		<br>
		<el-table :data="list" border>
			<el-table-column prop="id" label="ID" width="80px">
			</el-table-column>
			<el-table-column :label="$t('article.title')">
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
		<el-pagination background @current-change="changPage(page.p)" :current-page.sync="page.p" :page-size="page.page_size"
		 layout="total, prev, pager, next" :total="page.total">
		</el-pagination>

	</div>
</template>

<script>
	import {
		getSelectCateList
	} from '@/api/cate.js'
	import {
		getSelectTagsList
	} from '@/api/tags.js'
	import {
		getArticles,
		delArticle
	} from '@/api/article.js'
	export default {
		name: 'init',
		data() {
			return {
				page: {
					p: 1,
					page_size: 20,
					total: 0
				},
				list: [],
				cate_list: [],
				tags_list: [],
				search: {
					cate_id: '',
					tag_id: '',
					title: '',
					p: ''
				},
				search_data: {
					cate_id: '',
					tag_id: '',
					title: '',
					p: ''
				},
				initForm: {

				}
			}
		},
		methods: {
			refreshView() {
			  // In order to make the cached page re-rendered
			  this.$store.dispatch('delAllCachedViews', this.$route)
			
			  const {
			    fullPath
			  } = this.$route
			
			  this.$nextTick(() => {
			    this.$router.replace({
			      path: fullPath
			    })
			  })
			},
			getArticles(data) {
				getArticles(data).then(response => {
					if (!response.code) {
						var data = response.data
						this.list = data.list
						this.page.p = data.p
						this.page.total = data.total
					}
				}).catch(err => {
					console.log(err)
				})
			},
			getSelectCateList() {
				getSelectCateList().then(response => {
					if (!response.code) {
						var data = response.data
						this.cate_list = data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			getSelectTagsList() {
				getSelectTagsList().then(response => {
					if (!response.code) {
						this.tags_list = response.data.list
					}
				}).catch(err => {
					console.log(err)
				})
			},
			searchList() {
				this.page.p = 1
				for (var p in this.search) {
					this.search_data[p] = this.search[p]
				}
				this.search_data["p"] = 1
				this.search_data["page_size"] = this.page.page_size
				this.getArticles(this.search_data)
			},
			changPage(p) {
				this.search_data["p"] = p
				this.search_data["page_size"] = this.page.page_size
				this.getArticles(this.search_data)
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
			this.getArticles({
				page_size: this.page.page_size
			})
			this.getSelectCateList()
			this.getSelectTagsList()
		}

	}
</script>

<style scoped="scoped">
	a {
	    text-decoration: none;
		color: #FFFFFF;
	}
</style>
