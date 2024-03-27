<template>
  <div id="rightTopRight">
    <div class="bg-color-black">
      <div class="chart-title">{{ this.title }}</div>
      <div>
        <RightTopRightChart :loss_reason="this.loss_reason"/>
      </div>
    </div>
  </div>
</template>


<script>
import {defineComponent} from 'vue'
import RightTopRightChart from '_c/echart/oee/rightTop/rightTopRightChart'

export default defineComponent({
  name: "rightTopRight",
  data() {
    return {
      title: "设备运行时间损失原因分析",
    }
  },
  props: {
    loss_reason: {
      type: Array,
      required: true,
    },
  },
  components: {
    RightTopRightChart
  },
})
</script>

<style scoped lang="scss">
$box-height: 450px;
$box-width: 100%;
#rightTopRight {
  padding: 20px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }

  .chart-title {
    padding: 10px 0 0 20px;
    font-size: 24px;
    font-weight: bold;
    color: green;
  }
}
</style>
