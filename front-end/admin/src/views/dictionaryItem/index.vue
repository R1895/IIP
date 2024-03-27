<template>
  <div class="app-container">

    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="数据键值ID">
            <el-input v-model="searchBox.id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加数据字典</el-button>
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
    <el-table-column align="center" label="数据键值ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="数据键">
        <template slot-scope="scope">
          {{ scope.row.key }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="数据值">
        <template slot-scope="scope">
          {{ scope.row.value }}
        </template>
      </el-table-column>
       <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          {{ scope.row.description}}
        </template>
      </el-table-column>
      <el-table-column label="目录" align="center">
        <template slot-scope="scope">
          {{ scope.row.category}}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.id)" circle></el-button>
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

    <el-dialog width="30%" title="添加键值" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="数据键" >
                <el-input v-model="newCol.id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="数据值" >
                <el-input v-model="newCol.value" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="描述" >
                <el-input v-model="newCol.description" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="目录" >
                <el-input v-model="newCol.category" autocomplete="off"></el-input>
            </el-form-item>
            
        
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新数据" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="数据键值ID" >
                <el-input v-model="updateColData.id" autocomplete="off" disabled></el-input>
            </el-form-item>
            <el-form-item label="数据键" >
                <el-input v-model="updateColData.key" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="数据值" >
                <el-input v-model="updateColData.value" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="描述" >
                <el-input v-model="updateColData.description" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="目录" >
                <el-input v-model="updateColData.category" autocomplete="off"></el-input>
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
import { getDictionaryItemManagerCols, addDictionaryItemManagerCol, updateDictionaryItemManagerCol, deleteDictionaryItemManagerCol} from '@/api/others'

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
        key:"",
        value:"",
      },
      addColVisible:false,
      newCol:{
        id:"",
        key:"",
        value:"",
        category:"",
        description:""
      },
      updateColData:{
        id:"",
        key:"",
        value:"",
        category:"",
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
            let res = await deleteDictionaryItemManagerCol({id:id});
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
        sendParams.id = this.searchBox.id
        console.log(this.searchBox.id)
        let res = await getDictionaryItemManagerCols({id:this.searchBox.id})
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
            let res = await addDictionaryItemManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                key:"",
                value:"",
                category:"",
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
                id:this.updateColData.id,
                key:this.updateColData.key,
                value:this.updateColData.value,
                category:this.updateColData.category,
                description:this.updateColData.description
            }
    
            let res = await updateDictionaryItemManagerCol(sendP);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '更新成功!'
                });
            }
            this.updateColVisible = false
            this.fetchData(this.currentPage,this.colsPerPage)
        }).catch((err)=>{console.warn(err)})
    },
    fetchData(start=1,limit=10) {
      this.listLoading = true
      getDictionaryItemManagerCols({start:start,limit:limit}).then(response => {
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
