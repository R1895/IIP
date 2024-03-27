<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="异常ID">
            <el-input v-model="searchBox.anomaly_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item label="机器ID">
            <el-input v-model="searchBox.machine_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加异常</el-button>
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
    <el-table-column align="center" label="异常ID">
        <template slot-scope="scope">
          {{ scope.row.anomaly_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="异常机器ID">
        <template slot-scope="scope">
          {{ scope.row.machine_id }}
        </template>
      </el-table-column>
       <el-table-column label="时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.time}}
        </template>
      </el-table-column>
      <el-table-column label="机器异常状态码" align="center">
        <template slot-scope="scope">
          {{ scope.row.machine_status_id}}
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
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.anomaly_id)" circle></el-button>
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
    <el-dialog width="30%" title="上报异常" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="机器ID" >
                <el-input v-model="newCol.machine_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="异常类型" >
                <el-input v-model="newCol.machine_status_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="开始时间" >
                <el-input v-model="newCol.time" autocomplete="off"></el-input>
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
    <el-dialog width="30%" title="更新异常信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="异常ID" >
                <el-input v-model="updateColData.anomaly_id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="时间" >
                <el-input v-model="updateColData.time" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="机器ID" >
                <el-input v-model="updateColData.machine_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="机器异常状态码" >
                <el-input v-model="updateColData.machine_status_id" autocomplete="off"></el-input>
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
import { getMachineAnomalyManagerCols, addMachineAnomalyManagerCol, updateMachineAnomalyManagerCol, deleteMachineAnomalyManagerCol} from '@/api/industry'

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
        machine_id:"",
        machine_status_id:"",
        time:"",
        description:""
      },
      updateColData:{
        anomaly_id: "",
        time: "",
        machine_id: "",
        machine_status_id: "",
        description: ""
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
            let res = await deleteMachineAnomalyManagerCol({id:id});
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
        sendParams.anomaly_id = this.searchBox.anomaly_id
        console.log(this.searchBox.anomaly_id)
        let res = await getMachineAnomalyManagerCols({id:this.searchBox.anomaly_id})
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
        this.$confirm('确定添加新异常吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            this.newCol.machine_id = parseInt(this.newCol.machine_id)
            this.newCol.machine_status_id = parseInt(this.newCol.machine_status_id)
            let res = await addMachineAnomalyManagerCol(data);
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
        this.$confirm('确定更新此异常信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            let sendP= {
                id:this.updateColData.anomaly_id,
                machine_id:this.updateColData.machine_id,
                machine_status_id:this.updateColData.machine_status_id,
                description:this.updateColData.description,
                time:this.updateColData.time,
            }
            console.log(sendP)
            let res = await updateMachineAnomalyManagerCol(sendP);
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
      getMachineAnomalyManagerCols({start:start,limit:limit}).then(response => {
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
