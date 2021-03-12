<template>
  <div class="form-wrapper">
    <h1 class="title">注册B乎账户</h1>
    <el-form        :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="邮箱地址" prop="email">
        <el-input type="text" v-model="ruleForm.email" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="昵称" prop="nick">
        <el-input type="text" v-model="ruleForm.nick" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="pass">
        <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="确认密码" prop="checkPass">
        <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="年龄" prop="age">
        <el-input v-model.number="ruleForm.age"></el-input>
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
  name: "Register",
  data() {
    var checkAge = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('年龄不能为空'));
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值'));
        } else {
          if (value < 18) {
            callback(new Error('必须年满18岁'));
          } else {
            callback();
          }
        }
      }, 1000);
    };
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
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
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
    var checkNick = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('昵称不可为空'));
      } else {
        callback();
      }
    };
    return {
      ruleForm: {
        pass: '',
        checkPass: '',
        age: '',
        email:'',
        nick:''
      },
      accountExits:false,
      registerOk:false,
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, trigger: 'blur' }
        ],
        age: [
          { validator: checkAge, trigger: 'blur' }
        ],
        email:[
          { validator: checkEmail, trigger: 'blur' }
        ],
        nick: [
          { validator: checkNick, trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$store.state.axiosConfig.registerCfg.params={
            email:this.ruleForm.email,
            username:this.ruleForm.nick,
            password:this.ruleForm.pass,
            age:this.ruleForm.age
          }
          this.$axios.get(this.$store.state.url.registerUrl,this.$store.state.axiosConfig.registerCfg).then(res=>{
            console.log(res.data)
            this.accountExits=res.data["exits"]
            this.registerOk=res.data["isRegister"]
            if(this.accountExits){

              this.$message.error({
                message: '账号已被注册',

              });

              this.ruleForm= {}

            }
            if(this.registerOk){

              this.$message({
                message: '恭喜你，注册成功,正在跳转首页',
                type: 'success'
              });
              this.ruleForm= {}
              setTimeout(()=>{this.$router.push("/")},2000)

            }
          }).catch(err=>{
            console.log("注册失败!")
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