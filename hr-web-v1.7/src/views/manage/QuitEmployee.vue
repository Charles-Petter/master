<template>
  <div>
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
        <download-excel :data="emps" :fields="json_fields" :name="tableName">
          <el-button type="success" icon="el-icon-download" :disabled="emps.length == 0">
            导出数据
          </el-button>
        </download-excel>
      </el-button-group>
    </div>
    <div>
      <el-table :data="empsData" stripe border v-loading="loading" element-loading-text="正在加载..." element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)" style="width: 100%">
        <el-table-column prop="id" align="left" label="员工编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="name" align="left" label="姓名" show-overflow-tooltip></el-table-column>
        <el-table-column prop="sex" align="left" label="性别" show-overflow-tooltip></el-table-column>
        <el-table-column prop="resignation_date" align="left" label="离职日期" show-overflow-tooltip></el-table-column>
        <el-table-column prop="reasons_for_resignation" align="left" label="离职原因" show-overflow-tooltip></el-table-column>
        <el-table-column prop="department_number" align="left" label="部门编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="department_name" align="left" label="部门名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="post_number" align="left" label="岗位编号" show-overflow-tooltip></el-table-column>
        <el-table-column prop="post_name" align="left" label="岗位名称" show-overflow-tooltip></el-table-column>
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
  name: "QuitEmployee",
  data() {
    return {
      loading : false,
      emps : [],
      tableName : "离职员工表",
      currentPage : 1,
      pageSize : 10,
      json_fields: {
        //导出Excel表格的表头设置
        '员工编号': 'id',
        '姓名': 'name',
        '性别': 'sex',
        '离职日期': 'resignation_date',
        '离职原因': 'reasons_for_resignation',
        '部门编号': 'department_number',
        '部门名称': 'department_name',
        '岗位编号': 'post_number',
        '岗位名称': 'post_name',
      },
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
      let temp = {
        id : "",
      }
      if (localStorage.getItem("role") === "总裁") {
        temp.id = localStorage.getItem("id");
        this.$axios.post('/EmployeeQuit').then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            // this.total = resp.total;
          }
        })
      } else {
        temp.id = localStorage.getItem("id");
        this.getDepartmentName(temp);
        this.$axios.post('/EmployeeQuitByDirector', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            // this.total = resp.total;
          }
        })
      }

    },
    getDepartmentName(id) {
      this.$axios.post('/GetDepartmentName', id).then(resp => {
        this.tableName = resp.data.data + "离职员工表";
      });
    },
    sizeChange(currentSize) {
      this.pageSize = currentSize;
      this.initEmps();
    },
    currentChange(currentPage) {
      this.currentPage = currentPage;
      this.initEmps();
    },
    selectDate() {
      if ((this.value1 === "" || this.value1 === null) || (this.value2 === "" || this.value2 === null)) {
        Message.error({message : '请输入开始和结束日期'});
      } else {
        console.log("value1 = ", this.value1);
        var temp = {
          'id' : localStorage.getItem("id"),
          'start' : this.value1,
          'end' : this.value2,
        }
        var url;
        if (localStorage.getItem("role") === "总裁") {
          url = '/EmployeeQuit/SearchByDate';
        } else {
          url = '/EmployeeQuitByDirector/SearchByDate';
        }
        this.$axios.post(url, temp).then(resp => {
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