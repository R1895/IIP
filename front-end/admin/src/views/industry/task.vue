<!--添加产线不参与添加工序-->

<!-- issue: 没有返回task对应的ID -->
<!-- issue: 没有返回task对应的ID -->
<!-- 添加失败，数据库操作失败-->
<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="产线名">
        <el-input v-model="searchBox.task_name" placeholder="产线名"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="warning" @click="search">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addColVisible = true"
          >添加产线</el-button
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
      <el-table-column align="center" label="产线ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="所属车间ID">
        <template slot-scope="scope">
          {{ scope.row.workshop_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="产线名">
        <template slot-scope="scope">
          {{ scope.row.task_name }}
        </template>
      </el-table-column>
      <el-table-column label="产线所含工序数量" align="center">
        <template slot-scope="scope">
          {{ scope.row.procedures==null?0:scope.row.procedures.length }}
        </template>
      </el-table-column>
      <el-table-column label="产线描述" align="center">
        <template slot-scope="scope">
          {{ scope.row.description }}
        </template>
      </el-table-column>
      <el-table-column label="开始时间" align="center">
        <template slot-scope="scope">
          {{ scope.row.start_date }}
        </template>
      </el-table-column>
      

      <el-table-column label="是否完成" align="center">
        <template slot-scope="scope">
          {{ scope.row.is_finished == true ? "是" : "否" }}
        </template>
      </el-table-column>

      <el-table-column align="center" prop="created_at" label="操作">
        <template slot-scope="scope">
          <el-tooltip
            class="item"
            effect="dark"
            content="查看产线工序"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-s-order"
              @click="
                toUrl({
                  id: scope.row.id,
                  path: '/taskMan/addProcedureToTask',
                })
              "
              circle
            ></el-button>
          </el-tooltip>
          <el-tooltip
            class="item"
            effect="dark"
            content="修改产线信息"
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
            content="删除此产线信息"
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
    <el-dialog width="30%" title="添加新产线" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="产线名">
          <el-input v-model="newCol.task_name" autocomplete="off"></el-input>
        </el-form-item>

        <el-form-item label="选择车间">
          <el-select v-model="newCol.workshop_id" placeholder="请选择">
            <el-option
              :label="option.workshop_name"
              :value="option.id"
              v-for="(option, index) in workshopList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="选择产线状态">
          <el-select v-model="newCol.task_status_id" placeholder="请选择">
            <el-option
              :label="option.task_status"
              :value="option.id"
              v-for="(option, index) in taskStatusList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="选择产线类别">
          <el-select v-model="newCol.task_type_id" placeholder="请选择">
            <el-option
              :label="option.task_type"
              :value="option.task_type_id"
              v-for="(option, index) in taskTypeList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="产线描述">
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
      title="更新产线信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="产线ID">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="产线名">
          <el-input
            v-model="updateColData.task_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>

        <el-form-item label="更改车间">
          <el-select v-model="updateColData.workshop_id" placeholder="请选择">
            <el-option
              :label="option.workshop_name"
              :value="option.id"
              v-for="(option, index) in workshopList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="更改产线状态">
          <el-select v-model="updateColData.task_status_id" placeholder="请选择">
            <el-option
              :label="option.task_status"
              :value="option.id"
              v-for="(option, index) in taskStatusList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="更改产线类别">
          <el-select v-model="updateColData.task_type_id" placeholder="请选择">
            <el-option
              :label="option.task_type"
              :value="option.task_type_id"
              v-for="(option, index) in taskTypeList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="产线描述">
          <el-input
            v-model="updateColData.description"
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
  getTaskManagerCols,
  addTaskManagerCol,
  updateTaskManagerCol,
  deleteTaskManagerCol,
} from "@/api/team";
import { getWorkshopManagerCols,getTaskStatusManagerCols,getTaskTypeManagerCols } from "@/api/industry";
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
        task_name: "",
      },
      taskStatusList:[],
      taskTypeList:[],
      workshopList: [],
      addColVisible: false,
      newCol: {
        task_name: "",
        workshop_id: "",
        description: "",
        task_status_id:"",
        task_type_id:""
      },
      updateColData: {
        task_name: "",
        workshop_id: "",
        description: "",
        task_status_id:"",
        task_type_id:""
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
          let res = await deleteTaskManagerCol({ id: id });
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
      let res = await getTaskManagerCols({ id: this.searchBox.id });
      this.listLoading = true;
      if (res.data.status == "200") {
        this.list = [];
        this.list.push(res.data.data);
        this.searchBox = {
          id: "",
          task_name: "",
        };
        this.listLoading = false;
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
      this.$confirm("确定添加新产线吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addTaskManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          (this.newCol = {
            task_name: "",
            workshop_id: "",
            task_status_id:"",
            task_type_id:"",
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
      console.log(this.updateColData)
      // this.updateColData is already exist
      this.$confirm("确定更新此产线信息吗?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          let sendP = {
            id: this.updateColData.id,
            task_name: this.updateColData.task_name,
            workshop_id: this.updateColData.workshop_id,
            description: this.updateColData.description,
            task_status_id: this.updateColData.task_status_id,
            task_type_id: this.updateColData.task_type_id
          };
          console.log(sendP);
          let res = await updateTaskManagerCol(sendP);
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
      getTaskManagerCols({ start: start, limit: limit }).then((response) => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        this.list = response.data.data.item;
        this.total = response.data.data.total;
        this.listLoading = false;
      });
    },
    initData() {
      getWorkshopManagerCols({ start: 1, limit: 1000 }).then((response) => {
        this.workshopList = response.data.data.item;
      });

      getTaskStatusManagerCols({ start: 1, limit: 1000 }).then((response) => {
        this.taskStatusList = response.data.data.item;
      });

      getTaskTypeManagerCols({ start: 1, limit: 1000 }).then((response) => {
        this.taskTypeList = response.data.data.item;
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
