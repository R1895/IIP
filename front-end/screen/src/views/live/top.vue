<template>
  <div class="top-box">
    <div class="grid-box">
      <div v-for="item in items" :key="item.id" class="grid-item" :style="{backgroundColor: getBColor(item.id)}">
        <div class="item-top">{{ item.status }}</div>
        <div class="item-bottom">
          <countTo :startVal='old_item_num[item.id]' :endVal='item.num'
                   :duration='3000'></countTo>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'

export default defineComponent({
  name: "top",
  data() {
    return {
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
    getBColor(id) {
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

.top-box {
  height: 180px;
  width: 100%;
  display: flex;
  justify-content: center; /* 在水平轴上居中 */
  align-items: center; /* 在纵轴上居中 */
  .grid-box {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    margin: 20px 20px;
    height: 90%;
    width: 1800px;

    .grid-item {
      border-radius: 10px;
      margin: 10px 20px;
      background-color: #999999;
      display: flex;
      flex-direction: column;

      .item-top {
        flex: 1;
        display: flex;
        margin-left: 10px;
        justify-content: flex-start; /* 在水平轴上居中 */
        align-items: center; /* 在纵轴上居中 */
        font-size: 20px;
        font-weight: bold;
        color: white;
      }

      .item-bottom {
        flex: 2;
        display: flex;
        justify-content: center; /* 在水平轴上居中 */
        align-items: center; /* 在纵轴上居中 */
        font-size: 36px;
        font-weight: bold;
        color: white;
      }
    }
  }
}
</style>
