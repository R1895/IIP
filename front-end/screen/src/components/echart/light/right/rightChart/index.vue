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
        data: [
          {value: 2, name: '机械磨损'},
          {value: 4, name: '电气故障'},
          {value: 6, name: '程序出错'},
          {value: 5, name: '控制故障'},
          {value: 3, name: '操作错误'},
        ],
      }
    };
  },
  components: {
    Chart,
  },
  props: {
    number: {
      type: Array,
      required: true,
    },
  },
  // watch: {
  //   number: {
  //     handler: function (newVal) {
  //       console.log(newVal)
  //       this.cdata.data = []
  //       newVal.forEach((item, index) => {
  //         this.cdata.data[index] = {
  //           value: item.quantity,
  //           name: item.reason
  //         }
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
