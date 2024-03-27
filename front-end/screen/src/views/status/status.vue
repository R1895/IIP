<template>
  <div class="status" ref="appRef">
    <div class="bg">

      <MyHeader :title="title"></MyHeader>
      <div class="status-container">
        <!--        <dv-border-box12>-->
        <div class="select-container">
          <el-select v-model="workshop_value" placeholder="车间编号" @change="handleSelect">
            <el-option
                v-for="item in workshop_options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
          <el-select v-model="production_line_value" placeholder="产线编号">
            <el-option
                v-for="item in production_line_options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
          </el-select>
        </div>
        <div class="grid-container">
          <div v-for="(items,rowIndex) in rows" class="grid-row" :key="rowIndex">
            <div v-for="(item,colIndex) in items" class="grid-item" :key="colIndex">
              <div :class="getTopBarStatusClass(item.machine_status)">
                <span>{{ item.machine_name }}</span>
                <span v-if="item.procrdure_name"> &nbsp;{{ item.procrdure_name }}</span>
                <span v-if="item.worker_name"> {{ item.worker_name }}</span>
                <span v-if="item.task_id"> Task{{ item.task_id }}</span>
              </div>
              <div :class="getPercentageStatusClass(item.machine_status)">
                <div class="percent_chart">
                  <vue-awesome-progress
                      circle-color="#7a7777"
                      :circle-radius="60"
                      :circle-width="5"
                      :line-width="4"
                      line-color="white"
                      :font-size="22"
                      font-color="white"
                      :point-radius="0"
                      :duration="1.5"
                      :start-deg="0"
                      :percentage="(item.percent * 100).toFixed(1)"
                      :show-text="true"
                      easing="0,0,1,1"
                  />
                </div>
                <span>{{ getStatusName(item.machine_status) }}{{ calculateTimeDifference(item.time) }}</span>
              </div>
              <!--              <div :class="getBottomBarStatusClass(item.machine_status)">-->
              <!--                <span v-if="item.machine_status!==offStatus"> {{ getStatusName(item.machine_status) }} {{-->
              <!--                    calculateTimeDifference(item.time)-->
              <!--                  }}</span>-->
              <!--              </div>-->
            </div>
          </div>
        </div>
        <div class="block">
          <el-pagination
              layout="prev, pager, next"
              :total="total_page"
              :page-size="this.page_size"
              @current-change="handleCurrentChange">
          </el-pagination>
        </div>
        <!--        </dv-border-box12>-->
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MyHeader from "../../components/header/myHeader.vue";
import drawMixin from "@/utils/drawMixin";
// import mock from "@/api/mock";
import status from "@/api/status";
import VueAwesomeProgress from "vue-awesome-progress"

export default defineComponent({
  name: "status",
  mixins: [drawMixin],
  data() {
    return {
      title: "生产状态看板",
      statusData: [],
      totalItems: null,
      itemsPerRow: 5,
      offStatus: "关机",
      listLoading: false,
      workshop_options: [{
        value: '1',
        label: 'ws001'
      }, {
        value: '2',
        label: 'ws002'
      }, {
        value: '3',
        label: 'ws003'
      }, {
        value: '4',
        label: 'ws004'
      }, {
        value: '5',
        label: 'ws005'
      }],
      workshop_value: '',
      production_line_options: [{
        value: '1',
        label: 'L001'
      }, {
        value: '2',
        label: 'L002'
      }, {
        value: '3',
        label: 'L003'
      }, {
        value: '4',
        label: 'L004'
      }, {
        value: '5',
        label: 'L005'
      }],
      production_line_value: '',
      total_page: 0,
      page_size: 20,
      current_page: 1,
    }
  },
  components: {
    MyHeader, VueAwesomeProgress
  },
  mounted() {
    // setInterval(() => {
    //   this.fetchData();
    // }, 3000);
  },
  created() {
    this.fetchData();
  },
  computed: {
    rows() {
      this.totalItems = this.statusData.length;
      const rows = [];
      let remainingItems = this.totalItems;
      while (remainingItems > this.itemsPerRow) {
        rows.push(this.statusData.slice(rows.length * this.itemsPerRow, rows.length * this.itemsPerRow + this.itemsPerRow));
        remainingItems -= this.itemsPerRow;
      }
      if (remainingItems > 0)
        rows.push(this.statusData.slice(rows.length * this.itemsPerRow, rows.length * this.itemsPerRow + remainingItems));
      return rows;
    }
  },
  methods: {
    handleSelect() {
      this.fetchData();
    },
    getStatusName(status) {
      // 根据不同的 status 返回对应的类名
      switch (status) {
        case 1:
          return "运行";
        case 2:
          return "等待";
        case 3:
          return "离线";
        default:
          return "故障";
      }
    },
    getTopBarStatusClass(status) {
      // 根据不同的 status 返回对应的类名
      switch (status) {
        case 1:
          return "top-bar";
        case 2:
          return "top-bar-orange";
        case 3:
          return "top-bar-gray";
        default:
          return "top-bar-red";
      }
    },
    getPercentageStatusClass(status) {
      // 根据不同的 status 返回对应的类名
      switch (status) {
        case 1:
          return "percentage"; // 默认类名
        case 2:
          return "percentage-orange";
        case 3:
          return "percentage-gray";
        default:
          return "percentage-red";
      }
    },
    getBottomBarStatusClass(status) {
      // 根据不同的 status 返回对应的类名
      switch (status) {
        case 1:
          return "bottom-bar";
        case 2:
          return "bottom-bar-orange";
        case 3:
          return "bottom-bar-gray";
        default:
          return "bottom-bar-red"; // 默认类名
      }
    },
    calculateTimeDifference(time) {
      const timeDiffMilliseconds = Math.abs(new Date(time) - new Date());
      const hours = Math.floor(timeDiffMilliseconds / (1000 * 60 * 60));
      const minutes = Math.floor((timeDiffMilliseconds % (1000 * 60 * 60)) / (1000 * 60));
      if (hours === 0)
        return `${minutes}分钟`;
      else
        return `${hours}小时${minutes}分钟`;
    },
    fetchData() {
      status.getProductionStatus(this.page_size, this.current_page, this.workshop_value)
          .then(response => {
            this.statusData = [];
            this.total_page = 0;
            console.log(response)
            if (response.data.data !== null) {
              this.statusData = response.data.data;
              this.total_page = response.data.data.length;
            }
          })
          .catch(error => {
            // 处理错误
            console.error(error);
          });
    },
    handleCurrentChange(val) {
      this.current_page = val;
      this.fetchData();
    },
  }
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/status.scss';
</style>
