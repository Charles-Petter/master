import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from './views/Login.vue'
import Home from './views/Home.vue'
import EmployeeBasic_cx from "@/views/emp/EmployeeBasic_cx";
import CollectBasic_cx from "@/views/Dep/CollectBasic_cx";
import Addemploy_cx from "@/views/emp/Addemploy_cx";
import SalEmp_cx from "@/views/sal/SalEmp_cx";
import SalCount_cx from "./views/sal/SalCount_cx";


Vue.use(VueRouter)
//静态路由
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
        redirect : '/employee/basic_cx',
        children : [
            {
                path : '/employee/basic_cx',
                name : '基本资料',
                component : EmployeeBasic_cx,
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
                path : '/Dep/CollectBasic_cx',
                name : '查询部门信息',
                component : CollectBasic_cx
            },
            //员工工资路由
            {
                path : '/sal/SalCount_cx',
                name : '员工工资',
                component : SalCount_cx,
            },
        ]
    },



]
//员工登录
export const employeeLogin = [
    {
        path : '/home',
        name : '员工资料',
        component : Home,
        redirect : '/employee/basic_cx',
        children : [
            {
                path : '/employee/basic_cx',
                name : '基本资料',
                component : EmployeeBasic_cx,
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
                path : '/sal/SalEmp_cx',
                 name : '员工工资',
               component : SalEmp_cx,
             },


        ]
    },
]

const router = new VueRouter({
    mode : 'hash',
    routes : startLogin
})

export default router



