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
        data: [
          [12, 32, 42, 24, 52],
          [46, 36, 24, 42, 23],
        ],

      }
    };
  },
  components: {
    Chart,
  },
  mounted() {
  },
  props: {
    time: {
      type: Array,
      required: true,
    },
  },
  // watch: {
  //   time: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.category = []
  //       this.cdata.data = [[], []]
  //       newVal.forEach((item, index) => {
  //         this.cdata.category[index] = item.reason
  //         this.cdata.data[0][index] = item.process_time
  //         this.cdata.data[1][index] = item.wait_time
  //       })
  //     },
  //     deep: true
  //   }
  // },
  methods: {}
};
</script>

<style lang="scss" scoped>
</style>
