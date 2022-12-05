<template>
  <div class="note" :style="note">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark">
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">
      <el-col :span="12">
        <div class="from_cx">
              <el-form
                  :rules="rules"
                  ref="loginForm"
                  v-loading="loading"
                  element-loading-text="正在登录..."
                  element-loading-spinner="el-icon-loading"
                  element-loading-background="rgba(0, 0, 0, 0.8)"
                  :model="loginForm"
                  class="loginContainer">
                <h3 class="loginTitle">人事管理系统</h3>
                <el-form-item prop="id">
                  <el-input size="normal" type="text" v-model="loginForm.id"
                            placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                  <el-input size="normal" type="password" v-model="loginForm.password"
                            placeholder="请输入密码"></el-input>
                </el-form-item>
                <tr>
                  <td>
<!--                    权限分离-->
                    <el-radio v-model="type" :label="1" border >主管</el-radio>
                    <el-radio v-model="type" :label="2" border >员工</el-radio>
                  </td>
                </tr>
                <el-button-group style="width: 100%">
                  <el-button size="normal" type="primary" style="width: 100%;" @click="submitLogin" :round="true" :disabled="right">登录</el-button>
                </el-button-group>
              </el-form>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>

import {Message} from "element-ui";
import Vcode from 'vue-puzzle-vcode'

export default {
  name: "Login",
  data() {
    return {
      isShow : false,
      right : false,
      dialogAddVisible : false,
      bannerHeight : "",
      note: {
        backgroundImage: "url(" + require("../assets/pexels-photo-2088203.webp") + ")",//背景圖片
        backgroundRepeat: "no-repeat",
        backgroundPosition: "auto auto",
        backgroundSize: "100% 100%",
        height: "100%",
        width: "100%",
        position: "fixed",
        marginTop: "5px",
      },
      //登录状态
      loading: false,
      //登录类型
      type: 1,
      vcUrl: '/verifyCode?time='+new Date(),
      loginForm: {//登录表单
        id: '',
        password: '',
        employee_type: localStorage.getItem("role"),
      },
      checked: true,
      //验证
      rules: {
        id: [{required: true, message: '请输入用户名', trigger: 'blur'}],
        password: [{required: true, message: '请输入密码', trigger: 'blur'}],
        code: [{required: true, message: '请输入验证码', trigger: 'blur'}]
      }
    }
  },
  components : {
    Vcode
  },
  created() {
    this.loginForm.employee_type = this.type;
  },
  mounted(){
    this.initDepartment();
  },
  methods: {
    submit () {
      this.isShow = true;
    },
    submitLogin() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          let str = '';
          if (this.type === 1) {
            str = '主管';
          } else  {
            str = '员工';
          }
          localStorage.setItem("type", str);//设置登录人员的身份
          localStorage.setItem("id", this.loginForm.id);
          if ( str === '主管' || str === '员工') {
            localStorage.setItem("role", str);//设置登陆角色
            this.loginForm.employee_type = str;
            console.log("loginForm.type = ", this.loginForm.employee_type);
            console.log("即将提交的信息：", this.loginForm);
            this.$axios.post('/VerifyLogin_cx', this.loginForm).then((resp) => {
              this.loading = false;
              const _this = this;
              console.log("resp = ", resp.data)
              if (resp) {
                if (resp.data.msg === "登录成功") {
                  console.log("aaaa")
                  this.$store.commit("INIT_CURRENTHR", resp.data.data);
                  console.log("ccc")
                  window.sessionStorage.setItem("user", JSON.stringify(resp.data.data));
                  console.log("存储的数据："+JSON.stringify(resp.data.data));
                  let path = this.$route.query.redirect;
                  this.$router.replace((path === '/' || path === undefined) ? '/home' : path);
                  _this.$router.push({path : '/home', query : _this.loginForm})
                } else {
                  Message.error({message: resp.data.msg})
                }
              } else {
                this.vcUrl = '/verifyCode?time='+new Date();
              }
            })
          }
        } else {
          return false;
        }
      });
    },
    success() {
      this.isShow = false;
      this.right = true;
    },
    close() {
      this.isShow = false;
      this.right = false;
    },
    initDepartment() {
      this.$axios.post('/Department/Init_cx').then(resp => {
        this.department_names = resp.data;
        console.log("初始化部门：", this.department_names);
      })
    },
  }
}
</script>

<style>
.from_cx{
  width: 500px;
  height: 400px;
  position: absolute;
  left: 50%;
  transform: translate(-50%,0%);
}
.loginContainer {
  border-radius: 15px;
  background-clip: padding-box;
  margin: auto auto;
  width: 400px;
  padding: 15px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.loginTitle {
  margin: 15px auto 20px auto;
  text-align: center;
  color: #505458;
}
.login-title {
  text-align: center;
  margin: 0 auto auto auto;
  color: #303133;
}
.loginContainer {
  border-radius: 15px;
  background-clip: padding-box;
  margin: 180px auto;
  width: 350px;
  padding: 15px 35px 15px 35px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

</style>
