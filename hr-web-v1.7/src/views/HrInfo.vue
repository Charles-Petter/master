<template>
  <div>
    <!--      <div>-->
    <!--        <el-card>-->
    <!--          <el-form :rules="rules">-->
    <!--            <el-descriptions class="margin-top" :column="3" :size="size" border>-->
    <!--              &lt;!&ndash;        <template slot="extra">&ndash;&gt;-->
    <!--              &lt;!&ndash;          <el-button type="primary" size="small">操作</el-button>&ndash;&gt;-->
    <!--              &lt;!&ndash;        </template>&ndash;&gt;-->
    <!--&lt;!&ndash;              <el-form-item prop="id"></el-form-item>&ndash;&gt;-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  员工编号-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.id}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.id" style="width: 100%" disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--                <el-descriptions-item>-->
    <!--                  <template slot="label">-->
    <!--                    <i class="el-icon-user"></i>-->
    <!--                    密码-->
    <!--                  </template>-->
    <!--                  &lt;!&ndash;            {{// this.emp.password}}&ndash;&gt;-->
    <!--                  <el-input v-model="emp.password" style="width: 100%":disabled="disabled"></el-input>-->
    <!--                </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  员工类型-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.employee_type}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.employee_type" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  姓名-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.name}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.name" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  出生日期-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.birthday}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.birthday" style="width: 100%"></el-input>&ndash;&gt;-->
    <!--                <el-date-picker-->
    <!--                    v-model="emp.birthday"-->
    <!--                    size="mini"-->
    <!--                    type="date"-->
    <!--                    value-format="yyyy-MM-dd"-->
    <!--                    style="width: 150px;"-->
    <!--                    placeholder="出生日期"-->
    <!--                    disabled>-->
    <!--                </el-date-picker>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  身份证号-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.id_card}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.id_card" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  政治面貌-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.political}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.political" style="width: 100%"></el-input>&ndash;&gt;-->
    <!--                <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;">-->
    <!--                  <el-option-->
    <!--                      v-for="item in politicals"-->
    <!--                      :key="item"-->
    <!--                      :label="item"-->
    <!--                      :value="item"-->
    <!--                      :disabled="disabled">-->
    <!--                  </el-option>-->
    <!--                </el-select>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  民族-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.nation}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.nation" style="width: 100%"></el-input>&ndash;&gt;-->
    <!--                <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;">-->
    <!--                  <el-option-->
    <!--                      v-for="item in nations"-->
    <!--                      :key="item"-->
    <!--                      :label="item"-->
    <!--                      :value="item"-->
    <!--                      :disabled="disabled">-->
    <!--                  </el-option>-->
    <!--                </el-select>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  籍贯-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.native_place}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.native_place" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  电话-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.phone}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.phone" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  电子邮箱-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.email}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.email" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  身高-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.height}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.height" style="width: 100%" type="number":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  血型-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.blood_type}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.blood_type" style="width: 100%"></el-input>&ndash;&gt;-->
    <!--                <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;":disabled="disabled">-->
    <!--                  <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">-->
    <!--                  </el-option>-->
    <!--                </el-select>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  婚姻状况-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.marital_status}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.marital_status" style="width: 100%"></el-input>&ndash;&gt;-->
    <!--                <el-radio-group v-model="emp.marital_status":disabled="disabled">-->
    <!--                  <el-radio label="已婚">已婚</el-radio>-->
    <!--                  <el-radio label="未婚">未婚</el-radio>-->
    <!--                  <el-radio label="离异">离异</el-radio>-->
    <!--                </el-radio-group>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  出生地-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.birthplace}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.birthplace" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  户口所在地-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.registered_residence}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.registered_residence" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  部门编号-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.department_number}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.department_name" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  部门名称-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.department_name}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.department_name" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  岗位编号-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.post_number}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.post_number" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  岗位名称-->
    <!--                </template>-->
    <!--                &lt;!&ndash;              {{this.emp.post_name}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.post_name" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  入职日期-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.entry_date}}&ndash;&gt;-->
    <!--                &lt;!&ndash;            <el-input v-model="emp.entry_date" style="width: 100%" disabled></el-input>&ndash;&gt;-->
    <!--                <el-date-picker-->
    <!--                    v-model="emp.entry_date"-->
    <!--                    size="mini"-->
    <!--                    type="date"-->
    <!--                    value-format="yyyy-MM-dd"-->
    <!--                    style="width: 150px;"-->
    <!--                    placeholder="入职日期" disabled>-->
    <!--                </el-date-picker>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  用工形式-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.employment_form}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.employment_form" style="width: 100%" disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  人员来源-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.personnel_source}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.personnel_source" style="width: 100%"disabled></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  最高学历-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.highest_education}}&ndash;&gt;-->
    <!--                <el-select v-model="emp.highest_education" placeholder="最高学历" size="mini"-->
    <!--                           style="width: 150px;">-->
    <!--                  <el-option-->
    <!--                      v-for="item in highest_educations"-->
    <!--                      :key="item"-->
    <!--                      :label="item"-->
    <!--                      :value="item"-->
    <!--                      :disabled="disabled">-->
    <!--                  </el-option>-->
    <!--                </el-select>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  毕业院校-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.graduation_school}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.graduation_school" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  所学专业-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.major_studied}}&ndash;&gt;-->
    <!--                <el-input v-model="emp.major_studied" style="width: 100%":disabled="disabled"></el-input>-->
    <!--              </el-descriptions-item>-->
    <!--              <el-descriptions-item>-->
    <!--                <template slot="label">-->
    <!--                  <i class="el-icon-user"></i>-->
    <!--                  毕业日期-->
    <!--                </template>-->
    <!--                &lt;!&ndash;            {{this.emp.graduation_date}}&ndash;&gt;-->
    <!--                <el-date-picker-->
    <!--                    v-model="emp.graduation_date"-->
    <!--                    size="mini"-->
    <!--                    type="date"-->
    <!--                    value-format="yyyy-MM-dd"-->
    <!--                    style="width: 150px;"-->
    <!--                    placeholder="毕业日期"-->
    <!--                    :disabled="disabled">-->
    <!--                </el-date-picker>-->
    <!--              </el-descriptions-item>-->
    <!--            </el-descriptions>-->
    <!--          </el-form>-->
    <div>
      <el-card>
        <el-form :model="emp" :rules="rules" ref="empForm">
          <el-row>
            <el-col :span="6">
              <el-form-item label="编号:" prop="id">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit"
                          v-model="emp.id" placeholder="编号" disabled></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="密码:" prop="password">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-edit" v-model="emp.password"
                          placeholder="请输入密码" clearable :disabled="disabled"></el-input>
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
                          placeholder="请输入员工姓名"  clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="性别:" prop="sex" >
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
                    placeholder="出生日期"
                    :disabled="disabled">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="身份证号码:" prop="id_card">
                <el-input size="mini" style="width: 180px" prefix-icon="el-icon-edit"
                          v-model="emp.id_card" placeholder="请输入身份证号码" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="政治面貌:" prop="political">
                <el-select v-model="emp.political" placeholder="政治面貌" size="mini" style="width: 200px;" :disabled="disabled">
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
                <el-select v-model="emp.nation" placeholder="民族" size="mini" style="width: 150px;" filterable :disabled="disabled">
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
                          v-model="emp.native_place" placeholder="请输入籍贯" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电话号码:" prop="phone">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-phone"
                          v-model="emp.phone" placeholder="电话号码" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="电子邮箱:" prop="email">
                <el-input size="mini" style="width: 150px" prefix-icon="el-icon-message"
                          v-model="emp.email" placeholder="请输入电子邮箱" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="身高:" prop="height">
                <el-input size="mini" style="width: 100px" type="number" prefix-icon="el-icon-edit"
                          v-model="emp.height" placeholder="请输入身高" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="血型:" prop="blood_type">
                <el-select v-model="emp.blood_type" placeholder="血型" size="mini" style="width: 150px;" :disabled="disabled">
                  <el-option v-for="item in blood_types" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="婚姻状况:" prop="marital_status">
                <el-radio-group v-model="emp.marital_status" :disabled="disabled">
                  <el-radio label="已婚">已婚</el-radio>
                  <el-radio label="未婚">未婚</el-radio>
                  <el-radio label="离异">离异</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="出生地:" prop="birthplace">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                          v-model="emp.birthplace" placeholder="请输入出生地" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="6">
              <el-form-item label="户口所在地:" prop="registered_residence">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                          v-model="emp.registered_residence" placeholder="请输入户口所在地" clearable :disabled="disabled"></el-input>
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
                    placeholder="入职日期"
                    disabled>
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
                           style="width: 150px;" :disabled="disabled">
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
                          v-model="emp.graduation_school" placeholder="毕业院校名称" clearable :disabled="disabled"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="所学专业:" prop="major_studied">
                <el-input size="mini" style="width: 200px" prefix-icon="el-icon-edit"
                          v-model="emp.major_studied" placeholder="请输入专业名称" clearable :disabled="disabled"></el-input>
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
                    placeholder="毕业日期"
                    :disabled="disabled">
                </el-date-picker>

              </el-form-item>
            </el-col>
          </el-row>
          <!--            <el-row>-->
          <!--              <el-col :span="6">-->
          <!--                <el-form-item label="是否离职" prop="is_quit">-->
          <!--                  &lt;!&ndash;                <el-input size="mini" style="width: 100px" prefix-icon="el-icon-edit"&ndash;&gt;-->
          <!--                  &lt;!&ndash;                          v-model="emp.is_quit" placeholder="是否离职" clearable></el-input>&ndash;&gt;-->
          <!--                  <el-radio-group v-model="emp.is_quit">-->
          <!--                    <el-radio label="是">是</el-radio>-->
          <!--                    <el-radio label="否">否</el-radio>-->
          <!--                  </el-radio-group>-->
          <!--                </el-form-item>-->
          <!--              </el-col>-->
          <!--            </el-row>-->
        </el-form>
        <el-col style="margin-left: 70%">
          <el-button type="primary" size="medium" @click="disabled=false; btnDisabled=!btnDisabled" :disabled="btnDisabled">
            修改
          </el-button>
          <el-button type="success" size="medium" @click="disabled=true;dialogTipVisible=true" :disabled="!btnDisabled">
            提交
          </el-button>
          <el-dialog
              title="提示"
              :visible.sync="dialogTipVisible"
              width="30%">
            <span>是否提交修改？</span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="cancelEdit">取 消</el-button>
                <el-button type="primary" @click="edit">确 定</el-button>
              </span>
          </el-dialog>
        </el-col>
      </el-card>
    </div>

    <!--        </el-card>-->

  </div>
  <!--    </div>-->
</template>

<script>
import {Message} from "element-ui";

export default {
  name: "HrInfo",
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
      hr: null,
      hr2: null,
      dialogVisible: false,
      dialogTipVisible : false,
      btnDisabled : false,
      passwdDialogVisible: false,
      size : '',
      disabled : true,
      employee_types:['总裁', '主管', '员工'],
      post_names : {
        '开发部' : ['C++开发', 'Java开发', 'C#开发', 'Python开发', 'Go开发'],
        '运维部' : ['云运维', '服务器运维'],
        '测试部' : ['系统测试', 'Bug测试'],
        '设计部' : ['UI设计', '动画设计'],
        '策划部' : ['策划', '系统策划'],
      },
      post_type_options : [],
      employee_forms:['实习生', '正式职工'],
      department_names : ['开发部', '运维部', '测试部', '设计部', '策划部'],
      personnel_sources:['校招', '社招'],
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
      blood_types:['A型', 'B型', 'AB型', 'O型', '其他'],
      highest_educations: ['小学', '初中', '中专/高中', '专科', '本科', '硕士', '博士', '其他'],
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
      }
    }
  },
  mounted() {
    this.initHr();
  },
  methods: {
    onSuccess() {
      this.initHr();
    },
    showUpdateHrInfoView() {
      this.dialogVisible = true;
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.emp.id = localStorage.getItem("id")
          this.$axios.post('/DirectorDepartment', this.emp).then(resp => {
            if (resp) {
              this.emp = resp.data;
            }
          })
          // this.putRequest("/hr/pass", this.ruleForm).then(resp => {
          //     if (resp) {
          //         this.getRequest("/logout");
          //         window.sessionStorage.removeItem("user")
          //         this.$store.commit('initRoutes', []);
          //         this.$router.replace("/");
          //     }
          // })
        } else {
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    showUpdatePasswdView() {
      this.passwdDialogVisible = true;
    },
    initHr() {
      console.log("我执行了inithr")
      this.emp.id = localStorage.getItem("id")
      this.$axios.post('/DirectorDepartment', this.emp).then(resp => {
        if (resp) {
          console.log("我是修改后的data")
          console.log(resp.data)
          this.emp = resp.data;
        }
      })
    },
    edit() {
      this.dialogTipVisible = false;
      this.btnDisabled=!this.btnDisabled;
      console.log("我执行了")
      this.$axios.post('/Edit', this.emp).then(resp => {
        if (resp) {
          Message.success({message : resp.data.msg});
        } else {
          Message.error({message : resp.data.msg});
        }
      })
      // this.initHr();
    },
    cancelEdit() {
      this.dialogTipVisible = false;
      this.btnDisabled = !this.btnDisabled;
    },
    changeEditSelect() {
      // 清空部门内容
      this.emp.post_name = ''

      // 遍历部门的下拉选项数组
      for (const k in this.department_names) {
        // 岗位名称 是否等于 部门名称的下拉选择数组中的某一项
        if (this.emp.department_name === this.department_names[k]) {
          this.post_type_options = this.post_names[this.emp.department_name]
        }
      }
      console.log("post_type_options = "+this.post_type_options);
    },
  }
}
</script>

<style scoped>

</style>
