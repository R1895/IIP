
<!-- todo:缺少添加和修改 -->
<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="工序ID">
        <el-input v-model="searchBox.id" placeholder="ID"></el-input>
      </el-form-item>
      <el-form-item label="工序名">
        <el-input
          v-model="searchBox.workshop_name"
          placeholder="工序名"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="search">搜索</el-button>
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
      <el-table-column align="center" label="工序ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工序类别ID">
        <template slot-scope="scope">
          {{ scope.row.process_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="进行的任务ID">
        <template slot-scope="scope">
          {{ scope.row.task_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="使用的机器ID">
        <template slot-scope="scope">
          {{ scope.row.machine_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工件数">
        <template slot-scope="scope">
          {{ scope.row.workpieces_num }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="完成任务数">
        <template slot-scope="scope">
          {{ scope.row.processed_num }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工序序列号">
        <template slot-scope="scope">
          {{ scope.row.sequence }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工序任务开始时期">
        <template slot-scope="scope">
          {{ scope.row.start_date }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工序任务是否完成">
        <template slot-scope="scope">
          {{ scope.row.is_finished }}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="操作">
        <template slot-scope="scope">
          <el-button
            type="primary"
            icon="el-icon-edit"
            @click="updateColShow(scope.row)"
            circle
          ></el-button>
          <el-button
            type="danger"
            icon="el-icon-delete"
            @click="deleteCol(scope.row.id)"
            circle
          ></el-button>
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
    <el-dialog width="30%" title="添加新工序" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="工序名">
          <el-input
            v-model="newCol.workshop_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="选择工序负责人">
          <el-select v-model="newCol.worker_id" placeholder="请选择">
            <el-option
              :label="option.worker_name"
              :value="option.worker_id"
              v-for="(option, index) in workerList"
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
    <el-dialog
      width="30%"
      title="更新工序信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="工序id">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="工序名">
          <el-input
            v-model="updateColData.workshop_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="负责人id">
          <el-input
            v-model="updateColData.worker.worker_id"
            autocomplete="off"
          ></el-input>
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
import {
  getProcedureManagerCols,
  addProcedureManagerCol,
  updateProcedureManagerCol,
  deleteProcedureManagerCol,
} from "@/api/industry";
import { getWorkerManagerCols } from "@/api/team";
export default {
  filters: {
    statusFilter(status) {
      if (status) {
        return "success";
      } else {
        return "gray";
      }
    },
  },
  data() {
    return {
      list: [],
      listLoading: false,
      searchBox: {
        id: "",
        procedure_name: "",
      },
      workerList: [],
      addColVisible: false,
      newCol: {
        workshop_name: "",
        worker_id: "",
      },
      updateColData: {
        workshop_name: "",
        worker: {
          worker_id: "",
        },
      },
      updateColVisible: false,
      currentPage: 1,
      colsPerPage: 10,
      total: 0,
    };
  },
  created() {
    this.fetchData(this.currentPage, this.colsPerPage);
    this.initData();
  },
  methods: {
    changePage(newPage) {
      //处理分页
      this.currentPage = newPage;
      console.log(this.currentPage);
      this.fetchData(this.currentPage, this.colsPerPage);
    },
    preprocessData(list) {
      for (let i = 0; i < list.length; i++) {
        let str = "";
        if (list[i].roles == null || list[i].roles == void 0) {
          list[i].rolestring = "暂无";
          continue;
        }
        for (let j = 0; j < list[i].roles.length; j++) {
          if (j == list[i].roles.length - 1) {
            str += `${list[i].roles[j].role_name}`;
          } else {
            str += `${list[i].roles[j].role_name}-`;
          }
        }
        list[i].rolestring = str;
      }
      return list;
    },
    deleteCol(id) {
      this.$confirm("确定删除此行信息吗(删除后不可恢复)?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          console.log("confirm");
          let res = await deleteProcedureManagerCol({ id: id });
          console.log(res);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "删除成功",
            });
          }
          this.fetchData(this.currentPage, this.colsPerPage);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async search() {
      if (this.$areAllValuesEmpty(this.searchBox)) {
        this.fetchData(this.currentPage, this.colsPerPage);
        return;
      }
      let sendParams = this.searchBox;
      sendParams.id = this.searchBox.id;
      console.log(this.searchBox.id);
      let res = await getProcedureManagerCols({ id: this.searchBox.id });
      this.listLoading = true;
      if (res.data.status == "200") {
        this.list = [];
        this.list.push(res.data.data);
        this.searchBox = {
          id: "",
          workshop_name: "",
          email: "",
          telephone: "",
        };
        this.listLoading = false;
        console.log(this.listLoading);
        this.total = 1;
      } else {
        this.$message({
          type: "warning",
          message: "未搜索到此信息",
        });
        this.listLoading = false;
      }
    },
    addCol() {
      let data = this.newCol;
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
            workshop_name: "",
            email: "",
            telephone: "",
          };
          this.addColVisible = false;
          this.fetchData(this.currentPage, this.colsPerPage);
        })
        .catch(() => {});
    },
    updateColShow(data) {
      let dataStr = JSON.stringify(data);
      data = JSON.parse(dataStr);
      this.updateColVisible = true;
      this.updateColData = data;
    },
    updateCol() {
      // this.updateColData is already exist
      this.$confirm("确定更新此工序信息吗?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          let sendP = {
            id: this.updateColData.id,
            workshop_name: this.updateColData.workshop_name,
            worker_id: this.updateColData.worker.worker_id,
          };
          let res = await updateProcedureManagerCol(sendP);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "更新成功!",
            });
          }
          this.updateColVisible = false;
          this.fetchData(this.currentPage, this.colsPerPage);
        })
        .catch(() => {});
    },
    fetchData(start = 1, limit = 10) {
      this.listLoading = true;
      getProcedureManagerCols({ start: start, limit: limit }).then(
        (response) => {
          //这里的response视具体情况而定，这里由于是本地数据，
          //所以response的结构直接就是数据（可以查看/apimock.js的代码）
          //在使用http/request库时，response的数据在response.data中
          this.list = response.data.data.item;
          this.total = response.data.data.total;
          this.listLoading = false;
        }
      );
    },
    initData() {
      getWorkerManagerCols({ start: 1, limit: 1000 }).then((response) => {
        this.workerList = response.data.data.item;
      });
    },
  },
};
</script>
<style>
.block {
  display: flex;
  justify-content: center;
}
</style>
