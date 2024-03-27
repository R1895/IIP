<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="日志ID">
            <el-input v-model="searchBox.log_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加日志</el-button>
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
    <el-table-column align="center" label="日志ID">
        <template slot-scope="scope">
          {{ scope.row.log_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作用户ID">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
       <el-table-column label="操作时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.timestamp}}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center">
        <template slot-scope="scope">
          {{ scope.row.action==""?"暂无":scope.row.action}}
        </template>
      </el-table-column>
      <el-table-column label="方法" align="center">
        <template slot-scope="scope">
          {{scope.row.method==""?"暂无":scope.row.method}}
        </template>
      </el-table-column>
      <el-table-column label="详情" align="center">
        <template slot-scope="scope">
          {{scope.row.details==""?"暂无":scope.row.details}}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.log_id)" circle></el-button>
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
    <el-dialog width="30%" title="上报日志" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="日志操作" >
                <el-input v-model="newCol.action" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="日志方法" >
                <el-input v-model="newCol.method" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="操作用户ID" >
                <el-input v-model="newCol.user_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="操作日志详情" >
                <el-input v-model="newCol.details" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新日志信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="日志ID" >
                <el-input v-model="updateColData.log_id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="日志操作" >
                <el-input v-model="updateColData.action" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="日志方法" >
                <el-input v-model="updateColData.method" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="操作用户ID" >
                <el-input v-model="updateColData.user_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="操作日志详情" >
                <el-input v-model="updateColData.details" autocomplete="off"></el-input>
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
import { getLogManagerCols, addLogManagerCol, updateLogManagerCol, deleteLogManagerCol} from '@/api/others'

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
        machine_name:"",
      },
      addColVisible:false,
      newCol:{
        user_id: "",
        action: "",
        method: "",
        details: "",
      },
      updateColData:{
        log_id:"",
        user_id: "",
        action: "",
        method: "",
        details: "",
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
    preprocessData(list){
        for(let i=0;i<list.length;i++){
            let str = "";
            if(list[i].roles==null||list[i].roles==void 0){
                list[i].rolestring = "暂无"
                continue;
            }
            for(let j=0;j<list[i].roles.length;j++){
                if(j==list[i].roles.length-1){
                    str+= `${list[i].roles[j].role_name}`
                } else {
                    str += `${list[i].roles[j].role_name}-`
                }
            }
            list[i].rolestring = str
        }
        return list
    },
    deleteCol(id){
        this.$confirm('确定删除此行信息吗(删除后不可恢复)?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            console.log("confirm")
            let res = await deleteLogManagerCol({id:id});
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
        sendParams.log_id = this.searchBox.log_id
        console.log(this.searchBox.log_id)
        let res = await getLogManagerCols({id:this.searchBox.log_id})
        this.listLoading = true
        if(res.data.status=="200"){
            this.list = []
            this.list.push(res.data.data)
            this.searchBox = {
                id:"",
                machine_name:"",
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
        this.$confirm('确定添加新日志吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addLogManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                machine_id:"",
                machine_status_id:"",
                time:"",
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
        this.$confirm('确定更新此日志信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            let sendP= {
                id:this.updateColData.log_id,
                user_id: this.updateColData.user_id,
                action: this.updateColData.action,
                method: this.updateColData.method,
                details: this.updateColData.details,
                timestamp: this.updateColData.timestamp
            }
            console.log(sendP)
            let res = await updateLogManagerCol(sendP);
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
      getLogManagerCols({start:start,limit:limit}).then(response => {
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
