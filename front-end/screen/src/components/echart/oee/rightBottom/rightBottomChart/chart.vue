<template>
  <div>
    <Echart
        :options="options"
        id="RightBottomChart"
        height="400px"
        width="100%"
    ></Echart>
  </div>
</template>

<script>
import Echart from '@/common/echart/index.vue'

export default {
  data() {
    return {
      options: {},
    };
  },
  components: {
    Echart,
  },
  props: {
    cdata: {
      type: Object,
      default: () => ({})
    },
  },
  watch: {
    cdata: {
      handler(newData) {
        this.options = {
          xAxis: {
            type: 'category',
            data: newData.category,
            axisLabel: {
              textStyle: {// 设置字体颜色
                fontSize: 16 // 设置字体大小
              }
            },
          },
          yAxis: {
            type: 'value',
            axisLabel: {
              textStyle: {// 设置字体颜色
                fontSize: 14 // 设置字体大小
              }
            },
            name: '%百分比',
            nameTextStyle: {
              fontSize: 16
            }
          },
          series: [
            {
              data: newData.data,
              type: 'line',
              label: {
                show: true, // 显示标签
                position: 'top', // 标签位置，可以是 'top', 'insideTop', 'inside', 'insideBottom' 等
                color: 'white', // 标签颜色
                fontSize: 16, // 字体大小
              },
            }
          ]
        }
      },
      immediate: true,
      deep: true
    },
  },
}
</script>