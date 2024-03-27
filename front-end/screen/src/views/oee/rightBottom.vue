<template>
  <div id="rightBottom">
    <div class="bg-color-black">
      <div class="chart-title">{{ this.title }}</div>
      <div>
        <RightBottom :oee_trend="this.oee_trend"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import RightBottom from "_c/echart/oee/rightBottom/rightBottomChart";

export default defineComponent({
  name: "rightBottom",
  data() {
    return {
      title: "设备OEE综合趋势",
    }
  },
  props: {
    oee_trend: {
      type: Array,
      required: true,
    },
  },
  components: {
    RightBottom
  },
})
</script>

<style scoped lang="scss">
$box-height: 450px;
$box-width: 100%;
#rightBottom {
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
