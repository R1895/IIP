<template>
  <div>
    <Chart :cdata="cdata"/>
  </div>
</template>

<script>
import Chart from './chart.vue'

export default {
  data() {
    return {
      cdata: {
        category: [
          '机械磨损', '电气故障', '程序出错', '控制故障', '操作错误'
        ],
        data: [151, 97, 87, 165, 79, 64, 186, 120, 159],
      }
    };
  },
  components: {
    Chart,
  },
  props: {
    loss_reason: {
      type: Array,
      required: true,
    },
  },
  // watch: {
  //   loss_reason: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.category = []
  //       this.cdata.data = []
  //       newVal.forEach((item, index) => {
  //         this.cdata.category[index] = item.reason
  //         this.cdata.data[index] = item.loss
  //       })
  //     },
  //     deep: true
  //   }
  // },
  mounted() {
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
</style>
