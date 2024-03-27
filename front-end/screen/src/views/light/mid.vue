<template>
  <div class="mid-container">
    <div class="bg-color-black">
      <div class="mid-desc">
        <div class="mid-title">{{ this.title }}</div>
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
  name: "mid",
  data() {
    return {
      title: "设备异常",
      config: {
        header: ["设备", "发生时间", "等待时长", "处理时长", "状态", "类型"],
        data: [
          ["C119", "10:11:21", "1时08分", "1时08分", "<span  class='colorGrass'>等待中</span>", '操作错误'],
          ["C107", "09:10:11", "58分", "23分", "<span  class='colorGrass'>处理中</span>", '程序出错'],
          ["C108", "11:21:47", "1时28分", "28分", "<span  class='colorGrass'>处理中</span>", '机械磨损'],
          ["C114", "13:10:34", "2时10分", "---", "<span  class='colorGrass'>等待中</span>", '电气故障'],
          ["C121", "15:28:41", "42分", "24分", "<span  class='colorGrass'>等待中</span>", '程序出错'],
          ["C117", "17:10:51", "1时04分", "---", "<span  class='colorGrass'>等待中</span>", '控制故障'],
          ["C109", "19:33:23", "1时24分", "1时19分", "<span  class='colorGrass'>处理中</span>", '电气故障'],
          ["C110", "21:02:11", "2时10分", "---", "<span  class='colorGrass'>处理中</span>", '操作错误'],
        ],
        rowNum: 10, //表格行数
        headerHeight: 35,
        headerBGC: "#0f1325", //表头
        oddRowBGC: "#0f1325", //奇数行
        evenRowBGC: "#171c33", //偶数行
        index: false,
        columnWidth: [100, 100, 100, 100, 100, 100],
        align: ["center", "center", "center", "center", "center", "center"],
      },
    }
  },
  methods: {}
})
</script>

<style scoped lang="scss">
$box-height: 900px;
$box-width: 100%;
.mid-container {
  padding: 20px 16px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
  }

  .mid-desc {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 20px;

    .mid-title {
      padding: 5px 0 0 20px;
      font-size: 24px;
      font-weight: bold;
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
      height: $box-height - 100px;
      font-size: 20px;

    }
  }

}
</style>
