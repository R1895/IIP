<template>
  <div>
    <Echart
        :options="options"
        id="LeftChart"
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
        this.options = {
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              // Use axis to trigger tooltip
              type: 'shadow' // 'shadow' as default; can also be 'line' or 'shadow'
            }
          },
          legend: {
            top: '0',
            left: 'center',
            textStyle: {
              fontSize: 20,
            },
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          xAxis: {
            position: 'top',
            type: 'value',
            axisLabel: {
              textStyle: {
                fontSize: 16 // 设置字体大小
              }
            }
          },
          yAxis: {
            type: 'category',
            data: newData.category,
            axisLabel: {
              textStyle: {
                fontSize: 16 // 设置字体大小
              }
            }
          },
          series: [
            {
              name: '运行',
              color: '#28D094',
              type: 'bar',
              stack: 'total',
              emphasis: {
                focus: 'series'
              },
              data: newData.data[0]
            },
            {
              name: '故障',
              color: '#FA746B',
              type: 'bar',
              stack: 'total',
              emphasis: {
                focus: 'series'
              },
              data: newData.data[1]
            },
            {
              name: '等待',
              color: '#FFC71C',
              type: 'bar',
              stack: 'total',
              emphasis: {
                focus: 'series'
              },
              data: newData.data[2]
            },
            {
              name: '关机',
              color: '#BFBFBF',
              type: 'bar',
              stack: 'total',
              emphasis: {
                focus: 'series'
              },
              data: newData.data[3]
            },
            // {
            //   name: '关机',
            //   color: '#BFBFBF',
            //   type: 'bar',
            //   stack: 'total',
            //   emphasis: {
            //     focus: 'series'
            //   },
            //   data: newData.data[4]
            // }
          ]
        }
      },
      immediate: true,
      deep: true
    },
  },
}
</script>
