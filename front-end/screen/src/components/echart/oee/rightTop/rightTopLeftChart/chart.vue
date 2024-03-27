<template>
  <div>
    <Echart
        :options="options"
        id="RightTopLeftChart"
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
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            }
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
              axisTick: {
                alignWithLabel: true
              },
              axisLabel: {
                textStyle: {// 设置字体颜色
                  fontSize: 16 // 设置字体大小
                },
                interval:0,
              },
              // name: '日期',
            }
          ],
          yAxis: [
            {
              type: 'value',
              axisLabel: {
                textStyle: {// 设置字体颜色
                  fontSize: 16 // 设置字体大小
                }
              },
              name: 'min/分钟',
              nameTextStyle: {
                fontSize: 16
              }
            }
          ],
          series: [
            {
              name: 'Direct',
              type: 'bar',
              barWidth: '60%',
              data: newData.data,
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