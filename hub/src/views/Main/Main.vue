<template>
<div >
<!--  写作文章功能-->
  <el-row>
    <div class="imgWrapper">
      <img src="../../assets/img/newarticle.svg" class="articleImg">
      <h1 class="articleFont">随心写作,自由表达</h1>
      <el-button type="primary" @click="beginWriteArticle" class="btn">开始写文章</el-button>
    </div>
  </el-row>


  <!-- 查看其他用户博客标题栏-->

  <el-row>
    <div class="imgWrapper">
      <h1 class="findColor">发现精彩</h1>
    </div>
  </el-row>
  <!-- 查看其他用户博客功能-->
  <el-row >

    <el-col :span="8"  v-for="msg  in   this.$store.state.loadMoveUsers[0]">
      <el-card class="card">
        <el-row class="imgWrapper">
          <img src="@/assets/img/img.png" class="img_circle">
          <h1 class="userName">{{msg.username}}</h1>
        </el-row>
        <el-row>
          <p>{{msg.columnintro}}</p>
        </el-row>
        <el-row class="text_Wrapper">
          <el-button type="primary"  class="write_btn" @click="toColumnPage(msg.userid)">进入专栏</el-button>
        </el-row>
      </el-card>
    </el-col>

  </el-row>




<!--加载更多功能-->
  <el-row class="text_Wrapper">
    <el-button type="primary"  class="write_btn load_more" @click="loadMove">加载更多</el-button>
  </el-row>

</div>
</template>

<script>
import { Loading } from 'element-ui';
export default {
  name: "Main",
  data(){
    return{
      articleMessage:[
        {avatarUrl:require("../../assets/img/img.png"),
         userName:"夏夏",
         articleIntro:"123"},
        {avatarUrl:require("../../assets/img/img_1.png"),
          userName:"邹邹",
          articleIntro:"123"},
        {avatarUrl:require("../../assets/img/img_2.png"),
          userName:"盼盼",
          articleIntro:"123"},

      ],

    }
  },
  methods:{
    toColumnPage(userid){
      this.$store.commit("getUserInfo")
      this.$router.push({
        name:"Column",
        params:{
          id:userid
        }
      })
      // location.reload()
    },
    beginWriteArticle(){

      this.$router.push("/create")
    },
    loadMove(){

      this.$store.state.loadMoveUsers.forEach(e=>{
        console.log(e.username)
      })
      let loadingInstance = Loading.service({ fullscreen: true }); Loading.service({ fullscreen:true  });
      setTimeout(()=>{
        loadingInstance.close()
        this.$store.commit("loadMove")
        console.log(this.$store.state.loadMoveUsers[0][0])
      },1000)
    }
  }
}
</script>

<style scoped>
.load_more{
  width:300px ;
  margin-bottom: 20px
}
.write_btn:hover{
  background: #0D6EFD;
  color: white;
}
.write_btn{
  margin-top: 30px;
  background: white;
  border: 1px solid #0D6EFD;
  color: #0D6EFD;
}
.text_Wrapper{
  text-align: center;

}
.userName{
  margin-top: 10px;
  font-weight: bolder;
  font-size: 20px;
}
.articleImg{
  width: 270px;
  height: 225px;
}
.imgWrapper{
text-align: center;
  margin-top: 50px;
}
.articleFont{
  font-size: 34px;
}
.btn{
  background-color: #0D6EFD;
  margin-top: 20px;
}
.btn:hover{
  background-color: #0257d5;
}
.findColor{
  font-size: 24px;
  font-weight: bolder;
  margin-top: 40px;

}
.card{

 height: 240px;
  width: 400px;
  margin: 0 auto;
  margin-top: 20px;
  margin-bottom: 20px;

}
.img_circle{
  width: 50px;
  height: 50px;
  border-radius: 50%;
  margin-top: -20px;

}
</style>