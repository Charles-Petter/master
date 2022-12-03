import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from './views/Login.vue'
import Home from './views/Home.vue'
import EmployeeBasic from "@/views/emp/EmployeeBasic";
import DepartmentBasic from "@/views/manage/DepartmentBasic";
import SalSob from "@/views/sal/SalSob";
import Addemploy_cx from "@/views/emp/Addemploy_cx";



Vue.use(VueRouter)

export const startLogin = [
    {
        path : '/',
        name : 'Login',
        component : Login,
        hidden : true,
    },
]

export const directorLogin = [
    {
        path : '/home',
        name : '员工资料',
        component : Home,
        // hidden : true,
        // meta : ['admin', 'user'],
        redirect : '/employee/basic',
        children : [
            {
                path : '/employee/basic',
                name : '基本资料',
                component : EmployeeBasic,
            },
            //增加员工信息
            {
                path : '/emp/Addemploy_cx',
                name : '添加员工信息',
                component :Addemploy_cx,
            },
        ]
    },
    {
        path : '/home',
        name : '人事管理',
        component : Home,
        redirect : '/manage/post',
        children : [
            {
                path : '/department/basic',
                name : '查询部门信息',
                component : DepartmentBasic
            },

            //员工工资路由
            {
                path : '/sal/SalSob',
                name : '员工工资',
                component : SalSob,
            },
            // {
            //     path : '/sal/SalSobCfg',
            //     name : '员工工资',
            //     component : SalSobCfg,
            // },
        ]
    },
]
//员工登录
export const employeeLogin = [
    {
        path : '/home',
        name : '员工资料',
        component : Home,
        redirect : '/employee/basic',
        children : [
            {
                path : '/employee/basic',
                name : '基本资料',
                component : EmployeeBasic,
            },
        ]
    },
    {
        path : '/home',
        name : '人事管理',
        component : Home,
        redirect : '/manage/post',
        children : [
            {
                path : '/department/basic',
                name : '查询部门信息',
                component : DepartmentBasic
            },
        ]
    },
]
const router = new VueRouter({
    mode : 'hash',
    routes : startLogin
})

export default router



