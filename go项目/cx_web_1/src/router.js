import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from './views/Login.vue'
import Home from './views/Home.vue'
import HrInfo from './views/HrInfo.vue'
import EmployeeBasic from "@/views/emp/EmployeeBasic";
import PostTransfer from "@/views/transfer/PostTransfer";
import DepartmentTransfer from "@/views/transfer/DepartmentTransfer";
import PostExamine from "@/views/examine/PostExamine";
import ExpelEmployee from "@/views/manage/ExpelEmployee";
import QuitEmployee from "@/views/manage/QuitEmployee";
import ApplyQuit from "@/views/manage/ApplyQuit";
import ApplyPost from "@/views/manage/ApplyPost";
import PostTalent from "@/views/talent/PostTalent";
import DepartmentBasic from "@/views/manage/DepartmentBasic";

Vue.use(VueRouter)
//通用路由表
export const startLogin = [
    {
        path : '/',
        name : 'Login',
        component : Login,
        hidden : true,
    },
]



//主管登录
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
            // {
            //     path : '/employee/advance',
            //     name : '高级资料',
            //     component : EmployeeAdvance,
            // }
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
            //**********加入信息录入组件
            // {
            //     path : '/department/basic',
            //     name : '员工工资统计',
            //     component : SalaryBasic_cx
            // },
            {
                path : '/manage/post',
                name : '申请新岗位',
                component : ApplyPost,
            },
            {
                path : '/manage/apply',
                name : '申请离职',
                component : ApplyQuit,
            },
            {
                path : '/manage/quit',
                name : '离职员工',
                component : QuitEmployee,
            },
        ]
    },
    {
        path : '/home',
        name : '申请审核',
        component : Home,
        redirect : '/examine/post',
        children : [
            {
                path : '/examine/post',
                name : '查看岗位申请',
                component : PostExamine,
            },
            {
                path : '/talent/post',
                name : '查看人才库申请',
                component : PostTalent,
            },
        ]
    },
    {
        path: '/home',
        name: '人事管理',
        component: Home,
        redirect: '/hrinfo',
        hidden: true,
        children: [
            {
                path: '/hrinfo',
                name: '个人信息',
                component: HrInfo,
                hidden: true,
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
        // hidden : true,
        // meta : ['admin', 'user'],
        redirect : '/employee/basic',
        children : [
            {
                path : '/employee/basic',
                name : '基本资料',
                component : EmployeeBasic,
            },
            // {
            //     path : '/employee/advance',
            //     name : '高级资料',
            //     component : EmployeeAdvance,
            // }
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
            {
                path : '/manage/post',
                name : '申请新岗位',
                component : ApplyPost,
            },
            {
                path : '/manage/apply',
                name : '申请离职',
                component : ApplyQuit,
            },
        ]
    },
    {
        path: '/home',
        name: '人事管理',
        component: Home,
        redirect: '/hrinfo',
        hidden: true,
        children: [
            {
                path: '/hrinfo',
                name: '个人信息',
                component: HrInfo,
                hidden: true,
            },
        ]
    },
]

// export const hrinfo = [
//     {
//         path : '/hrinfo',
//         name : 'information',
//         component : HrInfo,
//         hidden : true,
//     },
// ]


const router = new VueRouter({
    mode : 'hash',
    routes : startLogin
})

export default router


