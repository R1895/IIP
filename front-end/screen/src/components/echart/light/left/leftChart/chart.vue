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
              type: 'shadow'
            }
          },
          legend: {
            top: '20',
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
          xAxis: [
            {
              type: 'category',
              data: newData.category,
              axisLabel: {
                textStyle: {
                  fontSize: 20, // 设置字体大小
                },
                interval:0,
              },
            }
          ],
          yAxis: [
            {
              type: 'value',
              axisLabel: {
                textStyle: {
                  fontSize: 20 // 设置字体大小
                },
              },
            }
          ],
          series: [
            {
              name: '平均等待时间',
              type: 'bar',
              emphasis: {
                focus: 'series'
              },
              stack: 'Ad',
              barWidth: '30%',   // 设置柱子宽度为类目宽度的 20%
              barGap: '10%',     // 设置间隔为柱子宽度的 10%
              data: newData.data[0]
            },
            {
              name: '平均处理时间',
              type: 'bar',
              stack: 'Ad',
              emphasis: {
                focus: 'series'
              },
              barWidth: '30%',   // 设置柱子宽度为类目宽度的 20%
              barGap: '10%',     // 设置间隔为柱子宽度的 10%
              data: newData.data[1]
            },
          ]
        }
      },
      immediate: true,
      deep: true
    },
  },
}
</script>
