<template>
  <div id="right">
    <div class="bg-color-black">
      <div class="right-title">{{ this.title }}</div>
      <div>
        <RightChart :rank="this.rank"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import RightChart from "_c/echart/rank/right/rightChart/index.vue";

export default defineComponent({
  name: "right",
  data() {
    return {
      title: "今日工时排行榜",
    }
  },
  props: {
    rank: {
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
