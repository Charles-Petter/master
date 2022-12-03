<template>
  <div>
    <!--    部门信息-->
    <div>

      <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
        <el-col :span="16">
          <el-form :model="emp" :rules="rules" ref="emp" @submit.native.prevent>
          </el-form>
        </el-col>
      </el-col>
      <el-table :data="empsData"
                stripe border
                v-loading="loading"
                element-loading-text="正在加载..."
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)"
                style="width: 100%">
          <el-table-column
              prop="id"
              fixed
              sortable
              align="left"
              label="账户"
              width="75"
              show-overflow-tooltip>
          </el-table-column>
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
              label="员工类型"
              width="75"
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
        <el-table-column  width="75"  prop="initsalay" label="初始工资">
          <template slot-scope="scope">
            <span class="input-group-addon" >税前工资</span>
            <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3"
                   v-model="TaxBefore"
                   @keyup.enter="search" @input="search($event)"
            >
            <span class="input-group-addon" >元</span>
            <el-input v-model.initsalay_cx="scope.row.initsalay_cx"  />
          </template>
        </el-table-column>
        <el-table-column  width="150" :show-overflow-tooltip="true" prop="taxafter" label="税后工资">
          <template slot-scope="scope">
            <el-input v-model.salayCount_cx="scope.row.salayCount_cx"  />
<!--            v-show不支持在template中使用-->
            <p class="TaxInfo">实发工资：<span v-show="FuckMoney > 0">{{TaxAfter}} 元</span></p>
          </template>
        </el-table-column>

        </el-table>
      </div>

    </div>
</template>

<script>
import {Message} from "element-ui";
export default {

  name: "SalSob",
  data() {
    return {
      emp: {
        id : "",
        name : "",
        employee_type : "",
        sex : "",
        initsalay:"",
        Taxafter:"",

      },
      FuckMoney: 0,	     //应参加缴税金额
      TaxAfter: 0 ,        //税后工资
      emps : [],
      post : {
        id : "",
        employee_type : "",
        name : "",
        sex : "",
        initsalay:"",
        taxafter:"",
      },
      posts : [],
      loading : false,
      currentPage: 1,
      pageSize: 10,
      title : "岗位信息",
      size : '',
      is_boss : true,
      dialogEditVisible : false,
      dialogAddVisible : false,
      dialogAddPostVisible : false,
      dialogEditPostVisible : false,
      open : false,
    }
  },
  computed : {
    empsData() {
      console.log("emps.length = ", this.emps.length);
      if (this.emps.length > 0) {
        return this.emps.slice( (this.currentPage -1) * this.pageSize, this.currentPage * this.pageSize) || [];
      }
      console.log("emps = ", this.emps);
      return this.emps;
    },
    postsData() {
      if (this.posts.length > 0) {
        return this.posts.slice((this.currentPage - 1) * this.pageSize, this.currentPage * this.pageSize) || [];
      }
      return this.posts;
    }
  },
  mounted() {
    this.initEmps();
    this.initSalaries();
  },
  methods: {
    initEmps(type) {
      this.loading = true;
      this.$axios.post('/SearchEmpSalary').then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
        }
      })
    },
    search(event){
      event.currentTarget.value
      console.log(event.currentTarget.value)
    },
    resetForm(data) {
      console.log("data = ", data, "emp = ", this.emp)
      this.$refs[data].resetFields();
      this.initEmps();
    },
    initSalaries() {
      this.getRequest("/EmpSalary/Init/").then(resp => {
        if (resp) {
          this.salaries = resp;
        }
      })
    },
  }
}
</script>

<style scoped>

</style>