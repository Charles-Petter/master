<template>
  <div>
    <div style="margin-top: 10px">
      <el-table
          :data="empsData"
          stripe
          border
          :default-sort = "{prop: 'date', order: 'descending'}"
          v-loading="loading"
          @selection-change="handleSelectionChange"
          element-loading-text="正在加载..."
          element-loading-spinner="el-icon-loading"
          element-loading-background="rgba(0, 0, 0, 0.8)"
          style="width: 100%">
        <el-table-column prop="name" fixed sortable align="left" label="姓名" width="75" show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="employee_type" fixed  sortable align="left" label="角色" width="75"show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="id" label="编号" align="left" width="85" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="password" label="密码" align="left" width="85" v-if="is_director" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="sex" label="性别" align="left" width="75" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="birthday" width="100" align="left" label="出生日期" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="id_card" width="150" align="left" label="身份证号码"v-if="is_director" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="political" width="100" label="政治面貌" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="nation" label="民族" align="left" width="85" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="native_place" label="籍贯" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="phone" label="电话" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="email" label="电子邮箱" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="height" label="身高" align="left" width="75" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="blood_type" label="血型" align="left" width="75" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="marital_status" label="婚姻状况" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="department_number" label="部门编号" align="left" width="125" sortable :filters="filterDepartmentText" :filter-method="filterHandler" show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="department_name" fixed width="100" align="left" label="部门名称" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="entry_date" label="入职日期" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="employment_form" label="用工形式" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="personnel_source" label="人员来源" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="highest_education" label="最高学历" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="graduation_school" label="毕业院校" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="major_studied" label="所学专业" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="graduation_date" label="毕业日期" align="left" width="100" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column fixed="right" width="200" label="修改" v-if="!is_director">
          <template slot-scope="scope">
<!--            员工可以修改自己的信息-->
            <el-button @click="showEditEmpView(scope.row)" v-if="!is_director" style="padding: 3px" size="mini" type="warning">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!--      员工登录会出现编辑框-->
    <el-dialog
        :title="title"
        @close="resetForm('emp')"
        :visible.sync="dialogEditVisible"
        width="80%">
      <div>
<!--        绑定额emp-->
        <el-form :model="emp" :rules="rules" ref="empForm">
          <el-row>
            <el-col :span="6">
              <el-form-item label="编号:" prop="id">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user-solid"
                          v-model="emp.id" placeholder="编号" disabled></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="密码:" prop="password">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-lock" v-model="emp.password"
                          placeholder="请输入密码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="员工类型:" prop="employee_type">
                <el-select v-model="emp.employee_type" placeholder="员工类型" size="mini" style="width: 150px;">
                  <el-option v-for="item in employee_types" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="姓名:" prop="name">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.name"
                          placeholder="请输入员工姓名"  clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
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
              <el-form-item label="身份证号码:" prop="id_card">
                <el-input size="mini" style="width: 180px" prefix-icon="el-icon-s-custom"
                          v-model="emp.id_card" placeholder="请输入身份证号码" clearable></el-input>
              </el-form-item>
            </el-col>
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
          </el-row>
          <el-row>
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
            <el-col :span="6">
              <el-form-item label="电子邮箱:" prop="email">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                          v-model="emp.email" placeholder="请输入电子邮箱" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
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
          </el-row>
          <el-row>

            <el-col :span="6">
              <el-form-item label="入职日期:" prop="entry_date">
                <el-date-picker
                    v-model="emp.entry_date"
                    size="mini"
                    type="date"
                    value-format="yyyy-MM-dd"
                    style="width: 150px;"
                    placeholder="入职日期">
                </el-date-picker>
              </el-form-item>
            </el-col>
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
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="最高学历:" prop="highest_education">
                <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"
                           style="width: 150px;">
                  <el-option v-for="item in highest_educations" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="毕业院校:" prop="graduation_school">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-school" v-model="emp.graduation_school" placeholder="毕业院校名称" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="所学专业:" prop="major_studied">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-notebook-1" v-model="emp.major_studied" placeholder="请输入专业名称" clearable></el-input>
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
        <template>
           <el-popover placement="bottom" title="注意事项" width="300" trigger="click" :content="tipContent">
          <el-button type="danger" size="mini" slot="reference">?</el-button>
        </el-popover>
        </template>
              <el-button @click="dialogEditVisible = false">取 消</el-button>
              <el-button type="primary" @click="doEditEmp">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import { Message } from "element-ui";

export default {
  name: "EmployeeBasic_cx",
  data() {
    return {
      tipContent : "提示",
      currentPage: 1,
      pageSize: 10,
      searchValue: {
      },
      isAdd : false,
      title: '',
      allDeps: [],
      emps: [],
      loading: false,
      popVisible: false,
      popVisible2: false,
      dialogEditVisible: false,
      dialogAddVisible : false,
      dialogShowVisible : false,
      is_director : true,
      total: 0,
      page: 1,
      keyword: '',
      size: 10,
      birthday: '',
      post_number: null,
      post_name: '',
      department_names : ['区块链学院', '智能科技学院', '新媒体学院', '设计部', '退役军人学院'],
      blood_type:'',
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      employee_type:3,
      employee_types:[ '主管', '员工'],
      nation: "",
      nations: [
        '汉族', '朝鲜族',"回族",'维吾尔族',
        '乌孜别克族', '鄂温克族', '保安族', '京族', '独龙族', '赫哲族', '珞巴族',
        '其他',
      ],
      joblevels: [],
      political: "群众",
      politicals: ['群众', '共青团员', '中共预备党员', '中共党员', '无党派人士', '其他'],
      positions: [],
      highest_educations: [ '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
      employee_forms:['实习生', '正式职工'],
      personnel_source:'',
      personnel_sources:['校招', '社招'],
      inputDepName: '所属部门',
      //员工数组
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
        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,
      },
      //请求子组件
      defaultProps: {
        children: 'children',
        label: 'name'
      },
      //判断输入内容是否规范
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
        phone : [{required : true, message : '请输入电话号码', trigger : 'blur'},
          {
            validator: function(rule, value, callback) {
              if (/^1[34578]\d{9}$/.test(value) == false) {
                callback(new Error("手机号格式不正确"));
              } else {
                callback();
              }
            },
            trigger: 'blur'
          }
        ],
        email : [{required : true, message : '请输入邮箱地址', trigger : 'blur'}, {
          type : 'email',
          message : '邮箱格式不正确',
          trigger : 'blur'
        }],
      },
      pickerOptions1: {
        disabledDate: (time) => {
          return time.getTime() < this.value1 || time.getTime() > Date.now();
        }
      },
      showDateSearchView : false,
    }
  },
  mounted() {
    this.initEmps();
    this.initDepartment();
    localStorage.setItem("type", this.emp.employee_type);
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
  methods: {
    emptyEmp() {
      this.emp = {
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
        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,
      }
      this.inputDepName = '';
    },
    showEditEmpView(data) {
      // this.initPositions();
      this.title = '编辑员工信息';
      this.emp = data;
      this.inputDepName = data.department_name;
      this.dialogEditVisible = true;
    },
    //编辑
    doEditEmp() {
      //组件完成后渲染数据
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          this.emp.height = parseInt(this.emp.height);
          //发送post请求
          this.$axios.post('/EmployeeBasic/Update_cx', this.emp).then(resp => {
            if (resp) {
              this.dialogEditVisible = false;
              this.initEmps();
              Message.success({message : '修改成功!'});
            } else {
              Message.error({message : resp.data.msg});
            }
          });
        }
      });
    },
    showDepView() {
      this.popVisible = !this.popVisible
    },
    sizeChange(currentSize) {
      this.pageSize = currentSize;
      this.initEmps();
    },
    currentChange(currentPage) {
      this.currentPage = currentPage;
      this.initEmps('advanced');
    },
    showAddEmpView() {
      this.emptyEmp();
      this.title = '添加员工信息';
      this.dialogAddVisible = true;
    },
    //异步 判断登录人的职位类型 请求不同的接口 渲染主管或员工
    async searchEmp(data) {
      console.log("name = ", this.emp.name);
      var url;
      if (localStorage.getItem("role") === "主管") {
        url = '/EmployeeBasic/SearchByDirector';
        var temp = {
          'id' : localStorage.getItem("id"),
          'name' : this.emp.name,
        };
        this.$refs[data].validate((valid) => {
          if (valid) {
            this.$axios.post(url, temp).then((resp) => {
              if (resp.data.msg === "查询成功") {
                this.emps = resp.data.data;
              } else {
                Message.error({message : resp.data.msg});
              }
            });
          }
        })
      } else {
        url = '/Employee/SearchEmp_cx';
        var temp = {
          'id' : localStorage.getItem("id"),
          'name' : this.emp.name,
        };
        this.$refs[data].validate((valid) => {
          if (valid) {
            this.$axios.post(url, temp).then((resp) => {
              if (resp.data.msg === "查询成功") {
                this.emps = resp.data.data;
              } else {
                Message.error({message : resp.data.msg});
              }
            });
          }
        })
      }
    },
    //刷新表单
    resetForm(data) {
      console.log("data = ", data, "emp = ", this.emp)
      this.$refs[data].resetFields();
      this.initEmps();
    },
    //初始化信息  不同身份登录显示信息不同
    initEmps(type) {
      this.loading = true;
      var tempRole = localStorage.getItem("role");
     if (tempRole === "主管") {
        let temp = {
          id : "",
        }
        temp.id = localStorage.getItem("id");
        this.$axios.post('/EmployeeBasicByDirector', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            this.total = resp.total;
          }
        })
      } else {
        this.is_director = false;
        let temp = {
          id : "",
        }
        temp.id = localStorage.getItem("id");
        this.$axios.post('/EmployeeBasic_cx', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            this.total = resp.total;
          }
        })
      }

    },

  }
}
</script>

<style>
/* 可以设置不同的进入和离开动画 */
/* 设置持续时间和动画函数 */
.slide-fade-enter-active {
  transition: all .8s ease;
}

.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter, .slide-fade-leave-to
  /* .slide-fade-leave-active for below version 2.1.8 */
{
  transform: translateX(10px);
  opacity: 0;
}
</style>
