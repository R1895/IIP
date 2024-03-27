<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="用户ID">
            <el-input v-model="searchBox.user_id " placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item label="用户名">
            <el-input v-model="searchBox.user_name " placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item label="手机号">
            <el-input v-model="searchBox.telephone" placeholder="手机号"></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
            <el-input v-model="searchBox.email" placeholder="邮箱"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加用户</el-button>
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
     <el-table-column align="center" label="用户ID">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="用户名">
        <template slot-scope="scope">
          {{ scope.row.user_name }}
        </template>
      </el-table-column>
      <el-table-column label="邮箱" align="center">
        <template slot-scope="scope">
          {{ scope.row.email  }}
        </template>
      </el-table-column>
      <el-table-column label="手机号"  align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.telephone }}</span>
        </template>
      </el-table-column>
      <el-table-column label="上次登录时间"  align="center">
        <template slot-scope="scope">
          {{ scope.row.last_login }}
        </template>
      </el-table-column>
        
      <el-table-column label="角色"  align="center">
        <template slot-scope="scope">
          {{scope.row.rolestring}}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.is_active?'success':'gray'">
            <span v-if="scope.row.is_active">在线</span>
            <span v-else>离线</span>
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="注册时间"  align="center">
        <template slot-scope="scope">
          {{ scope.row.registration_date}}
        </template>
      </el-table-column>


      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.user_id)" circle></el-button>
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
    <el-dialog width="30%" title="添加新用户" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="用户名" >
                <el-input v-model="newCol.user_name" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="邮箱" >
                <el-input v-model="newCol.email" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="手机号" >
                <el-input v-model="newCol.telephone " autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新用户信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="用户id" >
                <el-input v-model="updateColData.user_id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="用户名" >
                <el-input v-model="updateColData.user_name" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="邮箱" >
                <el-input v-model="updateColData.email" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="手机号" >
                <el-input v-model="updateColData.telephone" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="上次登录时间" >
                <el-input v-model="updateColData.last_login" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="注册时间" >
                <el-input v-model="updateColData.registration_date" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="用户角色" >
                <el-input v-model="updateColData.rolestring" autocomplete="off" disabled></el-input>
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
import { getUserManagerCols, addUserManagerCol, updateUserManagerCol, deleteUserManagerCol} from '@/api/user'

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
        user_id:"",
        user_name:"",
        email:"",
        telephone:"",
      },
      addColVisible:false,
      newCol:{
        user_name:"",
        email:"",
        telephone:"",
      },
      updateColData:{
        user_name:"",
        email:"",
        telephone:"",
        rolestring:"",
        last_login: "",
        IsActive: true,
        registration_date:"",
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
            let res = await deleteUserManagerCol({id:id});
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
        sendParams.id = this.searchBox.user_id
        let res = await getUserManagerCols({id:this.searchBox.user_id})
        this.listLoading = true
        if(res.data.status=="200"){
            this.list = []
            this.list.push(res.data.data)
            this.list = this.preprocessData(this.list)
            this.searchBox = {
                user_id:"",
                user_name:"",
                email:"",
                telephone:"",
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
        this.$confirm('确定添加新用户吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addUserManagerCol(data);
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
        this.$confirm('确定更新此用户信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            this.updateColData.id = this.updateColData.user_id
            let res = await updateUserManagerCol(this.updateColData);
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
      getUserManagerCols({start:start,limit:limit}).then(response => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
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
