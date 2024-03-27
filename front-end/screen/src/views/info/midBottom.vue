<template>
  <div id="midBottom">
    <div class="bg-color-black">
      <div class=" top-box">
        <div class="chart-title">{{ this.title }}</div>
        <el-select v-model="value" placeholder="近七天" @change="handleSelectChange">
          <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
          </el-option>
        </el-select>
      </div>
      <div>
        <MidBottomChart :env="env"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MidBottomChart from "_c/echart/info/mid/midBottomChart/index.vue";

export default defineComponent({
  name: "midBottom",
  data() {
    return {
      title: "环境指标折线图",
      options: [{
        value: '7',
        label: '近七天'
      }, {
        value: '30',
        label: '近一月'
      }],
      value: '7',
    }
  },
  props: {
    env: {
      type: Array,
      required: true,
    },
  },
  components: {
    MidBottomChart
  },
  methods: {
    handleSelectChange() {
      this.$emit('select-change', this.value);
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 375px;
$box-width: 100%;
#midBottom {
  padding: 10px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }

  .top-box {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .chart-title {
      padding: 0 0 0 20px;
      font-size: 24px;
      font-weight: bold;
      color: green;
    }
  }
}
</style>
