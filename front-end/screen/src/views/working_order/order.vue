<template>
  <div class="order" ref="appRef">
    <div class="bg">
      <MyHeader :title="title"></MyHeader>

      <div class="order-container">
        <div>
          <dv-border-box12>
            <div class="top">
              <div class="order-total">
                <div v-for="item in desc_tasks" :key="item.id" class="order-total-item">
                  <div class="task-name">{{ item.task_name }}</div>
                  <div class="task-info">
                    <div class="info-label">任务工件数:</div>
                    <div>{{ item.plan_item }}</div>
                  </div>
                  <div class="task-info">
                    <div class="info-label">已加工工件:</div>
                    <div>{{ item.processed_item }}</div>
                  </div>
                </div>
              </div>
              <div class="block">
                <el-pagination
                    layout="prev, pager, next"
                    :total="total_page"
                    :page-size="desc_page_size"
                    @current-change="handleDescChange">
                </el-pagination>
              </div>
              <div class="placeholder_block"></div>
            </div>
          </dv-border-box12>
        </div>

        <div class="table-desc">
          <div class="table-title">{{ table_title }}</div>
          <div class="table-legend">
            <div class="color-info">
              <div class="color-label">进行中</div>
              <div class="color-block-success"></div>
            </div>
            <div class="color-info">
              <div class="color-label">将截止</div>
              <div class="color-block-warning"></div>
            </div>
            <div class="color-info">
              <div class="color-label">已逾期</div>
              <div class="color-block-danger"></div>
            </div>
          </div>
        </div>

        <div class="order-table">
          <el-table
              :data="this.detail_tasks"
              border
              style="width: 100%"
              :row-class-name="tableRowClassName">
            <el-table-column
                prop="task_name"
                label="任务名称"
                width="180">
            </el-table-column>
            <el-table-column
                prop="workshop_id"
                label="加工车间"
                width="180">
            </el-table-column>
            <el-table-column
                prop="start_date"
                label="开始日期"
                width="180">
            </el-table-column>
            <el-table-column
                prop="effective_time"
                label="截止日期"
                width="180">
            </el-table-column>
            <el-table-column
                prop="plan_item"
                label="计划工件数"
                width="180">
            </el-table-column>
            <el-table-column
                prop="processed_item"
                label="已加工工件数"
                width="180">
            </el-table-column>
            <el-table-column
                prop="is_finished"
                label="是否完成"
                width="180">
            </el-table-column>
            <el-table-column
                prop="description"
                label="其他描述">
            </el-table-column>
          </el-table>
        </div>
        <div class="block">
          <el-pagination
              layout="prev, pager, next"
              :total="total_page"
              :page-size="this.detail_page_size"
              @current-change="handleDetailChange">
          </el-pagination>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import MyHeader from "../../components/header/myHeader.vue";
import drawMixin from "@/utils/drawMixin";
import order from "@/api/order";
// import mock from "@/api/mock";

export default {
  name: "order",
  mixins: [drawMixin],
  data() {
    return {
      title: "工单列表",
      table_title: "任务清单",
      detail_tasks: [],
      desc_tasks: [],
      detail_current_page: 1,
      desc_current_page: 1,
      desc_page_size: 4,
      detail_page_size: 12,
      total_page: 12,
    }
  },
  components: {
    MyHeader,
  },
  mounted() {
    // setInterval(() => {
    //   this.fetchDescData();
    // }, 3000);
    // setInterval(() => {
    //   this.fetchDetailData();
    // }, 3000);
  },
  created() {
    // this.fetchDescData();
    // this.fetchDetailData();
  },
  computed: {},
  methods: {
    tableRowClassName({row}) {
      // console.log(row,'--row--')
      if (row.task_type_id == 1) {
        return 'warning-row';
      } else if (row.task_type_id == 2) {
        return 'success-row';
      } else if (row.task_type_id == 3) {
        return 'expired-row';
      }
      return '';
    },
    handleDescChange(val) {
      this.desc_current_page = val;
      this.fetchDescData();
    },
    handleDetailChange(val) {
      this.detail_current_page = val;
      this.fetchDetailData();
    },
    fetchDescData() {
      order.getOrderInfoAnalysis(this.desc_page_size, this.desc_current_page)
          .then(response => {
            // console.log(response)
            this.total_page = response.data.data.total;
            this.desc_tasks = response.data.data.item;
            // console.log(this.desc_tasks)
          })
          .catch(error => {
            console.error(error);
          });
    },
    fetchDetailData() {
      order.getOrderInfoAnalysis(this.detail_page_size, this.detail_current_page)
          .then(response => {
            this.total_page = response.data.data.total;
            this.detail_tasks = response.data.data.item;
          })
          .catch(error => {
            console.error(error);
          });
    },

  },

}
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}

.el-table .expired-row {
  background: #f8e5ec;
}
</style>
<style scoped lang="scss">
@import '../../assets/scss/order.scss';
</style>



