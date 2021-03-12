import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from "@/views/Main/Main";
import Login from "@/views/Login/Login";
import Register from "@/views/Register/Register";
import Column from "@/views/Column/Column";
import Create from "@/views/Create/Create";
import Edit from "@/views/Edit/Edit";
import Posts from "@/views/Posts/Posts";
Vue.use(VueRouter)

const routes = [
  {
     path: '/',
    name: 'Main',
     component:Main
  },
  {
    path: '/login',
    name: "Login",
    component: Login
  },
  {
    path: "/register",
    name: "Register",
    component: Register
  },
  {
    path: "/column",
    name: "Column",
    component: Column
  },
  {
    path: "/create",
    name: "Create",
    component: Create
  },
  {
    path:"/edit",
    name:"Edit",
    component: Edit
  },
  {
    path:"/posts",
    name:"Posts",
    component: Posts,


  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
router.beforeEach((to,from,next)=>{
  if(!JSON.parse(localStorage.getItem("currentUser")))
  {
    if(to.path=="/login"||to.path=="/register"||to.path=="/"){
      next()
    }else{
      next("/")
    }
  }
  else {

    next()
  }
})
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
export default router
