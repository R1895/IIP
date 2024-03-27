<template>
  <div id="midTop">
    <div class="bg-color-black">
      <div class="chart-title">{{ this.title }}</div>
      <div>
        <MidTopChart :safe="safe"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MidTopChart from '_c/echart/info/mid/midTopChart'

export default defineComponent({
  name: "midTop",
  data() {
    return {
      title: "安全事故分布图",
    }
  },
  components: {
    MidTopChart
  },
  props: {
    safe: {
      type: Object,
      required: true,
    },
  },
})
</script>


<style scoped lang="scss">
$box-height: 375px;
$box-width: 100%;
#midTop {
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
    margin-bottom: 20px;
  }

}
</style>
