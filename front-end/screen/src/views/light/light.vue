<template>
  <div id="light" ref="appRef">
    <div class="bg">
      <div class="light-body">

        <MyHeader :title="title"></MyHeader>

        <div class="body-box">
          <!-- 左侧数据 -->
          <div class="left-box">
            <div>
              <dv-border-box-12>
                <left :time="this.time"/>
              </dv-border-box-12>
            </div>
          </div>
          <!-- 中间数据 -->
          <div class="mid-box">
            <div>
              <dv-border-box-12>
                <mid/>
              </dv-border-box-12>
            </div>
          </div>
          <!-- 右侧数据 -->
          <div class="right-box">
            <div>
              <dv-border-box-12>
                <right :number="this.number"/>
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
import drawMixin from "@/utils/drawMixin";
import MyHeader from "../../components/header/myHeader.vue";
import left from "@/views/light/left.vue";
import right from "@/views/light/right.vue";
import mid from "@/views/light/mid.vue";


export default defineComponent({
  name: "light",
  mixins: [drawMixin],
  data() {
    return {
      title: "设备异常分析",
      time: [],
      number: [],
    }
  },
  components: {left, right, MyHeader, mid},
  mounted() {
    // setInterval 轮询策略待实现
  },
  created() {
    // this.listLoading = true
    // mock.getLightEfficiency()
    //     .then(response => {
    //       // console.log(response)
    //       this.time = response.data.data.time;
    //       this.number = response.data.data.number;
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
@import '../../assets/scss/light.scss';
</style>
