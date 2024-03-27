<template>
  <div>
    <Echart
        :options="options"
        id="MidTopChart"
        height="300px"
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
            top: 20,
            left: 20,
            orient: 'vertical',
            textStyle: {
              fontSize: 20,
            },
          },

          series: [
            {
              name: '安全事故分布',
              type: 'pie',
              radius: ['40%', '70%'],
              center: ['65%', '50%'],
              avoidLabelOverlap: false,
              label: {
                show: true,
                formatter(param) {
                  // correct the percentage
                  return param.percent + '%';
                  // return param.name + ' (' + param.percent + '%)';
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
