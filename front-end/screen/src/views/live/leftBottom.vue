<template>
  <div class="lb-box">
    <div class="bg-color-black">
      <div class="lb-content">
        <div class="lb-desc">
          <div class="lb-title">{{ this.title }}</div>
          <div class="lb-detail" @click="navigateToStatus">详情</div>
        </div>
        <div class="grid-content">
          <div v-for="item in items" :key="item.id" class="grid-item" :style="{color: getColor(item.id)}">
            <div class="item-top">
              <countTo :startVal='old_item_num[item.id]' :endVal='item.num'
                       :duration='3000'></countTo>
            </div>
            <div class="item-bottom">{{ item.status }}</div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>
<script>
import {defineComponent} from 'vue'

export default defineComponent({
  name: "leftBottom",
  data() {
    return {
      title: "设备状态",
      items: [
        {id: 0, num: 0, status: '运行'},
        {id: 1, num: 0, status: '等待'},
        {id: 2, num: 0, status: '故障'},
        {id: 3, num: 0, status: '关机'},
      ],
      old_item_num: [0, 0, 0, 0],
    }
  },
  props: {
    status_num: {
      type: Object,
      required: true,
    },
  },
  watch: {
    status_num: {
      handler(newData) {
        this.old_item_num = this.items.map(item => item.num)
        this.items = [
          {id: 0, num: newData.num_of_run, status: '运行'},
          {id: 1, num: newData.num_of_waiting, status: '等待'},
          {id: 2, num: newData.num_of_error, status: '故障'},
          {id: 3, num: newData.num_of_close, status: '关机'},
        ]
      },
      deep: true
    }
  },
  methods: {
    navigateToStatus() {
      this.$router.push({path: '/status'})
    },
    getColor(id) {
      if (id === 3) {
        return '#999999'
      } else if (id === 2) {
        return '#e14a3b'
      } else if (id === 1) {
        return '#F29D38'
      } else if (id === 0) {
        return '#52962A'
      }
    }
  }
})
</script>

<style scoped lang="scss">
$box-height: 350px;
$box-width: 622px;
.lb-box {
  padding: 16px;
  padding-top: 20px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
  }

  .lb-content {
    display: flex;
    height: 300px;
    flex-direction: column;
    margin-bottom: 20px;

    .lb-desc {
      display: flex;
      flex: 1;
      flex-direction: row;
      justify-content: space-between;
      margin-bottom: 20px;

      .lb-title {
        padding: 5px 0 0 20px;
        font-size: 24px;
        font-weight: bold;
        color: green;
      }

      .lb-detail {
        cursor: pointer;
        padding: 5px 0 0 20px;
        //padding: 5px 20px 0 0;
        font-size: 20px;
        color: green;
      }
    }

    .grid-content {
      flex: 4;
      display: grid;
      grid-template-columns: repeat(4, 1fr);

      .grid-item {
        display: flex;
        flex-direction: column;

        .item-top {
          flex: 1;
          display: flex;
          font-weight: bold;
          font-size: 42px;
          align-items: center; /* 在纵轴上居中 */
          justify-content: center; /* 在水平轴上居中 */
          width: 100%;
          height: 100%;
        }

        .item-bottom {
          flex: 1;
          display: flex;
          align-items: flex-start; /* 在纵轴上居中 */
          justify-content: center; /* 在水平轴上居中 */
          font-size: 22px;
          font-weight: bold;
          width: 100%;
          height: 100%;
        }
      }
    }
  }

}
</style>
