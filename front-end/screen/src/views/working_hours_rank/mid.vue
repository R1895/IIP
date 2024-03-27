<template>
  <div class="mid">
    <div class="bg-color-black">
      <div class="mid-desc">
        <div class="mid-title">{{ this.title }}</div>
      </div>
      <div class="grid-box">
        <div class="grid-item">
          <div class="blank-icon"></div>
          <div class="item-info">
            <div class="item-desc">姓名</div>
          </div>
          <div class="item-info">
            <div class="item-desc">工时</div>
          </div>
          <div class="item-info">
            <div class="item-desc">排名</div>
          </div>
        </div>
        <div v-for="item in items" :key="item.id" class="grid-item">
          <div class="el-icon-user"></div>
          <div class="item-info">
            <div class="item-content">{{ item.name }}</div>
          </div>
          <div class="item-info">
            <div class="item-content">{{ item.hours }}</div>
          </div>
          <div class="item-info">
            <div class="item-content">{{ item.rank }}</div>
          </div>
        </div>
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
      title: "年度排行",
      items: [
        // {id: 1, name: '李一帆', hours: '1723时28分', rank: 1},
        // {id: 2, name: '徐子航', hours: '1662时43分', rank: 2},
        // {id: 3, name: '陈程', hours: '1623时52分', rank: 3},
        // {id: 4, name: '王刚', hours: '1602时41分', rank: 4},
        // {id: 5, name: '张伟', hours: '1600时01分', rank: 5},
      ]
    }
  },
  props: {
    rank: {
      type: Array,
      required: true,
    },
  },
  watch: {
    rank: {
      handler: function (newVal) {
        // console.log(newVal)
        this.items = []
        newVal.forEach((item, index) => {
          this.items.push({id: index, name: item.WorkerID, hours: item.TotalTime, rank: index + 1})
        })
        this.items = this.items.slice(0, 7)
      },
      deep: true
    }
  },
  methods: {}
})
</script>

<style scoped lang="scss">
$box-height: 900px;
$box-width: 100%;
.mid {
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

  .grid-box {
    display: grid;
    grid-template-rows: repeat(8, 100px);
    margin: 20px 20px;
    height: 800px;
    width: 100%;

    .grid-item {
      border-radius: 10px;
      margin: 10px 20px;
      display: flex;
      flex-direction: row;

      .el-icon-user {
        display: flex;
        justify-content: center; /* 在水平轴上居中 */
        align-items: center; /* 在纵轴上居中 */
        font-size: 30px;
        font-weight: bold;
        color: #ffffff;
      }

      .blank-icon {
        width: 30px;
        height: 30px;
      }

      .item-info {
        flex: 1;
        display: flex;
        flex-direction: column;
        margin-left: 10px;
        justify-content: center; /* 在水平轴上居中 */
        align-items: center; /* 在纵轴上居中 */
        color: #ffffff;

        .item-desc {
          flex: 1;
          font-size: 24px;
          font-weight: bold;
          align-items: center; /* 在纵轴上居中 */
          display: flex;
        }

        .item-content {
          flex: 1;
          font-size: 22px;
          font-weight: bold;
          align-items: center; /* 在纵轴上居中 */
          display: flex;
        }
      }
    }
  }

}
</style>
