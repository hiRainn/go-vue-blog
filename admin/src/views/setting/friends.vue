<template>
	<div class="app-container">
		<el-dialog :visible.sync="addDisabled" width="30%" center>
			<el-form ref="form" :model="form" label-width="80px" label-position="left">
				<el-form-item label="avatar">
					<input type="text" v-model="form.id" hidden>
					<el-avatar shape="circle" :size="100" fit="fill" :src="form.avatar"></el-avatar>
					<el-upload :on-success="(resopnse,file,fileList) => {avatatSucc(resopnse,file,fileList)}" :headers="{'X-Token':token}"
					 name="image" :action="url + 'friend/avatar'" class="upload-demo" :limit="1">
						<el-button size="small" type="primary">{{$t('user.upload')}}</el-button>
					</el-upload>
				</el-form-item>
				<el-form-item label="name" prop="name">
					<el-input v-model="form.name"></el-input>
				</el-form-item>

				<el-form-item label="link" prop="link">
					<el-input v-model="form.link"></el-input>
				</el-form-item>
				
				<el-form-item label="sort" prop="sort">
					<el-input v-model="form['sort']" placeholder="0~255"></el-input>
				</el-form-item>

				<el-form-item>
					<el-button @click="addDisabled = false">{{$t('os.cancle')}}</el-button>
					<el-button type="primary" @click="addFriend()">{{$t('article.save')}}</el-button>
				</el-form-item>
			</el-form>
		
		</el-dialog>

		<el-button type="success" style="float: right;margin-bottom: 10px;" @click="addDisabled = true">{{$t('os.add')}}</el-button>
		<br>
		<el-table :data="list" border>
			<el-table-column prop="id" label="ID" width="80px">
			</el-table-column>
			<el-table-column label="avatar" width="170px">
				<template slot-scope="scope">
					<el-avatar shape="circle" :size="50" fit="fill" :src="scope.row.avatar"></el-avatar>
				</template>
			</el-table-column>
			<el-table-column prop="name" label="name" width="170px">
			</el-table-column>
			<el-table-column prop="link" label="link">
			</el-table-column>
			<el-table-column prop="sort" label="sort" width="85px">
			</el-table-column>
			<el-table-column :label="$t('os.operate')" width="220px">
				<template slot-scope="scope">
					<el-button type="primary" size="small" icon="el-icon-edit" @click="editFriend(scope.row)">
						{{$t('os.edit')}}
					</el-button>
					<el-button type="danger" size="small" icon="el-icon-delete" @click="del(scope.$index,scope.row.id)">{{$t('os.delete')}}</el-button>
				</template>
			</el-table-column>
		</el-table>
		<el-pagination background layout="total" :total="page.total">
		</el-pagination>

	</div>
</template>

<script>
	import {
		getFriendsLink,
		addFriend,
		delFriend
	} from '@/api/setting.js'

	export default {
		name: 'init',
		data() {
			return {
				url: process.env.VUE_APP_URL,
				token:'',
				addDisabled: false,
				page: {
					total: 0
				},
				list: [],
				form:{
					id:0,
					name:'',
					link:'',
					avatar:'',
					sort:''
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
			avatatSucc(r,file,list) {
				this.form.avatar =r.data.url
				this.$set(this.form,'avatar',r.data.url)
			},
			getFriendsLink(data) {
				getFriendsLink(data).then(response => {
					if (!response.code) {
						var data = response.data
						this.list = data
						this.page.total = data.length
					}
				}).catch(err => {
					console.log(err)
				})
			},
			addFriend() {
				addFriend(this.form).then(r => {
					if(r.code) {
						this.$alert(error, this.$i18n.t('os.tip'))
						return
					}
					this.$message({
						message: this.$i18n.t('os.success'),
						type: 'success'
					});
					
					var data = {
						id : r.data,
						name : this.form.name,
						link: this.form.link,
						avatar:this.form.avatar,
						sort:this.form['sort']
					}
					// add
					if(this.form.id == 0) {
						this.list.push(data)
						this.page.total++
					}else {
						for (let p in this.list) {
							if(this.list[p].id == this.form.id) {
								this.$set(this.list,p,data)
								break;
							}
						}
					}
					this.reFreshForm()
					this.addDisabled = false
					
				}).catch( e => {
					console.log(e)
				})
			},
			editFriend(row){
				this.form.id = row.id
				this.form.name = row.name
				this.form.link = row.link
				this.form.avatar = row.avatar
				this.form['sort'] = row['sort']
				this.addDisabled = true
			},
			reFreshForm() {
				this.form.id = 0
				this.form.name = ''
				this.form.link = ''
				this.form.avatar = ''
				this.form['sort'] = ''
			},
			del(index, id) {
				this.$confirm(this.$i18n.t('article.confirm_delete'), this.$i18n.t('os.tip'), {
					confirmButtonText: this.$i18n.t('os.confirm'),
					cancelButtonText: this.$i18n.t('os.cancle'),
					type: 'warning'
				}).then(() => {
					delFriend({
						id: id
					}).then(response => {
						this.$message({
							message: this.$i18n.t('os.success'),
							type: 'success'
						});
						this.list.splice(index, 1)
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
			this.getFriendsLink()
			this.token = this.$store.getters.token;
		}

	}
</script>

<style scoped="scoped">
	a {
		text-decoration: none;
		color: #FFFFFF;
	}
</style>
