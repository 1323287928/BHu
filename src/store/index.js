import Vue from 'vue'
import Vuex from 'vuex'
import axios from "axios";

Vue.use(Vuex)
export default new Vuex.Store({
  state: {
    isLogin:"",
    loadMoveSet:{
      start:0,
      end:0
    },
    currentUser:{
      userId:"",
      email:"",
      userName:'',
      password:"",
      age:"",
      icon:"",
      columnname:"",
      columnintro:"",
      personnalintro:"",

    },
    currentArticle:{
      articleId:"",
      userId:"",
      articleIcon:"",
      articleTitle:"",
      articleDetail:""
    },
    editArticle:{
      articleId:"",
      userId:"",
      articleIcon:"",
      articleTitle:"",
      articleDetail:""
    },

    myColumn:
        [


        ]
        ,
    url:{
      registerUrl:"http://localhost:8080/register",
      loginUrl:"http://localhost:8080/login",
      createUrl:"http://localhost:8080/create",
      columnUrl:"http://localhost:8080/column",
      userinfoUrl:"http://localhost:8080/userinfo",
      editArticleUrl:"http://localhost:8080/editarticle",
      delArticleUrl:"http://localhost:8080/delarticle",
      editPersonalIntroUrl:"http://localhost:8080/editpersonalintro",
      editColumnIntroUrl:"http://localhost:8080/editcolumnintro",
      loadMoveUrl:"http://localhost:8080/loadmove"
    },
    axiosConfig:{
      registerCfg:{
        params:{
          email:"",
          username:"",
          password:"",
          age:""
        }
      },
    editArticleCfg:{
      params:{
        articledetail: "",
articleicon: "",
articleid: "",
articletitle: ""


      }
    },
      loginCfg:{
        params: {
          email: "",
          password: ""
        },


      },
      userInfoCfg:{
        params:{
          userid:"",
          email:"",
          username:"",
          password:"",
          age:"",
          icon:"",
          columnname: "",
          columnintro:""
        }
      },
      createCfg:{
        params:{
          userid:"",
          articleicon:"",
          articletitle:"",
          articledetail:""
        }
      },
      columnCfg:{
        params:{
          userid:""
        }
      },
      editPersonalIntroCfg:{
        params:{
          userid:"",
          username:"",
          personalintro:""
        }
      },
      editColumnIntroCfg:{
        params:{
          userid:"",

        }
      }
    },
    sqlLimitCfg:{
      start:"",
      limitNum:"",
    },
    loadMoveUsers:[],
    beforeloadMoveUsers:[],

  },

  mutations: {
  seteditArticle(state){
    state.editArticle=""
  },
  getMyColumns(state,id){
    let userid=""
    if(id){
      userid=id
    }else{
      userid =parseInt(JSON.parse(localStorage.getItem("currentUser")).userid)
    }
    state.axiosConfig.columnCfg.params.userid=userid
    axios.get(state.url.columnUrl,state.axiosConfig.columnCfg).then(res=>{
      state.myColumn=res.data.columns
      console.log( state.myColumn)


    })
  },
    getUserInfo(state,id){
      let userid
    if(id){
      userid=id
    }else {
      userid=parseInt(JSON.parse(localStorage.getItem("currentUser")).userid)
    }
      state.axiosConfig.userInfoCfg.params.userid=userid
      axios.get(state.url.userinfoUrl,state.axiosConfig.userInfoCfg).then(res=>{
      console.log(res)
            state.currentUser.userId=res.data.userid,
            state.currentUser.email=res.data.email,
            state.currentUser.userName=res.data.username
            state.currentUser.password=res.data.password
            state.currentUser.age=res.data.age
            state.currentUser.icon=res.data.icon
            state.currentUser.columnname=res.data.columnname
            state.currentUser.columnintro=res.data.columnintro
            state.currentUser.personnalintro=res.data.personalintro

      })

    },
    loadMove(state){
      if(state.loadMoveUsers.length!=state.beforeloadMoveUsers.length)
      {

         let count=0
         count=state.beforeloadMoveUsers.length/6!=0?6:state.beforeloadMoveUsers%6
        state.loadMoveUsers.push(state.beforeloadMoveUsers.slice(state.loadMoveSet.start,count))
        console.log(state.beforeloadMoveUsers)
        console.log(state.beforeloadMoveUsers.slice(state.loadMoveSet.start,count))
        console.log(state.loadMoveUsers)
      }

    },
    beforeLoadMove(state){
      axios.get(state.url.loadMoveUrl).then(res=>{
        console.log(res.data["moveuser"])
        res.data["moveuser"].forEach(element=>{
          state.beforeloadMoveUsers.push(element)
        })
      })

    }
  },
  actions: {
  },
  modules: {
  }
})
