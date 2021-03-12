<template>

  <div class="column_wrapper">
    <el-row >
      <el-col :span="5" >
        <img src="../../assets/img/img.png" class="column_avater">
      </el-col>
      <el-col :span="19">
        <h1 class="userName">{{$store.state.currentUser.userName}}</h1>
        <br>
        <p class="articleIntro">
          {{$store.state.currentUser.columnintro}}
        </p>
      </el-col>
    </el-row>
    <hr>


    <el-row>
      <el-card v-for="(item,index) in $store.state.myColumn " class="column_card">
        <a  @click="goPost(index)"  class="userName article_name">{{item.articletitle}}</a>
        <div class="article_top">
          <el-col :span="7" class="article_img">
          <img src="../../assets/img/img_3.png">
          </el-col>
          <el-col :span="17" >
            <p class="article_detail"> {{item.articledetail}}</p>
          </el-col>
          <el-col>
            <p class="article_top">2020-08-21 19:20:00</p>
          </el-col>
        </div>

      </el-card>
    </el-row>
  </div>


</template>

<script>
export default {
  name: "Column",
  currentUser:{
    userName:"",
    columnintro:"",
  },
  methods:{
    goPost(index){
      this.$store.state.currentArticle= this.$store.state.myColumn[index]
      sessionStorage.setItem("currentArticle",JSON.stringify( this.$store.state.currentArticle))
      sessionStorage.setItem("currentUser",JSON.stringify(this.$store.state.currentUser))
      this.$router.push("/posts")
      location.reload()
    }
  },
  created() {
    console.log(this.$route.params.id)
     this.$store.commit("getMyColumns",this.$route.params.id)
  }

}
</script>

<style scoped>
.article_detail{
  overflow: hidden;

}
.column_card{
  margin-top: 10px;
}
.article_name{
  text-decoration: none;
  color: black;
}
.article_name:hover{
  color: #0D6EFD;
  cursor: pointer;
}
.article_top{
  margin-top: 20px;
  margin-bottom: 20px;
}
.column_wrapper{
margin: 0 auto;
  width: 800px;
  margin-top: 30px;

}
.userName{
  font-size: 24px;
  font-weight: bolder;
}
.column_avater{
 border-radius: 50%;
  height: 150px;
  width: 150px;
}
</style>