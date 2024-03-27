<template>
  <div id="rank" ref="appRef">
    <div class="bg">
      <div class="rank-body">

        <MyHeader :title="title"></MyHeader>

        <div class="body-box">
          <!-- 左侧数据 -->
          <div class="left-box">
            <div>
              <left :rank="this.mouth_rank"/>
            </div>
          </div>
          <!-- 中间数据 -->
          <div class="mid-box">
            <div>
              <mid :rank="this.year_rank"/>
            </div>
          </div>
          <!-- 右侧数据 -->
          <div class="right-box">
            <div>
              <right :rank="this.day_rank"/>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>


<script>
import {defineComponent} from 'vue'
import drawMixin from "@/utils/drawMixin";
import MyHeader from "../../components/header/myHeader.vue";
import left from "@/views/working_hours_rank/left.vue";
import right from "@/views/working_hours_rank/right.vue";
import mid from "@/views/working_hours_rank/mid.vue";
import rank from "@/api/rank";
import {getFirstDayofMouthFormatDate, getFirstDayofYearFormatDate} from "@/utils/date";


export default defineComponent({
  name: "rank",
  mixins: [drawMixin],
  data() {
    return {
      title: "工时排行",
      mouth_rank: [],
      day_rank: [],
      year_rank: []
    }
  },
  components: {left, right, MyHeader, mid},
  mounted() {
    // setInterval(() => {
    //   this.fetchDayData();
    // }, 3000);
    // setInterval(() => {
    //   this.fetchMouthData();
    // }, 3000);
    // setInterval(() => {
    //   this.fetchYearData();
    // }, 10000);
  },
  created() {
    // this.fetchDayData();
    // this.fetchMouthData();
    // this.fetchYearData();
  },
  methods: {
    fetchDayData() {
      const now = '2023-12-14'
      // const now = this.getNowFormatDate()
      rank.getRankInfoAnalysis(now, now, 1)
          .then(response => {
            // console.log(response)
            this.day_rank = response.data.data;
          })
          .catch(error => {
            console.error(error);
          });
    },
    fetchMouthData() {
      const now = '2023-12-14'
      // const now = this.getNowFormatDate()
      rank.getRankInfoAnalysis(getFirstDayofMouthFormatDate(now), now, 1)
          .then(response => {
            // console.log(response)
            this.mouth_rank = response.data.data;
          })
          .catch(error => {
            console.error(error);
          });
    },
    fetchYearData() {
      const now = '2023-12-14'
      // const now = this.getNowFormatDate()
      rank.getRankInfoAnalysis(getFirstDayofYearFormatDate(now), now, 1)
          .then(response => {
            // console.log(response)
            this.year_rank = response.data.data;
          })
          .catch(error => {
            console.error(error);
          });
    },
  },
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/rank.scss';
</style>
