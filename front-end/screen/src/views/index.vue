<template>
  <div id="index" ref="appRef">
    <div class="bg">
      <dv-loading v-if="loading">Loading...</dv-loading>
      <div v-else class="host-body">

        <MyHeader title="智慧大屏"></MyHeader>

        <div class="body-box">
          <!-- 第一行数据 -->
          <div class="content-box">
            <div>
              <dv-border-box-12>
                <centerLeft1/>
              </dv-border-box-12>
            </div>
            <div>
              <dv-border-box-12>
                <centerLeft2/>
              </dv-border-box-12>
            </div>
            <!-- 中间 -->
            <div>
              <center/>
            </div>
            <!-- 中间 -->
            <div>
              <centerRight2/>
            </div>
            <div>
              <dv-border-box-13>
                <centerRight1/>
              </dv-border-box-13>
            </div>
          </div>

          <!-- 第二行数据 -->
          <div class="bottom-box">
            <dv-border-box-13>
              <bottomLeft/>
            </dv-border-box-13>
            <dv-border-box-12>
              <bottomRight/>
            </dv-border-box-12>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import drawMixin from "../utils/drawMixin";
import centerLeft1 from './centerLeft1'
import centerLeft2 from './centerLeft2'
import centerRight1 from './centerRight1'
import centerRight2 from './centerRight2'
import center from './center'
import bottomLeft from './bottomLeft'
import bottomRight from './bottomRight'
import MyHeader from "../components/header/myHeader.vue";

export default {
  mixins: [drawMixin],
  data() {
    return {
      loading: true,
    }
  },
  components: {
    centerLeft1,
    centerLeft2,
    centerRight1,
    centerRight2,
    center,
    bottomLeft,
    bottomRight,
    MyHeader
  },
  mounted() {
    this.cancelLoading()
  },
  methods: {
    cancelLoading() {
      setTimeout(() => {
        this.loading = false
      }, 500)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/scss/index.scss';
</style>
