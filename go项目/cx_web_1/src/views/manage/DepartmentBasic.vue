<template>
  <div>
    <!--    部门信息-->
    <div>
      <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
        <el-col :span="16">
          <el-form :model="emp" :rules="rules" ref="emp" @submit.native.prevent>
            <el-form-item prop="department_name">
              <el-input placeholder="请输入部门名称" prefix-icon="el-icon-search"
                        clearable
                        @clear="initEmps"
                        style="width: 350px;margin-right: 10px" v-model="emp.department_name"
                        @keydown.enter.native="searchEmp('emp')" :disabled="showAdvanceSearchView || showDateSearchView"></el-input>
              <el-button-group>
                <el-button icon="el-icon-search" type="primary" @click="searchEmp('emp')" :disabled="showAdvanceSearchView || showDateSearchView">搜索</el-button>
                <el-button icon="el-icon-refresh" type="info" @click="resetForm('emp')" :disabled="showAdvanceSearchView || showDateSearchView">重置</el-button>
              </el-button-group>
            </el-form-item>
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
            prop="department_number"
            fixed
            align="left"
            label="部门编号"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_name"
            fixed
            align="left"
            label="部门名称"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_type"
            label="部门类型"
            align="left"
            width="85"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_head"
            label="部门主管员工编号"
            align="left"
            width="150"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="phone"
            label="部门电话"
            align="left"
            width="150"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="fax"
            width="100"
            align="left"
            label="传真"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="describe"
            width="150"
            align="left"
            label="描述"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="superior_department"
            width="150"
            align="left"
            label="上级部门"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="incorporation_date"
            label="成立日期"
            align="left"
            width="150"
            show-overflow-tooltip>
        </el-table-column>
      </el-table>
    </div>
  </div>
 </template>


<script>
import {Message} from "element-ui";

export default {
  name: "DepartmentBasic",
  data() {
    return {
      emp: {
        department_number : "",
        department_name : "",
        department_type : "",
        department_head : "",
        phone : "",
        fax : "",
        describe : "",
        superior_department : "",
        incorporation_date : "",
      },
      emps : [],
      // post : {
      //   post_number : "",
      //   post_name : "",
      //   department_number : "",
      //   post_type : "",
      //   post_establishment : "",
      // },
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
      // department_types : ['一级部门', '二级部门', '三级部门'],
      // department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      // post_establishments : ['有编制', '无编制'],
    }
  },
  mounted() {
    this.initEmps();
    this.initDepartment();
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
  methods : {
    initEmps(type) {
      this.loading = true;
      this.$axios.post('/Department/Basic').then(resp => {
        this.loading = false;
        if (resp) {
          this.emps = resp.data;
        }
      })
    },
    resetForm(data) {
      console.log("data = ", data, "emp = ", this.emp)
      this.$refs[data].resetFields();
      this.initEmps();
    },
    // initPost(data) {
    //   this.open = true;
    //   this.$axios.post('/Department/Post', data).then(resp => {
    //     if (resp) {
    //       this.emp.department_number = data.department_number;
    //       this.post.department_number = this.emp.department_number;
    //       this.posts = resp.data;
    //       console.log("岗位信息：", this.posts);
    //     }
    //   })
    // },
    // showEditEmpView(data) {
    //   // this.initPositions();
    //   this.title = '编辑部门信息';
    //   this.emp = data;
    //   // this.inputDepName = data.department_name;
    //   this.dialogEditVisible = true;
    // },

    emptyEmp() {
      this.emp = {
        department_number : "",
        department_name : "",
        department_type : "",
        department_head : "",
        phone : "",
        fax : "",
        describe : "",
        superior_department : "",
        incorporation_date : "",
      }
    },
    showAddEmpView() {
      this.emptyEmp();
      this.title = '添加部门信息';
      this.dialogAddVisible = true;
    },

    // initDepartment() {
    //   this.$axios.post('/Department/Init').then(resp => {
    //     this.department_names = resp.data;
    //     console.log("初始化部门：", this.department_names);
    //   })
    // },


    //查询部门实现(搜索框)
    async searchEmp(data) {
      console.log("department_name = ", this.emp.department_name);
      var url;
      if (localStorage.getItem("role") === "主管") {
        // url = '/EmployeeBasic/SearchByDirector';
        url = '/DepartmentSearch';
        var temp = {
          'id' : localStorage.getItem("id"),
          'department_name' : this.emp.department_name,
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

  }
}
</script>

<style scoped>

</style>