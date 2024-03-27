<template>
  <div class="app-container">
    <!-- 表单 -->
    <h1></h1>
    <el-form :model="form" ref="form" label-width="100px">
      <!-- 生产任务字段 -->
      <el-form-item label="产线ID">
        {{this.taskForm.id}}
      </el-form-item>
      <el-form-item label="产线类型">
        {{this.taskForm.task_name}}
      </el-form-item>
      <el-form-item label="产线描述">
        {{this.taskForm.description}}
      </el-form-item>
      <el-form-item label="产线工序数量">
        {{this.taskForm.procedures==null?0:this.taskForm.procedures.length}}
      </el-form-item>
      <el-form-item label="产线开始时间">
        {{this.taskForm.start_date}}
      </el-form-item>

      <el-form-item label="添加工序">
        <el-button type="primary" @click="addColVisible = true"
          >点击添加</el-button
        >
      </el-form-item>

      <!-- 工序列表 -->
      <el-form-item label="工序列表">
        <el-table
          v-loading="listLoading"
          :data="procedures"
          element-loading-text="Loading"
          border
          fit
          highlight-current-row
        >
          <el-table-column align="center" label="工序序列">
            <template slot-scope="scope">
              {{ scope.row.sequence }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="工序名称">
            <template slot-scope="scope">
              {{ scope.row.procedure_name }}
            </template>
          </el-table-column>
          <el-table-column label="机器ID" align="center">
            <template slot-scope="scope">
              {{ scope.row.machine_id }}
            </template>
          </el-table-column>
          <el-table-column label="工件数量" align="center">
            <template slot-scope="scope">
              {{ scope.row.workpieces_num }}
            </template>
          </el-table-column>
          <el-table-column label="正在生产工件数量" align="center">
            <template slot-scope="scope">
              {{ scope.row.processed_num }}
            </template>
          </el-table-column>
          <el-table-column label="开始时间" align="center">
            <template slot-scope="scope">
              {{ scope.row.start_date }}
            </template>
          </el-table-column>
          <el-table-column label="是否完成" align="center">
            <template slot-scope="scope">
              {{ scope.row.is_finished==0?"否":"是" }}
            </template>
          </el-table-column>

          <el-table-column align="center" prop="created_at" label="操作">

            <template slot-scope="scope">
              <el-button
                type="primary"
                icon="el-icon-top"
                @click="updateOrder(scope.row.id,index=procedures.indexOf(scope.row),up=true)"
                circle
              ></el-button>
              <el-button
                type="primary"
                icon="el-icon-bottom"
                @click="updateOrder(scope.row.id,index=procedures.indexOf(scope.row),up=false)"
                circle
              ></el-button>
              <el-button
                type="danger"
                icon="el-icon-delete"
                @click="deleteCol(scope.row.id,procedures.indexOf(scope.row))"
                circle
              ></el-button>
            </template>
            
          </el-table-column>
        </el-table>
      </el-form-item>
      <el-dialog width="30%" title="添加新工序" :visible.sync="addColVisible">
        <el-form :model="newCol">
          <el-form-item label="工序名">
            <el-input v-model="newCol.procedure_name" autocomplete="off"></el-input>
          </el-form-item>
        <el-form-item label="选择工序状态">
          <el-select v-model="newCol.procedure_type" placeholder="请选择" @change="procedureTypeChange">
            <el-option
              :label="option.procedure_name"
              :value="option.procedure_type"
              v-for="(option, index) in procedureTypeList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="为工序选择机器">
          <el-select v-model="newCol.machine_id" placeholder="请选择">
            <el-option
              :label="option.machine_name"
              :value="option.id"
              v-for="(option, index) in machineList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="addColVisible = false">取消</el-button>
          <el-button type="primary" @click="addCol">确认</el-button>
        </div>
      </el-dialog>
      <!-- 提交按钮 -->
      <el-form-item>
        <el-button type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getTaskTypeManagerCols } from "@/api/industry";
import { getProcedureManagerCols,addProcedureManagerCol,updateProcedureManagerCol,deleteProcedureManagerCol } from "@/api/industry"
import { getTaskManagerCols } from "@/api/team"
import { getProceduretypeManagerCols,getDeviceManagerCols} from '@/api/industry'

export default {
  data() {
    return {
      listLoading:false,
      form: {
        machine: "", // 机器
        productionTask: "", // 生产任务
        processList: [], // 工序列表
      },
      newCol: {
        procedure_name: "",
        procedure_type: "",
        task_id: "",
        machine_id:"",
        sequence:1,
        start_date:"2006-01-02T15:04:05Z",
        procedure_status_id:1
      },
      addColVisible: false,
      task_id:"",
      taskForm:{},
      procedures:[],
      procedureTypeList:[],
      machineList:[],
      lock:false
    };
  },
  methods: {
    back(){
      this.$router.back()
    },
    async updateOrder(id,index,up){
      if(this.lock==true){
        console.log("点击太快，锁住了")
        return;
      }
      if(up==true&&index==0){
        console.log("边界情况")
      } else if (up==false&&index==this.procedures.length-1){
        console.log("边界情况")
      } else {
        this.lock=true
        //移动
        if(up==true){
          //和前一位交换
          this.procedures[index].sequence-=1;
          this.procedures[index-1].sequence+=1;
          console.log(this.procedures[index].sequence,this.procedures[index-1].sequence)
          let res1 = await updateProcedureManagerCol(this.procedures[index])
          let res2 = await updateProcedureManagerCol(this.procedures[index-1])
          this.fetchData()
        } else {
          this.procedures[index].sequence+=1;
          this.procedures[index+1].sequence-=1;
          let res1 = await updateProcedureManagerCol(this.procedures[index])
          let res2 = await updateProcedureManagerCol(this.procedures[index+1])
          this.fetchData()
        }
        this.lock=false
      }
    },
    procedureTypeChange(value){
      console.log(value)
      getDeviceManagerCols({start:1,limit:1000,procedure_type:value}).then((response) => {
        this.machineList = response.data.data.item==null?[]:response.data.data.item;
      });
    },
    addCol() {
      let data = this.newCol
      let lastIndex= this.procedures==[]?-1:this.procedures.length-1
      if(lastIndex==-1){
        data.sequence = 1
      } else {
        data.sequence = this.procedures[lastIndex].sequence +1
      }
      data.task_id = parseInt(this.task_id)
      this.$confirm("确定添加新工序吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addProcedureManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          this.newCol = {
            procedure_name: "",
            procedure_type: "",
            task_id: "",
            machine_id:"",
            sequence:0,
            start_date:"2006-01-02T15:04:05Z",
            procedure_status_id:1
          }
          this.addColVisible = false
          this.fetchData();
        })
        .catch((err) => {
          this.$message({
            type: "warning",
            message: "请检查字段是否为空",
          });
        });
    },
    addProcess() {
      // 添加工序
      this.form.processList.push(`工序${this.form.processList.length + 1}`);
    },
    removeProcess(index) {
      // 移除工序
      this.form.processList.splice(index, 1);
    },
    submitForm() {
      // 提交表单逻辑
      console.log("表单数据:", this.form);
    },
    deleteCol(id,index) {
      this.$confirm("确定删除此行信息吗(删除后不可恢复)?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          if(this.procedures[index].is_finished==2){
             this.$message({
              type: "warning",
              message: "无法删除，仍有工件正在加工中",
            });
            return;
          }
          console.log("confirm");
          let res = await deleteProcedureManagerCol({ id: id });

          for(let i=index+1;i<this.procedures.length;i++){
            this.procedures[i].sequence-=1
            let res = await updateProcedureManagerCol(this.procedures[i])
          }
          this.fetchData();
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "删除成功",
            });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    fetchData(){
      this.listLoading = true
      getProcedureManagerCols({start:1,limit:10,task_id:this.task_id}).then((response)=>{
        this.procedures = response.data.data.item==null?[]:response.data.data.item;
        if(this.procedures!=[]){
            this.procedures.sort((a,b)=>{
            return a.sequence>b.sequence
        })
        }
        this.listLoading = false
      });
    },
    initData(){
      getTaskManagerCols({id:this.task_id}).then((response) => {
        this.taskForm = response.data.data;
      });
      getProcedureManagerCols({start:1,limit:10,task_id:this.task_id}).then((response)=>{
        this.procedures = response.data.data.item==null?[]:response.data.data.item;
        this.procedures.sort((a,b)=>{
          return a.sequence>b.sequence
        })
      });
      getProceduretypeManagerCols({start:1,limit:1000}).then((response) => {
        this.procedureTypeList = response.data.data.item==null?[]:response.data.data.item;
      });
    }
  },
  created(){
    this.task_id = this.$route.query.id;
    console.log(this.task_id)
    this.initData()
  },
};
</script>

<style>
/* 可以添加一些样式进行美化 */
.cards {
  margin-top: 10px;
  margin-bottom: 10px;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
.process-card {
  margin-left: 10px;
  margin-top: 5px;
  margin-bottom: 5px;
  width: 200px;
}
.process-card > p {
  line-height: 20px;
}
.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}
</style>
