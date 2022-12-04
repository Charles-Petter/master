<template>
  <div>
    <!--    部门信息-->
    <div>
      <el-col :xs="12" :sm="12" :md="12" :lg="12" :xl="12">
        <el-col :span="16">
<!--          <el-dialog-->
<!--              :title="title"-->
<!--              @close="resetForm('emp')"-->
<!--              :visible.sync="dialogEditVisible"-->
<!--              width="80%">-->
<!--            <div>-->
<!--              <el-form :model="emp" :rules="rules" ref="empForm">-->

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

        <el-table-column  width="75"  prop="initsalay"  v-model="initsalay" label="初始工资">
        </el-table-column>

        <el-table-column  width="150" :show-overflow-tooltip="true" prop="classtime_cx" label="课时量" >
<!--          <input type="text" placeholder="input search text"  @change="onchange"  v-model="tims">-->
<!--          <a-input-search style="width: 15%;height:3%;margin-top: 15px;margin-left: 10px;"   @change="onchange">-->
<!--            <a-button slot="enterButton"  >-->
<!--              搜索-->
<!--            </a-button>-->
<!--          </a-input-search>-->
          <template slot-scope="scope">
<!--            scope.row.salayCount_cx-->
<!--            1.获取输入框的值-->
            <el-input v-model.salayCount_cx="scope.row.classtime_cx" @change="onchange" ref="taxafter" :formatter="state" />
<!--            获取prop=taxafter的值-->
<!--            <el-link :underline="false" icon="el-icon-edit" @click="clickChange(scope.row.taxafter)">修改</el-link>-->
<!--            v-show不支持在template中使用-->
<!--            <p>  {{}}</p>-->
<!--            <p class="TaxInfo">实发工资：<span>{{actsalary}} 元</span></p>-->
          </template>
        </el-table-column>
        <el-table-column
            prop="taxafter"
            label="实际工资"
            align="left"
            width="150"
            sortable
            show-overflow-tooltip>
          <p class="TaxInfo">实发工资：<span>{{actsalary}} 元</span></p>
        </el-table-column>
        <el-table-column
            fixed="right"
            width="200"
            label="修改"
          >
          <template slot-scope="scope">
            <el-button @click="doEditSalary(scope.row)"  style="padding: 3px" size="mini" type="warning">计算工资</el-button>
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
          //这里是值对应的处理
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
      // row.taxafter=this.actsalary
      console.log("ccc")
      console.log(row.taxafter);
    },
    //修改工资条
    doEditSalary() {
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          // this.emp.height = parseInt(this.emp.height);
          row.taxafter=this.actsalary
          console.log(this.emp)
          this.$axios.post('/Employee/Salary_cx', this.emp).then(resp => {
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
//传值给后端
//     async save(){
//       const  res=await this.$http.pos('接口地址',this.model)
//       console.log(res)//返回结果
//       this.$message({
//         type:'success',
//         message:'保存成功'
//       })
//     },

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
    clickChange(){
      this.TaxAfter = this.initsalay  + this.times*50;
      console.log("ccc");
    },
    FuckMoney () {

      this.TaxAfter = this.initsalay  + this.times*50;

 // this.emps
    },
  }
}


</script>

<style scoped>

</style>