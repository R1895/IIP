<template>
  <div class="microchain-container">
    <div class="microchain-search">
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
          @click="searchMicrochain"
        ></el-button>
      </el-input>
    </div>
    <div class="microchain-box">
      <div class="microchain-top">
        <h2>Microchain</h2>
        <div>
          <el-button
            class="microchain-add"
            type="success"
            @click="addMicrochain(scope.row.id)"
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
        highlight-current-row
        class="microchain-table"
      >
        <el-table-column align="center" label="ID" width="60">
          <template slot-scope="scope">
            {{ scope.$index }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="MCid" width="80">
          <template slot-scope="scope">
            {{ scope.row.id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="MicrochainName" width="180">
          <template slot-scope="scope">
            <span>{{ scope.row.title }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="client_id">
          <template slot-scope="scope">
            {{ scope.row.client_id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="nid">
          <template slot-scope="scope">
            {{ scope.row.nid }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="nodes">
          <template slot-scope="scope">
            {{ scope.row.nodes }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="num">
          <template slot-scope="scope">
            {{ scope.row.num }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="price">
          <template slot-scope="scope">
            {{ scope.row.price }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="client_id">
          <template slot-scope="scope">
            {{ scope.row.client_id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="init_repu">
          <template slot-scope="scope">
            {{ scope.row.init_repu }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="client_id">
          <template slot-scope="scope">
            {{ scope.row.client_id }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="init_storage">
          <template slot-scope="scope">
            {{ scope.row.init_storage }}
          </template>
        </el-table-column>
        <el-table-column align="center" label="init_evil">
          <template slot-scope="scope">
            {{ scope.row.init_evil }}
          </template>
        </el-table-column>

        <el-table-column
          class-name="status-col"
          label="Type"
          width="110"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag :type="scope.row.status | statusFilter">{{
              scope.row.status
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="center" prop="created_at" label="CreatedAt">
          <template slot-scope="scope">
            <i class="el-icon-time" />
            <span>{{ scope.row.created_at }}</span>
          </template>
        </el-table-column>
        <el-table-column align="center" label="Process" width="200">
          <template slot-scope="scope">
            <div>
              <el-button type="primary" @click="modityMicrochain(scope.row.id)"
                >修改</el-button
              >
              <el-button type="danger" @click="deleteMicrochain(scope.row.id)"
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
  
  <!-- TODO 修改添加一个dialog -->
  <!-- 删出直接就删出了 -->
    <script>
import { getShardingList } from "@/api/microchain";

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
      limit: 6,
      network_id: 1,
      input: "",
      select: "1",
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    searchMicrochain() {},
    modityMicrochain() {},
    deleteMicrochain(id) {},
    addMicrochain() {},
    handleSizeChange(val) {
      this.fetchData();
    },
    handleCurrentChange(val) {
      this.fetchData();
    },
    fetchData() {
      this.listLoading = true;
      var params = {
        network_id: this.network_id,
        limit: this.limit,
        start: this.currentPage,
      };
      getShardingList(params).then((response) => {
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
.microchain-container {
  padding: 20px 40px 0 40px;
  .microchain-search {
    height: 80px;
    margin-bottom: 10px;
    padding-left: 10px;
    .input-with-select {
      width: 800px;
    }
  }
  .microchain-box {
    .microchain-top {
      display: flex;
      justify-content: space-between;
      padding-left: 10px;
      padding-right: 10px;
      align-items: center;
      .microchain-add {
        width: 120px;
      }
    }
    .microchain-table {
    }
  }
}
</style>