<template>
  <div class="app-container">
    <h1>查看工人工作的任务</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      worker_id:""
    };
  },
  mounted(){
    this.worker_id = this.$route.params.id
    console.log(this.worker_id)
  },
  methods: {
    
  }
};
</script>

<style scoped>
</style>
