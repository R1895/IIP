<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="角色ID">
            <el-input v-model="searchBox.role_id" placeholder="角色ID"></el-input>
        </el-form-item>
        <el-form-item label="角色名">
            <el-input v-model="searchBox.role_name " placeholder="角色名"></el-input>
        </el-form-item>
        <el-form-item label="角色描述">
            <el-input v-model="searchBox.description" placeholder="角色描述"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加角色</el-button>
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
      <el-table-column align="center" label="角色ID">
        <template slot-scope="scope">
          {{ scope.row.role_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="角色名">
        <template slot-scope="scope">
          {{ scope.row.role_name }}
        </template>
      </el-table-column>
      <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          {{ scope.row.description  }}
        </template>
      </el-table-column>
        
      <el-table-column label="角色权限"  align="center">
        <template slot-scope="scope">
          {{scope.row.permissionstring}}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.role_id)" circle></el-button>
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
    <el-dialog width="30%" title="添加新角色" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="角色名" >
                <el-input v-model="newCol.role_name" autocomplete="off"></el-input>
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
    <el-dialog width="30%" title="更新用户信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="角色ID" >
                <el-input v-model="updateColData.role_name" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="角色名" >
                <el-input v-model="updateColData.role_name" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="描述" >
                <el-input v-model="updateColData.description" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="角色权限" >
                <el-input v-model="updateColData.permissionstring" autocomplete="off" disabled></el-input>
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
import { getRoleManagerCols, addRoleManagerCol, updateRoleManagerCol, deleteRoleManagerCol} from '@/api/user'
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
        role_id:"",
        role_name:"",
        description:"",
      },
      addColVisible:false,
      newCol:{
        role_name:"",
        description:"",
      },
      updateColData:{
        role_id:"",
        role_name:"",
        description:"",
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
            if(list[i].permissions==null||list[i].permissions==void 0){
                list[i].permissionstring = "暂无"
                continue;
            }
            for(let j=0;j<list[i].permissions.length;j++){
                if(j==list[i].permissions.length-1){
                    str+= `${list[i].permissions[j].permission_name}`
                } else {
                    str += `${list[i].permissions[j].permission_name}-`
                }
            }
            list[i].permissionstring = str
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
            let res = await deleteRoleManagerCol({id:id});
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
        sendParams.id = this.searchBox.role_id
        let res = await getRoleManagerCols(sendParams)
        this.listLoading = true
        if(res.data.status=="200"){
            this.list = []
            this.list.push(res.data.data)
            this.list = this.preprocessData(this.list)
            this.searchBox = {
                role_id:"",
                role_name:"",
                description:"",
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
        this.$confirm('确定添加新信息吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addRoleManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                user_name:"",
                email:"",
                telephone:"",
            }
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
        this.$confirm('确定更新此行信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            this.updateColData.id = this.updateColData.role_id
            let res = await updateRoleManagerCol(this.updateColData);
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
      getRoleManagerCols({start:start,limit:limit}).then(response => {
        this.list = response.data.data.item==null?[]:response.data.data.item
        this.total = response.data.data.total
        console.log(this.total)
        this.list = this.preprocessData(this.list)
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
