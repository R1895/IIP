<template>
  <div>
    <Echart
        :options="options"
        id="RightChart"
        height="800px"
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
        // console.log(newData)
        this.options = {
          xAxis: {
            max: 'dataMax',
            inverse: true,
            show: false
          },
          yAxis: {
            type: 'category',
            data: newData.category,
            position: 'right',
            inverse: true,
            animationDuration: 300,
            animationDurationUpdate: 300,
            max: 9, // only the largest 3 bars will be displayed
            axisLabel: {
              margin: 10,
              textStyle: {
                fontSize: 16,
                fontWeight: 'bold'
              }
            }
          },
          series: [
            {
              realtimeSort: true,
              name: '工时',
              type: 'bar',
              data: newData.data,
              label: {
                show: true,
                position: 'left',
                valueAnimation: true,
                color: 'white', // 标签颜色
                fontSize: 16, // 字体大小
              }
            }
          ],
          legend: {
            show: true
          },
          animationDuration: 0,
          animationDurationUpdate: 3000,
          animationEasing: 'linear',
          animationEasingUpdate: 'linear'
        }
      },
      immediate: true,
      deep: true
    },
  },
}
</script>
