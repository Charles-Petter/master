<template>
  <div>
    <!--    个人信息-->
    <div>
      <span>
          您的信息
      </span>
      <el-table :data="empsData"
                stripe border
                v-loading="loading"
                element-loading-text="正在加载..."
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)"
                style="width: 100%">
        <el-table-column
            prop="name"
            fixed
            align="left"
            label="姓名"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="employee_type"
            fixed
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
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="password"
            label="密码"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="sex"
            label="性别"
            align="left"
            width="50"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="birthday"
            width="100"
            align="left"
            label="出生日期"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="id_card"
            width="150"
            align="left"
            label="身份证号码"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="political"
            label="政治面貌"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="nation"
            label="民族"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="native_place"
            label="籍贯"
            align="left"
            width="100"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="phone"
            label="电话"
            align="left"
            width="100"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="email"
            label="电子邮箱"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="height"
            label="身高"
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="blood_type"
            label="血型"
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="marital_status"
            label="婚姻状况"
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="birthplace"
            label="出生地"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="registered_residence"
            label="户口所在地"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_number"
            label="部门编号"
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_name"
            fixed
            width="75"
            align="left"
            label="部门名称"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_number"
            label="岗位编号"
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_name"
            label="岗位名称"
            fixed
            align="left"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="entry_date"
            label="入职日期"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="employment_form"
            label="用工形式"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="personnel_source"
            label="人员来源"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="highest_education"
            label="最高学历"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_school"
            label="毕业院校"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="major_studied"
            label="所学专业"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_date"
            label="毕业日期"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="is_quit"
            label="是否离职"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            fixed="right"
            width="200"
            label="操作">
          <template slot-scope="scope">
            <el-button @click="showMyData(scope.row)" style="padding: 3px" size="mini" type="primary">查看完整资料</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!--    申请离职-->
    <div>
      <span>
        申请离职
      </span>
      <el-form :model="resigned" :rules="resignedRules" ref="resigned">
        <el-form-item label="离职原因:" prop="reasons_for_resignation">
          <el-input type="textarea" v-model="resigned.reasons_for_resignation" style="width: 300px" clearable maxlength="50" show-word-limit></el-input>
        </el-form-item>
        <el-button-group>
          <el-button type="warning" size="mini" @click="applyResigned">提交</el-button>
          <el-button type="info" size="mini" @click="resetForm('resigned')">重置</el-button>
        </el-button-group>
      </el-form>
<!--      <el-form :model="application" :rules="applyRules" ref="application">-->
<!--        <el-form-item label="新部门名称:" prop="new_department_name">-->
<!--          <el-select @change="changeEditSelect" v-model="application.new_department_name" placeholder="请输入部门名称" size="mini" prefix-icon="el-icon-edit">-->
<!--            <el-option-->
<!--                v-for="(item, index) in department_names"-->
<!--                :key="index"-->
<!--                :label="item"-->
<!--                :value="item">-->
<!--            </el-option>-->
<!--          </el-select>-->
<!--          &lt;!&ndash;                            <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"&ndash;&gt;-->
<!--          &lt;!&ndash;                                      v-model="emp.department_name" disabled></el-input>&ndash;&gt;-->
<!--        </el-form-item>-->
<!--        <el-form-item label="新岗位名称:" prop="new_post_name">-->
<!--          <el-select v-model="application.new_post_name" placeholder="请输入岗位" size="mini" prefix-icon="el-icon-edit">-->
<!--            <el-option-->
<!--                v-for="(item, index) in post_type_options"-->
<!--                :key="index"-->
<!--                :label="item"-->
<!--                :value="item">-->
<!--            </el-option>-->
<!--          </el-select>-->
<!--          &lt;!&ndash;                            <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"&ndash;&gt;-->
<!--          &lt;!&ndash;                                      v-model="emp.post_name" disabled></el-input>&ndash;&gt;-->
<!--        </el-form-item>-->
<!--        <el-form-item label="申请原因:" prop="reason_for_application">-->
<!--          <el-input type="textarea" v-model="application.reason_for_application" style="width: 300px"></el-input>-->
<!--        </el-form-item>-->
<!--        <el-button-group>-->
<!--          <el-button type="success" size="mini" @click="applyNewPost">申请</el-button>-->
<!--          <el-button type="info" size="mini" @click="resetForm('application')">重置</el-button>-->
<!--        </el-button-group>-->
<!--      </el-form>-->


    </div>
    <div>
      <!--      查看框-->
      <el-dialog
          :title="title"
          :visible.sync="dialogShowVisible"
          width="80%">
        <div>
          <el-form :model="emp" :rules="rules" ref="empForm">
            <el-row>
              <el-col :span="6">
                <el-form-item label="编号:" prop="id">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit"
                            v-model="emp.id" placeholder="编号" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="密码:" prop="password">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit" v-model="emp.password"
                            placeholder="请输入密码" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="员工类型:" prop="employee_type">
                  <el-select v-model="emp.employee_type" placeholder="员工类型" size="mini" style="width: 150px;" disabled>
                    <el-option v-for="item in employee_types" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="姓名:" prop="name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit" v-model="emp.name"
                            placeholder="请输入员工姓名" disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="性别:" prop="sex">
                  <el-radio-group v-model="emp.sex" disabled>
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
                      placeholder="出生日期" disabled>
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="身份证号码:" prop="id_card">
                  <el-input size="mini" style="width: 180px" prefix-icon="el-icon-edit"
                            v-model="emp.id_card" placeholder="请输入身份证号码" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="政治面貌:" prop="political">
                  <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;" disabled>
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
                  <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;" disabled>
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
                  <el-input size="mini" style="width: 120px" prefix-icon="el-icon-edit"
                            v-model="emp.native_place" placeholder="请输入籍贯" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="电话号码:" prop="phone">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-phone"
                            v-model="emp.phone" placeholder="电话号码" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="电子邮箱:" prop="email">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                            v-model="emp.email" placeholder="请输入电子邮箱" disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="身高:" prop="height">
                  <el-input size="mini" style="width: 100px" type="number" prefix-icon="el-icon-edit"
                            v-model="emp.height" placeholder="请输入身高" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="血型:" prop="blood_type">
                  <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;" disabled>
                    <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="婚姻状况:" prop="marital_status">
                  <el-radio-group v-model="emp.marital_status" disabled>
                    <el-radio label="已婚">已婚</el-radio>
                    <el-radio label="未婚">未婚</el-radio>
                    <el-radio label="离异">离异</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="出生地:" prop="birthplace">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                            v-model="emp.birthplace" placeholder="请输入出生地" disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="户口所在地:" prop="registered_residence">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                            v-model="emp.registered_residence" placeholder="请输入户口所在地" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门编号:" prop="department_number">
                  <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"
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
                  <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"
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
                  <!--                            <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                  <!--                                      v-model="emp.post_name" disabled></el-input>-->
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
                      placeholder="入职日期" disabled>
                  </el-date-picker>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="用工形式:" prop="employment_form">
                  <el-select v-model="emp.employment_form" placeholder="用工形式" size="mini" style="width: 150px;" disabled>
                    <el-option v-for="item in employee_forms" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="人员来源:" prop="personnel_source">
                  <el-select v-model="emp.personnel_source" placeholder="人员来源" size="mini" style="width: 150px;" disabled>
                    <el-option v-for="item in personnel_sources" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="最高学历:" prop="highest_education">
                  <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"
                             style="width: 150px;" disabled>
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
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit"
                            v-model="emp.graduation_school" placeholder="毕业院校名称" disabled></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="所学专业:" prop="major_studied">
                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                            v-model="emp.major_studied" placeholder="请输入专业名称" disabled></el-input>
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
                      placeholder="毕业日期" disabled>
                  </el-date-picker>

                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="是否离职" prop="is_quit">
                  <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"
                            v-model="emp.is_quit" placeholder="是否离职" disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
              <el-button @click="dialogShowVisible = false">关 闭</el-button>
          <!--              <el-button type="primary" @click="doEditEmp">确 定</el-button>-->
            </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {Message} from "element-ui";

export default {
  name: "ApplyQuit",
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
      },
      resigned : {
        id : '',
        name : '',
        sex : '',
        resignation_date : null,
        reasons_for_resignation : null,
        department_number : null,
        department_name : null,
        post_number : null,
        post_name : null,
      },
      resignedRules : {
        reasons_for_resignation : [{required : true, message : "请输入离职原因", trigger : 'blur'}],
      },
      employee_types:['总裁', '主管', '员工'],
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
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      employee_forms:['实习生', '正式职工'],
      highest_educations: ['小学', '初中', '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
      personnel_sources:['校招', '社招'],
      emps : [],
      loading : false,
      currentPage: 1,
      pageSize: 10,
      title : '',
      dialogShowVisible : false,
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
    initEmps(type) {
      this.loading = true;
      this.emp.id = localStorage.getItem("id");
      this.$axios.post('/ApplyResignedInformation', this.emp).then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
          this.emp = this.emps[0];
          this.total = resp.total;
        }
      })
    },
    showMyData(data) {
      this.title = '查看员工信息';
      this.emp = data;
      this.dialogShowVisible = true;
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
    resetForm(data) {
      this.$refs[data].resetFields();
      this.initEmps();
    },
    applyResigned() {
      this.initEmps();
      this.$refs['resigned'].validate(valid => {
        if (valid) {
          if (this.emp.employee_type === "总裁") {
            Message.error({message : "您是总裁，不能离职"});
          } else {
            this.$confirm(this.emp.name + "，确定要离职吗？", {
              confirmButtonText : '确定',
              cancelButtonText : '取消',
              type : 'warning',
            }).then(() => {
              this.resigned.id = this.emp.id;
              this.resigned.name = this.emp.name;
              this.resigned.sex = this.emp.sex;
              this.resigned.department_number = this.emp.department_number;
              this.resigned.department_name = this.emp.department_name;
              this.resigned.post_number = this.emp.post_number;
              this.resigned.post_name = this.emp.post_name;

              console.log("离职信息：", this.resigned);
              this.$axios.post('/ApplyResigned', this.resigned).then((resp) => {
                if (resp.data.msg === "离职成功") {
                  Message.success({message : resp.data.msg});
                  this.$router.push('/');
                } else {
                  Message.error({message : resp.data.msg});
                }
              })
            }).catch(() => {
              Message.info({message : "已取消离职"});
            })

          }
        }
      });

    }
  }
}
</script>

<style scoped>

</style>