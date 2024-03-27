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
        data: [['2023/01/01', 42], ['2023/01/02', 43],
          ['2023/01/03', 50], ['2023/01/04', 70],]
      }
    };
  },
  props: {
    smell: {
      type: Object,
      required: true,
    },
  },
  watch: {
    smell: {
      handler: function (newVal) {
        // console.log(newVal)
        this.cdata.data = []
        newVal.smells.forEach((item, index) => {
          this.cdata.data[index] = [item.date, item.value]
        })
        // console.log(this.cdata.data)
      },
      deep: true
    }
  },
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
