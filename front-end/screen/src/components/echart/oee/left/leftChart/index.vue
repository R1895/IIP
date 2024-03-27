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
          "C015",
          "C016",
          "C003",
          "C007",
          "C005",
          "C019",
          "C002",
          "C006",
          "C009",
          "C017",
          "C010",
          "C011",
          "C012",
          "C008",
          "C013",
          "C020",
          "C004",
          "C001",
          "C014",
          "C018",

        ],
        data: [
          [12, 17, 21, 23, 24, 24, 32, 34, 36, 43, 54, 59, 63, 65, 67, 72, 73, 76, 80, 82],
          [43, 40, 44, 40, 36, 34, 37, 31, 26, 24, 23, 21, 14, 5, 3, 13, 11, 10, 5, 6],
          [23, 15, 22, 12, 13, 22, 11, 15, 21, 13, 13, 3, 3, 12, 12, 3, 6, 4, 5, 6],
          [22, 28, 13, 25, 27, 20, 20, 20, 17, 20, 10,17, 20, 18, 18, 12, 10, 10, 10, 6],
          // [12, 7, 10, 13, 15, 10, 2, 10, 10, 6, 7, 11, 7, 9, 10, 5, 7, 9, 5, 0],
        ]
        ,

      }
    };
  },
  components: {
    Chart,
  },
  props: {
    oee_rank: {
      type: Array,
      required: true,
    },
  },
  mounted() {
  },
  methods: {},
  // watch: {
  //   oee_rank: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.category = []
  //       this.cdata.data = [[], [], [], []]
  //       newVal.forEach((item, index) => {
  //         this.cdata.category[index] = item.machine_id
  //         /*To do*/
  //         this.cdata.data[0][index] = item.percentages[0].percentage
  //         this.cdata.data[1][index] = item.percentages[1].percentage
  //         this.cdata.data[2][index] = item.percentages[2].percentage
  //         this.cdata.data[3][index] = item.percentages[3].percentage
  //       })
  //     },
  //     deep: true
  //   }
  // }
};
</script>

<style lang="scss" scoped>
</style>
