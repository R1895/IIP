<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="ID">
            <el-input v-model="searchBox.id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item label="Title">
            <el-input v-model="searchBox.title" placeholder="title"></el-input>
        </el-form-item>
        <el-form-item label="Author">
            <el-input v-model="searchBox.author" placeholder="author"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="search">Search</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">Add</el-button>
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
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>
      <el-table-column label="Author"  align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Pageviews"  align="center">
        <template slot-scope="scope">
          {{ scope.row.pageviews }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="Status" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作" >
        <template slot-scope="scope">
            <el-button type="primary" icon="el-icon-edit" @click="updateColShow(scope.row)" circle></el-button>
            <el-button type="danger" icon="el-icon-delete" @click="deleteCol(scope.row.id)" circle></el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog width="30%" title="Add new column" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="ID" >
                <el-input v-model="newCol.id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Title" >
                <el-input v-model="newCol.title" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Author" >
                <el-input v-model="newCol.author" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">Cancel</el-button>
            <el-button type="primary" @click="addCol">Confirm</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="Update column" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="ID" >
                <el-input v-model="updateColData.id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Title" >
                <el-input v-model="updateColData.title" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Author" >
                <el-input v-model="updateColData.author" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Page views" >
                <el-input v-model="updateColData.pageviews" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="Status" >
                <el-input v-model="updateColData.status" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="updateColVisible = false">Cancel</el-button>
            <el-button type="primary" @click="updateCol">Confirm</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import { getList } from '@/api/table'
import { getTables, addTable, updateTable, deleteTable} from '@/api/mock'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      searchBox:{
        user:"",
        region:"",
        author:""
      },
      addColVisible:false,
      newCol:{
        id:"",
        title:"",
        author:"",
        pageviews:0,
        status:0
      },
      updateColData:{
        id:"",
        title:"",
        author:"",
        pageviews:0,
        status:0
      },
      updateColVisible:false,
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    deleteCol(id){
        this.$confirm('Are you sure to delete this column?', 'Notice', {
            confirmButtonText: 'Confirm',
            cancelButtonText: 'Cancel',
            type: 'warning'
        }).then(async () => {
            console.log("confirm")
            let res = await deleteTable(id);
            console.log(res)
            if(res.Status==200){
                this.$message({
                    type: 'success',
                    message: 'Successfully delete column!'
                });
            }
            this.fetchData()
        }).catch((err)=>{console.log(err)})
    },
    addCol(){
        let data = this.newCol;
        this.$confirm('Are you sure to add a new column?', 'Notice', {
            confirmButtonText: 'Confirm',
            cancelButtonText: 'Cancel',
            type: 'warning'
        }).then(async () => {
            let res = await addTable(data);
            if(res.Status==200){
                this.$message({
                    type: 'success',
                    message: 'Successfully add a new column!'
                });
            }
            this.addColVisible = false
            this.fetchData()
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
        this.$confirm('Are you sure to update this column?', 'Notice', {
            confirmButtonText: 'Confirm',
            cancelButtonText: 'Cancel',
            type: 'warning'
        }).then(async () => {
            let res = await updateTable(this.updateColData);
            if(res.Status==200){
                this.$message({
                    type: 'success',
                    message: 'Successfully update this column!'
                });
            }
            this.updateColVisible = false
            this.fetchData()
        }).catch(()=>{})
    },
    fetchData() {
      this.listLoading = true
      getTables().then(response => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        console.log(response)
        this.list = response.Data
        this.listLoading = false
      })
    },

  }
}
</script>
