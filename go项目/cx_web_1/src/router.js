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

export const startLogin = [
    {
        path : '/',
        name : 'Login',
        component : Login,
        hidden : true,
    },
]

export const bossLogin = [
    {
        path : '/home',
        name : '员工资料',
        component : Home,
        // hidden : true,
        // meta : ['admin', 'user'],
        meta : {title : '员工资料', icon : 'EmployeeBasic.svg', affix : true},
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
                name : '部门信息',
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
            {
                path : '/manage/quit',
                name : '离职员工',
                component : QuitEmployee,
            },
            {
                path : '/manage/expel',
                name : '开除员工',
                component : ExpelEmployee,
            },
        ]
    },
    {
        path : '/home',
        name : '调动数据',
        component : Home,
        redirect : '/transfer/department',
        children : [
            {
                path : '/transfer/department',
                name : '部门调动',
                component : DepartmentTransfer,
            },
            {
                path : '/transfer/post',
                name : '岗位调动',
                component : PostTransfer,
            }
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



// Vue.use(Router)
//
// export default new Router({
//     routes: [
//         {
//             path: '/',
//             name: 'Login',
//             component: Login,
//             hidden: true
//         },
//         {
//             path: '/home',
//             name: 'Home',
//             component: Home,
//             hidden: true,
//             meta: {
//                 roles: ['admin', 'user']
//             },
//             children: [
//                 {
//                     path: '/chat',
//                     name: '在线聊天',
//                     component: FriendChat,
//                     hidden: true
//                 }, {
//                     path: '/hrinfo',
//                     name: '个人中心',
//                     component: HrInfo,
//                     hidden: true
//                 }
//             ]
//         },
//         {
//             path: '*',
//             redirect: '/home'
//         }
//     ]
// })
