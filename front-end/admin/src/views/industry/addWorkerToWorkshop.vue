<template>
  <div class="app-container">
    <!-- 表单 -->
    <el-form :model="form" ref="form" label-width="100px">
      <!-- 生产任务字段 -->
      <el-form-item label="车间名称">
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
          <el-table-column align="center" label="技能等级">
            <template slot-scope="scope">
              {{ scope.row.skill_level }}
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
          <el-form-item label="选择工人">
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
import { getWorkerManagerCols } from "@/api/team";
import { getWorkshopManagerCols, addWorkerToWorkshop, deleteWorkerToWorkshop } from "@/api/industry";

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
      workshop_id: "",
      workshopForm: {},
      workers: [],
      procedureTypeList: [],
      machineList: [],
      lock: false,
      workerOptionList:[]
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
      data.workshop_id = parseInt(this.workshop_id);
      this.$confirm("确定分配新工人吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addWorkerToWorkshop(data);
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
    deleteCol(id, index) {
      this.$confirm("确定删除此行信息吗(删除后不可恢复)?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          console.log("confirm");
          let res = await deleteWorkerToWorkshop({ worker_id: id,workshop_id:this.workshop_id });

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
    fetchData() {
      this.listLoading = true;
      getWorkshopManagerCols({ id: this.workshop_id }).then((response) => {
        this.workshopForm = response.data.data;
        this.workers = response.data.data.workers;
        this.listLoading = false;
      });
       getWorkerManagerCols({start:1,limit:100,type:1}).then((response) => {
        this.workerOptionList = response.data.data.item;
        console.log(this.workerOptionList)
      });
    },
    initData() {
      getWorkshopManagerCols({ id: this.workshop_id }).then((response) => {
        this.workshopForm = response.data.data;
        this.workers = response.data.data.workers;
      });
      getWorkerManagerCols({start:1,limit:100,type:1}).then((response) => {
        this.workerOptionList = response.data.data.item;
        console.log(this.workerOptionList)
      });

    },
  },
  created() {
    this.workshop_id = this.$route.query.id;
    console.log(this.workshop_id);
    this.initData();
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
