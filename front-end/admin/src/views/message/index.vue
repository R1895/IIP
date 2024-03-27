<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <!-- <el-form-item label="消息ID">
            <el-input v-model="searchBox.anomaly_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item label="机器ID">
            <el-input v-model="searchBox.message_id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item> -->
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加消息</el-button>
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
    <el-table-column align="center" label="消息ID">
        <template slot-scope="scope">
          {{ scope.row.message_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="消息机器ID">
        <template slot-scope="scope">
          {{ scope.row.device_id }}
        </template>
      </el-table-column>
       <el-table-column label="用户ID" align="center">
        <template slot-scope="scope">
          {{ scope.row.user_id}}
        </template>
      </el-table-column>

      <el-table-column label="已读" align="center">
        <template slot-scope="scope">
          {{ scope.row.is_read==true?"是":"否"}}
        </template>
      </el-table-column>

      <el-table-column label="消息标题" align="center">
        <template slot-scope="scope">
          {{ scope.row.title}}
        </template>
      </el-table-column>
        <el-table-column label="消息内容" align="center">
            <template slot-scope="scope">
            {{ scope.row.content}}
            </template>
        </el-table-column>
        <el-table-column label="消息类型" align="center">
            <template slot-scope="scope">
            {{ scope.row.type}}
            </template>
        </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.message_id)" circle></el-button>
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
    <el-dialog width="30%" title="添加消息" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="机器ID" >
                <el-input v-model="newCol.device_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="用户ID" >
                <el-input v-model="newCol.user_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="标题" >
                <el-input v-model="newCol.title" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="内容" >
                <el-input v-model="newCol.content" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="类型" >
                <el-input v-model="newCol.type" autocomplete="off"></el-input>
            </el-form-item>

        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新消息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="消息ID" >
                <el-input v-model="updateColData.message_id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="已读" >
                <el-switch v-model="updateColData.is_read" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
            </el-form-item>
            <el-form-item label="内容" >
                <el-input v-model="updateColData.content" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="类型" >
                <el-input v-model="updateColData.type" autocomplete="off"></el-input>
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
import { getMessageManagerCols, addMessageManagerCol, updateMessageManagerCol, deleteMessageManagerCol} from '@/api/others'

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
        message_id:"",
        device_id:"",
        user_id:"",
        content:"",
        title:"",
        type:""
      },
      updateColData:{
        message_id:"",
        device_id:"",
        is_read:0,
        user_id:"",
        content:"",
        title:"",
        type:""
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
            let res = await deleteMessageManagerCol({id:id});
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
        let res = await getMessageManagerCols({id:this.searchBox.anomaly_id})
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
        this.$confirm('确定添加新消息吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addMessageManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                message_id:"",
                device_id:"",
                user_id:"",
                content:"",
                title:"",
                type:""
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
        this.$confirm('确定更新此消息信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            let sendP= {
                id:this.updateColData.message_id,
                device_id:this.updateColData.device_id,
                user_id:this.updateColData.user_id,
                content:this.updateColData.content,
                title:this.updateColData.title,
                type:this.updateColData.type,
                is_read:this.updateColData.is_read==true?1:0
            }
            console.log(sendP)
            let res = await updateMessageManagerCol(sendP);
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
      getMessageManagerCols({start:start,limit:limit}).then(response => {
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
