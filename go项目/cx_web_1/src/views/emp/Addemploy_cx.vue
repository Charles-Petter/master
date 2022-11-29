<template>
  <div>
    <div>
      <div style="display: flex;justify-content: space-between">
        <div>
          <el-form :model="emp" :rules="rules" ref="emp" @submit.native.prevent>
          </el-form>
        </div>
        <div>
<!--          添加用户按钮-->
          <el-button-group  >
            <el-button type="primary" icon="el-icon-plus" style="background: #36353c;" @click="showAddEmpView" :disabled="!is_director">
              <div class="container" >
                <a data-animation="ripple">添加 </a>
              </div>
            </el-button>
          </el-button-group>
        </div>
      </div>
<!--      <transition name="slide-fade">-->
<!--        <div v-show="showDateSearchView" style="position: center; border: 1px solid #40ff76;border-radius: 5px;box-sizing: border-box;padding: 5px;margin: 10px 0px;">-->
<!--          <div>-->
<!--            <span class="demonstration"></span>-->
<!--            <el-date-picker-->
<!--                v-model="value1"-->
<!--                type="date"-->
<!--                size="mini"-->
<!--                value-format="yyyy-MM-dd"-->
<!--                placeholder="开始日期"-->
<!--                :picker-options="pickerOptions0">-->
<!--            </el-date-picker>-->
<!--            <el-date-picker-->
<!--                v-model="value2"-->
<!--                type="date"-->
<!--                size="mini"-->
<!--                value-format="yyyy-MM-dd"-->
<!--                placeholder="结束日期"-->
<!--                :picker-options="pickerOptions1">-->
<!--            </el-date-picker>-->
<!--            <el-button-group>-->
<!--              <el-button type="primary" icon="el-icon-search" @click="selectDate"></el-button>-->
<!--              <el-button type="info" icon="el-icon-refresh" @click="resetDate"></el-button>-->
<!--            </el-button-group>-->
<!--          </div>-->
<!--        </div>-->
<!--      </transition>-->
<!--      <transition name="slide-fade">-->
<!--        <div v-show="showAdvanceSearchView"-->
<!--             style="border: 1px solid #409eff;border-radius: 5px;box-sizing: border-box;padding: 5px;margin: 10px 0px;">-->
<!--          <el-form :model="searchValue" :rules="rules" ref="searchValue">-->
<!--            <el-row>-->

<!--              <el-col :span="6">-->
<!--                <el-form-item label="政治面貌:" prop="political">-->
<!--                  <el-select v-model="searchValue.political" placeholder="政治面貌" size="mini"-->
<!--                             style="width: 130px;">-->
<!--                    <el-option-->
<!--                        v-for="item in politicals"-->
<!--                        :key="item"-->
<!--                        :label="item"-->
<!--                        :value="item">-->
<!--                    </el-option>-->
<!--                  </el-select>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--              <el-col :span="6">-->
<!--                <el-form-item label="民族:" prop="nation">-->
<!--                  <el-select v-model="searchValue.nation" placeholder="民族" size="mini"-->
<!--                             style="width: 130px;" filterable>-->
<!--                    <el-option-->
<!--                        v-for="item in nations"-->
<!--                        :key="item"-->
<!--                        :label="item"-->
<!--                        :value="item">-->
<!--                    </el-option>-->
<!--                  </el-select>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--              <el-col :span="6">-->
<!--                <el-form-item label="部门:" prop="department_name">-->
<!--                  <el-select @change="changeSelect" v-model="searchValue.department_name" placeholder="部门" size="mini" style="width: 130px;">-->
<!--                    <el-option-->
<!--                        v-for="(item, index) in department_names"-->
<!--                        :key="index"-->
<!--                        :label="item"-->
<!--                        :value="item">-->
<!--                    </el-option>-->
<!--                  </el-select>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--              <el-col :span="6">-->
<!--                <el-form-item label="岗位:" prop="post_name">-->
<!--                  <el-select v-model="searchValue.post_name" placeholder="岗位" size="mini"-->
<!--                             style="width: 130px;">-->
<!--                    <el-option-->
<!--                        v-for="(item, index) in post_type_options"-->
<!--                        :key="index"-->
<!--                        :label="item"-->
<!--                        :value="item">-->
<!--                    </el-option>-->
<!--                  </el-select>-->
<!--                </el-form-item>-->
<!--              </el-col>-->

<!--            </el-row>-->
<!--            <el-row style="margin-top: 10px">-->

<!--              <el-col :span="8">-->
<!--                <el-form-item label="毕业院校:" prop="graduation_school">-->
<!--                  <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit"-->
<!--                            v-model="searchValue.graduation_school" placeholder="毕业院校名称" clearable></el-input>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--              <el-col :span="8">-->
<!--                <el-form-item label="所学专业:" prop="major_studied">-->
<!--                  <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"-->
<!--                            v-model="searchValue.major_studied" placeholder="请输入专业名称" clearable></el-input>-->
<!--                </el-form-item>-->
<!--              </el-col>-->
<!--              <el-col :span="8" :offset="4">-->
<!--                <el-button icon="el-icon-refresh" type="info" size="mini" @click="resetForm('searchValue')">取消</el-button>-->
<!--                <el-button size="mini" icon="el-icon-search" type="primary" @click="searchEmpAdvance('searchValue')">搜索</el-button>-->
<!--              </el-col>-->

<!--            </el-row>-->
<!--          </el-form>-->
<!--        </div>-->
<!--      </transition>-->
    </div>
    <div style="margin-top: 10px">
      <el-table
          :data="empsData"
          stripe
          border
          :default-sort = "{prop: 'date', order: 'descending'}"
          v-loading="loading"
          @selection-change="handleSelectionChange"
          element-loading-text="正在加载..."
          element-loading-spinner="el-icon-loading"
          element-loading-background="rgba(0, 0, 0, 0.8)"
          style="width: 100%">
        <el-table-column
            type="selection"
            width="55"
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
            label="角色"
            width="75"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="id"
            label="编号"
            align="left"
            width="85"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="password"
            label="密码"
            align="left"
            width="85"
            v-if="is_director"
            sortable
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
        <el-table-column
            prop="birthday"
            width="100"
            align="left"
            label="出生日期"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="id_card"
            width="150"
            align="left"
            label="身份证号码"
            v-if="is_director"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="political"
            width="100"
            label="政治面貌"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="nation"
            label="民族"
            align="left"
            width="85"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="native_place"
            label="籍贯"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="phone"
            label="电话"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="email"
            label="电子邮箱"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="height"
            label="身高"
            align="left"
            width="75"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="blood_type"
            label="血型"
            align="left"
            width="75"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="marital_status"
            label="婚姻状况"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="birthplace"
            label="出生地"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="registered_residence"
            label="户口所在地"
            align="left"
            width="125"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_number"
            label="部门编号"
            align="left"
            width="125"
            sortable
            :filters="filterDepartmentText"
            :filter-method="filterHandler"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="department_name"
            fixed
            width="100"
            align="left"
            label="部门名称"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_number"
            label="岗位编号"
            align="left"
            width="125"
            sortable
            :filters="filterPostText"
            :filter-method="filterHandler"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="post_name"
            label="岗位名称"
            fixed
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="entry_date"
            label="入职日期"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="employment_form"
            label="用工形式"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="personnel_source"
            label="人员来源"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="highest_education"
            label="最高学历"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_school"
            label="毕业院校"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="major_studied"
            label="所学专业"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="graduation_date"
            label="毕业日期"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            prop="is_quit"
            label="是否离职"
            align="left"
            width="100"
            sortable
            show-overflow-tooltip>
        </el-table-column>


        <!--                <el-table-column-->
        <!--                        prop="position.name"-->
        <!--                        width="100"-->
        <!--                        label="职位">-->
        <!--                </el-table-column>-->
        <!--                <el-table-column-->
        <!--                        prop="jobLevel.name"-->
        <!--                        width="100"-->
        <!--                        label="职称">-->
        <!--                </el-table-column>-->
        <!--                <el-table-column-->
        <!--                        prop="conversionTime"-->
        <!--                        width="95"-->
        <!--                        align="left"-->
        <!--                        label="转正日期">-->
        <!--                </el-table-column>-->
        <!--                <el-table-column-->
        <!--                        prop="beginContract"-->
        <!--                        width="95"-->
        <!--                        align="left"-->
        <!--                        label="合同起始日期">-->
        <!--                </el-table-column>-->
        <!--                <el-table-column-->
        <!--                        prop="endContract"-->
        <!--                        width="95"-->
        <!--                        align="left"-->
        <!--                        label="合同截止日期">-->
        <!--                </el-table-column>-->
        <!--                <el-table-column-->
        <!--                        width="100"-->
        <!--                        align="left"-->
        <!--                        label="合同期限">-->
        <!--                    <template slot-scope="scope">-->
        <!--                        <el-tag>{{scope.row.contractTerm}}</el-tag>-->
        <!--                        年-->
        <!--                    </template>-->
        <!--                </el-table-column>-->
<!--        <el-table-column-->
<!--            fixed="right"-->
<!--            width="200"-->
<!--            label="操作">-->
<!--          <template slot-scope="scope">-->
<!--            <el-button @click="showEditEmpView(scope.row)" v-if="is_director" style="padding: 3px" size="mini" type="warning">编辑</el-button>-->
<!--            <el-button @click="showAllData(scope.row)" style="padding: 3px" size="mini" type="primary">查看完整资料</el-button>-->
<!--            <el-button @click="deleteEmp(scope.row)" v-if="is_director" style="padding: 3px" size="mini" type="danger">删除-->
<!--            </el-button>-->
<!--          </template>-->
<!--        </el-table-column>-->
      </el-table>

<!--      <el-dialog :visible.sync="multiDeleteVisible" title="提示" width="30%">-->
<!--        <span>确定要删除吗?</span>-->
<!--        <span slot="footer">-->
<!--          <el-button type="danger" @click="multiDelete">确 定</el-button>-->
<!--          <el-button @click="dialogVisible = false; multiDeleteVisible = false">取 消</el-button>-->
<!--        </span>-->
<!--      </el-dialog>-->
<!--      <div style="display: flex;justify-content: flex-end">-->
<!--        <el-pagination-->
<!--            background-->
<!--            @current-change="currentChange"-->
<!--            @size-change="sizeChange"-->
<!--            :page-size="pageSize"-->
<!--            :page-sizes="[1,5,10,20,100]"-->
<!--            :current-page.sync="currentPage"-->
<!--            layout="sizes, prev, pager, next, jumper, ->, total, slot"-->
<!--            :total="emps.length">-->
<!--          &lt;!&ndash;                        :total="total">&ndash;&gt;-->
<!--        </el-pagination>-->
<!--      </div>-->
<!--    </div>-->
    <!--      编辑框-->
    <el-dialog
        :title="title"
        @close="resetForm('emp')"
        :visible.sync="dialogEditVisible"
        width="80%">
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
              <el-form-item label="密码:" prop="password">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-lock" v-model="emp.password"
                          placeholder="请输入密码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="员工类型:" prop="employee_type">
                <el-select v-model="emp.employee_type" placeholder="员工类型" size="mini" style="width: 150px;">
                  <el-option v-for="item in employee_types" :key="item" :label="item" :value="item"></el-option>
                </el-select>
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
              <el-form-item label="性别:" prop="sex">
                <el-radio-group v-model="emp.sex">
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
                    placeholder="出生日期">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="身份证号码:" prop="id_card">
                <el-input size="mini" style="width: 180px" prefix-icon="el-icon-s-custom"
                          v-model="emp.id_card" placeholder="请输入身份证号码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="政治面貌:" prop="political">
                <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;">
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
                <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;" filterable>
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
                <el-input size="mini" style="width: 120px" prefix-icon="el-icon-suitcase"
                          v-model="emp.native_place" placeholder="请输入籍贯" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电话号码:" prop="phone">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-phone"
                          v-model="emp.phone" placeholder="电话号码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电子邮箱:" prop="email">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                          v-model="emp.email" placeholder="请输入电子邮箱" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="身高:" prop="height">
                <el-input size="mini" style="width: 100px" type="number" prefix-icon="el-icon-edit"
                          v-model="emp.height" placeholder="请输入身高" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="血型:" prop="blood_type">
                <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;">
                  <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="婚姻状况:" prop="marital_status">
                <el-radio-group v-model="emp.marital_status">
                  <el-radio label="已婚">已婚</el-radio>
                  <el-radio label="未婚">未婚</el-radio>
                  <el-radio label="离异">离异</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="出生地:" prop="birthplace">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-s-opportunity"
                          v-model="emp.birthplace" placeholder="请输入出生地" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="户口所在地:" prop="registered_residence">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-house"
                          v-model="emp.registered_residence" placeholder="请输入户口所在地" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="部门编号:" prop="department_number">
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
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
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
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
                    placeholder="入职日期">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="用工形式:" prop="employment_form">
                <el-select v-model="emp.employment_form" placeholder="用工形式" size="mini" style="width: 150px;">
                  <el-option v-for="item in employee_forms" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="人员来源:" prop="personnel_source">
                <el-select v-model="emp.personnel_source" placeholder="人员来源" size="mini" style="width: 150px;">
                  <el-option v-for="item in personnel_sources" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="最高学历:" prop="highest_education">
                <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"
                           style="width: 150px;">
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
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-school"
                          v-model="emp.graduation_school" placeholder="毕业院校名称" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="所学专业:" prop="major_studied">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-notebook-1"
                          v-model="emp.major_studied" placeholder="请输入专业名称" clearable></el-input>
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
                    placeholder="毕业日期">
                </el-date-picker>

              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="是否离职" prop="is_quit">
                <!--                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                <!--                          v-model="emp.is_quit" placeholder="是否离职" clearable></el-input>-->
                <el-radio-group v-model="emp.is_quit">
                  <el-radio label="是">是</el-radio>
                  <el-radio label="否">否</el-radio>
                </el-radio-group>
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
              <el-button type="primary" @click="doEditEmp">确 定</el-button>
      </span>
    </el-dialog>
    <!--      输入框-->
    <el-dialog
        :title="title"
        @close="resetForm('emp')"
        :visible.sync="dialogAddVisible"
        width="80%">
      <div>
        <el-form :model="emp" :rules="rules" ref="empForm">
          <el-row>
            <el-col :span="6">
              <el-form-item label="编号:" prop="id">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user-solid"
                          v-model="emp.id" placeholder="编号" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="密码:" prop="password">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-lock" v-model="emp.password"
                          placeholder="请输入密码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="员工类型:" prop="employee_type">
                <el-select v-model="emp.employee_type" placeholder="员工类型" size="mini" style="width: 150px;">
                  <el-option v-for="item in employee_types" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="姓名:" prop="name">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-user" v-model="emp.name"
                          placeholder="请输入员工姓名" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="性别:" prop="sex">
                <el-radio-group v-model="emp.sex">
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
                    placeholder="出生日期">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="身份证号码:" prop="id_card">
                <el-input size="mini" style="width: 180px" prefix-icon="el-icon-s-custom"
                          v-model="emp.id_card" placeholder="请输入身份证号码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="政治面貌:" prop="political">
                <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;">
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
                <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;" filterable>
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
                <el-input size="mini" style="width: 120px" prefix-icon="el-icon-suitcase"
                          v-model="emp.native_place" placeholder="请输入籍贯" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电话号码:" prop="phone">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-phone"
                          v-model="emp.phone" placeholder="电话号码" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电子邮箱:" prop="email">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                          v-model="emp.email" placeholder="请输入电子邮箱" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="身高:" prop="height">
                <el-input size="mini" style="width: 100px" type="number" prefix-icon="el-icon-edit"
                          v-model="emp.height" placeholder="请输入身高" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="血型:" prop="blood_type">
                <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;">
                  <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="婚姻状况:" prop="marital_status">
                <el-radio-group v-model="emp.marital_status">
                  <el-radio label="已婚">已婚</el-radio>
                  <el-radio label="未婚">未婚</el-radio>
                  <el-radio label="离异">离异</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="出生地:" prop="birthplace">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-s-opportunity"
                          v-model="emp.birthplace" placeholder="请输入出生地" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="户口所在地:" prop="registered_residence">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-house"
                          v-model="emp.registered_residence" placeholder="请输入户口所在地" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="部门编号:" prop="department_number">
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
                          v-model="emp.department_number" type="number" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="部门名称:" prop="department_name">
                <el-select @change="changeEditSelect" v-model="emp.department_name" placeholder="请输入部门名称" size="mini" prefix-icon="el-icon-edit">
                  <el-option
                      v-for="(item, index) in department_names"
                      :key="index"
                      :label="item"
                      :value="item">
                  </el-option>
                </el-select>
                <!--                  <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                <!--                            v-model="emp.department_name"></el-input>-->
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="岗位编号:" prop="post_number">
                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-s-flag"
                          v-model="emp.post_number" type="number" clearable></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="岗位名称:" prop="post_name">
                <el-select v-model="emp.post_name" placeholder="请输入岗位名称" size="mini" prefix-icon="el-icon-edit">
                  <el-option
                      v-for="(item, index) in post_type_options"
                      :key="index"
                      :label="item"
                      :value="item">
                  </el-option>
                </el-select>
                <!--                  <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                <!--                            v-model="emp.post_name"></el-input>-->
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
                    placeholder="入职日期">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="用工形式:" prop="employment_form">
                <el-select v-model="emp.employment_form" placeholder="用工形式" size="mini" style="width: 150px;">
                  <el-option v-for="item in employee_forms" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="人员来源:" prop="personnel_source">
                <el-select v-model="emp.personnel_source" placeholder="人员来源" size="mini" style="width: 150px;">
                  <el-option v-for="item in personnel_sources" :key="item" :label="item" :value="item"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="最高学历:" prop="highest_education">
                <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"
                           style="width: 150px;">
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
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-school"
                          v-model="emp.graduation_school" placeholder="毕业院校名称" clearable></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="所学专业:" prop="major_studied">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-notebook-1"
                          v-model="emp.major_studied" placeholder="请输入专业名称" clearable></el-input>
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
                    placeholder="毕业日期">
                </el-date-picker>

              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="是否离职" prop="is_quit">
                <!--                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"-->
                <!--                          v-model="emp.is_quit" placeholder="是否离职" clearable></el-input>-->
                <el-radio-group v-model="emp.is_quit">
                  <el-radio label="是">是</el-radio>
                  <el-radio label="否">否</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>
      <span slot="footer" class="dialog-footer">
<!--          <el-tooltip :content="tipContent" effect="dark" placement="top-start">-->
        <!--            <el-button type="danger" size="mini">?</el-button>-->
        <!--            </el-tooltip>-->
        <template>
           <el-popover placement="bottom" title="注意事项" width="300" trigger="click" :content="tipContent">
          <el-button type="danger" size="mini" slot="reference">?</el-button>
        </el-popover>
        </template>

              <el-button @click="dialogAddVisible = false; resetForm('emp')">取 消</el-button>
              <el-button type="primary" @click="doAddEmp">确 定</el-button>
            </span>
    </el-dialog>
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
              <el-form-item label="密码:" prop="password" v-if="is_director">
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
              <el-form-item label="身份证号码:" prop="id_card" v-if="is_director">
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

import { Message } from "element-ui";
export default {
  name: "Addemploy_cx",
  data() {
    return {
      tipContent : "提示",
      currentPage: 1,
      pageSize: 10,
      searchValue: {
        political: null,
        nation: null,
        department_name : null,
        post_name : null,
        employment_form: null,
        graduation_school: null,
        major_studied: null
      },
      isAdd : false,
      title: '',
      emps: [],
      loading: false,
      popVisible: false,
      popVisible2: false,
      dialogEditVisible: false,
      dialogAddVisible : false,
      dialogShowVisible : false,
      is_director : true,
      total: 0,
      page: 1,
      keyword: '',
      size: 10,
      birthday: '',
      post_number: null,
      post_name: '',
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      blood_type:'',
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      employee_type:3,
      employee_types:['总裁', '主管', '员工'],
      nation: "",
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
      joblevels: [],
      political: "群众",
      politicals: ['群众', '共青团员', '中共预备党员', '中共党员', '无党派人士', '其他'],
      positions: [],
      highest_educations: ['小学', '初中', '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
      employee_forms:['实习生', '正式职工'],
      personnel_source:'',
      personnel_sources:['校招', '社招'],
      // options: [{
      //   value: '选项1',
      //   label: '黄金糕'
      // }, {
      //   value: '选项2',
      //   label: '双皮奶'
      // }, {
      //   value: '选项3',
      //   label: '蚵仔煎'
      // }, {
      //   value: '选项4',
      //   label: '龙须面'
      // }, {
      //   value: '选项5',
      //   label: '北京烤鸭'
      // }],
      inputDepName: '所属部门',
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


        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,
      },
      defaultProps: {
        children: 'children',
        label: 'name'
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
        phone : [{required : true, message : '请输入电话号码', trigger : 'blur'},
          {
            validator: function(rule, value, callback) {
              if (/^1[34578]\d{9}$/.test(value) == false) {
                callback(new Error("手机号格式不正确"));
              } else {
                callback();
              }
            },
            trigger: 'blur'
          }
        ],
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

        // workAge: [{required: true, message: '请输入工龄', trigger: 'blur'}],
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
      showDateSearchView : false,
    }
  },
  mounted() {
    this.initEmps();
    this.initData();
    this.initDepartment();
    this.initPost();
    this.initSearchText();
    this.tipContent = '部门和岗位的编号与名称请一一对应！！！ 例如:部门编号-部门名称(岗位1编号-岗位1名称, 岗位2编号-岗位2名称) 1000-开发部(10-C++开发, 11-Java开发, 12-C#开发, 13-Python开发, 14-Go开发), 1001-运维部(20-云运维, 21-服务器运维), 1002-测试部(30-Bug测试, 31-系统测试), 1003-设计部(40-UI设计, 41-动画设计), 1004-策划部(50-策划, 51-系统策划)';

    localStorage.setItem("type", this.emp.employee_type);
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
  methods: {
    searvhViewHandleNodeClick(data) {
      this.inputDepName = data.name;
      this.searchValue.department_number = data.id;
      this.popVisible2 = !this.popVisible2
    },

    uploadHttpRequest(param) {
      const formData = new FormData()
      formData.append('file', param.file)
      formData.append('uploadType', this.rules)
      const url = `/EmployeeBasic/Import`
      console.log("param = ", param);
      console.log("formData = ", formData);
      this.$axios.post(url, formData).then(resp => {
        const {data : {code, mark}} = resp
      })
    },
    exportData() {
      window.open('/EmployeeBasic/Export', '_parent');
    },

    emptyEmp() {
      this.emp = {
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


        nationId: 1,
        politicId: 13,
        jobLevelId: 9,
        posId: 29,
      }
      this.inputDepName = '';

    },
    doAddEmp() {
      this.$refs['empForm'].validate(valid => {
        if (valid) {
          console.log(this.emp);
          this.emp.height = parseInt(this.emp.height);
          this.emp.department_number = parseInt(this.emp.department_number);
          this.emp.post_number = parseInt(this.emp.post_number);
          this.$axios.post('/EmployeeBasic/Add', this.emp).then(resp => {
            if (resp) {
              this.dialogAddVisible = false;
              this.initEmps();
              if (resp.data.msg === "添加成功") {
                Message.success({message : '添加成功'});
              } else {
                Message.error({message : resp.data.msg})
              }

            } else {
              Message.error({message : resp.data.msg});
            }
          })
        }
      });
    },
    // handleNodeClick(data) {
    //   this.inputDepName = data.name;
    //   this.emp.department_number = data.id;
    //   this.popVisible = !this.popVisible
    // },
    // showDepView() {
    //   this.popVisible = !this.popVisible
    // },
    // showDepView2() {
    //   this.popVisible2 = !this.popVisible2
    // },
    // initPositions() {
    //   // this.getRequest('/employee/basic/positions').then(resp => {
    //   //     if (resp) {
    //   //         this.positions = resp;
    //   //     }
    //   // })
    //   this.$axios.post('/EmployeeBasic/Position').then(resp => {
    //     if (resp) {
    //       this.positions = resp;
    //     }
    //   })
    // },
    // getMaxWordID() {
    //   this.getRequest("/employee/basic/maxid").then(resp => {
    //     if (resp) {
    //       this.emp.id = resp.obj;
    //     }
    //   })
    // },
    // initData() {
    //
    // },
    sizeChange(currentSize) {
      this.pageSize = currentSize;
      this.initEmps();
    },
    currentChange(currentPage) {
      this.currentPage = currentPage;
      this.initEmps('advanced');
    },
    showAddEmpView() {
      this.emptyEmp();
      this.title = '添加员工信息';
      this.dialogAddVisible = true;
    },
    // async searchEmp(data) {
    //   console.log("name = ", this.emp.name);
    //   var url;
    //   if (localStorage.getItem("role") === "主管") {
    //     url = '/EmployeeBasic/SearchByDirector';
    //     var temp = {
    //       'id' : localStorage.getItem("id"),
    //       'name' : this.emp.name,
    //     };
    //     this.$refs[data].validate((valid) => {
    //       if (valid) {
    //         this.$axios.post(url, temp).then((resp) => {
    //           if (resp.data.msg === "查询成功") {
    //             this.emps = resp.data.data;
    //           } else {
    //             Message.error({message : resp.data.msg});
    //           }
    //         });
    //       }
    //     })
    //   } else {
    //     url = '/EmployeeBasic/SearchByEmployee';
    //     var temp = {
    //       'id' : localStorage.getItem("id"),
    //       'name' : this.emp.name,
    //     };
    //     this.$refs[data].validate((valid) => {
    //       if (valid) {
    //         this.$axios.post(url, temp).then((resp) => {
    //           if (resp.data.msg === "查询成功") {
    //             this.emps = resp.data.data;
    //           } else {
    //             Message.error({message : resp.data.msg});
    //           }
    //         });
    //       }
    //     })
    //   }
    //
    //
    // },
    // resetForm(data) {
    //   console.log("data = ", data, "emp = ", this.emp)
    //   this.$refs[data].resetFields();
    //   this.initEmps();
    // },

    // //高级搜索  根据部门查询
    // async searchEmpAdvance(data) {
    //   this.$refs[data].validate((valid) => {
    //     if (valid) {
    //       this.$axios.post('/EmployeeBasic/SearchAdvance', this.searchValue).then((resp) => {
    //         if (resp.data.msg == "查询成功") {
    //           this.emps = resp.data.data;
    //         } else {
    //           Message.error({message : resp.data.msg});
    //         }
    //       })
    //     }
    //   })
    // },
    initEmps(type) {
      this.loading = true;
      var tempRole = localStorage.getItem("role");
    if (tempRole === "主管") {
        let temp = {
          id : "",
        }
        temp.id = localStorage.getItem("id");
        this.$axios.post('/EmployeeBasicByDirector', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            this.total = resp.total;
          }
        })
      } else {
        this.is_director = false;
        let temp = {
          id : "",
        }
        temp.id = localStorage.getItem("id");
        this.$axios.post('/EmployeeBasicByEmployee', temp).then(resp => {
          this.loading = false;
          if (resp) {
            this.emps = resp.data;
            console.log("resp = ", this.emps);
            this.total = resp.total;
          }
        })
      }

    },
    popDelete() {
      this.multiDeleteVisible = true;
    },
    // multiDelete() {
    //   this.multiDeleteVisible = false;
    //   let checkArr = this.multipleSelection;   // multipleSelection存储了勾选到的数据
    //   let params = [];
    //   let self = this;
    //   let flag = true;
    //   checkArr.forEach(function (item) {
    //     if (localStorage.getItem("id") === item.id) {//不能删除自己
    //       console.log("不能删除自己");
    //       flag = false;
    //     } else if (localStorage.getItem("role") === "主管" && item.employee_type === "主管") {//身份为主管，不能删除其他主管
    //       console.log("执行力第二个false");
    //       flag = false;
    //     } else {
    //       params.push(item.id);       // 添加所有需要删除数据的id到一个数组，post提交过去
    //     }
    //   });
    //   if (flag === true) {
    //     console.log("批量删除的id：", params);
    //
    //     //  $http即是axios，可以在main.js里面设置 Vue.prototype.$http = axios;
    //     this.$axios.post('/EmployeeBasic/MultiDelete', params).then(function (res) {
    //       if (res.data.msg == "删除成功") {
    //         Message.success({message : "批量删除成功"});
    //       } else {
    //         Message.error({message : resp.data.msg});
    //       }
    //       // self.getFashionList(1, 1, 5);
    //     })
    //     this.initEmps();
    //   } else {
    //     Message.error({message : "不能删除自己和其他主管"});
    //     this.initEmps();
    //   }
    //   this.initEmps();
    //
    // },
    initDepartment() {
      this.$axios.post('/Department/Init').then(resp => {
        this.department_names = resp.data;
        console.log("初始化部门：", this.department_names);
      })
    },
    initPost() {
      var temp = {
        "department_name" : this.department_name
      }
      this.$axios.post('/Post/Init', temp).then(resp => {
        this.post_names = resp.data;
        console.log("初始化岗位：", this.post_names);
      })
    },
    // filterHandler(value, row, column) {
    //   const property = column['property'];
    //   return row[property] === value;
    // },
    // initSearchText() {
    //   this.$axios.post('/SearchByDepartmentNumber').then(resp => {
    //     this.filterDepartmentText = resp.data;
    //     console.log("Text : ", this.filterDepartmentText);
    //   })
    //   this.$axios.post('/SearchByPostNumber').then(resp => {
    //     this.filterPostText = resp.data;
    //     console.log("Text : ", this.filterPostText);
    //   })
    // },
    // selectDate() {
    //   if ((this.value1 === "" || this.value1 === null) || (this.value2 === "" || this.value2 === null)) {
    //     Message.error({message : '请输入开始和结束日期'});
    //   } else {
    //     console.log("value1 = ", this.value1);
    //     var url;
    //  if (localStorage.getItem("role") === "主管") {
    //       url = '/EmployeeBasic/SearchDateByDirector';
    //     } else {
    //       url = '/EmployeeBasic/SearchDateByEmployee';
    //     }
    //     var temp = {
    //       'id' : localStorage.getItem("id"),
    //       'start' : this.value1,
    //       'end' : this.value2,
    //     }
    //     this.$axios.post(url, temp).then(resp => {
    //       if (resp.data.data) {
    //         this.emps = resp.data.data;
    //         Message.success({message : '查询成功'});
    //       } else {
    //         Message.warning({message : '未找到符合条件的信息'});
    //       }
    //     })
    //     // this.initEmps();
    //   }
    //
    // },
    // resetDate() {
    //   this.value1 = ""
    //   this.value2 = ""
    //   this.initEmps();
    // }
  }
}
window["tmripple"] =
    /******/ (function(modules) { // webpackBootstrap
  /******/ 	// The module cache
  /******/ 	var installedModules = {};
  /******/
  /******/ 	// The require function
  /******/ 	function __webpack_require__(moduleId) {
    /******/
    /******/ 		// Check if module is in cache
    /******/ 		if(installedModules[moduleId]) {
      /******/ 			return installedModules[moduleId].exports;
      /******/ 		}
    /******/ 		// Create a new module (and put it into the cache)
    /******/ 		var module = installedModules[moduleId] = {
      /******/ 			i: moduleId,
      /******/ 			l: false,
      /******/ 			exports: {}
      /******/ 		};
    /******/
    /******/ 		// Execute the module function
    /******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
    /******/
    /******/ 		// Flag the module as loaded
    /******/ 		module.l = true;
    /******/
    /******/ 		// Return the exports of the module
    /******/ 		return module.exports;
    /******/ 	}
  /******/
  /******/
  /******/ 	// expose the modules object (__webpack_modules__)
  /******/ 	__webpack_require__.m = modules;
  /******/
  /******/ 	// expose the module cache
  /******/ 	__webpack_require__.c = installedModules;
  /******/
  /******/ 	// define getter function for harmony exports
  /******/ 	__webpack_require__.d = function(exports, name, getter) {
    /******/ 		if(!__webpack_require__.o(exports, name)) {
      /******/ 			Object.defineProperty(exports, name, {
        /******/ 				configurable: false,
        /******/ 				enumerable: true,
        /******/ 				get: getter
        /******/ 			});
      /******/ 		}
    /******/ 	};
  /******/
  /******/ 	// getDefaultExport function for compatibility with non-harmony modules
  /******/ 	__webpack_require__.n = function(module) {
    /******/ 		var getter = module && module.__esModule ?
        /******/ 			function getDefault() { return module['default']; } :
        /******/ 			function getModuleExports() { return module; };
    /******/ 		__webpack_require__.d(getter, 'a', getter);
    /******/ 		return getter;
    /******/ 	};
  /******/
  /******/ 	// Object.prototype.hasOwnProperty.call
  /******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
  /******/
  /******/ 	// __webpack_public_path__
  /******/ 	__webpack_require__.p = "";
  /******/
  /******/ 	// Load entry module and return exports
  /******/ 	return __webpack_require__(__webpack_require__.s = 0);
  /******/ })
    /************************************************************************/
    /******/ ([
      /* 0 */
      /***/ (function(module, exports, __webpack_require__) {

        "use strict";


// Default Settings
        var settings = {
          area: "",
          color: "rgba(255, 255, 255, 0.4)",
          offsetEl: null,
          eventListener: "click",
          mouseMove: false
        };

        /**
         * @description Where the magic happens
         * @param {object} e
         * @param {string} rippleColor
         * @param {string} eventListener
         */
        function ripple(e, rippleColor, eventListener) {
          var clickedEl = e.currentTarget;
          var PageX = eventListener.match(/touch/) ? e.changedTouches[0].pageX : e.clientX;
          var PageY = eventListener.match(/touch/) ? e.changedTouches[0].pageY : e.clientY;
          var btnWidth = clickedEl.clientWidth;
          var el = clickedEl.getBoundingClientRect();
          var rippleOffset = settings.offsetEl ? settings.offsetEl.clientHeight : 0;
          var btnOffsetTop = el.top + rippleOffset;
          var btnOffsetLeft = el.left;
          var posMouseX = PageX;
          var posMouseY = PageY + rippleOffset;
          var rippleX = posMouseX - btnOffsetLeft;
          var rippleY = posMouseY - btnOffsetTop;

          var baseCSS = "\n          position: absolute;\n          width: " + btnWidth * 2 + "px;\n          height: " + btnWidth * 2 + "px;\n          border-radius: 50%;\n          transition: transform 700ms, opacity 700ms;\n          transition-timing-function: cubic-bezier(0.250, 0.460, 0.450, 0.940);\n          background: " + rippleColor + ";\n          background-position: center;\n          background-repeat: no-repeat;\n          background-size: 100%;\n          left: " + (rippleX - btnWidth) + "px;\n          top: " + (rippleY - btnWidth) + "px;\n          transform: scale(0);\n          pointer-events: none;\n      ";

          // Prepare the dom
          var rippleEffect = document.createElement("span");
          rippleEffect.style.cssText = baseCSS;

          // Add some css for prevent overflow errors
          clickedEl.style.overflow = "hidden";

          // Check if the element is not static because the ripple is in absolute
          if (window.getComputedStyle(clickedEl).position === "static") {
            clickedEl.style.position = "relative";
          }

          // Check for the mousemove event
          if (settings.mouseMove) {
            settings.mouseMove = false;
            return;
          }

          clickedEl.appendChild(rippleEffect);

          // start animation
          requestAnimationFrame(function () {
            rippleEffect.style.cssText = baseCSS + " transform: scale(1); opacity: 0;";
          });

          setTimeout(function () {
            rippleEffect.remove();
          }, 700);
        }

        /**
         * @description Prevent ripple when scrolling (Mobile Only)
         * @param {string} eventListener
         */
        function onDrag(eventListener) {
          if (eventListener === "touchend") {
            document.getElementsByTagName("body")[0].addEventListener("touchmove", function () {
              settings.mouseMove = true;
            });
          }
        }

        function attachRipple(els, rippleColor, eventListener) {
          for (var i = 0; i < els.length; i += 1) {
            var currentBtn = els[i];
            currentBtn.addEventListener(eventListener, function (e) {
              return ripple(e, rippleColor, eventListener);
            });
          }
        }

        function attachRippleToAttribute(area, rippleColor, eventListener) {
          var attributeEl = document.querySelectorAll(area + " [data-animation='ripple']");

          if (attributeEl.length > 0) {
            attachRipple(attributeEl, rippleColor, eventListener);
          } else {
            throw new Error('not found any element with data-animation="ripple"');
          }
        }

        function attachRippleToSelectors(selectors, rippleColor, eventListener) {
          if (selectors) {
            var selectorsEl = document.querySelectorAll(selectors);
          } else {
            throw new Error("You have to enter at least 1 selector");
          }

          if (selectorsEl.length > 0) {
            attachRipple(selectorsEl, rippleColor, eventListener);
          } else {
            console.warn("No element found with this selector: ", selectors);
          }
        }

        module.exports = {
          init: function init() {
            var data = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};

            try {
              var area = settings.area,
                  color = settings.color,
                  offsetEl = settings.offsetEl,
                  eventListener = settings.eventListener;


              area = data.area || area;
              color = data.color || color;
              offsetEl = data.offsetEl ? this.setOffsetEl(data.offsetEl) : offsetEl;
              eventListener = data.eventListener || eventListener;

              onDrag(eventListener);
              attachRippleToAttribute(area, color, eventListener);
            } catch (e) {
              console.warn(e.message);
            }
          },
          attachToSelectors: function attachToSelectors() {
            var data = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {};

            try {
              var elSetting = {
                color: data.color || settings.color,
                eventListener: data.eventListener || settings.eventListener
              };
              var color = elSetting.color,
                  eventListener = elSetting.eventListener;


              attachRippleToSelectors(data.selectors, color, eventListener);
            } catch (e) {
              console.warn(e.message);
            }
          },
          setOffsetEl: function setOffsetEl(el) {
            settings.offsetEl = document.querySelector(el);
          },

          ripple: ripple
        };

        /***/ })
      /******/ ]);

tmripple.init()
</script>

<style>
/* 可以设置不同的进入和离开动画 */
/* 设置持续时间和动画函数 */
.slide-fade-enter-active {
  transition: all .8s ease;
}

.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter, .slide-fade-leave-to
  /* .slide-fade-leave-active for below version 2.1.8 */
{
  transform: translateX(10px);
  opacity: 0;
}
body {
  background: #36353c;
}

.container {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  height: 50px;
  width: 52px;
  margin: auto;
}
*[data-animation="ripple"] {
  height: 100%;
  width: 100%;
  display: block;
  outline: none;
  padding: 20px;
  color: #fff;
  text-transform: uppercase;
  background: linear-gradient(135deg, #e570e7 0%,#79f1fc 100%);
  box-sizing: border-box;
  text-align: center;
  line-height: 14px;
  font-family: roboto, helvetica;
  font-weight: 200;
  letter-spacing: 1px;
  text-decoration: none;
  box-shadow: 0 5px 3px rgba(0, 0, 0, 0.3);
  cursor: pointer;
  /*border-radius: 50px;*/
  -webkit-tap-highlight-color: transparent;
  border-radius: 5px;
}

*[data-animation="ripple"]:focus {
  outline: none;
}

*[data-animation="ripple"]::selection {
  background: transparent;
  pointer-events: none;
}
</style>
