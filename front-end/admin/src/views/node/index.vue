<template>
  <div class="node-container">
    <div class="node-search">
      <h2>节点检索</h2>
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
          @click="searchNode"
        ></el-button>
      </el-input>
    </div>
    <div class="node-box">
      <div class="node-top">
        <h2>Node</h2>
        <div>
          <el-button
            class="node-add"
            type="success"
            @click="addNode(scope.row.id)"
            >新增</el-button
          >
        </div>
      </div>
      <el-table
        v-loading="listLoading"
        :data="list"
        :header-cell-style="tableHeaderColor"
        element-loading-text="Loading"
 
        fit
        highlight-current-row
        class="node-table"
      >
      <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="reputation">
                <span>{{ props.row.reputation }}</span>
              </el-form-item>
              <el-form-item label="hashrate">
                <span>{{ props.row.hashrate }}</span>
              </el-form-item>
              <el-form-item label="storage">
                <span>{{ props.row.storage }}</span>
              </el-form-item>
              <el-form-item label="bandwidth">
                <span>{{ props.row.bandwidth }}</span>
              </el-form-item>
              <el-form-item label="bandwidth">
                <span>{{ props.row.bandwidth }}</span>
              </el-form-item>
              <el-form-item label="evil_rate">
                <span>{{ props.row.evil_rate }}</span>
              </el-form-item>
              <el-form-item label="crash_rate">
                <span>{{ props.row.crash_rate }}</span>
              </el-form-item>

              <el-form-item label="pub_key">
                <span>{{ props.row.pub_key }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>

        <el-table-column align="center" label="NodeId" width="80">
          <template slot-scope="scope">
            {{ scope.row.id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="NodeName">
          <template slot-scope="scope">
            <span>{{ scope.row.node_name }}</span>
          </template>
        </el-table-column>
        <!-- <el-table-column align="center" label="NetworkId">
          <template slot-scope="scope">
            <span>{{ scope.row.network_id }}</span>
          </template>
        </el-table-column> -->
        <el-table-column align="center" label="Ip">
          <template slot-scope="scope">
            {{ scope.row.ip }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="Port">
          <template slot-scope="scope">
            {{ scope.row.port }}
          </template>
        </el-table-column>

        <el-table-column
          class-name="status-col"
          label="Status"
          width="110"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag :type="scope.row.status | statusFilter">{{
              scope.row.status
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          prop="created_at"
          label="CreatedAt"
          width="150"
        >
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.created_at }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="Process" width="280">
          <template slot-scope="scope">
            <div>
              <el-button type="primary" @click="modityNode(scope.row.id)"
                >修改</el-button
              >
              <el-button type="danger" @click="deleteNode(scope.row.id)"
                >删除</el-button
              >
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
  </div>
</template>

    <script>
import { getNodeList } from "@/api/node";

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
      networkId: 1,
      currentPage: 1,
      pageSize: 10,
      limit: 10,
      input: "",
      select: "1",
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    searchNode() {},
    modityNode() {},
    deleteNode(id) {},
    addNode() {},
    handleSizeChange(val) {
      this.fetchData();
    },
    handleCurrentChange(val) {
      this.fetchData();
    },
    fetchData() {
      this.listLoading = true;
      var params = {
        limit: this.limit,
        start: this.currentPage,
        network_id: this.networkId,
      };
      getNodeList(params).then((response) => {
        this.list = response.data.data.item;
        this.total = response.data.data.total;
        this.listLoading = false;
      });
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
.node-container {
  padding: 20px 40px 0 40px;
  .node-search {
    height: 80px;
    margin-bottom: 10px;
    padding-left: 10px;
    .input-with-select {
      width: 800px;
    }
  }
  .node-box {
    .node-top {
      display: flex;
      justify-content: space-between;
      padding-left: 10px;
      padding-right: 10px;
      align-items: center;
      .node-add {
        width: 120px;
      }
    }
    .node-table {
    }
  }
}
.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>