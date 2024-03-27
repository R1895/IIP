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
          "3月1日",
          "3月2日",
          "3月3日",
          "3月4日",
          "3月5日",
          "3月6日",
          "3月7日",
        ],
        data: [
          11, 10.86, 20.61,19.8, 21.57, 24.97, 39.64
        ],

      }
    };
  },
  props: {
    oee_trend: {
      type: Array,
      required: true,
    },
  },
  // watch: {
  //   oee_trend: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.category = []
  //       this.cdata.data = []
  //       newVal.forEach((item, index) => {
  //         this.cdata.category[index] = item.date
  //         this.cdata.data[index] = item.percentage
  //       })
  //     },
  //     deep: true
  //   }
  // },
  components: {
    Chart,
  },
  mounted() {
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
</style>
