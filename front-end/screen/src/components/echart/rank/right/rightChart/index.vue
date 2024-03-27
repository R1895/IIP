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
        data: [],
        category: [],
      }
    };
  },
  components: {
    Chart,
  },
  props: {
    rank: {
      type: Array,
      required: true,
    },
  },
  watch: {
    rank: {
      handler: function (newVal) {
        // console.log(newVal, '---newVal---')
        this.cdata.data = []
        this.cdata.category = []
        newVal.forEach((item) => {
          this.cdata.data.push(item.TotalTime)
          this.cdata.category.push(item.WorkerID)
        })
      },
      deep: true
    }
  },
  mounted() {
  },
  methods: {
    // run() {
    //   for (let i = 0; i < this.cdata.category.length; ++i) {
    //     if (Math.random() > 0.9) {
    //       // 修改数组中的元素时使用 this.$set
    //       this.$set(this.cdata.data, i, this.cdata.data[i] + Math.round(Math.random() * 50));
    //     } else {
    //       this.$set(this.cdata.data, i, this.cdata.data[i] + Math.round(Math.random() * 10));
    //     }
    //   }
    // }
  },
  created() {
    // setInterval(this.run, 3000);
  }
};
</script>

<style lang="scss" scoped>
</style>
