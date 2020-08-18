<template>
	<div style="margin-top: 15px;">
		<div v-if="filingHide">
			<a-skeleton :paragraph="{ rows: 5 }" />
		</div>
		<a-row v-if="!filingHide" class="content">
			<a-collapse>
				<a-collapse-panel :header="item.year + ' (' + item.data.length + ')'" v-for="(item,index) in list">
					<a-timeline>
						<a-timeline-item v-for="(v,i) in item.data">
							<span>
								<router-link :to="'/article/' + v.id">{{v.title}}  --  {{v.created_at}}</router-link>
							</span>
						</a-timeline-item>
					</a-timeline>
				</a-collapse-panel>
			</a-collapse>
		</a-row>
	</div>
</template>


<script>
	import {
		getAllArticles
	} from '@/api/article.js'
	export default {

		data() {
			return {
				filingHide: true,
				list: [],
			}
		},
		methods: {
			getAllArticles() {
				getAllArticles().then(r => {
					if (r.code) {
						console.log(r)
						return;
					}
					let arr = []
					let data = []
					for (let p in r.data) {
						let key = parseInt(r.data[p].created_at.slice(0, 4))
						if(isNaN(key)) {
							continue
						}
						if (arr[key] == undefined) {
							arr[key] = true
							data.push({year:key,data:[r.data[p]]})
						}
						else {
							for(let q in data) {
								if(data[q].year == key) {
									data[q].data.push(r.data[p])
									break
								}
							}
						}
					}
					this.list = data
					this.filingHide = false
				})
			}
		},
		created() {
			this.getAllArticles()
		}
	}
</script>

<style scoped="scoped">
	.content {
		text-align: left;
	}
	
	li.ant-timeline:last-child{
		margin-bottom: 0px !important;
	}
</style>
