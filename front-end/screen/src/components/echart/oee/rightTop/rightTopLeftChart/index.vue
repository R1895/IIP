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
          "3/1",
          "3/2",
          "3/3",
          "3/4",
          "3/5",
          "3/6",
          "3/7",
        ],
        data: [
          72, 68, 90, 86, 160, 210, 211
        ],

      }
    };
  },
  props: {
    daily_loss: {
      type: Array,
      required: true,
    },
  },
  // watch: {
  //   daily_loss: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.category = []
  //       this.cdata.data = []
  //       newVal.forEach((item, index) => {
  //         this.cdata.category[index] = item.date
  //         this.cdata.data[index] = item.loss
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
