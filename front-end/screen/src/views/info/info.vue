<template>
  <div id="info" ref="appRef">
    <div class="bg">
      <div class="info-body">

        <MyHeader :title=title></MyHeader>

        <div class="body-box">
          <!-- 顶部数据 -->
          <div class="top-box">
            <dv-border-box-12>
              <top :latest_info="latest_info"/>
            </dv-border-box-12>
          </div>
          <!--主要数据-->
          <div class="content-box">

            <div class="left-box">
              <dv-border-box-12>
                <left :noise="noise" @select-change="handleNoiseSelectChange"/>
              </dv-border-box-12>
            </div>

            <div class="mid-box">
              <div class="mid-top-box">
                <dv-border-box-12>
                  <midTop :safe="safe"/>
                </dv-border-box-12>
              </div>
              <div class="mid-bottom-box">
                <dv-border-box-12>
                  <midBottom :env="env" @select-change="handleEnvSelectChange"/>
                </dv-border-box-12>
              </div>
            </div>

            <div class="right-box">
              <dv-border-box-12>
                <right :smell="smell" @select-change="handleSmellSelectChange"/>
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
import Top from "@/views/info/top.vue";
import Left from "@/views/info/left.vue";
import Right from "@/views/info/right.vue";
import MidBottom from "@/views/info/midBottom.vue";
import MidTop from "@/views/info/midTop.vue";
import info from "@/api/info";
import {getNowFormatDate, getOneMonthAgoDate, getSevenDaysAgoDate} from "@/utils/date";

export default defineComponent({
  name: "live",
  mixins: [drawMixin],
  components: {
    Top, Left, MidTop, MidBottom, Right,
    MyHeader
  },
  data() {
    return {
      title: "车间信息",
      env: [],
      smell: {},
      noise: {},
      safe: {},
      latest_info: {
        env: -1,
        noise: -1,
        smell: -1,
        safe: -1,
      },
      envSelectedValue: '',
      smellSelectedValue: '',
      noiseSelectedValue: '',
    }
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
    handleEnvSelectChange(value) {
      // 监听子组件的自定义事件，更新父组件中的值
      // console.log(value)
      this.envSelectedValue = value;
      this.fetchData();
    },
    handleSmellSelectChange(value) {
      // console.log(value)
      this.smellSelectedValue = value;
      this.fetchData();
    },
    handleNoiseSelectChange(value) {
      // console.log(value)
      this.noiseSelectedValue = value;
      this.fetchData();
    },

    fetchData() {
      const now = '2023-12-14'
      // const now = this.getNowFormatDate()

      let envStartTime = getSevenDaysAgoDate(now);
      if (this.envSelectedValue == 30) {
        envStartTime = getOneMonthAgoDate(now);
      }
      let smellStartTime = getNowFormatDate(now);
      if (this.smellSelectedValue == 7) {
        smellStartTime = getSevenDaysAgoDate(now);
      } else if (this.smellSelectedValue == 30) {
        smellStartTime = getOneMonthAgoDate(now);
      }
      let noiseStartTime = getNowFormatDate(now);
      if (this.noiseSelectedValue == 7) {
        noiseStartTime = getSevenDaysAgoDate(now);
      } else if (this.noiseSelectedValue == 30) {
        noiseStartTime = getOneMonthAgoDate(now);
      }

      info.getEnvInfoAnalysis(envStartTime, now, 1)
          .then(response => {
            this.env = response.data.data.item;
            this.latest_info.env = this.env[this.env.length - 1].level;
          })
          .catch(error => {
            console.error(error);
          });
      info.getSmellInfoAnalysis(smellStartTime, now, 1)
          .then(response => {
            this.smell = response.data.data;
            this.latest_info.smell = this.smell.smells[this.smell.smells.length - 1].value;
          })
          .catch(error => {
            console.error(error);
          });
      info.getNoiseInfoAnalysis(noiseStartTime, now, 1)
          .then(response => {
            this.noise = response.data.data;
            this.latest_info.noise = this.noise.noises[this.noise.noises.length - 1].value;
          })
          .catch(error => {
            console.error(error);
          });
      info.getSafeInfoAnalysis('2023-1-1', now, 1)
          .then(response => {
            this.safe = response.data.data;
            this.latest_info.safe = this.safe.count;
          })
          .catch(error => {
            console.error(error);
          });
    },
  }
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/info.scss';
</style>
