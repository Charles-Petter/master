<template>
  <div>
    <div>
      <el-form :model="emp" :rules="rules" ref="emp" @submit.native.prevent>
        <el-form-item prop="name">
          <el-input placeholder="请输入员工姓名进行搜索，可以直接回车搜索..." prefix-icon="el-icon-search"
                    clearable
                    @clear="initEmps"
                    style="width: 350px;margin-right: 10px" v-model="emp.name"
                    @keydown.enter.native="searchEmp('emp')"></el-input>
          <el-button-group>
            <el-button icon="el-icon-search" type="primary" @click="searchEmp('emp')">搜索</el-button>
            <el-button icon="el-icon-refresh" type="info" @click="resetForm('emp')">重置</el-button>
          </el-button-group>
        </el-form-item>
      </el-form>
    </div>
    <div style="margin-right:50px; float:right;display:flex;justify-content:space-between">
      <div class="block"  style="float: right; margin-right: 50px">
        <span class="demonstration"></span>
        <el-date-picker
            v-model="value1"
            type="date"
            size="mini"
            value-format="yyyy-MM-dd"
            placeholder="开始日期"
            :picker-options="pickerOptions0">
        </el-date-picker>
        <el-date-picker
            v-model="value2"
            type="date"
            size="mini"
            value-format="yyyy-MM-dd"
            placeholder="结束日期"
            :picker-options="pickerOptions1">
        </el-date-picker>
        <el-button-group>
          <el-button type="primary" icon="el-icon-search" @click="selectDate"></el-button>
          <el-button type="info" icon="el-icon-refresh" @click="resetDate"></el-button>
        </el-button-group>
      </div>
      <el-button-group>
        <download-excel :data="emps" :fields="json_fields" name="部门调动员工表">
          <el-button type="success" icon="el-icon-download" :disabled="emps.length == 0">
            导出数据
          </el-button>
        </download-excel>
      </el-button-group>
    </div>
    <div style="margin-top: 10px">
      <el-table :data="empsData" stripe border v-loading="loading" element-loading-text="正在加载..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)" style="width: 100%">
        <el-table-column prop="transfer_id" align="left" label="部门调动编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="id" align="left" label="员工编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="name" align="left" label="姓名" show-overflow-tooltip></el-table-column>
        <el-table-column prop="sex" align="left" label="性别" show-overflow-tooltip></el-table-column>
        <el-table-column prop="original_department_number" align="left" label="原部门编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="original_department_name" align="left" label="原部门名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="new_department_number" align="left" label="新部门编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="new_department_name" align="left" label="新部门名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="transfer_date" align="left" label="调动日期" show-overflow-tooltip></el-table-column>
        <el-table-column prop="reasons_for_transfer" align="left" label="调动原因" show-overflow-tooltip></el-table-column>
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
  </div>
</template>

<script>
import {Message} from "element-ui";

export default {
  name: "DepartmentTransfer",
  data() {
    return {
      loading : false,
      emps : [],
      currentPage : 1,
      pageSize : 10,
      json_fields: {
        //导出Excel表格的表头设置
        '部门调动编号': 'transfer_id',
        '员工编号': 'id',
        '姓名': 'name',
        '性别': 'sex',
        '原部门编号': 'original_department_number',
        '原部门名称': 'original_department_name',
        '新部门编号': 'new_department_number',
        '新部门名称': 'new_department_name',
        '调动日期': 'transfer_date',
        '调动原因' : 'reasons_for_transfer',
      },
      emp: {
        transfer_id : "",
        id : "",
        name : "",
        sex : "",
        original_department_number : null,
        original_department_name : "",
        new_department_number : null,
        new_department_name : "",
        transfer_date : "",
        reasons_for_transfer : "",
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
    }
  },
  mounted() {
    this.initEmps();
    // console.log("value1 = ", this.value1);
  },
  computed : {
    empsData() {
      if (this.emps.length > 0) {
        return this.emps.slice( (this.currentPage -1) * this.pageSize, this.currentPage * this.pageSize) || [];
      }
      console.log("emps = ", this.emps);
      return this.emps;
    }
  },
  methods : {
    initEmps() {
      this.loading = true;
      this.$axios.post('/DepartmentTransfer').then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
          console.log("resp = ", this.emps);
          // this.total = resp.total;
        }
      })
    },
    sizeChange(currentSize) {
      this.pageSize = currentSize;
      this.initEmps();
    },
    currentChange(currentPage) {
      this.currentPage = currentPage;
      this.initEmps();
    },
    async searchEmp(data) {
      console.log("name = ", this.emp.name);
      this.$refs[data].validate((valid) => {
        if (valid) {
          this.$axios.post('/DepartmentTransfer/Search', this.emp).then((resp) => {
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
    selectDate() {
      if ((this.value1 === "" || this.value1 === null) || (this.value2 === "" || this.value2 === null)) {
        Message.error({message : '请输入开始和结束日期'});
      } else {
        console.log("value1 = ", this.value1);
        var temp = {
          'start' : this.value1,
          'end' : this.value2,
        }
        this.$axios.post('/DepartmentTransfer/SearchByDate', temp).then(resp => {
          if (resp.data.data) {
            this.emps = resp.data.data;
            Message.success({message : '查询成功'});
          } else {
            Message.warning({message : '未找到符合条件的信息'});
          }
        })
        // this.initEmps();
      }

    },
    resetDate() {
      this.value1 = ""
      this.value2 = ""
      this.initEmps();
    }
  }
}
</script>

<style scoped>

</style>