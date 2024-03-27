<template>
  <div id="left">
    <div class="bg-color-black">
      <div class="left-title">{{ this.title }}</div>
      <div>
        <LeftChart :rank="this.rank"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import LeftChart from "_c/echart/rank/left/leftChart/index.vue";

export default defineComponent({
  name: "right",
  data() {
    return {
      title: "本月工时排行榜",
    }
  },
  props: {
    rank: {
      type: Array,
      required: true,
    },
  },
  components: {LeftChart}
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
