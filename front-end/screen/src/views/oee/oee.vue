<template>
  <div id="oee" ref="appRef">
    <div class="bg">
      <div class="oee-body">

        <MyHeader :title="title"></MyHeader>

        <div class="body-box">
          <!-- 左侧数据 -->
          <div class="left-box">
            <div>
              <dv-border-box-12>
                <left :oee_rank="this.oee_rank"/>
              </dv-border-box-12>
            </div>
          </div>

          <!-- 右侧数据 -->
          <div class="right-box">
            <div class="right-top-box">
              <div>
                <dv-border-box-12>
                  <rightTopLeft :daily_loss="this.daily_loss"/>
                </dv-border-box-12>
              </div>
              <div>
                <dv-border-box-12>
                  <rightTopRight :loss_reason="this.loss_reason"/>
                </dv-border-box-12>
              </div>
            </div>
            <div class="right-bottom-box">
              <div>
                <dv-border-box-12>
                  <rightBottom :oee_trend="this.oee_trend"/>
                </dv-border-box-12>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MyHeader from "../../components/header/myHeader.vue";
import drawMixin from "@/utils/drawMixin";
import centerRight2 from "@/views/centerRight2.vue";
import bottomRight from "@/views/bottomRight.vue";
import centerLeft1 from "@/views/centerLeft1.vue";
import centerLeft2 from "@/views/centerLeft2.vue";
import center from "@/views/center.vue";
import centerRight1 from "@/views/centerRight1.vue";
import bottomLeft from "@/views/bottomLeft.vue";
import left from "@/views/oee/left.vue";
import rightTopLeft from "@/views/oee/rightTopLeft.vue";
import rightTopRight from "@/views/oee/rightTopRight.vue";
import rightBottom from "@/views/oee/rightBottom.vue";

export default defineComponent({
  name: "oee",
  mixins: [drawMixin],
  data() {
    return {
      title: "OEE看板",
      daily_loss: [],
      oee_rank: [],
      loss_reason: [],
      oee_trend: [],
    }
  },
  components: {
    rightBottom, left, rightTopLeft, rightTopRight,
    bottomLeft, centerRight1, center, centerLeft2, centerLeft1, bottomRight, centerRight2,
    MyHeader,/**/
  },
  mounted() {
    // setInterval 轮询策略待实现
  },
  created() {
    // mock.getOeeAnalysis()
    //     .then(response => {
    //       // console.log(response)
    //       this.oee_rank = response.data.data.oee_rank;
    //       this.daily_loss = response.data.data.daily_loss;
    //       this.loss_reason = response.data.data.loss_reason;
    //       this.oee_trend = response.data.data.oee_trend;
    //       this.listLoading = false
    //     })
    //     .catch(error => {
    //       // 处理错误
    //       console.error(error);
    //     });
  },
  methods: {},
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/oee.scss';
</style>
