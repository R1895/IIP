<template>
  <div class="lt-box">
    <div class="bg-color-black">
      <div class="lt-desc">
        <div class="lt-title">{{ this.title }}</div>
        <div class="lt-detail" @click="navigateToOee">详情</div>
      </div>
      <div class="board-content">
        <dv-capsule-chart class="dv-capsule-chart" :config="config"/>
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
      title: "设备OEE",
      config: {
        data: [
          {
            name: '80%以上',
            value: 5
          },
          {
            name: '50%-80%',
            value: 18
          },
          {
            name: '50%以下',
            value: 7
          },
        ],
        colors: ['#e062ae', '#fb7293', '#96bfff'],
        unit: '个数',
        showValue: true
      }
    }
  },
  props: {
    oee: {
      type: Object,
      required: true,
    },
  },
  watch: {
    oee: {
      handler: function (newVal) {
        console.log(newVal)
        this.config.data[0].value = newVal.oee_uo_80
        this.config.data[1].value = newVal.oee_50_to_80
        this.config.data[2].value = newVal.oee_down_50
        this.config = {...this.config}
      },
      deep: true
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
.lt-box {
  padding: 16px;
  padding-top: 20px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 30px;
    border-radius: 10px;
  }

  .lt-desc {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 20px;

    .lt-title {
      padding: 5px 0 0 20px;
      font-size: 24px;
      font-weight: bold;
      color: green;
    }

    .lt-detail {
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

    .dv-capsule-chart {
      width: 100%;
      height: 275px;
    }
  }
}
</style>
