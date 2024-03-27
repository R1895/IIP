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
        data: [{value: 2, name: '跌倒和绊倒事故'},
          {value: 4, name: '机械伤害'},
          {value: 6, name: '人为错误'},
          {value: 5, name: '火灾和爆炸'},
          {value: 3, name: '化学品泄露'},],
      }
    };
  },
  props: {
    safe: {
      type: Object,
      required: true,
    },
  },
  watch: {
    safe: {
      handler: function (newVal) {
        // console.log(newVal)
        this.cdata.data = []
        newVal.safes.forEach((item, index) => {
          this.cdata.data[index] = {value: item.total_value, name: this.getSafeNameByType(item.type)}
        })
        // console.log(this.cdata.data)
      },
      deep: true
    }
  },
  methods: {
    getSafeNameByType(type) {
      if (type == 1) {
        return "跌倒和绊倒事故"
      } else if (type == 2) {
        return "机械伤害"
      } else if (type == 3) {
        return "人为错误"
      } else if (type == 4) {
        return "火灾和爆炸"
      } else if (type == 5) {
        return "化学品泄露"
      }
    }
  },
  components: {
    Chart,
  },
  mounted() {
  },
};
</script>

<style lang="scss" scoped>
</style>
