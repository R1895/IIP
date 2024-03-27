
<!-- 添加和修改车间的负责人ID加下拉列表 -->
<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="车间名">
        <el-input
          v-model="searchBox.workshop_name"
          placeholder="车间名"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="search">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addColVisible = true"
          >添加车间</el-button
        >
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
      <el-table-column align="center" label="车间ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="车间名">
        <template slot-scope="scope">
          {{ scope.row.workshop_name }}
        </template>
      </el-table-column>
      <el-table-column label="车间负责人" align="center">
        <template slot-scope="scope">
          {{
            scope.row.worker != void 0 ? scope.row.worker.worker_name : "暂无"
          }}
        </template>
      </el-table-column>
      <el-table-column label="车间工人数" align="center">
        <template slot-scope="scope">
          <span>{{
            scope.row.workers != void 0 ? scope.row.workers.length : 0
          }}</span>
        </template>
      </el-table-column>
      <el-table-column label="车间设备数" align="center">
        <template slot-scope="scope">
          {{ scope.row.machines != void 0 ? scope.row.machines.length : 0 }}
        </template>
      </el-table-column>
      <el-table-column label="车间任务数" align="center">
        <template slot-scope="scope">
          {{ scope.row.tasks != void 0 ? scope.row.tasks.length : 0 }}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="操作">
        <template slot-scope="scope">
          <el-tooltip
            class="item"
            effect="dark"
            content="为车间分配工人"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-s-fold"
              @click="
                toUrl({
                  id: scope.row.id,
                  path: '/taskMan/addWorkerToWorkshop',
                })
              "
              circle
            ></el-button>
          </el-tooltip>
          <el-tooltip
            class="item"
            effect="dark"
            content="修改设备其他信息"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-edit"
              @click="updateColShow(scope.row)"
              circle
            ></el-button>
          </el-tooltip>
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
    <el-dialog width="30%" title="添加新车间" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="车间名">
          <el-input
            v-model="newCol.workshop_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="选择车间负责人">
          <el-select v-model="newCol.worker_id" placeholder="请选择">
            <el-option
              :label="option.worker_name"
              :value="option.id"
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
      title="更新车间信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="车间id">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="车间名">
          <el-input
            v-model="updateColData.workshop_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="更换负责人">
          <el-select v-model="updateColData.worker_id" placeholder="请选择">
            <el-option
              :label="option.worker_name"
              :value="option.id"
              v-for="(option, index) in workerList"
              :key="index"
            ></el-option>
          </el-select>
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
import { getList } from "@/api/table";
import {
  getWorkshopManagerCols,
  addWorkshopManagerCol,
  updateWorkshopManagerCol,
  deleteWorkshopManagerCol,
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
        workshop_name: "",
      },
      addColVisible: false,
      newCol: {
        workshop_name: "",
        worker_id: "",
      },
      updateColData: {
        workshop_name: "",
        worker_id: "",
        worker: {
          worker_id: "",
        },
      },
      workerList: [],
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
    toUrl(args) {
      console.log(args);
      this.$router.push({
        path: args.path,
        query: {
          id: args.id,
        },
      });
    },
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
          let res = await deleteWorkshopManagerCol({ id: id });
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
          this.$message({
              type: "warning",
              message: "删除失败，请检查车间依赖是否删除干净",
            });
        });
    },
    async search() {
      if (this.$areAllValuesEmpty(this.searchBox)) {
        this.fetchData(this.currentPage, this.colsPerPage);
        return;
      }
      let sendParams = this.searchBox;
      sendParams.start = 1;
      sendParams.limit = 100;
      let res = await getWorkshopManagerCols(sendParams);
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
      this.$confirm("确定添加新车间吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          console.log(data);
          let res = await addWorkshopManagerCol(data);
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
      data.worker_id = data.worker.id;
      this.updateColVisible = true;
      this.updateColData = data;
    },
    updateCol() {
      // this.updateColData is already exist
      this.$confirm("确定更新此车间信息吗?", "注意", {
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
          let res = await updateWorkshopManagerCol(sendP);
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
      getWorkshopManagerCols({ start: start, limit: limit }).then(
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
