<!-- 车间ID和任务ID(更改中的)做下拉列表-->
<!-- 加一个分配工人的按钮,跳转到新页面展示设备所含的工人-->

<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="设备ID">
        <el-input v-model="searchBox.id" placeholder="ID"></el-input>
      </el-form-item>
      <el-form-item label="设备名">
        <el-input
          v-model="searchBox.machine_name"
          placeholder="设备名"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="search">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addColVisible = true"
          >添加设备</el-button
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
      <el-table-column align="center" label="机器ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="机器名">
        <template slot-scope="scope">
          {{ scope.row.machine_name }}
        </template>
      </el-table-column>
      <el-table-column label="所属车间ID(未返回名字)" align="center">
        <template slot-scope="scope">
          {{ scope.row.workshop_id }}
        </template>
      </el-table-column>
      <el-table-column label="状态(未返回名字)" align="center">
        <template slot-scope="scope">
          {{ scope.row.status_name }}
        </template>
      </el-table-column>

      <el-table-column
        label="当前任务ID(0为未分配，应返回任务名)"
        align="center"
      >
        <template slot-scope="scope">
          {{
            scope.row.current_task_id == 0
              ? "未分配"
              : scope.row.current_task_id
          }}
        </template>
      </el-table-column>

      <!-- <el-table-column label="购买时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.purchase_date}}
        </template>
      </el-table-column>
      <el-table-column label="保修截止日期" align="center">
        <template slot-scope="scope">
          {{ scope.row.warranty_until}}
        </template>
      </el-table-column> -->

      <el-table-column label="操作工人数" align="center">
        <template slot-scope="scope">
          {{
            scope.row.machine_workers == null
              ? "暂无分配"
              : scope.row.machine_workers.length
          }}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="操作">
        <template slot-scope="scope">
          <el-tooltip
            class="item"
            effect="dark"
            content="为机器分配工人及负责人"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-s-fold"
              @click="
                toUrl({
                  id: scope.row.id,
                  path: '/taskMan/allocateWorkerToMachine',
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
          <el-tooltip
            class="item"
            effect="dark"
            content="修改设备其他信息"
            placement="top"
          >
            <el-button
              type="danger"
              icon="el-icon-delete"
              @click="deleteCol(scope.row.id)"
              circle
            ></el-button>
          </el-tooltip>
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
    <el-dialog width="30%" title="添加新设备" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="设备名">
          <el-input v-model="newCol.machine_name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="选择车间">
          <el-select v-model="newCol.workshop_id" placeholder="请选择">
            <el-option
              :label="option.workshop_name"
              :value="option.id"
              v-for="(option, index) in workshopList"
              :key="index"
            ></el-option>
            <!-- 添加其他机器选项 -->
          </el-select>
        </el-form-item>
        <el-form-item label="选择机器状态">
          <el-select v-model="newCol.machine_status_id" placeholder="请选择">
            <el-option
              :label="option.status_name"
              :value="option.id"
              v-for="(option, index) in machineStatusList"
              :key="index"
            ></el-option>
            <!-- 添加其他机器选项 -->
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
      title="更新设备信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="设备ID">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="设备名">
          <el-input
            v-model="updateColData.machine_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="选择车间">
          <el-select v-model="updateColData.workshop_id" placeholder="请选择">
            <el-option
              :label="option.workshop_name"
              :value="option.id"
              v-for="(option, index) in workshopList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="购买时间" >
                <el-input v-model="updateColData.purchase_date" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="保修截止日期" >
                <el-input v-model="updateColData.warranty_until" autocomplete="off"></el-input>
            </el-form-item> -->
        <el-form-item label="选择机器状态">
          <el-select
            v-model="updateColData.machine_status_id"
            placeholder="请选择"
          >
            <el-option
              :label="option.status_name"
              :value="option.id"
              v-for="(option, index) in machineStatusList"
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
  getDeviceManagerCols,
  addDeviceManagerCol,
  updateDeviceManagerCol,
  deleteDeviceManagerCol,
} from "@/api/industry";
import {
  getWorkshopManagerCols,
  getMachineStatusManagerCols,
} from "@/api/industry";
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
        machine_name: "",
      },
      workshopList: [],
      machineStatusList: [],
      addColVisible: false,
      newCol: {
        machine_name: "",
        workshop_id: "",
        machine_status_id: "",
      },
      updateColData: {
        machine_name: "",
        workshop_id: "",
        purchase_date: "",
        warranty_until: "",
        machine_status_id: "",
        current_task_id: "",
        machine_status_id: "",
        status_id: "",
      },
      updateColVisible: false,
      currentPage: 1,
      colsPerPage: 10,
      total: 0,
    };
  },
  created() {
    this.fetchData(this.currentPage, this.colsPerPage);
    this.initSelectOptions();
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
          let res = await deleteDeviceManagerCol({ id: id });
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
      let res = await getDeviceManagerCols({ id: this.searchBox.id });
      this.listLoading = true;
      if (res.data.status == "200") {
        this.list = [];
        this.list.push(res.data.data);
        this.searchBox = {
          id: "",
          machine_name: "",
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
      this.$confirm("确定添加新设备吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addDeviceManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          (this.newCol = {
            machine_name: "",
            workshop_id: "",
            machine_status_id: "",
          }),
            (this.addColVisible = false);
          this.fetchData(this.currentPage, this.colsPerPage);
        })
        .catch(() => {});
    },
    updateColShow(data) {
      let dataStr = JSON.stringify(data);
      data = JSON.parse(dataStr);
      data.machine_status_id = data.status_id;
      this.updateColVisible = true;
      this.updateColData = data;
    },
    updateCol() {
      // this.updateColData is already exist
      this.$confirm("确定更新此设备信息吗?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          let sendP = {
            id: this.updateColData.id,
            machine_name: this.updateColData.machine_name,
            workshop_id: this.updateColData.workshop_id,
            machine_status_id: this.updateColData.machine_status_id,
            current_task_id: this.updateColData.current_task_id,
            procedure_type: this.updateColData.procedure_type,
          };
          let res = await updateDeviceManagerCol(sendP);
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
      getDeviceManagerCols({ start: start, limit: limit }).then((response) => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        this.list = response.data.data.item;
        this.total = response.data.data.total;
        this.listLoading = false;
      });
    },
    initSelectOptions() {
      getWorkshopManagerCols({ start: 1, limit: 1000 }).then((response) => {
        this.workshopList = response.data.data.item;
      });
      getMachineStatusManagerCols({ start: 1, limit: 1000 }).then(
        (response) => {
          this.machineStatusList = response.data.data.item;
        }
      );
    },
  },
};
</script>
<style lang="scss" scoped>
.block {
  display: flex;
  justify-content: center;
}

.box {
  width: 400px;

  .top {
    text-align: center;
  }

  .left {
    float: left;
    width: 60px;
  }

  .right {
    float: right;
    width: 60px;
  }

  .bottom {
    clear: both;
    text-align: center;
  }

  .item {
    margin: 4px;
  }

  .left .el-tooltip__popper,
  .right .el-tooltip__popper {
    padding: 8px 10px;
  }
}
</style>
