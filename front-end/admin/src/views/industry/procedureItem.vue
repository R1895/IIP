<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
        <el-form-item label="工序产品ID">
            <el-input v-model="searchBox.id" placeholder="ID"></el-input>
        </el-form-item>
        <el-form-item>
            <el-button type="warning" @click="search">搜索</el-button>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="addColVisible = true">添加工序产品</el-button>
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
    <el-table-column align="center" label="工序产品ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工序ID">
        <template slot-scope="scope">
          {{ scope.row.procedure_id }}
        </template>
      </el-table-column>
       <el-table-column label="设备ID" align="center">
        <template slot-scope="scope">
          {{ scope.row.machine_id }}
        </template>
      </el-table-column> 
      <el-table-column label="工序产品值" align="center">
        <template slot-scope="scope">
          {{ scope.row.worker_id}}
        </template>
      </el-table-column> 
      <el-table-column label="产品类别" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_type}}
        </template>
      </el-table-column> 
      <el-table-column label="产品名称" align="center">
        <template slot-scope="scope">
          {{ scope.row.item_type}}
        </template>
      </el-table-column> 
      <el-table-column label="产品品质" align="center">
        <template slot-scope="scope">
          {{ scope.row.quality}}
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
    <el-dialog width="30%" title="上报工序产品" :visible.sync="addColVisible">
        <el-form :model="newCol">
            <el-form-item label="选择工序">
                <el-select v-model="newCol.procedure_id" placeholder="请选择">
                    <el-option
                    :label="option.procedure_name"
                    :value="option.id"
                    v-for="(option, index) in procedureList"
                    :key="index"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="工人ID" >
                <el-input v-model="newCol.worker_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="物品名称" >
                <el-input v-model="newCol.item_name" autocomplete="off"></el-input>
            </el-form-item>
             <el-form-item label="物品类型" >
                <el-input v-model="newCol.item_type" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="物品品质" >
                <el-input v-model="newCol.quality" autocomplete="off"></el-input>
            </el-form-item>


        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="addColVisible = false">取消</el-button>
            <el-button type="primary" @click="addCol">确认</el-button>
        </div>
    </el-dialog>
    <el-dialog width="30%" title="更新工序产品信息" :visible.sync="updateColVisible">
        <el-form :model="updateColData">
            <el-form-item label="选择工序">
                <el-select v-model="updateColData.procedure_id" placeholder="请选择">
                    <el-option
                    :label="option.procedure_name"
                    :value="option.id"
                    v-for="(option, index) in procedureList"
                    :key="index"
                    ></el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="工人ID" >
                <el-input v-model="updateColData.worker_id" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="物品名称" >
                <el-input v-model="updateColData.item_name" autocomplete="off"></el-input>
            </el-form-item>
             <el-form-item label="物品类型" >
                <el-input v-model="updateColData.item_type" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="物品品质" >
                <el-input v-model="updateColData.quality" autocomplete="off"></el-input>
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
import { getProcedureItemManagerCols, addProcedureItemManagerCol, updateProcedureItemManagerCol, deleteProcedureItemManagerCol, getProcedureManagerCols} from '@/api/industry'

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
        procedure_id:"",
        machine_id:"",
        worker_id:"",
        item_type:"",
        item_name:"",
        quality:""
      },
      updateColData:{
        id: "",
        procedure_id:"",
        machine_id:"",
        worker_id:"",
        item_type:"",
        item_name:"",
        quality:""
      },
      updateColVisible:false,
      currentPage:1,
      colsPerPage:10,
      total:0,
      procedureList:[]
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
            let res = await deleteProcedureItemManagerCol({id:id});
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
        let res = await getProcedureItemManagerCols({id:this.searchBox.id})
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
        data.procedure_id = parseInt(data.procedure_id)
        let index = this.procedureList.findIndex(item=>item.id==data.procedure_id)
        if(index==-1){
            this.$message({
                type: 'warning',
                message: '绑定有误，请稍后再试'
            });
            return;
        }
        data.machine_id = parseInt(this.procedureList[index].machine_id)
        data.worker_id = parseInt(data.worker_id)
        this.$confirm('确定添加新工序产品吗', '提示', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'primary'
        }).then(async () => {
            let res = await addProcedureItemManagerCol(data);
            if(res.data.status==200){
                this.$message({
                    type: 'success',
                    message: '添加成功!'
                });
            }
            this.newCol = {
                procedure_id:"",
                machine_id:"",
                worker_id:"",
                item_type:"",
                item_name:"",
                quality:""
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
        this.$confirm('确定更新此工序产品信息吗?', '注意', {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning'
        }).then(async () => {
            let sendP= {
                id:this.updateColData.id,
                procedure_id:parseInt(this.updateColData.procedure_id),
                machine_id:parseInt(this.updateColData.machine_id),
                worker_id:parseInt(this.updateColData.worker_id),
                item_type:this.updateColData.item_type,
                item_name:this.updateColData.item_name,
                quality:this.updateColData.quality
            }
            let res = await updateProcedureItemManagerCol(sendP);
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
      getProcedureItemManagerCols({start:start,limit:limit}).then(response => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        this.list = response.data.data.item
        this.total = response.data.data.total
        this.listLoading = false
        })

        getProcedureManagerCols({start:1,limit:100}).then(response => {
            this.procedureList = response.data.data.item
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
