<template>
  <div class="app-container">
    <!-- 表单 -->
    <el-form :model="form" ref="form" label-width="100px">
      <!-- 生产任务字段 -->
      <el-form-item label="设备名称">
        {{ this.machineForm.machine_name }}
      </el-form-item>
      <el-form-item label="所属车间名称">
        {{ this.workshopForm.workshop_name }}
      </el-form-item>

      <el-form-item label="分配工人">
        <el-button type="primary" @click="addColVisible = true"
          >点击添加</el-button
        >
      </el-form-item>

      <!-- 工序列表 -->
      <el-form-item label="车间下分配的工人列表">
        <el-table
          v-loading="listLoading"
          :data="workers"
          element-loading-text="Loading"
          border
          fit
          highlight-current-row
        >
          <el-table-column align="center" label="工人名">
            <template slot-scope="scope">
              {{ scope.row.worker_name }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="工人类别">
            <template slot-scope="scope">
              {{ scope.row.worker_type }}
            </template>
          </el-table-column>

          <el-table-column align="center" prop="created_at" label="操作">
            <template slot-scope="scope">
              <el-button
                type="danger"
                icon="el-icon-delete"
                @click="deleteCol(scope.row.id, workers.indexOf(scope.row))"
                circle
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-form-item>
      <el-dialog width="30%" title="分配工人" :visible.sync="addColVisible">
        <el-form :model="newCol">
          <el-form-item label="选择车间下的工人">
            <el-select
              v-model="newCol.worker_id"
              placeholder="请选择"
            >
              <el-option
                :label="option.worker_name"
                :value="option.id"
                v-for="(option, index) in workerOptionList"
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
import { getWorkshopManagerCols,getMachineWorkerManagerCols,addMachineWorkerManagerCol,deleteMachineWorkerManagerCol } from "@/api/industry";
import { getDeviceManagerCols, addWorkerToDevice, deleteWorkerToDevice } from "@/api/industry";

export default {
  data() {
    return {
      listLoading: false,
      form: {
        machine: "", // 机器
        productionTask: "", // 生产任务
        processList: [], // 工序列表
      },
      newCol: {
        worker_id:"",
      },
      addColVisible: false,
      machine_id: "",
      machineForm: {},
      workers: [],
      procedureTypeList: [],
      machineList: [],
      lock: false,
      workerOptionList:[],
      workshopForm:{}
    };
  },
  methods: {
    back() {
      this.$router.back();
    },
    procedureTypeChange(value) {
      console.log(value);
      getDeviceManagerCols({
        start: 1,
        limit: 1000,
        procedure_type: value,
      }).then((response) => {
        this.machineList =
          response.data.data.item == null ? [] : response.data.data.item;
      });
    },
    addCol() {
      let data = this.newCol;
      data.worker_id = parseInt(data.worker_id);
      data.machine_id = parseInt(this.machine_id);
      data.start_date = new Date();
      this.$confirm("确定分配新工人吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addMachineWorkerManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          this.newCol = {
            worker_id:""
          };
          this.addColVisible = false;
          this.fetchData();
        })
        .catch((err) => {
          this.$message({
            type: "warning",
            message: "请检查字段是否为空",
          });
        });
    },
    deleteCol(id, index) {
      this.$confirm("确定删除此行信息吗(删除后不可恢复)?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          console.log("confirm");
          let res = await deleteMachineWorkerManagerCol({ id:id});

          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "删除成功",
            });
          }
          this.fetchData();
        })
        .catch((err) => {
          console.log(err);
        });
    },
    fetchData() {
      this.listLoading = true;
      getDeviceManagerCols({ id: this.machine_id }).then(async (response) => {
        this.machineForm = response.data.data;
        let workerRes = await getMachineWorkerManagerCols({start:1,limit:1000,machine_id:this.machineForm.id})
        console.log('workerRes',workerRes)
        this.workers = workerRes.data.data.item==null?[]:workerRes.data.data.item
        
        let res = await getWorkshopManagerCols({id:this.machineForm.workshop_id }).then((response)=>{
            this.workshopForm = response.data.data;
            let workerOptionList = response.data.data.workers;
            
            workerOptionList.forEach(item=>{
                let isDul = false;
                for(let i=0;i<this.workers.length;i++){
                    if(this.workers[i].id==item.id){
                        isDul = true;
                        break;
                    }
                }
                if(isDul==false){
                    this.workerOptionList.push(item)
                }
            })



            this.listLoading = false;
        })
      });
    },
    initData() {
      

    },
  },
  created() {
    this.machine_id = this.$route.query.id;
    console.log(this.machine_id);
    this.fetchData();
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
