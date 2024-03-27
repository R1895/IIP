<template>
  <div id="live" ref="appRef">
    <div class="bg">
      <div class="live-body">

        <MyHeader :title=title></MyHeader>

        <div class="body-box">
          <!-- 顶部数据 -->
          <div class="top-box">
            <dv-border-box-12>
              <top :status_num="status_num"/>
            </dv-border-box-12>
          </div>
          <!--主要数据-->
          <div class="content-box">
            <div class="left-box">
              <div class="left-top-box">
                <dv-border-box-12>
                  <leftTop :oee="oee"/>
                </dv-border-box-12>
              </div>
              <div class="left-bottom-box">
                <dv-border-box-12>
                  <leftBottom :status_num="status_num"/>
                </dv-border-box-12>
              </div>
            </div>
            <div class="center-box">
              <dv-border-box12>
                <mid :machines="machines"/>
              </dv-border-box12>
            </div>
            <div class="right-box">
              <dv-border-box-12>
                <right :anomalys="anomalys"/>
              </dv-border-box-12>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import drawMixin from "../../utils/drawMixin";
import MyHeader from "../../components/header/myHeader.vue";
import top from './top.vue'
import leftTop from './leftTop.vue'
import leftBottom from './leftBottom.vue'
import mid from './mid.vue'
import right from './right.vue'
import live from "@/api/live";

export default defineComponent({
  name: "live",
  data() {
    return {
      title: "生产实况",
      status_num: {},
      oee: {},
      machines: [],
      anomalys: [],

    }
  },
  mixins: [drawMixin],
  components: {
    MyHeader, top, leftTop, leftBottom, mid, right
  },
  mounted() {
    // setInterval(() => {
    //   this.fetchData();
    // }, 3000);
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      live.getLiveInformation(1).then(
          (response) => {
            console.log(response.data.data)
            this.status_num = response.data.data.status_num;
            this.oee = response.data.data.oee;
            this.machines = response.data.data.machines;
            this.anomalys = response.data.data.anomalys;
          }
      ).catch(
          (error) => {
            console.log(error)
          }
      )
    },
  },

})
</script>

<style scoped lang="scss">
@import '../../assets/scss/live.scss';
</style>
