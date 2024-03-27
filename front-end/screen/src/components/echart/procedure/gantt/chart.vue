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
        // console.log(newData)
        this.options = {
          // title: {
          //   text: 'Waterfall Chart',
          //   subtext: 'Living Expenses in Shenzhen'
          // },
          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'shadow'
            },
            formatter: function (params) {
              var tar = params[1];
              return tar.name + '<br/>' + tar.seriesName + ' : ' + tar.value;
            }
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          yAxis: {
            type: 'category',
            splitLine: {show: false},
            // data: ['Total', 'Rent', 'Utilities', 'Transportation', 'Meals', 'Other']
          },
          xAxis: {
            type: 'time'
          },
          series: [
            {
              name: 'Placeholder',
              type: 'bar',
              stack: '1',

              // itemStyle: {
              //   borderColor: 'transparent',
              //   color: 'transparent'
              // },
              // emphasis: {
              //   itemStyle: {
              //     borderColor: 'transparent',
              //     color: 'transparent'
              //   }
              // },
              data: [['2021-1-1','Task1'], ['2021-1-2','Task2'], ['2021-1-3','Task3'], ['2021-1-4','Task4'], ['2021-1-5','Task5'], ['2021-1-6','Task6']]
            },
            {
              name: 'Life Cost',
              type: 'bar',
              stack: '1',
              z: -1,
              label: {
                show: true,
                position: 'inside'
              },
              data: [ ['2021-1-2','Task1'], ['2021-1-3','Task2'], ['2021-1-4','Task3'], ['2021-1-5','Task4'], ['2021-1-6','Task5'],['2021-1-7','Task6']]
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
