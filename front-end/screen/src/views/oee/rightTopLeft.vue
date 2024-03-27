<template>
  <div id="rightTopLeft">
    <div class="bg-color-black">
      <div class="chart-title">{{ this.title }}</div>
      <div>
        <RightTopLeftChart :daily_loss="this.daily_loss"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import RightTopLeftChart from '_c/echart/oee/rightTop/rightTopLeftChart'

export default defineComponent({
  name: "rightTopLeft",
  data() {
    return {
      title: "设备运行时长每日损失量",
    }
  },
  props: {
    daily_loss: {
      type: Array,
      required: true,
    },
  },
  components: {
    RightTopLeftChart
  },
})
</script>

<style scoped lang="scss">
$box-height: 450px;
$box-width: 100%;
#rightTopLeft {
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
