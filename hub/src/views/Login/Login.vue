<template>
  <div class="form-wrapper">
    <h1 class="title">登录到B乎</h1>
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="邮箱地址" prop="email">
        <el-input type="text" v-model="ruleForm.email" autocomplete="off"></el-input>
      </el-form-item>

      <el-form-item label="密码" prop="pass">
        <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item class="btn_wrapper">
        <el-button class="btn_submit"  @click="submitForm('ruleForm')">提交</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {

    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };

    var checkEmail = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入邮箱账号'));
      } else if (0){
        callback(new Error('邮箱格式不正确!'));
      } else {
        callback();
      }
    };

    return {
      ruleForm: {
        pass: '',
        email:'',

      },
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        email:[
          { validator: checkEmail, trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$store.state.axiosConfig.loginCfg.params={
            email: this.ruleForm.email,
            password: this.ruleForm.pass
          }
          this.$axios.get(this.$store.state.url.loginUrl,this.$store.state.axiosConfig.loginCfg).then(res=>{

           if( res.data.isLogin){
             this.$message({
               message: '恭喜你，登录成功,正在跳转首页',
               type: 'success'
             })
              this.$store.state.currentUser=res.data["user"]
             console.log(this.$store.state.currentUser)
             localStorage.setItem("isLogin",JSON.stringify(res.data.isLogin))
             localStorage.setItem("currentUser",JSON.stringify( this.$store.state.currentUser))
             this.ruleForm={}
             this.$router.push("/")
             this.$store.commit("getUserInfo")
             console.log(this.$store.state.currentUser)
             location.reload()



           }else{
             this.$message.error({
               message: '账号或密码错误',
             })
             this.ruleForm={}
           }
          }).catch(error=>{
            console.log("登录失败")
          })


        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }


}
</script>

<style scoped>
.btn_wrapper{
  text-align: center;
}
.title{
  text-align: center;
  margin-top: 30px;
  margin-bottom: 30px;
  font-weight: bolder;
  font-size: 20px;
}
.form-wrapper{

  margin: 0 auto;
  width: 400px;
}
.btn_submit{
  background-color:#0D6EFD;
  color: white;
}
.btn_submit:hover{
  background-color: #0257d5;

}
</style>