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
        this.options = {
          tooltip: {
            trigger: 'item'
          },
          legend: {
            top: '20',
            left: 'center',
            textStyle: {
              fontSize: 20,
            },
          },
          series: [
            {
              name: '异常请求类别数',
              type: 'pie',
              radius: ['40%', '70%'],
              avoidLabelOverlap: false,
              label: {
                show: true,
                formatter(param) {
                  // correct the percentage
                  // return param.name + ' (' + param.percent  + '%)';
                  return  param.percent  + '%';
                },
                fontSize: 20,
                color: '#fff',
                position: 'outside',
              },
              emphasis: {
                label: {
                  show: false,
                  fontSize: 40,
                  fontWeight: 'bold'
                }
              },
              labelLine: {
                show: false
              },
              data: newData.data,
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
