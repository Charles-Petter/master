<template>
  <div>
    <div>
      <span v-if="is_director">您是{{ this.emp.department_name }}部门的主管</span>
      <el-descriptions class="margin-top" :title="title" :column="3" :size="size" border>
        <!--        <template slot="extra">-->
        <!--          <el-button type="primary" size="small">操作</el-button>-->
        <!--        </template>-->
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            员工编号
          </template>
          {{this.emp.id}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            密码
          </template>
          {{this.emp.password}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            员工类型
          </template>
          {{this.emp.employee_type}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            姓名
          </template>
          {{this.emp.name}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            出生日期
          </template>
          {{this.emp.birthday}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            身份证号
          </template>
          {{this.emp.id_card}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            政治面貌
          </template>
          {{this.emp.political}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            民族
          </template>
          {{this.emp.nation}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            籍贯
          </template>
          {{this.emp.native_place}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            电话
          </template>
          {{this.emp.phone}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            电子邮箱
          </template>
          {{this.emp.email}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            身高
          </template>
          {{this.emp.height}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            血型
          </template>
          {{this.emp.blood_type}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            婚姻状况
          </template>
          {{this.emp.marital_status}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            出生地
          </template>
          {{this.emp.birthplace}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            户口所在地
          </template>
          {{this.emp.registered_residence}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            部门编号
          </template>
          {{this.emp.department_number}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            部门名称
          </template>
          {{this.emp.department_name}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            岗位编号
          </template>
          {{this.emp.post_number}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            岗位名称
          </template>
          {{this.emp.post_name}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            入职日期
          </template>
          {{this.emp.entry_date}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            用工形式
          </template>
          {{this.emp.employment_form}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            人员来源
          </template>
          {{this.emp.personnel_source}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            最高学历
          </template>
          {{this.emp.highest_education}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            毕业院校
          </template>
          {{this.emp.graduation_school}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            所学专业
          </template>
          {{this.emp.major_studied}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            毕业日期
          </template>
          {{this.emp.graduation_date}}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            是否离职
          </template>
          {{this.emp.is_quit}}
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <!--  主管审核自己部门的申请-->
    <div>
      <span>人才申请信息</span>
      <el-table :data="applysData"
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
            prop="talent_id"
            label="人才编号"
            fixed
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
            prop="department_name"
            width="75"
            align="left"
            label="部门名称"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_name"
            label="岗位名称"
            align="left"
            width="75"
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
        <el-table-column fixed="right" width="200" label="操作">
          <template slot-scope="scope">
            <el-button @click="agreeApply(scope.row)" style="padding: 3px" size="mini" type="success">同意</el-button>
            <el-button @click="refuseApply(scope.row)" style="padding: 3px" size="mini" type="danger">拒绝</el-button>
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
            :total="applys.length">
          <!--                        :total="total">-->
        </el-pagination>
      </div>
    </div>

  </div>
</template>

<script>
import {Message} from "element-ui";

export default {
  name: "PostTalent",
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
      is_director : true,
      title : "主管信息",
      emps : [],
      apply : {
        talent_id : "",
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
        entry_date : "",
        employment_form : "",
        personnel_source : "",
        highest_education : "",
        graduation_school : "",
        major_studied : "",
        graduation_date : "",
        is_agree : "",
      },
      applys : [],
      size : '',
      loading : false,
      currentPage : 1,
      pageSize : 10,
    }
  },
  mounted() {
    this.initEmps();
    localStorage.setItem("type", this.emp.employee_type);
  },
  computed : {
    applysData() {
      console.log("applys.length = ", this.applys.length);
      if (this.applys.length > 0) {
        return this.applys.slice((this.currentPage - 1) * this.pageSize, this.currentPage * this.pageSize) || [];
      }
      return this.applys;
    }

  },
  methods : {
    initEmps(type) {
      this.loading = true;
      this.emp.id = localStorage.getItem("id");

      this.$axios.post('/DirectorDepartment', this.emp).then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
          this.emp = resp.data;
          console.log("员工类型:", this.emp.employee_type);
          if (this.emp.employee_type === "总裁") {
            this.is_director = false;
            this.title = "总裁信息";
          }
          this.total = resp.total;
          console.log("emp = ", this.emp);
          this.initApplyData();
        }
      })

    },
    initApplyData() {
      // this.initEmps();
      console.log("initApplyDataEmp:", this.emp);
      // console.log("emps:", this.emps);
      if (this.emp.employee_type === "总裁") {
        this.$axios.post('/TalentApply', this.emp).then(resp => {
          if (resp) {
            this.applys = resp.data;
            console.log("applys = ", this.applys);
          }
        })
      } else {
        console.log("")
        this.$axios.post('/TalentApplyPart', this.emp).then(resp => {
          if (resp) {
            this.applys = resp.data;
            console.log("applys = ", this.applys);
          }
        })
      }
    },
    agreeApply(data) {
      this.$axios.post('/TalentAgreeApply', data).then(resp => {
        if (resp.data.data) {
          Message.success({message : resp.data.msg});
        } else {
          Message.error({message : resp.data.msg});
        }
        this.initApplyData();
      })
    },
    refuseApply(data) {
      this.$axios.post('/TalentRefuseApply', data).then(resp => {
        if (resp.data.data) {
          Message.info({message : resp.data.msg});
        } else {
          Message.error({message : resp.data.msg});
        }
        this.initApplyData();
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
  }
}
</script>

<style scoped>

</style>
