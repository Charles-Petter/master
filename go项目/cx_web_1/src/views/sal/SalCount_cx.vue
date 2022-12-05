<template>
  <div>
    <div>
      <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
        <el-col :span="16">
        </el-col>
      </el-col>
      <el-table :data="empsData"
                stripe border
                v-loading="loading"
                element-loading-text="正在加载..."
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)"
                style="width: 100%"
                :model="emp"  ref="empForm"
                @row-click="handleEdit">
          <el-table-column prop="id" fixed sortable align="left" label="账户" width="75" show-overflow-tooltip>
          </el-table-column>
          <el-table-column prop="name" fixed sortable align="left" label="姓名" width="75" show-overflow-tooltip>
          </el-table-column>
          <el-table-column prop="employee_type" fixed sortable align="left" label="员工类型" width="75" show-overflow-tooltip>
          </el-table-column>
          <el-table-column prop="sex" label="性别" align="left" width="75" sortable show-overflow-tooltip>
          </el-table-column>
        <el-table-column  width="75"  prop="initsalay"  v-model="initsalay" label="初始工资">
        </el-table-column>
        <el-table-column  width="150" :show-overflow-tooltip="true" prop="classtime_cx" label="课时量" >
          <template slot-scope="scope">
<!--            1.获取输入框的值-->
            <el-input v-model.salayCount_cx="scope.row.classtime_cx" @change="onchange" ref="taxafter" :formatter="state" />
          </template>
        </el-table-column>
        <el-table-column label="计算过程" align="left" width="150" sortable show-overflow-tooltip>
          <p class="TaxInfo">实发工资：<span>{{actsalary}} 元</span></p>
        </el-table-column>
        <el-table-column fixed="right" width="200" label="上传" v-if="!is_director">
        <template slot-scope="scope">
          <!--            主管上传工资-->
          <el-button  plain @click="showEditEmpView(scope.row)" v-if="!is_director" style="padding: 3px" size="mini" type="warning">上传工资信息</el-button>
        </template>
        </el-table-column>
        </el-table>
    </div>
    <!--      上传工资实际工资的编辑框-->
    <el-dialog :title="title" @close="resetForm('emp')" :visible.sync="dialogEditVisible" width="80%">
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
              <el-form-item label="姓名:" prop="name">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.name"
                          placeholder="请输入员工姓名"  clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="实际工资:" prop="taxafter">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-school" v-model="emp.taxafter" placeholder="准备上传工资" clearable></el-input>
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
              <el-button type="primary" @click="doEditSalary">确 定</el-button>
      </span>
    </el-dialog>
      </div>

</template>

<script>
import {Message} from "element-ui";
export default {

  name: "SalCount",
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
      TaxAfter: 0 , //税后工资
      actsalary:0,//实际工资
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
      loading : false,
      currentPage: 1,
      pageSize: 10,
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
  //模板渲染成html后调用
  mounted() {
    this.initEmps();
    this.initSalaries();
  },
  methods: {
//计算工资
    handleEdit (row) {
      console.log(row.initsalay);
      console.log(row.salayCount_cx);
      let a=row.initsalay-5000<0
      let b=row.initsalay>=0 && row.initsalay-5000<=3000
      let c = row.initsalay-5000>5000 && row.initsalay-5000<=8000
      console.log("a:",a)
      console.log("b:",b)
      // this.actsalary=(parseInt(row.initsalay)-5000)*(1-0.03)+5000
      console.log( this.actsalary)
      switch(true){
        case a:
          this.actsalary-=0   //工资-5000小于0 不用交税
          break
        case b:
          this.actsalary=(parseInt(row.initsalay)-5000)*(1-0.03)+5000+ (parseInt(row.classtime_cx)*100)
          break
        case c:
          this.actsalary=(parseInt(row.initsalay)-5000)*(1-0.1)+5000+ (parseInt(row.classtime_cx)*100)
          break
        default:
          console.log("未执行")
          break
      }
      console.log("ccc")
      console.log(row.taxafter);
    },
    //修改工资条
    doEditSalary() {
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          console.log(this.emp)
          this.$axios.post('/Employee/Salary_cx', this.emp).then(resp => {
            if (resp) {
              this.dialogEditVisible = false;
              this.initEmps();
              Message.success({message : '上传成功!'});
            } else {
              Message.error({message : resp.data.msg});
            }
          });
        }
      });
    },
    showEditEmpView(data) {
      this.title = '编辑信息';
      this.emp = data;
      this.inputDepName = data.department_name;
      this.dialogEditVisible = true;
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
    // search(event){
    //   event.currentTarget.value
    //   console.log(event.currentTarget.value)
    // },
    //刷新表单
    resetForm(data) {
      console.log("data = ", data, "emp = ", this.emp)
      this.$refs[data].resetFields();
      this.initEmps();
    },
    //初始化工资表
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