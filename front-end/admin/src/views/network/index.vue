<template>
  <div class="network-container">
    <div class="network-search">
      <h2>网络检索</h2>
      <el-input
        placeholder="请输入内容"
        v-model="input"
        class="input-with-select"
      >
        <el-select
          v-model="select"
          style="width: 110px"
          class="flex-center"
          slot="prepend"
          placeholder="请选择"
        >
          <el-option label="网络名" value="1"></el-option>
          <el-option label="网络状态" value="2"></el-option>
          <el-option label="网络Id" value="3"></el-option>
        </el-select>
        <el-button
          slot="append"
          icon="el-icon-search"
          @click="searchNetwork"
        ></el-button>
      </el-input>
    </div>
    <div class="network-box">
      <div class="network-top">
        <h2>Network</h2>
        <div>
          <el-button class="network-add" type="success" @click="addNetwork()"
            >新增</el-button
          >
        </div>
      </div>
      <el-table
        v-loading="listLoading"
        :data="list"
        :header-cell-style="tableHeaderColor"
        element-loading-text="Loading"
        border
        fit
        style="border-radius: 5px"
        highlight-current-row
        class="network-table"
      >
        <el-table-column align="center" label="ID" width="60">
          <template slot-scope="scope">
            {{ scope.$index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="NetworkId" width="120">
          <template slot-scope="scope">
            {{ scope.row.id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="NetworkName" width="250">
          <template slot-scope="scope">
            <span>{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="Gateway" width="200">
          <template slot-scope="scope">
            {{ scope.row.gateway }}
          </template>
        </el-table-column>
        <el-table-column
          class-name="status-col"
          label="Type"
          width="110"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag :type="scope.row.type | statusFilter">{{
              scope.row.type
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="created_at"
          label="CreatedAt"
          width="200"
        >
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.created_at }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="Process">
          <template slot-scope="scope">
            <div>
              <el-button type="primary" @click="modifyNetwork(scope.row.id)"
                >修改</el-button
              >
              <el-button type="danger" @click="deleteNetwork(scope.row.id)"
                >删除</el-button
              >
              <el-button @click="createNodes(scope.row.id)"
                >批量创建节点
              </el-button>
              <el-button @click="showNetwork(scope.row.id)"
                >查看网络状态
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        style="margin: 10px"
        class="flex-center"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page.sync="currentPage"
        :page-size="pageSize"
        layout="total, prev, pager, next"
        :total="total"
      >
      </el-pagination>
    </div>
    <el-dialog title="创建网络" :visible.sync="createVisible">
      <el-form :model="createForm">
        <el-form-item label="网络名称" :label-width="formLabelWidth">
          <el-input
            v-model="createForm.network_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="网关" :label-width="formLabelWidth">
          <el-input v-model="createForm.gateway" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="类型" :label-width="formLabelWidth">
          <el-input v-model="createForm.type" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="createVisible = false">取 消</el-button>
        <el-button type="primary" @click="createNetwork">确 定</el-button>
      </div>
    </el-dialog>
    <el-dialog title="修改网络" :visible.sync="modifyVisible" width="550">
      <el-form :model="modifyForm">
        <el-form-item label="网络Id" :label-width="formLabelWidth">
          <div>{{ modifyForm.id }}</div>
        </el-form-item>
        <el-form-item label="网络名称" :label-width="formLabelWidth">
          <el-input
            v-model="modifyForm.network_name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="网关" :label-width="formLabelWidth">
          <el-input v-model="modifyForm.gateway" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="类型" :label-width="formLabelWidth">
          <el-input v-model="modifyForm.type" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="modifyVisible = false">取 消</el-button>
        <el-button type="primary" @click="updateNetwork">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<!-- TODO 修改添加一个dialog -->
<!-- 删出直接就删出了 -->
  <script>
import {
  getNetworkList,
  CreateNetwork,
  DeleteNetwork,
  UpdateNetwork,
  GetNetwork,
} from "@/api/network";

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: "success",
        draft: "gray",
        deleted: "danger",
      };
      return statusMap[status];
    },
  },
  data() {
    return {
      list: null,
      listLoading: true,
      total: 0,
      currentPage: 1,
      pageSize: 10,
      limit: 10,
      input: "",
      select: "1",
      modifyVisible: false,
      createVisible: false,
      formLabelWidth: "120px",
      nodes: [],
      createForm: {
        network_name: "",
        gateway: "",
        type: "",
      },
      modifyForm: {
        id: "",
        network_name: "",
        gateway: "",
        type: 0,
      },
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      var params = {
        limit: this.limit,
        start: this.currentPage,
        network_name: "",
        network_id: "",
        type: "",
      };
      if (this.select == "1") {
        params.network_name = this.input;
      } else if (this.select == "2") {
        params.network_id = this.input;
      } else {
        params.type = this.input;
      }
      getNetworkList(params).then((response) => {
        this.list = response.data.data.item;
        this.total = response.data.data.total;
        this.listLoading = false;
      });
    },
    showNetwork(id) {},
    createNodes(id) {},
    addNetwork() {
      this.createVisible = true;
    },
    createNetwork() {
      var params = this.createForm;
      CreateNetwork(params).then((res) => {
        if (res.data.status == 200) {
          this.$message("sucess");
          this.createForm = {};
          this.createVisible = false;
          this.fetchData();
        } else {
          this.$message("error");
        }
        console.log(res);
      });
    },
    searchNetwork() {
      this.fetchData();
    },
    modifyNetwork(id) {
      this.modifyVisible = true;
      this.modifyForm.id = id;
    },
    updateNetwork() {
      var params = this.modifyForm;
      UpdateNetwork(params).then((res) => {
        if (res.data.status == 200) {
          this.$message("sucess");
          this.modifyForm = {};
          this.modifyVisible = false;
          this.fetchData();
        } else {
          this.$message("error");
        }
        console.log(res);
      });
    },

    deleteNetwork(id) {
      DeleteNetwork({ id }).then((res) => {
        if (res.data.status == 200) {
          this.$message("sucess");
          this.fetchData();
        } else {
          this.$message("error");
        }
      });
    },

    handleSizeChange(val) {
      this.fetchData();
    },
    handleCurrentChange(val) {
      this.fetchData();
    },

    tableHeaderColor() {
      return "background-color: #323546;color: white;font-weight: 550;font-size:16px";
    },
  },
};
</script>
  

<style lang="scss" scoped>
.flex-center {
  display: flex;
  justify-content: center;
  align-items: center;
}
.network-container {
  padding: 20px 40px 0 40px;
  .network-search {
    height: 80px;
    margin-bottom: 10px;
    padding-left: 10px;
    .input-with-select {
      width: 800px;
    }
  }
  .network-box {
    .network-top {
      display: flex;
      justify-content: space-between;
      padding-left: 10px;
      padding-right: 10px;
      align-items: center;
      .network-add {
        width: 120px;
      }
    }
    .network-table {
    }
  }
}
</style>