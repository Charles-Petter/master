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

const routes = [
  { name: "login", path: "/login", meta: { title: "login" }, component: () => import("../views/login/index"), hidden: true },
//   { name: "/", path: "/", component: Layout, redirect: "/index", hidden: false, children: [
//       { name: "index", path: "/index", meta: { title: "index" }, component: () => import("../views/index/index") },
//   ]},



//   { name: "/form", path: "/form", component: Layout, redirect: "/form/index", hidden: false, children: [
//       { name: "/form/index", path: "/form/index", meta: { title: "form" }, component: () => import("../views/form/index") }
//   ]},
//   { name: "/example", path: "/example", component: Layout, redirect: "/example/tree", meta: { title: "example" }, hidden: false, children: [
// 	  { name: "/tree", path: "/example/tree", meta: { title: "tree" }, component: () => import("../views/tree/index") },
// 	  { name: "/copy", path: "/example/copy", meta: { title: "copy" }, component: () => import("../views/tree/copy") }
//   ] },
//   { name: "/table", path: "/table", component: Layout, redirect: "/table/index", hidden: false, children: [
// 	  { name: "/table/index", path: "/table/index", meta: { title: "table" }, component: () => import("../views/table/index") }
//   ] },
//   { name: "/admin", path: "/admin", component: Layout, redirect: "/admin/index", hidden: false, children: [
// 	{ name: "/admin/index", path: "/admin/index", meta: { title: "admin" }, component: () => import("../views/admin/index") }
// ] },
// { name: "/people", path: "/people", component: Layout, redirect: "/people/index", hidden: false, children: [
// 	{ name: "/people/index", path: "/people/index", meta: { title: "people" }, component: () => import("../views/people/index") }
// ] }
]

const router = new VueRouter({ routes })

export default router
