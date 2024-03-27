<template>
  <div class="procedure" ref="appRef">
    <div class="bg">

      <MyHeader :title="title"></MyHeader>
      <div class="procedure-container">
        <div>
          <dv-border-box-12>
            <gantt/>
          </dv-border-box-12>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import MyHeader from "../../components/header/myHeader.vue";
import drawMixin from "@/utils/drawMixin";
import Gantt from "@/views/procedure/gantt.vue";
// import mock from "@/api/mock";

export default defineComponent({
  name: "procedure",
  mixins: [drawMixin],
  data() {
    return {
      title: "工序甘特图",
    }
  },
  components: {
    Gantt,
    MyHeader,
  },
  mounted() {
    // setInterval 轮询策略待实现
  },
  // created() {
  //   this.listLoading = true
  //   mock.getProductionStatus()
  //       .then(response => {
  //         console.log(response)
  //         console.log(response.data.data.list)
  //         this.statusData = response.data.data.list;
  //         this.listLoading = false
  //       })
  //       .catch(error => {
  //         // 处理错误
  //         console.error(error);
  //       });
  // },
  methods: {}
})
</script>

<style scoped lang="scss">
@import '../../assets/scss/procedure.scss';
</style>
