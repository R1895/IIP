<template>
  <div class="rb-box">
    <div class="bg-color-black">
      <div class="rb-desc">
        <div class="rb-title">{{ this.title }}</div>
        <div class="rb-detail" @click="navigateToOee">详情</div>
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
  name: "right",
  data() {
    return {
      title: "设备异常",
      config: {
        header: ["设备", "发生时间", "异常原因", "持续时长"],
        data: [
          // ["C104", "01:11:21", '控制故障', "10分8秒"],
          // ["C102", "10:10:11", '操作错误', "2时3分"],
          // ["C101", "12:21:34", '程序出错', "3时12分"],
          // ["C103", "14:11:31", '电气故障', "1时43分"],
          // ["C102", "16:53:27", '机械磨损', "7分14分"],
        ],
        rowNum: 9, //表格行数
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
  props: {
    anomalys: {
      type: Array,
      required: true,
    },
  },
  watch: {
    anomalys: {
      handler: function (newVal) {
        console.log(newVal)
        this.config.data = newVal.map(item => {
          return [item.machine_name, this.formatDateTime(item.date), item.reason, this.calculateTimeDifference(item.date)]
        })
        this.config = {...this.config}
      },
      deep: true
    }
  },
  methods: {
    navigateToOee() {
      this.$router.push({path: '/oee'})
    },
    calculateTimeDifference(time) {
      const timeDiffMilliseconds = Math.abs(new Date(time) - new Date());
      const hours = Math.floor(timeDiffMilliseconds / (1000 * 60 * 60));
      const minutes = Math.floor((timeDiffMilliseconds % (1000 * 60 * 60)) / (1000 * 60));
      if (hours === 0)
        return `${minutes}分钟`;
      else
        return `${hours}小时${minutes}分钟`;
    },
    formatDateTime(dateTimeString) {
      // 创建一个 Date 对象，将时间字符串转换为日期对象
      const dateTime = new Date(dateTimeString);

      // 获取日期、小时和分钟
      const mouth = dateTime.getMonth() + 1;
      const date = dateTime.getDate();
      const hours = dateTime.getHours();
      const minutes = dateTime.getMinutes();

      // 格式化输出
      const formattedDateTime = `${mouth}-${date}  ${hours}:${minutes}`;

      return formattedDateTime;
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 700px;
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

    .rb-detail {
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
      height: 630px;
    }
  }
}
</style>
