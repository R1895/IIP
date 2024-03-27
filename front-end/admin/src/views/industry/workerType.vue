<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="工人类别ID">
        <el-input v-model="searchBox.id" placeholder="ID"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="search">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addColVisible = true"
          >添加工人类别</el-button
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
      <el-table-column align="center" label="工人类别ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工人类别名字">
        <template slot-scope="scope">
          {{ scope.row.worker_type }}
        </template>
      </el-table-column>
      <el-table-column label="描述" align="center">
        <template slot-scope="scope">
          {{ scope.row.description }}
        </template>
      </el-table-column>
      <el-table-column label="此类别下的工人数" align="center">
        <template slot-scope="scope">
          {{ scope.row.workers==null?0:scope.row.workers.length }}
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
    <el-dialog width="30%" title="上报工人类别" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="工人类型名">
          <el-input v-model="newCol.worker_type" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="newCol.description" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addColVisible = false">取消</el-button>
        <el-button type="primary" @click="addCol">确认</el-button>
      </div>
    </el-dialog>
    <el-dialog
      width="30%"
      title="更新工人类别信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="工人类别ID">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="工人类型名">
          <el-input
            v-model="updateColData.worker_type"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input
            v-model="updateColData.description"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="时间">
          <el-input v-model="updateColData.date" autocomplete="off"></el-input>
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
  getWorkerTypeManagerCols,
  addWorkerTypeManagerCol,
  updateWorkerTypeManagerCol,
  deleteWorkerTypeManagerCol,
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
      addColVisible: false,
      newCol: {
        worker_type: "",
        description: "",
        time: "",
      },
      updateColData: {
        id: "",
        worker_type: "",
        description: "",
      },
      updateColVisible: false,
      currentPage: 1,
      colsPerPage: 10,
      total: 0,
    };
  },
  created() {
    this.fetchData(this.currentPage, this.colsPerPage);
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
          let res = await deleteWorkerTypeManagerCol({ id: id });
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
      let res = await getWorkerTypeManagerCols({ id: this.searchBox.id });
      this.listLoading = true;
      if (res.data.status == "200") {
        this.list = [];
        this.list.push(res.data.data);
        this.searchBox = {
          id: "",
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
      this.$confirm("确定添加新工人类别吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addWorkerTypeManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          (this.newCol = {
            id: "",
            worker_type: "",
            description: "",
          }),
            (this.addColVisible = false);
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
      this.$confirm("确定更新此工人类别信息吗?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          let sendP = {
            id: this.updateColData.id,
            worker_type: this.updateColData.worker_type,
            description: this.updateColData.description,
          };
          let res = await updateWorkerTypeManagerCol(sendP);
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
      getWorkerTypeManagerCols({ start: start, limit: limit }).then(
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
  },
};
</script>
<style>
.block {
  display: flex;
  justify-content: center;
}
</style>
