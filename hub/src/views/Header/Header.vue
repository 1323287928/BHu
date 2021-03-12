<template>
  <div>
    <el-row>
      <el-col :span="8">
        <div class="header">
          <span @click="toHomePage" class="home">帅气专栏</span>
        </div>
      </el-col>
      <el-col :span="8">
        <div class="header"></div>
      </el-col>
      <el-col :span="8">
        <div class="header" v-if="isLogin">
          <el-dropdown>
            <el-button type="primary" class="btn" >
              你好 <span class="userMargin"></span>{{currentUser.username}}   <i class="el-icon-arrow-down el-icon--right"></i>
            </el-button>

            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item ><span @click="toCreateArticle">新建文章</span></el-dropdown-item>
              <el-dropdown-item><span @click=" toColumnPage">我的专栏</span></el-dropdown-item>
              <el-dropdown-item> <span @click="toEditPage">编辑资料</span></el-dropdown-item>
              <el-dropdown-item> <span @click="logOut">退出登录</span></el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
        <div class="header" v-if="!isLogin">
          <el-button type="primary" class="btn" @click="toLoginPage">登录</el-button>
          <el-button type="primary"  class="btn" @click="toRegisterPage">注册</el-button>
        </div>


      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "Header",
  data(){
    return{

      isLogin:""  ,
      currentUser:""
    }
  },
  methods:{
    getColumn(){
console.log(1)
    },
    toLoginPage(){
      this.$router.push("/login")
    },
    toRegisterPage(){
      this.$router.push("/register")
    },
    toHomePage(){
      this.$router.push("/")
    },
    toCreateArticle(){

     this.$router.push("/create")


    },
    toColumnPage(){
        this.$router.push("/column")
      location.reload()
    },
    toEditPage(){
      this.$router.push("/edit")
    },
    logOut(){
      this.$message({
        message: '退出登录成功，正在转向首页',
        type: 'success'
      });
      localStorage.removeItem("isLogin")
      localStorage.removeItem("currentUser")
      this.$router.push("/")
      location.reload()

    }
  },
  created() {
    this.isLogin=JSON.parse(localStorage.getItem("isLogin")) ,
      this.currentUser=JSON.parse(localStorage.getItem("currentUser"))
  }



}
</script>

<style scoped>
.userMargin{
  margin-left: 3px;
}
.header{
  background-color: #0D6EFD;
  height: 80px;
  width: 100%;
line-height: 80px;
  text-align: center;
  color: white;
  font-size: 20px;
}
.btn{
  background-color: #0D6EFD;
}
.btn:hover{
  background: white;
  color: black;
}
.home{
  display: block;
  height: 100%;
  cursor: pointer;
}
</style>