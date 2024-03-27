<template>
  <div class="stock" ref="appRef">
    <div class="bg">

      <MyHeader :title="title"></MyHeader>
      <dv-decoration-6 style="width:100%;height:10px;"/>
      <div class="stock-container">
        <div class="left-box">
          <div v-for="item in itemsData" :key="item.name" class="item">
            <div class="item-title">
              <dv-decoration-3 style="width:50px;height:30px;"/> &nbsp;&nbsp;
              {{ item.name }}
            </div>
            <div class="item-detail">
              <div v-for="(model, index) in item.models" :key="model" class="detail-content">
                <!--              <div class="item-img">图片</div>-->
                <img :src="item.images[index]" width="120px" height="120px">
                <div class="item-type">{{ model }}</div>
                <div class="item-stock_num">库存数: {{ item.stock[index] }}</div>
              </div>
            </div>
            <dv-decoration-10 style="width:100%;height:5px;"/>
          </div>
          <div class="block">
            <el-pagination
                layout="prev, pager, next"
                :total="this.total_size"
                :page-size="this.page_size"
                @current-change="handleCurrentChange">
            </el-pagination>
          </div>
        </div>
        <div class="right-box">
          <div class="workpiece-table">
            <el-table
                :data="this.stock_details"
                border
                style="width: 100%">
              <el-table-column
                  prop="workpiece_name"
                  label="工件名称"
                  width="180">
              </el-table-column>
              <el-table-column
                  prop="workpiece_type"
                  label="工件类型"
                  width="180">
              </el-table-column>
              <el-table-column
                  prop="stock_number"
                  label="库存数量"
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
                :total="this.table_total_size"
                :page-size="this.table_page_size"
                @current-change="handleTableChange">
            </el-pagination>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MyHeader from "../../components/header/myHeader.vue";
import drawMixin from "@/utils/drawMixin";
// import mock from "@/api/mock";

export default defineComponent({
  name: "stock",
  mixins: [drawMixin],
  data() {
    return {
      title: "库存看板",
      itemsData: [],
      itemsData1: [
        {
          name: '联轴器',
          models: ['刚性联轴器', '挠性联轴器'],
          images: [require('@/assets/img/001.png'), require("@/assets/img/002.png")],
          stock: [100, 225]
        },
        {
          name: '轴承',
          models: ['滚珠轴承', '圆柱滚子轴承', '滚针轴承'],
          images: [require("@/assets/img/003.png"), require("@/assets/img/004.png"), require("@/assets/img/005.png")],
          stock: [218, 412, 213]
        },

        // Add more items as needed
      ],
      itemsData2: [
        {
          name: '联轴器',
          models: ['刚性联轴器', '挠性联轴器'],
          images: [require('@/assets/img/001.png'), require("@/assets/img/002.png")],
          stock: [100, 225]
        },
        // Add more items as needed
      ],
      total_size: 3,
      page_size: 2,
      current_page: 1,

      table_total_size: 5,
      table_page_size: 4,
      table_current_page: 1,

      stock_details: [],
      stock_details1: [{
        workpiece_name: '联轴器',
        workpiece_type: '刚性联轴器',
        stock_number: 100,
        description: '无'
      }, {
        workpiece_name: '联轴器',
        workpiece_type: '挠性联轴器',
        stock_number: 225,
        description: '无'
      }, {
        workpiece_name: '轴承',
        workpiece_type: '滚珠轴承',
        stock_number: 218,
        description: '无'
      }, {
        workpiece_name: '轴承',
        workpiece_type: '圆柱滚子轴承',
        stock_number: 412,
        description: '无'
      }],
      stock_details2: [{
        workpiece_name: '联轴器',
        workpiece_type: '刚性联轴器',
        stock_number: 100,
        description: '无'
      }, {
        workpiece_name: '联轴器',
        workpiece_type: '挠性联轴器',
        stock_number: 225,
        description: '无'
      }],

    }
  },
  components: {
    MyHeader,
  },
  mounted() {
    setInterval(() => {
      this.fetchData();
    }, 3000);
    setInterval(() => {
      this.fetchTableData();
    }, 3000);
  },
  created() {
    this.fetchData();
    this.fetchTableData();
  },
  methods: {
    fetchData() {
      // stock.getStockInfoAnalysis().then(response => {
      //   console.log(response);
      // })
      if (this.current_page == 1) {
        this.itemsData = this.itemsData1;
      } else if (this.current_page == 2) {
        this.itemsData = this.itemsData2;
      }
    },
    fetchTableData() {
      // stock.getStockInfoAnalysis().then(response => {
      //   console.log(response);
      // })
      if (this.table_current_page == 1) {
        this.stock_details = this.stock_details1;
      } else if (this.table_current_page == 2) {
        this.stock_details = this.stock_details2;
      }
    },
    handleCurrentChange(val) {
      this.current_page = val;
      this.fetchData();
    },
    handleTableChange(val) {
      this.table_current_page = val;
      this.fetchTableData();
    },

  }
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/stock.scss';
</style>
