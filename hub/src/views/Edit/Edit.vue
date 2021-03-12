<template>
  <el-tabs  type="card"  class="card">
    <h1 class="editTitle">编辑个人资料</h1>


    <el-tab-pane label="更新个人资料" >
      <el-col :span="24">
      <img src="../../assets/img/newarticle.svg" class="icon">
      </el-col>
      <el-col :span="24">
        <el-input v-model="userName" placeholder="请输入姓名" class="marginTop"></el-input>
      </el-col>
      <el-col :span="24">
        <el-input v-model="selfIntro" placeholder="请输入个人简介" class="marginTop" :rows="10" type="textarea" ></el-input>
      </el-col>
      <el-button type="primary"  class="marginTop btn" @click="editPersonInfo">提交修改</el-button>
    </el-tab-pane>




    <el-tab-pane label="更新专栏信息" name="second">
      <el-col :span="24">
        <img src="" class="icon">
      </el-col>
      <el-col :span="24">
        <el-input v-model="columnName" placeholder="请输入姓名" class="marginTop"></el-input>
      </el-col>
      <el-col :span="24">
        <el-input v-model="columnIntro" placeholder="请输入个人简介" class="marginTop" :rows="10" type="textarea" ></el-input>
      </el-col>
      <el-button type="primary"  class="marginTop btn" @click="editColumnInfo">提交修改</el-button>
    </el-tab-pane>

  </el-tabs>
</template>

<script>
export default {
name: "Edit",
  data(){
  return{
    userName:"1",
    selfIntro: "2",
    columnName:"3",
    columnIntro:"4"
  }
  },
  methods:{
    editPersonInfo(){
      this.$axios(this.$store.state.url.editPersonalIntroUrl,{
        params:{
          userid:this.$store.state.currentUser.userId,
          username:this.userName,
          personalintro:this.selfIntro,
          icon:this.$store.state.currentUser.icon
        }
      }).then(res=>{
      if(res.data.isEditPersonalIntro){
        this.$notify({
          title: '成功',
          message: '修改成功',
          type: 'success'
        });
        this.$router.push("/")

      }
      }).catch(err=>{
        console.log("修改失败",err)
        this.$notify({
          title: '警告',
          message: '修改失败',
          type: 'warning'
        });

      })
    },
    editColumnInfo(){
      this.$axios(this.$store.state.url.editColumnIntroUrl,{
        params:{
          userid:this.$store.state.currentUser.userId,
          columnname:this.columnName,
          columnintro:this.columnIntro
        }
      }).then(res=>{
      if(res.data.isEditColumnIntro){
        this.$notify({
          title: '成功',
          message: '修改成功',
          type: 'success'
        });
        setTimeout(()=>{
          this.$router.push("/")
        },2000)
      }
      }).catch(err=>{
        this.$notify({
          title: '警告',
          message: '修改失败',
          type: 'warning'
        });
      })
    }
  },
  created() {
    if(JSON.parse(localStorage.getItem("currentUser")))
    {
      this.$store.commit("getUserInfo")
      this.userName=this.$store.state.currentUser.userName
      this.selfIntro=this.$store.state.currentUser.personnalintro
      this.columnName=this.$store.state.currentUser.columnname
      this.columnIntro=this.$store.state.currentUser.columnintro
      console.log(this.$store.state.currentUser)
    }
  },
  beforeDestroy() {
    this.$store.commit("getUserInfo")
    console.log(1)
  }
}
</script>

<style scoped>
.card{
  width: 800px;
  margin: 0 auto;
  margin-top: 20px;
}
.editTitle{
  font-size: 24px;
}
.marginTop{
  margin-top: 20px;
}
.btn{
  background-color: #0D6EFD;
}
.icon{
  width: 200px;
  height: 200px;
  border-radius: 50%;
  border: 0.1px  dashed #6C757D;
  cursor: pointer;
  display: block;
  margin: 0 auto;
}
</style>