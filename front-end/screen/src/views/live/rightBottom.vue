<template>
  <div class="rb-box">
    <div class="bg-color-black">
      <div class="rb-desc">
        <div class="rb-title">{{ this.title }}</div>
        <div class="rb-detail" @click="navigateToOee">详情...</div>
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
  name: "rightBottom",
  data() {
    return {
      title: "生产异常",
      config: {
        header: ["设备", "发生时间", "异常原因", "持续时长"],
        data: [
          ["C104", "01:11:21", "设备故障", "10分8秒"],
          ["C102", "10:10:11", "换人", "2时3分"],
          ["C101", "12:21:34", "换料", "3时12分"],
          ["C103", "14:11:31", "待人", "1时43分"],
          ["C102", "16:53:27", "程序调整", "7分14分"],
        ],
        rowNum: 4, //表格行数
        headerHeight: 35,
        headerBGC: "#0f1325", //表头
        oddRowBGC: "#0f1325", //奇数行
        evenRowBGC: "#171c33", //偶数行
        index: false,
        columnWidth: [150, 150, 150, 150],
        align: ["center", "center", "center", "center"],
      },
    }
  },
  methods: {
    navigateToOee() {
      this.$router.push({path: '/oee'})
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 350px;
$box-width: 622px;
.rb-box {
  padding: 16px;
  padding-top: 20px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
  }

  .rb-desc {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 20px;
    .rb-title {
      padding: 5px 0 0 20px;
      font-size: 24px;
      font-weight: bold;
      color: green;
    }
    .rb-detail{
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
      height: 275px;
    }
  }
}
</style>