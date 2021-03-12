<template>
  <div id="app">
    <Header></Header>
  <keep-alive>
    <router-view></router-view>
  </keep-alive>

  </div>

</template>
<script>
import Header from "@/views/Header/Header";
import Main from "@/views/Main/Main";
import Footer from "@/views/Footer/Footer";
import Register from "@/views/Register/Register";
import Login from "@/views/Login/Login";
import Column from "@/views/Column/Column";
import Create from "@/views/Create/Create";
import Edit from "@/views/Edit/Edit";
import Posts from "@/views/Posts/Posts";
export  default {
  name:"App",
  components:{
    Edit,
    Header,
    Main,
    Footer,
    Register,
    Login,
    Column,
    Create,
    Posts
  },
  created() {
    this.$store.state.editArticle=""
    this.$store.state.beforeloadMoveUsers=[]
    if(JSON.parse(localStorage.getItem("currentUser")))
    {
      this.$store.commit("getMyColumns")
      this.$store.commit("getUserInfo")
      this.$store.commit("beforeLoadMove")
    }


  }
}
</script>
<style scoped>
@import "assets/reset-css/reset.css";
#app{
  height: 100%;
}
</style>
