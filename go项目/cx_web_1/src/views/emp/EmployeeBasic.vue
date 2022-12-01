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

        <el-table-column
            prop="name"
            fixed
            sortable
            align="left"
            label="姓名"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="employee_type"
            fixed
            sortable
            align="left"
            label="角色"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="id"
            label="编号"
            align="left"
            width="85"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="password"
            label="密码"
            align="left"
            width="85"
            v-if="is_director"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="sex"
            label="性别"
            align="left"
            width="75"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="birthday"
            width="100"
            align="left"
            label="出生日期"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="id_card"
            width="150"
            align="left"
            label="身份证号码"
            v-if="is_director"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="political"
            width="100"
            label="政治面貌"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="nation"
            label="民族"
            align="left"
            width="85"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="native_place"
            label="籍贯"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="phone"
            label="电话"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="email"
            label="电子邮箱"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="height"
            label="身高"
            align="left"
            width="75"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="blood_type"
            label="血型"
            align="left"
            width="75"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="marital_status"
            label="婚姻状况"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="birthplace"
            label="出生地"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="registered_residence"
            label="户口所在地"
            align="left"
            width="125"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_number"
            label="部门编号"
            align="left"
            width="125"
            sortable
            :filters="filterDepartmentText"
            :filter-method="filterHandler"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_name"
            fixed
            width="100"
            align="left"
            label="部门名称"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_number"
            label="岗位编号"
            align="left"
            width="125"
            sortable
            :filters="filterPostText"
            :filter-method="filterHandler"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_name"
            label="岗位名称"
            fixed
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="entry_date"
            label="入职日期"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="employment_form"
            label="用工形式"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="personnel_source"
            label="人员来源"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="highest_education"
            label="最高学历"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_school"
            label="毕业院校"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="major_studied"
            label="所学专业"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_date"
            label="毕业日期"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="is_quit"
            label="是否离职"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            fixed="right"
            width="200"
            label="操作">
          <template slot-scope="scope">
            <el-button @click="showEditEmpView(scope.row)" v-if="is_director" style="padding: 3px" size="mini" type="warning">编辑</el-button>

          </template>
        </el-table-column>
      </el-table>

      <el-dialog :visible.sync="multiDeleteVisible" title="提示" width="30%">
        <span>确定要删除吗?</span>
        <span slot="footer">
          <el-button type="danger" @click="multiDelete">确 定</el-button>
          <el-button @click="dialogVisible = false; multiDeleteVisible = false">取 消</el-button>
        </span>
      </el-dialog>
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
    <!--      编辑框-->
    <el-dialog
        :title="title"
        @close="resetForm('emp')"
        :visible.sync="dialogEditVisible"
        width="80%">
      <div>
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
            <el-col :span="6">
              <el-form-item label="出生地:" prop="birthplace">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-s-opportunity"
                          v-model="emp.birthplace" placeholder="请输入出生地" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="户口所在地:" prop="registered_residence">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-house"
                          v-model="emp.registered_residence" placeholder="请输入户口所在地" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="部门编号:" prop="department_number">
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
                          v-model="emp.department_number" disabled></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="部门名称:" prop="department_name">
                <el-select @change="changeEditSelect" v-model="emp.department_name" placeholder="请输入部门名称" size="mini" prefix-icon="el-icon-edit" disabled>
                  <el-option
                      v-for="(item, index) in department_names"
                      :key="index"
                      :label="item"
                      :value="item">
                  </el-option>
                </el-select>
                <!--                            <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                <!--                                      v-model="emp.department_name" disabled></el-input>-->
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="岗位编号:" prop="post_number">
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
                          v-model="emp.post_number" disabled></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="岗位名称:" prop="post_name">
                <el-select v-model="emp.post_name" placeholder="请输入岗位" size="mini" prefix-icon="el-icon-edit" disabled>
                  <el-option
                      v-for="(item, index) in post_type_options"
                      :key="index"
                      :label="item"
                      :value="item">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
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
          <el-row>
            <el-col :span="6">
              <el-form-item label="是否离职" prop="is_quit">
                <el-radio-group v-model="emp.is_quit">
                  <el-radio label="是">是</el-radio>
                  <el-radio label="否">否</el-radio>
                </el-radio-group>
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
  name: "EmpBasic",
  data() {
    return {

      tipContent : "提示",
      currentPage: 1,
      pageSize: 10,
      searchValue: {
        political: null,
        nation: null,
        department_name : null,
        post_name : null,
        employment_form: null,
        graduation_school: null,
        major_studied: null
      },
      isAdd : false,
      title: '',

      importDataBtnIcon: 'el-icon-upload2',
      importDataDisabled: false,
      showAdvanceSearchView: false,
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
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      blood_type:'',
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      employee_type:3,
      employee_types:['总裁', '主管', '员工'],
      nation: "",
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
      joblevels: [],
      political: "群众",
      politicals: ['群众', '共青团员', '中共预备党员', '中共党员', '无党派人士', '其他'],
      positions: [],
      highest_educations: ['小学', '初中', '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
      employee_forms:['实习生', '正式职工'],
      personnel_source:'',
      personnel_sources:['校招', '社招'],
      inputDepName: '所属部门',
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

      },
      defaultProps: {
        children: 'children',
        label: 'name'
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

        // workAge: [{required: true, message: '请输入工龄', trigger: 'blur'}],
      },
      multipleSelectionFlag : false,
      multiDeleteVisible : false,
      multipleSelection : '',
      filterDepartmentText : [],
      filterPostText : [],
      value1: "",
      value2: "",
      dateRules : {
        value : [{required : true, message : '请输入开始和结束日期', trigger : 'blur'}],
        value1 : [{required : true, message : '请输入开始日期', trigger : 'blur'}],
        value2 : [{required : true, message : '请输入结束日期', trigger : 'blur'}],
      },
      pickerOptions0: {
        disabledDate: (time) => {
          if (this.value2 !== "") {
            return time.getTime() > Date.now() || time.getTime() > this.value2;
          } else {
            return time.getTime() > Date.now();
          }
        }
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
    this.initData();
    this.initDepartment();
    this.initPost();
    this.initSearchText();
    this.tipContent = '部门和岗位的编号与名称请一一对应！！！ 例如:部门编号-部门名称(岗位1编号-岗位1名称, 岗位2编号-岗位2名称) 1000-开发部(10-C++开发, 11-Java开发, 12-C#开发, 13-Python开发, 14-Go开发), 1001-运维部(20-云运维, 21-服务器运维), 1002-测试部(30-Bug测试, 31-系统测试), 1003-设计部(40-UI设计, 41-动画设计), 1004-策划部(50-策划, 51-系统策划)';

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
    searvhViewHandleNodeClick(data) {
      this.inputDepName = data.name;
      this.searchValue.department_number = data.id;
      this.popVisible2 = !this.popVisible2
    },
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
        is_quit : "",


        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,

      }
      this.inputDepName = '';

    },
    showAllData(data) {
      this.title = '查看员工信息';
      this.emp = data;
      this.dialogShowVisible = true;
    },
    showEditEmpView(data) {
      // this.initPositions();
      this.title = '编辑员工信息';
      this.emp = data;
      this.inputDepName = data.department_name;
      this.dialogEditVisible = true;
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

    doEditEmp() {
      //编辑
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          this.emp.height = parseInt(this.emp.height);
          this.$axios.post('/EmployeeBasic/Update', this.emp).then(resp => {
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
    doAddEmp() {
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          console.log(this.emp);
          this.emp.height = parseInt(this.emp.height);
          this.emp.department_number = parseInt(this.emp.department_number);
          this.emp.post_number = parseInt(this.emp.post_number);
          this.$axios.post('/EmployeeBasic/Add', this.emp).then(resp => {
            if (resp) {
              this.dialogAddVisible = false;
              this.initEmps();
              if (resp.data.msg === "添加成功") {
                Message.success({message : '添加成功'});
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

    showDepView() {
      this.popVisible = !this.popVisible
    },

    initData() {

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
      //this.getMaxWordID();
      this.dialogAddVisible = true;
    },
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
              if (resp.data.msg == "查询成功") {
                this.emps = resp.data.data;
              } else {
                Message.error({message : resp.data.msg});
              }
            });
          }
        })
      } else {
        url = '/EmployeeBasic/SearchByEmployee';
        var temp = {
          'id' : localStorage.getItem("id"),
          'name' : this.emp.name,
        };
        this.$refs[data].validate((valid) => {
          if (valid) {
            this.$axios.post(url, temp).then((resp) => {
              if (resp.data.msg == "查询成功") {
                this.emps = resp.data.data;
              } else {
                Message.error({message : resp.data.msg});
              }
            });
          }
        })
      }


    },
    resetForm(data) {
      console.log("data = ", data, "emp = ", this.emp)
      this.$refs[data].resetFields();
      this.initEmps();
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
        this.$axios.post('/EmployeeBasicByEmployee', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            this.total = resp.total;
          }
        })
      }

    },
    handleSelectionChange(val) {
      // console.log(val);
      this.multipleSelection = val;
      this.multipleSelectionFlag = true;
      if (this.multipleSelection.length == 0) {
        // 如不进行判断则勾选完毕后批量删除按钮还是会在
        this.multipleSelectionFlag = false;
      }
    },
    popDelete() {
      this.multiDeleteVisible = true;
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
    },


    resetDate() {
      this.value1 = ""
      this.value2 = ""
      this.initEmps();
    }
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
