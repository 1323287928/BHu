<template>

<div class="wrapper">
  <input id="fileData" name="checkPizhu" type="hidden">
  <el-col :span="24" class="marginTop">
    <h1 class="title">{{title}}</h1>
  </el-col>

  <el-col :span="24" class="marginTop">
    <div class="cover"  >

      <UploadIcon class="uploadIcon" ref="UploadIcon"></UploadIcon>


    </div>

  </el-col>
  <el-col :span="24" class="marginTop">


    <h1>文章标题:</h1>
    <br>
    <el-input v-model="articleTitle"  placeholder="请输入文章标题" ></el-input>

  </el-col>

  <el-col :span="24" class="marginTop">
    <h1>文章详情:</h1>
    <br>
    <el-input  class="articleDetail" v-model="articleDeatail"  :rows="10" type="textarea"  placeholder="请输入文章详情"></el-input>
  </el-col>
  <el-col :span="24" class="marginTop">
    <el-button class="btn_submit" @click="submitArticle" >{{btnMsg}}</el-button>
  </el-col>

</div>
</template>

<script>
import UploadIcon from "@/components/UploadIcon/UploadIcon";
export default {
name: "Create",
  components:{
  UploadIcon
  },
  data(){
  return{
    articleDeatail:'',
    articleTitle:'',
    articleId:"",
    btnMsg:"发表文章",
    title:"新建文章",
  }
  },
  methods:{
    submitArticle(){
      if(this.articleDeatail.trim()==""||this.articleTitle.trim()==""){

        this.openNotificationError()
        return;
      }
      else {
        if(this.btnMsg=="发表文章")
        {
          this.openNotificationSuccess()
          return;
        }else{
         this.editColumn(this.articleId)
        }


      }
    },
    openNotificationError() {
      const h = this.$createElement;

      this.$notify.error({
        title: '注意!',
        message: h('i', { style: 'color: teal'}, '文章标题或者文章详情不能为空'),
        duration:2000
      });
    },
    openNotificationSuccess(){

      this.$store.state.axiosConfig.createCfg.params={
        userid:JSON.parse(localStorage.getItem("currentUser")).userid,
        articleicon:this.$refs.UploadIcon.currentFile.url,
        articletitle:this.articleTitle,
        articledetail:this.articleDeatail
      }
    this.$axios(this.$store.state.url.createUrl,this.$store.state.axiosConfig.createCfg).then(res=>{
        console.log(res)
      if(res.data["created"]){
        this.$notify({
          title: '成功',
          message: '文章提交成功',
          type: 'success',
          duration:2000
        });
        this.articleDeatail=''
            this.articleTitle=''
      }else{
        this.$notify.error({
          title: '注意!',
          message: h('i', { style: 'color: teal'}, '提交失败,请联系管理员'),
          duration:2000
        });
        this.articleDeatail=''
        this.articleTitle=''
      }
    })
      console.log(this.articleTitle,this.articleDeatail,this.$refs.UploadIcon.currentFile.url)
    },
    editColumn(id){
      console.log( this.$store.state.editArticle)
      this.$store.state.axiosConfig.editArticleCfg.params.articleid=id
      this.$store.state.axiosConfig.editArticleCfg.params.articledetail=this.articleDeatail
      this.$store.state.axiosConfig.editArticleCfg.params.articletitle=this.articleTitle
      this.$store.state.axiosConfig.editArticleCfg.params.articleicon=""
      this.$axios(this.$store.state.url.editArticleUrl,this.$store.state.axiosConfig.editArticleCfg).then(res=>{
      if(res.data["isEdit"]){
        this.$notify({
          title: '成功',
          message: '文章修改成功',
          type: 'success',
          duration:2000
        });
      }else{
        this.$notify.error({
          title: '注意!',
          message: h('i', { style: 'color: teal'}, '修改失败,请联系管理员'),
          duration:2000
        });
      }
     })
    }
  },
  created() {
    if(this.$store.state.editArticle){
    this.articleTitle=this.$store.state.editArticle.articletitle
      this.articleDeatail=this.$store.state.editArticle.articledetail
      this.articleId=this.$store.state.editArticle.articleid
      this.btnMsg="修改文章"
      this.title="修改文章"
    }
  },
  destroyed() {
   console.log(11)
    this.btnMsg="发表文章"
    this.title="新建文章"
  }
}
</script>

<style scoped>

.filepath {
  width: 100%;
  height: 100%;
  opacity: 0;
}
.articleDetail{
  width: 100%;

}
.uploadTitle{
  font-size: 30px;
  color: #6C757D;
  cursor: pointer;
}
.title{
  font-weight: bold;
  font-size: 26px;
}
.cover{
  height: 200px;
  background: #F8F9FA;
 line-height: 200px;
  text-align: center;
  padding-top: 50px;
}
.wrapper{
  width: 1200px;
  margin: 0 auto;
  margin-top: 20px;

}
.btn_submit{
  background-color:#0D6EFD;
  color: white;
}
.marginTop{
  margin-top: 20px;
}
</style>