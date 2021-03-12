<template>

    <div class="wrapper">
      <el-col :span="24" class="nav">
        <a  class="a">首页</a>/<a  class="a">专栏页</a>/{{this.currentArticle.articletitle}}
      </el-col>
      <el-col :span="24" class="title">
        <h1>{{this.currentArticle.articletitle}}</h1>
      </el-col>
      <el-col >
        <hr>
        <div class="user">
          <el-col :span="2">
            <img src="../../assets/logo.png" class="img">
          </el-col>

          <el-col :span="12">
            <el-row>{{this.currentUser.userName}}</el-row>
            <el-row class="font">{{this.currentUser.columnintro}}</el-row>
          </el-col>

        </div>
        <hr>
      </el-col>
      <el-col :span="24" class="detail">
        {{this.currentArticle.articledetail}}
      </el-col>

      <el-col :span="24">
        <el-row class="btngroup">
          <el-button type="primary" icon="el-icon-edit"  @click="toCreate" circle :title="'修改'" ></el-button>

          <el-button type="danger" icon="el-icon-delete"  @click="del" circle :title="'删除'"></el-button>
        </el-row>
      </el-col>

    </div>


</template>

<script>
export default {
name: "Posts",
  data(){
   return{
     currentUser:"",
     currentArticle:""
   }
  },
  methods:{
  toCreate(){
   this.$store.commit("seteditArticle")
 this.$store.state.editArticle=this.currentArticle
    console.log( this.$store.state.editArticle)
   this.$router.push("/create")
  },
    del(){
   console.log(this.currentArticle.articleid)
      this.$axios(this.$store.state.url.delArticleUrl,{params:{
          articleid:this.currentArticle.articleid
        }}).then(res=>{
          if(res.data["isDel"]){
            this.$message({
              message: '恭喜你，删除成功',
              type: 'success'
            });
            this.$router.push("/column")
            location.reload()
          }
          else{
            this.$message({
              message: '删除失败，请联系管理员',
              type: 'warning'
            });
          }

      })
    }
  },

  created() {

  this.currentArticle=JSON.parse(sessionStorage.getItem("currentArticle"))
    this.currentUser=JSON.parse( sessionStorage.getItem("currentUser"))


  }
}
</script>

<style scoped>
.font{
  color: gray;
}
a:hover{
  color: #0D6EFD;
  cursor: pointer;
}
.detail{
  margin-top: 50px;
  margin-bottom: 50px;
}
.btngroup{
 text-align: center;
}
.wrapper{
  width: 800px;
  height: auto;
  margin: 0 auto;
  margin-top: 10px;

}
.nav{
  background-color: #E9ECEF ;
  padding-left:20px;
 padding-top: 20px;
  padding-bottom: 20px;
  border-radius: 10px;

}
.title{
  width: 800px;
  padding-left:20px;
  padding-top: 20px;
  padding-bottom: 20px;
  border-radius: 10px;
  font-size: 25px;
  font-width: bolder;
}
.user{
  padding-top: 20px;
  padding-bottom: 20px;
  height: 20px;

}
.img{
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: 1px gray dashed;
}
</style>