<template>
  <div>
<!--    部门信息-->
    <div>
      <span>
        部门信息
      </span>
      <el-button-group style="float: right">
        <el-button type="primary" icon="el-icon-plus" @click="showAddPostView" v-if="is_boss" :disabled="!open">
          添加岗位
        </el-button>
      </el-button-group>
      <el-button-group style="float: right">
        <el-button type="primary" icon="el-icon-plus" @click="showAddEmpView" :disabled="!is_boss">
          添加部门
        </el-button>
      </el-button-group>
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

        <el-table-column
            fixed="right"
            width="200"
            label="操作">
          <template slot-scope="scope">
            <el-button @click="showEditEmpView(scope.row)" v-if="is_boss" style="padding: 3px" size="mini" type="warning">编辑</el-button>
            <el-button @click="initPost(scope.row)" style="padding: 3px" size="mini" type="primary">查看岗位信息</el-button>
            <el-button @click="deleteEmp(scope.row)" v-if="is_boss" style="padding: 3px" size="mini" type="danger">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div>
      <span>
        岗位信息
      </span>
<!--      岗位信息-->
      <el-table :data="postsData" stripe border
                v-loading="loading"
                element-loading-text="正在加载..."
                element-loading-spinner="el-icon-loading"
                element-loading-background="rgba(0, 0, 0, 0.8)"
                style="width: 100%">
        <el-table-column prop="post_number" label="岗位编号" align="left" width="100"></el-table-column>
        <el-table-column prop="post_name" label="岗位名称" align="left" width="100"></el-table-column>
        <el-table-column prop="department_number" label="部门编号" align="left" width="100"></el-table-column>
        <el-table-column prop="post_type" label="岗位类型" align="left" width="100"></el-table-column>
        <el-table-column prop="post_establishment" label="岗位编制" align="left" width="100"></el-table-column>
        <el-table-column fixed="right" width="200" label="操作">
          <template slot-scope="scope">
            <el-button @click="showEditPostView(scope.row)" v-if="is_boss" style="padding: 3px" size="mini" type="warning">编辑</el-button>
            <el-button @click="deletePost(scope.row)" v-if="is_boss" style="padding: 3px" size="mini" type="danger">删除</el-button>
          </template>
        </el-table-column>

      </el-table>
    </div>
<!--    部门编辑-->
    <div>
      <el-dialog :title="title" @close="resetForm('emp')" :visible.sync="dialogEditVisible" width="80%">
        <div>
          <el-form :model="emp" :rules="rules" ref="emp">
            <el-row>
              <el-col :span="6">
                <el-form-item label="部门名称:" prop="department_name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.department_name" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门类型:" prop="department_type">
                  <el-select v-model="emp.department_type" placeholder="请选择部门类型" size="mini" style="width: 150px;">
                    <el-option v-for="item in department_types" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门电话:" prop="phone">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.phone" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门传真:" prop="fax">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.fax" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="部门描述:" prop="describe">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.describe" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="上级部门:" prop="superior_department">
                  <el-select v-model="emp.superior_department" placeholder="请选择上级部门" size="mini" style="width: 150px;">
                    <el-option v-for="item in department_names" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="成立日期:" prop="incorporation_date">
                  <el-date-picker
                      v-model="emp.incorporation_date"
                      size="mini"
                      type="date"
                      value-format="yyyy-MM-dd"
                      style="width: 150px;"
                      placeholder="请选择成立日期">
                  </el-date-picker>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogEditVisible = false">取 消</el-button>
          <el-button type="primary" @click="doEditEmp">确 定</el-button>
      </span>
      </el-dialog>
    </div>
<!--    部门添加-->
    <div>
      <el-dialog :title="title" @close="resetForm('emp')" :visible.sync="dialogAddVisible" width="80%">
        <div>
          <el-form :model="emp" :rules="rules" ref="emp">
            <el-row>
              <el-col :span="6">
                <el-form-item label="部门名称:" prop="department_name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.department_name" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门类型:" prop="department_type">
                  <el-select v-model="emp.department_type" placeholder="请选择部门类型" size="mini" style="width: 150px;">
                    <el-option v-for="item in department_types" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门电话:" prop="phone">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.phone" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="部门传真:" prop="fax">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.fax" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="6">
                <el-form-item label="部门描述:" prop="describe">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.describe" placeholder="请输入部门名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="上级部门:" prop="superior_department">
                  <el-select v-model="emp.superior_department" placeholder="请选择上级部门" size="mini" style="width: 150px;">
                    <el-option v-for="item in department_names" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="6">
                <el-form-item label="成立日期:" prop="incorporation_date">
                  <el-date-picker
                      v-model="emp.incorporation_date"
                      size="mini"
                      type="date"
                      value-format="yyyy-MM-dd"
                      style="width: 150px;"
                      placeholder="请选择成立日期">
                  </el-date-picker>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogAddVisible = false">取 消</el-button>
          <el-button type="primary" @click="doAddEmp">确 定</el-button>
      </span>
      </el-dialog>
    </div>
<!--    岗位添加-->
    <div>
      <el-dialog :title="title" @close="resetForm('post')" :visible.sync="dialogAddPostVisible" width="80%">
        <div>
          <el-form :model="post" :rules="postRules" ref="post">
            <el-row>
              <el-col :span="12">
                <el-form-item label="岗位名称:" prop="post_name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.post_name" placeholder="请输入岗位名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="部门编号:" prop="department_number">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.department_number" placeholder="请输入部门编号" clearable disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-form-item label="岗位类型:" prop="post_type">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.post_type" placeholder="请输入岗位类型" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="岗位编制:" prop="post_establishment">
                  <el-select v-model="post.post_establishment" placeholder="请选择岗位编制" size="mini" style="width: 150px;">
                    <el-option v-for="item in post_establishments" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogAddPostVisible = false">取 消</el-button>
          <el-button type="primary" @click="doAddPost">确 定</el-button>
        </span>
      </el-dialog>
    </div>
<!--    岗位编辑-->
    <div>
      <el-dialog :title="title" @close="resetForm('post')" :visible.sync="dialogEditPostVisible" width="80%">
        <div>
          <el-form :model="post" :rules="postRules" ref="post">
            <el-row>
              <el-col :span="12">
                <el-form-item label="岗位名称:" prop="post_name">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.post_name" placeholder="请输入岗位名称" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="部门编号:" prop="department_number">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.department_number" placeholder="请输入部门编号" clearable disabled></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="12">
                <el-form-item label="岗位类型:" prop="post_type">
                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="post.post_type" placeholder="请输入岗位类型" clearable></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="岗位编制:" prop="post_establishment">
                  <el-select v-model="post.post_establishment" placeholder="请选择岗位编制" size="mini" style="width: 150px;">
                    <el-option v-for="item in post_establishments" :key="item" :label="item" :value="item"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogEditPostVisible = false">取 消</el-button>
          <el-button type="primary" @click="doEditPost">确 定</el-button>
        </span>
      </el-dialog>
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
      rules: {
        department_number : [{required : true, message : '请输入部门编号', trigger : 'blur'}],
        department_name : [{required : true, message : '请输入部门名称', trigger : 'blur'}],
        department_type : [{required : true, message : '请输入部门类型', trigger : 'blur'}],
        department_head : [{require : true, message : '请输入部门主管员工编号', trigger : 'blur'}],
        phone : [{required : true, message : '请输入部门电话', trigger : 'blur'}],
        fax : [{required : true, message : '请输入部门传真', trigger : 'blur'}],
        describe : [{required : true, message : '请输入部门描述', trigger : 'blur'}],
        superior_department : [{required : true, message : '请输入上级部门', trigger : 'blur'}],
        incorporation_date : [{required : true, message : '请输入成立日期', trigger : 'blur'}],
      },
      postRules : {
        post_number : [{required : true, message : '请输入岗位编号', trigger : 'blur'}],
        post_name : [{required : true, message : '请输入岗位名称', trigger : 'blur'}],
        department_number : [{required : true, message : '请输入部门编号', trigger : 'blur'}],
        post_type : [{required : true, message : '请输入岗位类型', trigger : 'blur'}],
        post_establishment : [{required : true, message : '请输入岗位编制', trigger : 'blur'}],
      },
      emps : [],
      post : {
        post_number : "",
        post_name : "",
        department_number : "",
        post_type : "",
        post_establishment : "",
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
      department_types : ['一级部门', '二级部门', '三级部门'],
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_establishments : ['有编制', '无编制'],
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
    showMyData(data) {
      this.title = '查看部门所有岗位信息';
      this.emp = data;
      this.dialogShowVisible = true;
    },
    initEmps(type) {
      this.loading = true;
      var tempRole = localStorage.getItem("role");
      if (tempRole === "总裁") {
        this.is_boss = true;
      } else {
        this.is_boss = false;
      }
      // this.emp.id = localStorage.getItem("id");
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
    initPost(data) {
      this.open = true;
      this.$axios.post('/Department/Post', data).then(resp => {
        if (resp) {
          this.emp.department_number = data.department_number;
          this.post.department_number = this.emp.department_number;
          this.posts = resp.data;
          console.log("岗位信息：", this.posts);
        }
      })
    },
    showEditEmpView(data) {
      // this.initPositions();
      this.title = '编辑部门信息';
      this.emp = data;
      // this.inputDepName = data.department_name;
      this.dialogEditVisible = true;
    },
    deleteEmp(data) {
      this.$confirm('此操作将永久删除【' + data.department_name + '】, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$axios.post("/Department/Delete", data).then(resp => {
          if (resp.data.msg === "删除成功") {
            this.initEmps();
            Message.success({message : "删除成功"});
          } else {
            Message.error({message : resp.data.msg});
          }
        })
      }).catch(() => {
        Message.info({message : "已取消删除"});
      });
    },
    doEditEmp() {
      //编辑
      this.$refs['emp'].validate(valid => {
        if (valid) {
          // this.emp.height = parseInt(this.emp.height);
          this.$axios.post('/Department/Update', this.emp).then(resp => {
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
    doAddEmp() {
      this.$refs['emp'].validate(valid => {
        if (valid) {
          this.$axios.post('/Department/Add', this.emp).then(resp => {
            if (resp.data.msg === "新增成功") {
              this.dialogAddVisible = false;
              this.initEmps();
              Message.success({message : '新增成功'});
            } else {
              Message.error({message : resp.data.msg});
            }
          })
        }
      })
    },
    showAddPostView() {
      this.title = '添加' + this.emp.department_name + '岗位信息';
      this.dialogAddPostVisible = true;
    },
    showEditPostView(data) {
      this.title = '编辑岗位信息';
      this.post = data;
      this.dialogEditPostVisible = true;
    },
    deletePost(data) {
      this.$confirm('此操作将永久删除【' + data.post_name + '】, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$axios.post("/Post/Delete", data).then(resp => {
          if (resp.data.msg === "删除成功") {
            this.initEmps();
            this.initPost(this.emp);
            Message.success({message : "删除成功"});
          } else {
            Message.error({message : resp.data.msg});
          }
        })
      }).catch(() => {
        Message.info({message : "已取消删除"});
      });
    },
    doAddPost() {
      this.$refs['post'].validate(valid => {
        if (valid) {
          this.$axios.post('/Post/Add', this.post).then(resp => {
            if (resp.data.msg === "新增成功") {
              this.dialogAddPostVisible = false;
              this.initEmps();
              this.initPost(this.emp);
              Message.success({message : '新增成功'});
            } else {
              Message.error({message : resp.data.msg});
            }
          })
        }
      })
    },
    doEditPost() {
      this.$refs['post'].validate(valid => {
        if (valid) {
          // this.emp.height = parseInt(this.emp.height);
          this.$axios.post('/Post/Update', this.post).then(resp => {
            if (resp.data.msg === "修改成功") {
              this.dialogEditPostVisible = false;
              this.initEmps();
              this.initPost(this.emp);
              Message.success({message : '修改成功'});
            } else {
              Message.error({message : resp.data.msg});
            }
          });
        }
      });
    },
    initDepartment() {
      this.$axios.post('/Department/Init').then(resp => {
        this.department_names = resp.data;
        console.log("初始化部门：", this.department_names);
      })
    },
  }
}
</script>

<style scoped>

</style>