<template>
  <div id="left">
    <div class="bg-color-black">
      <div class="left-title">{{ this.title }}</div>
      <div>
        <LeftChart :time="this.time"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import LeftChart from '_c/echart/light/left/leftChart'

export default defineComponent({
  name: "left",
  data() {
    return {
      title: "异常处理过程用时分析",
    }
  },
  components: {
    LeftChart
  },
  props: {
    time: {
      type: Array,
      required: true,
    },
  },
})
</script>


<style scoped lang="scss">
$box-height: 900px;
$box-width: 100%;
#left {
  padding: 20px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }

  .left-title {
    padding: 10px 0 0 20px;
    font-size: 24px;
    font-weight: bold;
    color: green;
  }
}
</style>
