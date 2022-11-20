<template>
  <div class="note" :style="note">
    <el-row>
      <el-col :span="24">
        <div class="grid-content bg-purple-dark">
<!--          <p class="hunter">人事管理系统</p>-->
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="20">

      <el-col :span="12">
        <div>
          <el-tabs type="border-card" class="tabs">
            <el-tab-pane label="加入我们">
              <p class="font">区块链人才计划</p><br/>
              <p class="font2">点击下方按钮，加入本公司人才库</p>
              <el-button class="joinUs" size="normal" type="success" @click="joinUs" style="width: 100%">加入我们</el-button>
            </el-tab-pane>
            <el-tab-pane label="登录">
              <el-form
                  :rules="rules"
                  ref="loginForm"
                  v-loading="loading"
                  element-loading-text="正在登录..."
                  element-loading-spinner="el-icon-loading"
                  element-loading-background="rgba(0, 0, 0, 0.8)"
                  :model="loginForm"
                  class="loginContainer">
                <h3 class="loginTitle">系统登录</h3>
                <el-form-item prop="id">
                  <el-input size="normal" type="text" v-model="loginForm.id"
                            placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item prop="password">
                  <el-input size="normal" type="password" v-model="loginForm.password"
                            placeholder="请输入密码"></el-input>
                </el-form-item>
                <!--            <el-form-item prop="code">-->
                <!--                <el-input size="normal" type="text" v-model="loginForm.code" auto-complete="off"-->
                <!--                          placeholder="点击图片更换验证码" @keydown.enter.native="submitLogin" style="width: 250px"></el-input>-->
                <!--                <img :src="vcUrl" @click="updateVerifyCode" alt="" style="cursor: pointer">-->
                <!--            </el-form-item>-->
<!--                <el-checkbox size="normal" class="loginRemember" v-model="checked"></el-checkbox>-->
                <tr>
                  <td class="login-title" style="justify-content: center">
                    身份
                  </td>
                  <td>
                    <el-radio v-model="type" :label="1" border>主管</el-radio>
                    <el-radio v-model="type" :label="2" border>员工</el-radio>
                  </td>
                </tr>
                <Vcode :show="isShow" @success="success" @close="close"></Vcode>
                <el-button-group style="width: 100%">
                  <el-button size="normal" type="danger" style="width: 50%;" @click="submit" :round="true" :disabled="right">人机验证</el-button>
                  <el-button size="normal" type="primary" style="width: 50%;" @click="submitLogin" :round="true" :disabled="!right">登录</el-button>
                </el-button-group>

              </el-form>
            </el-tab-pane>
          </el-tabs>

        </div>
      </el-col>
    </el-row>
    <div>
        <!--      输入框-->
        <el-dialog
            :title="title"
          @close="resetForm"
            append-to-body
          :visible.sync="dialogAddVisible"
          width="80%">
        <div>
          <el-form :model="emp" :rules="talentRules" ref="talentForm">
            <el-row>
              <el-col :span="6">
                <el-form-item label="姓名:" prop="name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user-solid" v-model="emp.name"
                            placeholder="请输入员工姓名" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="性别:" prop="sex">
                  <el-radio-group v-model="emp.sex">
                    <el-radio label="男">男</el-radio>
                    <el-radio label="女">女</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="出生日期:" prop="birthday">
                  <el-date-picker
                      v-model="emp.birthday"
                      size="mini"
                      type="date"
                      value-format="yyyy-MM-dd"
                      style="width: 150px;"
                      placeholder="出生日期">
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="身份证号:" prop="id_card">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-s-custom" v-model="emp.id_card"
                            placeholder="请输入身份证号" clearable></el-input>
                </el-form-item>
              </el-col>

            </el-row>

            <el-row>
              <el-col :span="6">
                <el-form-item label="政治面貌:" prop="political">
                  <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;">
                    <el-option
                        v-for="item in politicals"
                        :key="item"
                        :label="item"
                        :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="民族:" prop="nation">
                  <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;" filterable>
                    <el-option
                        v-for="item in nations"
                        :key="item"
                        :label="item"
                        :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="籍贯:" prop="native_place">
                  <el-input size="mini" style="width: 120px" prefix-icon="el-icon-suitcase"
                            v-model="emp.native_place" placeholder="请输入籍贯" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="电话号码:" prop="phone">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-phone"
                            v-model="emp.phone" placeholder="电话号码" clearable></el-input>
                </el-form-item>
              </el-col>

            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="电子邮箱:" prop="email">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                            v-model="emp.email" placeholder="请输入电子邮箱" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="身高:" prop="height">
                  <el-input size="mini" style="width: 100px" type="number" prefix-icon="el-icon-edit"
                            v-model="emp.height" placeholder="请输入身高" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="血型:" prop="blood_type">
                  <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;">
                    <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="婚姻状况:" prop="marital_status">
                  <el-radio-group v-model="emp.marital_status">
                    <el-radio label="已婚">已婚</el-radio>
                    <el-radio label="未婚">未婚</el-radio>
                    <el-radio label="离异">离异</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-col>

            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="出生地:" prop="birthplace">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-s-opportunity"
                            v-model="emp.birthplace" placeholder="请输入出生地" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="户口所在地:" prop="registered_residence">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-house"
                            v-model="emp.registered_residence" placeholder="请输入户口所在地" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门名称:" prop="department_name">
                  <el-select @change="changeEditSelect" v-model="emp.department_name" placeholder="请输入部门名称" size="mini" prefix-icon="el-icon-edit">
                    <el-option
                        v-for="(item, index) in department_names"
                        :key="index"
                        :label="item"
                        :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="岗位名称:" prop="post_name">
                  <el-select v-model="emp.post_name" placeholder="请输入岗位名称" size="mini" prefix-icon="el-icon-edit">
                    <el-option
                        v-for="(item, index) in post_type_options"
                        :key="index"
                        :label="item"
                        :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>

            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="用工形式:" prop="employment_form">
                  <el-select v-model="emp.employment_form" placeholder="用工形式" size="mini" style="width: 150px;">
                    <el-option v-for="item in employee_forms" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="人员来源:" prop="personnel_source">
                  <el-select v-model="emp.personnel_source" placeholder="人员来源" size="mini" style="width: 150px;">
                    <el-option v-for="item in personnel_sources" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="最高学历:" prop="highest_education">
                  <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"
                             style="width: 150px;">
                    <el-option
                        v-for="item in highest_educations"
                        :key="item"
                        :label="item"
                        :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="毕业院校:" prop="graduation_school">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-school"
                            v-model="emp.graduation_school" placeholder="毕业院校名称" clearable></el-input>
                </el-form-item>
              </el-col>

            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="所学专业:" prop="major_studied">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-notebook-1"
                            v-model="emp.major_studied" placeholder="请输入专业名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="毕业日期" prop="graduation_date">
                  <el-date-picker
                      v-model="emp.graduation_date"
                      size="mini"
                      type="date"
                      value-format="yyyy-MM-dd"
                      style="width: 150px;"
                      placeholder="毕业日期">
                  </el-date-picker>

                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogAddVisible = false; resetForm">取 消</el-button>
          <el-button type="primary" @click="doAddTalent">提 交</el-button>
        </span>
      </el-dialog>
    </div>
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
      title : '加入我们',
      dialogAddVisible : false,
      emp: {
        id : "",
        name : "",
        sex : "",
        birthday : "",
        id_card : "",
        political : "",
        nation : "",
        native_place : "",
        phone : "",
        email : "",
        height : null,
        blood_type :"",
        marital_status : "",
        birthplace : '',
        registered_residence : "",
        department_name : "",
        post_name : "",
        employment_form : "",
        personnel_source : "",
        highest_education : "",
        graduation_school : "",
        major_studied : "",
        graduation_date : "",
      },
      talentRules: {
        id : [{required : true, message : '请输入编号', trigger : 'blur'}],
        password : [{required : true, message : '请输入密码', trigger : 'blur'}],
        employee_type : [{required : true, message : '请输入员工类型', trigger : 'blur'}],
        name: [{required : true, message : '请输入姓名', trigger : 'blur'}],
        sex: [{required : true, message : '请输入性别', trigger : 'blur'}],
        birthday: [{required : true, message : '请输入出生日期', trigger : 'blur'}],
        id_card: [{required : true, message : '请输入身份证号码', trigger : 'blur'}, {
          pattern : /(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$)/,
          message : '身份证号码格式不正确',
          trigger : 'blur'
        }],
        political : [{required : true, message : '请输入政治面貌', trigger : 'blur'}],
        nation : [{required : true, message : '请输入民族', trigger : 'blur'}],
        native_place : [{required : true, message : '请输入籍贯', trigger : 'blur'}],
        phone : [{required : true, message : '请输入电话号码', trigger : 'blur'}],
        email : [{required : true, message : '请输入邮箱地址', trigger : 'blur'}, {
          type : 'email',
          message : '邮箱格式不正确',
          trigger : 'blur'
        }],
        height : [{required : true, message : '请输入身高', trigger : 'blur'}],
        blood_type : [{required : true, message : '请输入血型', trigger : 'blur'}],
        marital_status : [{required : true, message : '请输入婚姻状况', trigger : 'blur'}],
        birthplace : [{required : true, message : '请输入出生地', trigger : 'blur'}],
        registered_residence : [{required : true, message: '请输入户口所在地', trigger : 'blur'}],
        department_number : [{required : true, message : '请输入部门编号', trigger : 'blur'}],
        department_name : [{required : true, message : '请输入部门名称', trigger : 'blur'}],
        post_number : [{required : true, message : '请输入岗位编号', trigger : 'blur'}],
        post_name : [{required : true, message : '请输入岗位名称', trigger : 'blur'}],
        entry_date : [{required : true, message : '请输入入职日期', trigger : 'blur'}],
        employment_form : [{required : true, message : '请输入用工形式', trigger: 'blur'}],
        personnel_source : [{required : true, message : '请输入人员来源', trigger : 'blur'}],
        highest_education : [{required : true, message : '请输入最高学历', trigger : 'blur'}],
        graduation_school : [{required : true, message : '请输入毕业院校', trigger : 'blur'}],
        major_studied : [{required : true, message : '请输入所学专业', trigger : 'blur'}],
        graduation_date : [{required : true, message : '请输入毕业日期', trigger : 'blur'}],
        is_quit : [{required : true, message : '请输入是否离职', trigger : 'blur'}],
      },
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      employee_types:['总裁', '主管', '员工'],
      nations: [
        '蒙古族', '藏族', '苗族', '壮族', '回族', '维吾尔族', '彝族', '布依族',
        '朝鲜族', '侗族', '白族', '哈尼族', '傣族', '傈僳族', '畲族', '拉祜族',
        '满族', '瑶族', '土家族', '哈萨克族', '黎族', '佤族', '高山族', '水族',
        '东乡族', '景颇族', '土族', '仫佬族', '布朗族', '毛南族', '锡伯族', '普米族',
        '纳西族', '柯尔克孜族', '达斡尔族', '羌族', '撒拉族', '仡佬族', '阿昌族', '塔吉克族',
        '怒族', '俄罗斯族', '德昂族', '裕固族', '塔塔尔族', '鄂伦春族', '门巴族', '基诺族',
        '乌孜别克族', '鄂温克族', '保安族', '京族', '独龙族', '赫哲族', '珞巴族', '汉族',
        '其他',
      ],
      politicals: ['群众', '共青团员', '中共预备党员', '中共党员', '无党派人士', '其他'],
      highest_educations: ['小学', '初中', '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
      employee_forms:['实习生', '正式职工'],
      personnel_sources:['校招', '社招'],
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      // images : [
      //   // {url : require("../assets/tabs/tab1.png")},
      //   // {url : require("../assets/tabs/tab2.png")},
      //   // {url : require("../assets/tabs/tab3.png")},
      //   // {url : require("../assets/tabs/tab4.png")},
      // ],
      bannerHeight : "",
      note: {
        // backgroundImage: "url(" + require("../assets/loginBackgroundImg3.png") + ")",//背景圖片
        backgroundRepeat: "no-repeat",
        backgroundPosition: "auto auto",
        backgroundSize: "100% 100%",
        height: "100%",
        width: "100%",
        position: "fixed",
        marginTop: "5px",
      },
      loading: false,
      type: 1,
      vcUrl: '/verifyCode?time='+new Date(),
      loginForm: {//登录表单
        id: '',
        password: '',
        employee_type: localStorage.getItem("role"),
      },
      checked: true,
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
    // this.imgLoad();
    window.addEventListener('resize',() => {
      this.bannerHeight=this.$refs.bannerHeight[0].height;
      this.imgLoad();
    },false)
    this.initDepartment();
    this.initPost();
  },
  methods: {
    updateVerifyCode() {
      this.vcUrl = '/verifyCode?time='+new Date();
    },
    submit () {
      this.isShow = true;
    },
    submitLogin() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          // type:1是老板 2是主管 3是员工
          let str = '';
          if (this.type == 1) {
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

            this.$axios.post('/AdminLogin', this.loginForm).then((resp) => {
              this.loading = false;
              const _this = this;
              console.log("resp = ", resp.data)
              if (resp) {
                if (resp.data.msg == "登录成功") {
                  console.log("222")
                  this.$store.commit("INIT_CURRENTHR", resp.data.data);
                  console.log("333")
                  window.sessionStorage.setItem("user", JSON.stringify(resp.data.data));
                  console.log("存储的数据："+JSON.stringify(resp.data.data));
                  let path = this.$route.query.redirect;
                  this.$router.replace((path == '/' || path == undefined) ? '/home' : path);
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
    imgLoad(){
      this.$nextTick(()=>{
        this.bannerHeight=this.$refs.bannerHeight[0].height;
        console.log(this.$refs.bannerHeight[0].height);
        // document.getElementsByClassName拿到的是数组并非某一个对象
        var testH=document.getElementById("test-div");
        // testH.style.height= this.bannerHeight+'px'; //报错
      })
    },
    joinUs() {
      this.dialogAddVisible = true;
    },
    resetForm(data) {
      this.$nextTick(() => {
        this.$refs['talentForm'].resetFields();
      })
    },
    changeEditSelect() {
      // 清空部门内容
      this.emp.post_name = ''

      // 遍历部门的下拉选项数组
      for (const k in this.department_names) {
        // 岗位名称 是否等于 部门名称的下拉选择数组中的某一项
        if (this.emp.department_name === this.department_names[k]) {
          this.post_type_options = this.post_names[this.emp.department_name]
        }
      }
      console.log("post_type_options = "+this.post_type_options);
    },
    doAddTalent() {
      this.$refs['talentForm'].validate(valid => {
        if (valid) {
          console.log(this.emp);
          this.emp.height = parseInt(this.emp.height);
          this.$axios.post('/TalentPool/Add', this.emp).then(resp => {
            if (resp) {
              this.dialogAddVisible = false;
              if (resp) {
                Message.success({message : resp.data.msg});
              } else {
                Message.error({message : resp.data.msg})
              }
            } else {
              Message.error({message : resp.data.msg});
            }
          })
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
      this.$axios.post('/Department/Init').then(resp => {
        this.department_names = resp.data;
        console.log("初始化部门：", this.department_names);
      })
    },
    initPost() {
      var temp = {
        "department_name" : this.department_name
      }
      this.$axios.post('/Post/Init', temp).then(resp => {
        this.post_names = resp.data;
        console.log("初始化岗位：", this.post_names);
      })
    }
  }
}
</script>

<style>
.tabs {
  border-radius: 15px;
  background-clip: padding-box;
  /*margin:  auto;*/
  /*width: 500px;*/
  padding: 75px 35px 15px 35px;
  background: #ffffff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
  text-align: center;

  width: 500px;
  height: 350px;
  position: absolute;
  left: 50%;
  /*top: 50%;*/
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

.loginRemember {
  text-align: left;
  margin: 0px 0px 15px 0px;
}
.el-form-item__content{
  display: flex;
  align-items: center;
}



.font {
  font-family : 华文行楷;
  color: #6e6572;
  text-align: center;
  font-size: 30px;
}
.font2 {
  font-family : "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  color: #0fb9e3;
  text-align: center;
  font-size: medium;
}
.hunter {
  font-family : 幼圆;
  letter-spacing : 50px;
  color: #ff0000;
  text-align: center;
  font-size: 50px;
  font-weight: bolder;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.joinUs {
  font-family: 华文隶书;
  color: #600fe3;
  font-size: larger;
}
</style>
