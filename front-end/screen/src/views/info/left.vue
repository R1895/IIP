<template>
  <div id="left">
    <div class="bg-color-black">
      <div class=" top-box">
        <div class="left-title">{{ this.title }}</div>
        <el-select v-model="value" placeholder="当天" @change="handleSelectChange">
          <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">
          </el-option>
        </el-select>
      </div>
      <div class="show-box">
        <div class="show-item">
          <div class="show-item-title">最小分贝值</div>
          <div class="show-item-min-content">
            <countTo class="num" :startVal='old_min_moise' :endVal='min_noise'
                     :duration='5000'></countTo>
          </div>
        </div>
        <div class="show-item">
          <div class="show-item-title">最大分贝值</div>
          <div class="show-item-max-content">
            <countTo class="num" :startVal='old_max_moise' :endVal='max_noise'
                     :duration='5000'></countTo>
          </div>
        </div>
      </div>
      <div>
        <LeftChart :noise="noise"/>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import LeftChart from '_c/echart/info/left/leftChart'

export default defineComponent({
  name: "left",
  data() {
    return {
      title: "噪音分贝值统计图",
      options: [{
        value: '1',
        label: '当天'
      }, {
        value: '7',
        label: '近七天'
      }, {
        value: '30',
        label: '近一月'
      }],
      value: '1',
      min_noise: 0,
      max_noise: 0,
      old_min_moise: 0,
      old_max_moise: 0,
    }
  },
  props: {
    noise: {
      type: Object,
      required: true,
    },
  },
  watch: {
    noise: {
      handler: function (newVal) {
        this.old_max_moise = this.max_noise
        this.old_min_moise = this.min_noise
        // console.log(newVal)
        this.min_noise = newVal.min
        this.max_noise = newVal.max
      },
      deep: true
    }
  },
  components: {
    LeftChart
  },
  methods: {
    handleSelectChange() {
      this.$emit('select-change', this.value);
    }
  },
})
</script>


<style scoped lang="scss">
$box-height: 750px;
$box-width: 100%;
#left {
  padding: 20px 16px;
  height: $box-height;
  width: $box-width;
  border-radius: 5px;

  .bg-color-black {
    height: $box-height - 35px;
    border-radius: 10px;
  }

  .top-box {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .left-title {
      padding: 0 0 0 20px;
      font-size: 24px;
      font-weight: bold;
      color: green;
    }
  }


  .show-box {
    display: flex;
    flex-direction: row;
    justify-content: space-around;
    margin-top: 10px;
    height: 150px;

    .show-item {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .show-item-title {
        height: 40px;
        font-size: 20px;
        font-weight: bold;
        color: white;
      }

      .show-item-min-content {
        font-size: 32px;
        font-weight: bold;
        color: #c0dca5;
      }

      .show-item-max-content {
        font-size: 32px;
        font-weight: bold;
        color: #ea9393;
      }
    }
  }
}
</style>
