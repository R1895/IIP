<template>
  <div class="app-container">
    <el-button class="button" type="primary" @click="addCamera">添加监控</el-button>
    <div class="blockList">
      <div class="block" v-for="(item,index) in cameraList" :key="index">
      <p class="header">设备:{{index}}</p>
      <img
      class="camera"
      style="
        display: block;
        -webkit-user-select: none;
        margin: auto;
        background-color: hsl(0, 0%, 25%);
      "
      :src="item.videoUrl"
      />
      <div class="bottom">
        <label class="label">IP地址</label>
        <el-input class="input" v-model="item.ip" placeholder="输入监控设备的ip地址"></el-input>
        <label class="label">端口号</label>
        <el-input
          class="input"
          v-model="item.port"
          placeholder="输入监控设备的ip端口号"
        ></el-input>
        <el-button class="button" type="primary" @click="startCam(index)">点击开始监控</el-button>
        <el-button class="button" type="warning" @click="item.videoUrl = '' "
          >断开连接</el-button
        >
        <el-button class="button" type="danger" @click="cameraList.splice(index,1)"
          >删除此监控</el-button
        >
      </div>
    </div> 
    </div>
  </div>
</template>

<script>
export default {
  filters: {},
  data() {
    return {
      ip: "",
      port: "",
      videoUrl: "",
      cameraList:[]
    };
  },
  created() {},
  methods: {
    startCam(index) {
      this.cameraList[index].videoUrl = `http://${this.cameraList[index].ip}:${this.cameraList[index].port}/video`
    },
    addCamera(){
      this.cameraList.push({
        ip:"",
        port:"",
        videoUrl:""
      })
    }
  },
};
</script>
<style lang="scss" scoped>
.blockList{
  display:flex;
  flex-direction: row ;
  flex-wrap:wrap;
}
.block{
  width:520px;
  padding:10px;
  margin:10px;
  border: 1px dashed #000;
  >.header{
    text-align: center;
    font-size:20px;
    font-weight: 600;
  }
  >.camera{
    width:480px;
    height:300px;
    margin:10px;
  }
  >.bottom{
    width:480px;
    margin:10px auto;
    >.input{
      margin-top:10px;
    }
    >.label{
      display:block;
      font-size:16px;
      margin-top:10px;
    }
    >.button{
      margin-top:10px;
    }
  }
}
</style>

