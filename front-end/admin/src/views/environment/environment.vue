<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="环境ID">
            <el-input v-model="searchBox.environment_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加环境</el-button>
        </el-form-item>
    </el-form>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
    <el-table-column align="center" label="环境ID">
        <template slot-scope="scope">
          {{ scope.row.environment_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="车间ID">
        <template slot-scope="scope">
          {{ scope.row.workshop_id }}
        </template>
      </el-table-column>
       <el-table-column label="时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.timestamp}}
        </template>
      </el-table-column>
      <el-table-column label="环境水平" align="center">
        <template slot-scope="scope">
          {{ scope.row.level}}
        </template>
      </el-table-column> 

      <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          {{ scope.row.description}}
        </template>
      </el-table-column> 

      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.environment_id)" circle></el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="block">
        <el-pagination
            layout="prev, pager, next"
            :total="total"
            :page-size="colsPerPage"
            :current-page="currentPage"
            @current-change="changePage"
            >
        </el-pagination>
    </div>
    <el-dialog width="30%" title="上报环境" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="车间ID" >
                <el-input v-model="newCol.workshop_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="环境水平" >
                <el-input v-model="newCol.level" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="描述" >
                <el-input v-model="newCol.description" autocomplete="off"></el-input>
            </el-form-item>

        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新环境信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="环境ID" >
                <el-input v-model="updateColData.environment_id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="车间ID" >
                <el-input v-model="updateColData.workshop_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="环境水平" >
                <el-input v-model="updateColData.level" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="时间" >
                <el-input v-model="updateColData.timestamp" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="描述" >
                <el-input v-model="updateColData.description" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="updateColVisible = false">取消</el-button>
            <el-button type="primary" @click="updateCol">确认</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import { getList } from '@/api/table'
import { getEnvironmentManagerCols, addEnvironmentManagerCol, updateEnvironmentManagerCol, deleteEnvironmentManagerCol} from '@/api/environment'

export default {
  filters: {
    statusFilter(status) {
      if(status){
        return 'success'
      } else {
        return 'gray'
      }
    }
  },
  data() {
    return {
      list: [
      ],
      listLoading: false,
      searchBox:{
        id:"",
      },
      addColVisible:false,
      newCol:{
        workshop_id:"",
        level:"",
        timestamp:"",
        description:""
      },
      updateColData:{
        id:"",
        workshop_id:"",
        level:"",
        timestamp:"",
        description:""
      },
      updateColVisible:false,
      currentPage:1,
      colsPerPage:10,
      total:0
    }
  },
  created() {
    this.fetchData(this.currentPage,this.colsPerPage)
  },
  methods: {
    changePage(newPage){
        //处理分页
        this.currentPage = newPage
        console.log(this.currentPage)
        this.fetchData(this.currentPage,this.colsPerPage)
    },
    deleteCol(id){
        this.$confirm('确定删除此行信息吗(删除后不可恢复)?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            console.log("confirm")
            let res = await deleteEnvironmentManagerCol({id:id});
            console.log(res)
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '删除成功'
                });
            }
            this.fetchData(this.currentPage,this.colsPerPage)
        }).catch((err)=>{console.log(err)})
    },
    async search(){
        if(this.$areAllValuesEmpty(this.searchBox)){
            this.fetchData(this.currentPage,this.colsPerPage)
            return
        }
        let sendParams = this.searchBox
        sendParams.id = this.searchBox.environment_id
        console.log(this.searchBox.environment_id)
        let res = await getEnvironmentManagerCols({id:this.searchBox.environment_id})
        this.listLoading = true
        if(res.data.status=="200"){
            this.list = []
            this.list.push(res.data.data)
            this.searchBox = {
                id:""
            }
             this.listLoading = false
             console.log(this.listLoading)
             this.total = 1
        } else {
            this.$message({
                type: 'warning',
                message: '未搜索到此信息'
            });
             this.listLoading = false
        }
        
    },
    addCol(){
        let data = this.newCol;
        this.$confirm('确定添加新环境吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addEnvironmentManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                workshop_id:"",
                level:"",
                timestamp:"",
                description:""
            },
            this.addColVisible = false
            this.fetchData(this.currentPage,this.colsPerPage)
        }).catch(()=>{})
    },
    updateColShow(data){
        let dataStr = JSON.stringify(data);
        data = JSON.parse(dataStr)
        this.updateColVisible = true;
        this.updateColData = data
    },
    updateCol(){
        // this.updateColData is already exist
        this.$confirm('确定更新此环境信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            let sendP= {
                id:this.updateColData.environment_id,
                workshop_id:this.updateColData.workshop_id,
                level:parseInt(this.updateColData.level),
                timestamp:this.updateColData.timestamp,
                description:this.updateColData.description
            }
            let res = await updateEnvironmentManagerCol(sendP);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '更新成功!'
                });
            }
            this.updateColVisible = false
            this.fetchData(this.currentPage,this.colsPerPage)
        }).catch(()=>{})
    },
    fetchData(start=1,limit=10) {
      this.listLoading = true
      getEnvironmentManagerCols({start:start,limit:limit}).then(response => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        this.list = response.data.data.item
        this.total = response.data.data.total
        this.listLoading = false
        })
    },

  }
}
</script>
<style>
.block{
    display:flex;
    justify-content: center;
}
</style>
