<template>
  <div>
    <!--    部门信息-->
    <div>
      <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
        <el-col :span="16">
        </el-col>
      </el-col>
      <el-table :data="empsData" stripe border v-loading="loading" element-loading-text="正在加载..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)" style="width: 100%" :model="emp"  ref="empForm" @row-click="handleEdit">
        <el-table-column prop="id" fixed sortable align="left" label="账户" width="75" show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="name" fixed sortable align="left" label="姓名" width="75" show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="employee_type" fixed sortable align="left" label="员工类型" width="75" show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="sex" label="性别" align="left" width="75" sortable show-overflow-tooltip>
        </el-table-column>
        <el-table-column prop="taxafter" label="实发工资" align="left" width="150" sortable show-overflow-tooltip>
        </el-table-column>
      </el-table>
    </div>
  </div>

</template>

<script>
import {Message} from "element-ui";
export default {

  name: "SalEpm",
  data() {
    //税前工资
    return {
      emp: {
        id : "",
        name : "",
        employee_type : "",
        sex : "",
        initsalay:"",
        taxafter:0,
        times:0,
        classtime_cx:0,
      },
      FuckMoney: 0,	     //应参加缴税金额
      TaxAfter: 0 , //税后工资
      actsalary:0,
      // taxafter:0,
      map1:[],
      emps : [],
      post : {
        id : "",
        employee_type : "",
        name : "",
        sex : "",
        initsalay:"",
        taxafter:"",
        classtime_cx:0,
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
    isShow() {
      return this.emp.taxafter
    },
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
  watch: {
    isShow(newV,oldV) {
      // do something
      console.log(newV,oldV)
    },
    times:function (newValue,oldValue){
      console.log(newValue)
    },
  },
  mounted() {
    this.initEmps();
    this.initSalaries();
  },
  methods: {

    changeVal(val,params){
      console.log(val,params);
    },
    initEmps(type) {
      this.loading = true;
      this.$axios.post('/SearchEmpSalary_cx').then(resp => {
        this.loading = false;
        if (resp) {
          console.log('aaa')
          console.log(resp);
          this.emps = resp.data;
        }
      })
    },
    state(row, column) {
      return row.scale * 100
    },
    onchange(e){
      // this.$refs.taxafter.value=e.ta
      let val= e.target
      console.log("aaaa")
      console.log(val)
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
      this.getRequest("/EmpSalary/Init_cx/").then(resp => {
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