<template>
  <div class="rt-box">
    <div class="bg-color-black">
      <div class="rt-desc">
        <div class="rt-title">{{ this.title }}</div>
        <div class="rt-detail" @click="navigateToLight">详情...</div>
      </div>
      <div class="board-content">
        <dv-scroll-board class="dv-scr-board" :config="config"/>
      </div>
    </div>
  </div>
</template>
<script>
import {defineComponent} from 'vue'

export default defineComponent({
  name: "rightTop",
  data() {
    return {
      title: "安灯状态",
      config: {
        header: ["设备", "安灯", "发生时间", "安灯时长", "状态"],
        data: [
          ["C119", "修模", "10:11:21", "1时08分", "<span  class='colorGrass'>加工</span>"],
          ["C107", "修机", "09:10:11", "2时10分", "<span  class='colorGrass'>加工</span>"],

        ],
        rowNum: 5, //表格行数
        headerHeight: 35,
        headerBGC: "#0f1325", //表头
        oddRowBGC: "#0f1325", //奇数行
        evenRowBGC: "#171c33", //偶数行
        index: false,
        columnWidth: [120, 120, 120, 120, 120],
        align: ["center", "center", "center", "center", "center"],
      },
    }
  },
  methods: {
    navigateToLight() {
      this.$router.push({path: '/light'})
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 350px;
$box-width: 622px;
.rt-box {
  padding: 16px;
  padding-top: 20px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
  }

  .rt-desc {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 20px;

    .rt-title {
      padding: 5px 0 0 20px;
      font-size: 24px;
      font-weight: bold;
      color: green;
    }

    .rt-detail {
      cursor: pointer;
      padding: 5px 0 0 20px;
      //padding: 5px 20px 0 0;
      font-size: 20px;
      color: green;
    }
  }

  .board-content {
    border-radius: 10px;
    overflow: hidden;
    display: flex;
    justify-content: center;

    .dv-scr-board {
      width: $box-width -50px;
      height: $box-height - 50px;
      font-size: 20px;

    }
  }

}
</style>