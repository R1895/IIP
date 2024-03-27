<template>
  <div id="right">
    <div class="bg-color-black">
      <div class="right-title">{{ this.title }}</div>
      <div>
        <RightChart :number="this.number"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import RightChart from "_c/echart/light/right/rightChart/index.vue";

export default defineComponent({
  name: "right",
  data() {
    return {
      title: "异常处理类别数量分布",
    }
  },
  props: {
    number: {
      type: Array,
      required: true,
    },
  },
  components: {RightChart}
})
</script>

<style scoped lang="scss">
$box-height: 900px;
$box-width: 100%;
#right {
  padding: 20px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }

  .right-title {
    padding: 10px 0 0 20px;
    font-size: 24px;
    font-weight: bold;
    color: green;
  }
}
</style>
