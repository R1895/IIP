<template>
  <div class="mid-box">
    <div class="bg-color-black">
      <div class="mid-desc">
        <div class="mid-title">{{ this.title }}</div>
        <div class="mid-detail" @click="navigateToStatus">详情</div>
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
      title: "设备运行状态",
      config: {
        header: ["设备", "计划数量", "加工数", "进度", "运行状态"],
        data: [
          // ["C101", "1000", "700", "70%", "<span  class='colorGrass'>加工</span>"],
          // ["C102", "2000", "1800", "90%", "<span  class='colorGrass'>加工</span>"],
          // ["C103", "3000", "2300", "76%", "<span  class='colorGrass'>加工</span>"],
          // ["C104", "1100", "900", "81%", "<span  class='colorGrass'>加工</span>"],
          // ["C105", "1300", "1280", "97%", "<span  class='colorGrass'>加工</span>"],
          // ["C106", "2400", "2100", "87%", "<span  class='colorGrass'>加工</span>"],
          // ["C107", "800", "700", "82%", "<span  class='colorGrass'>加工</span>"],
          // ["C108", "780", "670", "86%", "<span  class='colorRed'>故障</span>"],
          // ["C109", "2100", "1900", "91%", "<span  class='colorRed'>故障</span>"],
          // ["C110", "930", "890", "95%", "<span  class='colorGrass'>加工</span>"],
        ],
        rowNum: 9, //表格行数
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
  props: {
    machines: {
      type: Array,
      required: true,
    },
  },
  watch: {
    machines: {
      handler: function (newVal) {
        console.log(newVal)
        this.config.data = newVal.map(item => {
          // if(item.machine_status == '运行') {
          //   item.machine_status = "<span  class='colorGrass'>运行</span>"
          // } else if(item.machine_status == '等待') {
          //   item.machine_status = "<span  class='colorOrange'>等待</span>"
          // } else {
          //   item.machine_status = "<span  class='colorRed'>故障</span>"
          // }
          return [item.machine_name, item.workpieces_num, item.processed_num, item.progress+'%', item.machine_status]
        })
        this.config = {...this.config}
      },
      deep: true
    }
  },
  methods: {
    navigateToStatus() {
      this.$router.push({path: '/status'})
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 700px;
$box-width: 622px;
.mid-box {
  padding: 16px;
  padding-top: 20px;
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

    .mid-detail {
      padding: 5px 0 0 20px;
      cursor: pointer;
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
      height: 630px;
    }
  }
}
</style>
