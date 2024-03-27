<template>
  <div>
    <Echart
        :options="options"
        id="LeftChart"
        height="550px"
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
            // formatter: function (params) {
            //   params = params[0];
            //   var date = new Date(params.name);
            //   return (
            //       date.getDate() +
            //       '/' +
            //       (date.getMonth() + 1) +
            //       '/' +
            //       date.getFullYear() +
            //       ' : ' +
            //       params.value[1]
            //   );
            // },
            axisPointer: {
              animation: false
            }
          },
          xAxis: {
            type: 'time',
            axisLabel: {
              textStyle: {// 设置字体颜色
                fontSize: 16 // 设置字体大小
              }
            },
            splitLine: {
              show: false
            },
          },
          yAxis: {
            type: 'value',
            axisLabel: {
              textStyle: {// 设置字体颜色
                fontSize: 14 // 设置字体大小
              }
            },
            // name: 'level',
            // nameTextStyle: {
            //   fontSize: 16
            // }
            splitLine: {
              show: false
            }
          },
          series: [
            {
              data: newData.data,
              type: 'line',
              showSymbol: false,
              // label: {
              //   show: true, // 显示标签
              //   position: 'top', // 标签位置，可以是 'top', 'insideTop', 'inside', 'insideBottom' 等
              //   color: 'white', // 标签颜色
              //   fontSize: 16, // 字体大小
              // },
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
