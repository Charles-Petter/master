<template>
  <div>
    <span>开除员工</span>
    <div>
      <el-form :model="emp" :rules="rules" ref="emp" @submit.native.prevent>
        <el-form-item prop="name">
          <el-input placeholder="请输入员工姓名进行搜索，可以直接回车搜索..." prefix-icon="el-icon-search"
                    clearable
                    @clear="initEmps"
                    style="width: 350px;margin-right: 10px" v-model="emp.name"
                    @keydown.enter.native="searchEmp('emp')" :disabled="showAdvanceSearchView"></el-input>
          <el-button-group>
            <el-button icon="el-icon-search" type="primary" @click="searchEmp('emp')" :disabled="showAdvanceSearchView">搜索</el-button>
            <el-button icon="el-icon-refresh" type="info" @click="resetForm('emp')" :disabled="showAdvanceSearchView">重置</el-button>
          </el-button-group>

          <el-button-group>
            <el-button style="margin-left: 50px;"  type="primary" @click="showAdvanceSearchView = !showAdvanceSearchView">
              <i :class="showAdvanceSearchView?'fa fa-angle-double-up':'fa fa-angle-double-down'"
                 aria-hidden="true"></i>
              高级搜索
            </el-button>
          </el-button-group>

        </el-form-item>
      </el-form>
    </div>
    <div>
      <transition name="slide-fade">
        <div v-show="showAdvanceSearchView"
             style="border: 1px solid #409eff;border-radius: 5px;box-sizing: border-box;padding: 5px;margin: 10px 0px;">
          <el-form :model="searchValue" :rules="rules" ref="searchValue">
            <el-row>

              <el-col :span="6">
                <el-form-item label="政治面貌:" prop="political">
                  <el-select v-model="searchValue.political" placeholder="政治面貌" size="mini"
                             style="width: 130px;">
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
                  <el-select v-model="searchValue.nation" placeholder="民族" size="mini"
                             style="width: 130px;" filterable>
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
                <el-form-item label="部门:" prop="department_name">
                  <el-select @change="changeSelect" v-model="searchValue.department_name" placeholder="部门" size="mini" style="width: 130px;">
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
                <el-form-item label="岗位:" prop="post_name">
                  <el-select v-model="searchValue.post_name" placeholder="岗位" size="mini"
                             style="width: 130px;">
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
            <el-row style="margin-top: 10px">

              <el-col :span="8">
                <el-form-item label="毕业院校:" prop="graduation_school">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit"
                            v-model="searchValue.graduation_school" placeholder="毕业院校名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-form-item label="所学专业:" prop="major_studied">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                            v-model="searchValue.major_studied" placeholder="请输入专业名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="8" :offset="4">
                <el-button icon="el-icon-refresh" type="info" size="mini" @click="resetForm('searchValue')">取消</el-button>
                <el-button size="mini" icon="el-icon-search" type="primary" @click="searchEmpAdvance('searchValue')">搜索</el-button>
              </el-col>

            </el-row>
          </el-form>
        </div>
      </transition>
    </div>
    <el-table :data="empsData" stripe border v-loading="loading" element-loading-text="正在加载..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)" style="width: 100%">
      <el-table-column prop="id" align="left" label="员工编号" show-overflow-tooltip></el-table-column>
      <el-table-column prop="password" align="left" label="密码" show-overflow-tooltip></el-table-column>
      <el-table-column prop="employee_type" align="left" label="员工类型" fixed show-overflow-tooltip></el-table-column>
      <el-table-column prop="name" align="left" label="姓名" fixed show-overflow-tooltip></el-table-column>
      <el-table-column prop="sex" align="left" label="性别" show-overflow-tooltip></el-table-column>
      <el-table-column prop="birthday" align="left" label="出生日期" show-overflow-tooltip></el-table-column>
      <el-table-column prop="id_card" align="left" label="身份证号" show-overflow-tooltip></el-table-column>
      <el-table-column prop="political" align="left" label="政治面貌" show-overflow-tooltip></el-table-column>
      <el-table-column prop="nation" align="left" label="民族" show-overflow-tooltip></el-table-column>
      <el-table-column prop="native_place" align="left" label="籍贯" show-overflow-tooltip></el-table-column>
      <el-table-column prop="phone" align="left" label="电话" show-overflow-tooltip></el-table-column>
      <el-table-column prop="email" align="left" label="电子邮箱" show-overflow-tooltip></el-table-column>
      <el-table-column prop="height" align="left" label="身高" show-overflow-tooltip></el-table-column>
      <el-table-column prop="blood_type" align="left" label="血型" show-overflow-tooltip></el-table-column>
      <el-table-column prop="marital_status" align="left" label="婚姻状况" show-overflow-tooltip></el-table-column>
      <el-table-column prop="birthplace" align="left" label="出生地" show-overflow-tooltip></el-table-column>
      <el-table-column prop="registered_residence" align="left" label="户口所在地" show-overflow-tooltip></el-table-column>
      <el-table-column prop="department_number" align="left" label="部门编号" show-overflow-tooltip></el-table-column>
      <el-table-column prop="department_name" align="left" label="部门名称" fixed show-overflow-tooltip></el-table-column>
      <el-table-column prop="post_number" align="left" label="岗位编号" show-overflow-tooltip></el-table-column>
      <el-table-column prop="post_name" align="left" label="岗位名称" fixed show-overflow-tooltip></el-table-column>
      <el-table-column prop="entry_date" align="left" label="入职日期" show-overflow-tooltip></el-table-column>
      <el-table-column prop="employment_form" align="left" label="用工形式" show-overflow-tooltip></el-table-column>
      <el-table-column prop="personnel_source" align="left" label="人员来源" show-overflow-tooltip></el-table-column>
      <el-table-column prop="highest_education" align="left" label="最高学历" show-overflow-tooltip></el-table-column>
      <el-table-column prop="graduation_school" align="left" label="毕业院校" show-overflow-tooltip></el-table-column>
      <el-table-column prop="major_studied" align="left" label="所学专业" show-overflow-tooltip></el-table-column>
      <el-table-column prop="graduation_date" align="left" label="毕业日期" show-overflow-tooltip></el-table-column>
      <el-table-column prop="is_quit" align="left" label="是否离职" show-overflow-tooltip></el-table-column>

      <el-table-column fixed="right" width="200" label="操作">
        <template slot-scope="scope">
          <el-button @click="expelEmp(scope.row)" style="padding: 3px" size="mini" type="danger" :disabled="scope.row.is_quit==='是'">开除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div style="display: flex;justify-content: flex-end">
      <el-pagination
          background
          @current-change="currentChange"
          @size-change="sizeChange"
          :page-size="pageSize"
          :page-sizes="[1,5,10,20,100]"
          :current-page.sync="currentPage"
          layout="sizes, prev, pager, next, jumper, ->, total, slot"
          :total="emps.length">
        <!--                        :total="total">-->
      </el-pagination>
    </div>
  </div>
</template>

<script>
import {Message} from "element-ui";

export default {
  name: "ExpelEmployee",
  data() {
    return {
      emp: {
        id : "",
        password : "",
        employee_type : "",
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
        department_number : null,
        department_name : "",
        post_number : null,
        post_name : "",
        entry_date : "",
        employment_form : "",
        personnel_source : "",
        highest_education : "",
        graduation_school : "",
        major_studied : "",
        graduation_date : "",
        is_quit : "",


        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,
        // workState: "在职",
        // contractTerm: 2,
        // conversionTime: "2018-03-31",
        // notworkDate: null,
        // beginContract: "2017-12-31",
        // endContract: "2019-12-31",
        // workAge: null,
      },
      rules: {
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

        nationId: [{required: true, message: '请输入您组', trigger: 'blur'}],
        jobLevelId: [{required: true, message: '请输入职称', trigger: 'blur'}],
        posId: [{required: true, message: '请输入职位', trigger: 'blur'}],
        // workState: [{required: true, message: '请输入工作状态', trigger: 'blur'}],
        // contractTerm: [{required: true, message: '请输入合同期限', trigger: 'blur'}],
        // conversionTime: [{required: true, message: '请输入转正日期', trigger: 'blur'}],
        // notworkDate: [{required: true, message: '请输入离职日期', trigger: 'blur'}],
        // beginContract: [{required: true, message: '请输入合同起始日期', trigger: 'blur'}],
        // endContract: [{required: true, message: '请输入合同结束日期', trigger: 'blur'}],
        // workAge: [{required: true, message: '请输入工龄', trigger: 'blur'}],
      },
      searchValue: {
        political: null,
        nation: null,
        department_name : null,
        post_name : null,
        employment_form: null,
        graduation_school: null,
        major_studied: null
      },
      politicals: ['群众', '共青团员', '中共预备党员', '中共党员', '无党派人士', '其他'],
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
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      showAdvanceSearchView: false,
      loading : false,
      emps : [],
      currentPage: 1,
      pageSize: 10,
    }
  },
  mounted() {
    this.initEmps();
  },
  computed : {
    empsData() {
      console.log("emps.length = ", this.emps.length);
      if (this.emps.length > 0) {
        return this.emps.slice( (this.currentPage -1) * this.pageSize, this.currentPage * this.pageSize) || [];
      }
      console.log("emps = ", this.emps);
      return this.emps;
    }
  },
  methods : {
    sizeChange(currentSize) {
      this.pageSize = currentSize;
      this.initEmps();
    },
    currentChange(currentPage) {
      this.currentPage = currentPage;
      this.initEmps('advanced');
    },
    initEmps() {
      this.loading = true;
      this.$axios.post('/EmployeeBasic').then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
          console.log("resp = ", this.emps);
          this.total = resp.total;
        }
      })
    },
    expelEmp(data) {
      if (data.is_quit !== "是") {
        var tempId = localStorage.getItem("id");
        console.log("当前登录ID：" + tempId);
        if (data.id === localStorage.getItem("id")) {
          Message.error({message : "不能开除自己"});
        } else {
          this.$confirm('此操作将开除【' + data.name + '】, 是否继续?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            this.$axios.post("/ExpelEmployee", data).then(resp => {
              if (resp) {
                this.initEmps();
                Message.success({message : "开除成功"});
              } else {
                Message.error({message : resp.data.msg});
              }
            })
          }).catch(() => {
            Message.info({message : "已取消开除"});
          });
        }
      } else {
        Message.error({message : "此员工已离职"});
      }
      this.initEmps();
    },
    async searchEmp(data) {
      console.log("name = ", this.emp.name);
      this.$refs[data].validate((valid) => {
        if (valid) {
          this.$axios.post('/EmployeeBasic/Search', this.emp).then((resp) => {
            if (resp.data.msg == "查询成功") {
              this.emps = resp.data.data;
            } else {
              Message.error({message : resp.data.msg});
            }
          });
        }
      })
    },
    resetForm(data) {
      this.$refs[data].resetFields();
      this.initEmps();
    },
    changeSelect() {
      // 清空部门内容
      this.searchValue.post_name = ''

      // 遍历部门的下拉选项数组
      for (const k in this.department_names) {
        // 岗位名称 是否等于 部门名称的下拉选择数组中的某一项
        if (this.searchValue.department_name === this.department_names[k]) {
          this.post_type_options = this.post_names[this.searchValue.department_name]
        }
      }
      console.log("post_type_options = "+this.post_type_options);
    },
    async searchEmpAdvance(data) {
      this.$refs[data].validate((valid) => {
        if (valid) {
          this.$axios.post('/EmployeeBasic/SearchAdvance', this.searchValue).then((resp) => {
            if (resp.data.msg == "查询成功") {
              this.emps = resp.data.data;
            } else {
              Message.error({message : resp.data.msg});
            }
          })
        }
      })
    },
  }
}
</script>

<style scoped>

</style>