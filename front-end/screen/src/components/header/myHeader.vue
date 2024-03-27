<template>
  <div class="myHeader">
    <div class="d-flex jc-center">
      <dv-decoration-10 class="dv-dec-10"/>
      <div class="d-flex jc-center">
        <dv-decoration-8 class="dv-dec-8" :color="decorationColor"/>
        <div class="title">
          <span class="title-text">{{ this.title }}</span>
          <dv-decoration-6
              class="dv-dec-6"
              :reverse="true"
              :color="['#50e3c2', '#67a1e5']"
          />
        </div>
        <dv-decoration-8
            class="dv-dec-8"
            :reverse="true"
            :color="decorationColor"/>
      </div>
      <div class="time-container">
        <dv-decoration-10 class="dv-dec-10-s"/>
        <!--        <div class="time-box">111111111</div>-->
        <div class="time-box"
        >{{ dateYear }} {{ dateWeek }} {{ dateDay }}
        </div
        >
      </div>

    </div>

    <!-- 第二行 -->
    <div class="d-flex jc-between px-2">
      <div class="d-flex aside-width">
        <div class="react-left ml-4 react-l-s">
          <span class="react-left"></span>
          <span class="text" @click="navigateToStatus">生产状态</span>
        </div>
        <div class="react-left ml-3">
          <span class="text" @click="navigateToOee">OEE</span>
        </div>
        <div class="react-left ml-3">
          <span class="text" @click="navigateToLight">异常处理</span>
        </div>
        <div class="react-left ml-3">
          <span class="text" @click="navigateToLive">生产实况</span>
        </div>
        <div class="react-left ml-3">
          <span class="text" @click="navigateToInfo">工厂环境</span>
        </div>
      </div>
      <div class="d-flex aside-width">
        <div class="react-right  react-l-s mr-3">
          <span class="text fw-b" @click="navigateToOrder">工单统计</span>
        </div>
        <div class="react-right  react-l-s mr-3">
          <span class="text fw-b" @click="navigateToProcedure">工序顺序</span>
        </div>
        <div class="react-right react-l-s mr-3">
          <span class="text fw-b" @click="navigateToRank">工时排行</span>
        </div>
        <div class="react-right react-l-s mr-3">
          <span class="react-after"></span>
          <span class="text fw-b" @click="navigateToStock">库存</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import {formatTime} from "@/utils";

export default defineComponent({
  name: "myHeader",
  props: {title: String},
  data() {
    return {
      timing: null,
      loading: true,
      dateDay: null,
      dateYear: null,
      dateWeek: null,
      weekday: ['周日', '周一', '周二', '周三', '周四', '周五', '周六'],
      decorationColor: ['#568aea', '#000000'],
    }
  },
  mounted() {
    this.timeFn()
  },
  beforeDestroy() {
    clearInterval(this.timing)
  },
  methods: {
    timeFn() {
      this.timing = setInterval(() => {
        this.dateDay = formatTime(new Date(), 'HH: mm: ss')
        this.dateYear = formatTime(new Date(), 'yyyy-MM-dd')
        this.dateWeek = this.weekday[new Date().getDay()]
      }, 1000)
    },
    navigateToStatus() {
      this.$router.push('/status').catch(err => {
        console.log(err)
      });
    },
    navigateToOee() {
      this.$router.push('/oee').catch(err => {
        console.log(err)
      });
    },
    navigateToLight() {
      this.$router.push('/light').catch(err => {
        console.log(err)
      });
    },
    navigateToLive() {
      this.$router.push('/').catch(err => {
        console.log(err)
      });
    },
    navigateToInfo() {
      this.$router.push('/info').catch(err => {
        console.log(err)
      });
    },
    navigateToRank() {
      this.$router.push('/rank').catch(err => {
        console.log(err)
      });
    },
    navigateToOrder() {
      this.$router.push('/order').catch(err => {
        console.log(err)
      });
    },
    navigateToStock() {
      this.$router.push('/stock').catch(err => {
        console.log(err)
      });
    },
    navigateToProcedure() {
      this.$router.push('/procedure').catch(err => {
        console.log(err)
      });
    },
  }
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/header.scss';
</style>
