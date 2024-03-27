<template>
  <div class="top-box">
    <div class="grid-box">
      <div v-for="item in items" :key="item.id" class="grid-item">
        <div class="item-left">
          <i :class="getIcon(item.id)"></i>
        </div>
        <div class="item-right">
          <div class="item-top">{{ item.status }}</div>
          <div class="item-bottom">
            <countTo :startVal='old_item_num[item.id]' :endVal='item.num'
                     :duration='3000'></countTo>
            {{ item.unit }}
          </div>
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
      old_value: 0,
      old_item_num: [0, 0, 0, 0],
      items: [
        {id: 0, num: 0, unit: '', status: '环境评级'},
        {id: 1, num: 0, unit: 'dB', status: '噪音指标'},
        {id: 2, num: 0, unit: '', status: '气味指标'},
        {id: 3, num: 0, unit: '', status: '安全事故'},
      ]
    }
  },
  props: {
    latest_info: {
      type: Object,
      required: true,
    },
  },
  watch: {
    latest_info: {
      handler: function (newVal) {
        if (Object.values(newVal).every(value => value !== -1)) {
          this.old_item_num = this.items.map(item => item.num)
          // console.log(this.old_item_num, '---------old')
          // console.log(newVal, '---------new')
          // console.log(newVal.env, newVal.noise, newVal.smell, newVal.safe)
          this.items[0].num = newVal.env
          this.items[1].num = newVal.noise
          this.items[2].num = newVal.smell
          this.items[3].num = newVal.safe
          // console.log(this.items.map(item => item.num))
        }
      },
      deep: true
    }
  },
  methods: {
    getIcon(id) {
      if (id === 0) {
        return "iconfont icon-huanjing"
      } else if (id === 1) {
        return "iconfont icon-zaoyin"
      } else if (id === 2) {
        return "iconfont icon-qiti1"
      } else if (id === 3) {
        return "iconfont icon-UI_icon_shigu"
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
    margin: 10px 10px;
    height: 90%;
    width: 1800px;

    .grid-item {
      border-radius: 10px;
      margin: 10px;
      display: flex;
      flex-direction: row;

      .item-left {
        flex: 2;
        display: flex;
        justify-content: center; /* 在水平轴上居中 */
        align-items: center; /* 在纵轴上居中 */
        font-size: 40px;
        font-weight: bold;
        color: white;
      }

      .item-right {
        flex: 3;
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
          font-size: 40px;
          font-weight: bold;
          color: white;
        }
      }


    }
  }
}
</style>
