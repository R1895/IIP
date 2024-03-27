<!-- 加两个，查看工人干的任务和工人操作的机器，跳转新页面-->
<!-- 按钮加tooltip-->
<template>
  <div class="app-container">
    <el-form :inline="true" :model="searchBox" class="demo-form-inline">
      <el-form-item label="工人名">
        <el-input
          v-model="searchBox.worker_name"
          placeholder="工人名"
        ></el-input>
      </el-form-item>
      <el-form-item>     
        <el-button type="warning" @click="search">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addColVisible = true"
          >添加工人</el-button
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
      <el-table-column align="center" label="工人ID">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="绑定用户ID">
        <template slot-scope="scope">
          {{ scope.row.user_id }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工人类型">
        <template slot-scope="scope">
          {{ scope.row.worker_type }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="工人名">
        <template slot-scope="scope">
          {{ scope.row.worker_name }}
        </template>
      </el-table-column>
      <el-table-column label="技术等级" align="center">
        <template slot-scope="scope">
          {{ scope.row.skill_level }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="操作">
        <template slot-scope="scope">
          <!-- <el-tooltip
            class="item"
            effect="dark"
            content="查看所工作的机器详情"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-s-fold"
              @click="toUrl({id:scope.row.id,path:'/taskMan/workerMachineDetail'})"
              circle
            ></el-button>
          </el-tooltip>
          <el-tooltip
            class="item"
            effect="dark"
            content="查看任务详情"
            placement="top"
          >
            <el-button
              type="primary"
              icon="el-icon-s-order"
              @click="toUrl({id:scope.row.id,path:'/taskMan/workerTaskDetail'})"
              circle
            ></el-button>
          </el-tooltip> -->
          <el-tooltip
            class="item"
            effect="dark"
            content="修改工人信息"
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
            content="删除信息"
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
    <el-dialog width="30%" title="添加新工人" :visible.sync="addColVisible">
      <el-form :model="newCol">
        <el-form-item label="选择绑定的用户">
          <el-select v-model="newCol.user_id" filterable placeholder="请选择">
            <el-option
              :label="option.user_name"
              :value="option.user_id"
              v-for="(option, index) in userList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
         <el-form-item label="选择工人类型">
          <el-select v-model="newCol.worker_type_id" filterable placeholder="请选择">
            <el-option
              :label="option.worker_type"
              :value="option.id"
              v-for="(option, index) in workerTypeList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工人名字">
          <el-input v-model="newCol.worker_name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="技术等级">
          <el-input v-model="newCol.skill_level" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addColVisible = false">取消</el-button>
        <el-button type="primary" @click="addCol">确认</el-button>
      </div>
    </el-dialog>
    <el-dialog
      width="30%"
      title="更新工人信息"
      :visible.sync="updateColVisible"
    >
      <el-form :model="updateColData">
        <el-form-item label="工人ID">
          <el-input
            v-model="updateColData.id"
            autocomplete="off"
            disabled
          ></el-input>
        </el-form-item>
        <el-form-item label="选择绑定的用户">
          <el-select v-model="updateColData.user_id" filterable placeholder="请选择">
            <el-option
              :label="option.user_name"
              :value="option.user_id"
              v-for="(option, index) in userList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
         <el-form-item label="选择工人类型">
          <el-select v-model="updateColData.worker_type_id" filterable placeholder="请选择">
            <el-option
              :label="option.worker_type"
              :value="option.id"
              v-for="(option, index) in workerTypeList"
              :key="index"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="工人名">
          <el-input
            v-model="updateColData.worker_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="技能等级">
          <el-input
            v-model="updateColData.skill_level"
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
  getWorkerManagerCols,
  addWorkerManagerCol,
  updateWorkerManagerCol,
  deleteWorkerManagerCol,
} from "@/api/team";
import{
  getUserManagerCols
} from '@/api/user'
import{
  getWorkerTypeManagerCols
} from '@/api/industry'


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
        worker_name: "",
      },
      workerTypeList:[],
      addColVisible: false,
      newCol: {
        worker_name: "",
        worker_type_id: "",
        user_id: "",
        skill_level: "",
      },
      updateColData: {
        id: "",
        worker_name: "",
        worker_type_id: "",
        user_id: "",
        skill_level: "",
      },
      updateColVisible: false,
      currentPage: 1,
      userList:[],
      colsPerPage: 10,
      total: 0,
    };
  },
  created() {
    this.fetchData(this.currentPage, this.colsPerPage);
  },
  methods: {
    toUrl(args){
      console.log(args)
      this.$router.push({
        path:args.path,
        params:{
          id:args.id
        }
      })
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
          let res = await deleteWorkerManagerCol({ id: id });
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
      let res = await getWorkerManagerCols({ id: this.searchBox.id });
      this.listLoading = true;
      if (res.data.status == "200") {
        this.list = [];
        this.list.push(res.data.data);
        this.searchBox = {
          id: "",
          worker_name: "",
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
      data.worker_type = data.worker_type_id+""
      data.skill_level = parseInt(data.skill_level)
      this.$confirm("确定添加新工人吗", "提示", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "primary",
      })
        .then(async () => {
          let res = await addWorkerManagerCol(data);
          if (res.data.status == 200) {
            this.$message({
              type: "success",
              message: "添加成功!",
            });
          }
          (this.newCol = {
            worker_name: "",
            worker_type_id: "",
            user_id: "",
            skill_level: "",
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
      this.$confirm("确定更新此工人信息吗?", "注意", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(async () => {
          let sendP = {
            id: this.updateColData.id,
            worker_name: this.updateColData.worker_name,
            workshop_id: this.updateColData.workshop_id,
            description: this.updateColData.Description,
            start_date: this.updateColData.start_date,
            due_date: this.updateColData.due_date,
          };
          console.log(sendP);
          let res = await updateWorkerManagerCol(sendP);
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
      getWorkerManagerCols({ start: start, limit: limit }).then((response) => {
        //这里的response视具体情况而定，这里由于是本地数据，
        //所以response的结构直接就是数据（可以查看/apimock.js的代码）
        //在使用http/request库时，response的数据在response.data中
        this.list = response.data.data.item;
        this.total = response.data.data.total;
        this.listLoading = false;
      });
      getUserManagerCols({start:1,limit:1000}).then(response=>{
        this.userList = response.data.data.item;
      })
      getWorkerTypeManagerCols({start:1,limit:1000}).then(response=>{
        this.workerTypeList = response.data.data.item;

      })
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
