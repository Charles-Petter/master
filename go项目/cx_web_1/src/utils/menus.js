// import {getRequest, postRequest} from "./api";
//
// export const initMenu = (router, store) => {
//     if (store.state.routes.length > 0) {
//         return;
//     }
//     postRequest("/AllMenu").then(data => {
//         console.log("getRequest.data = ", data);
//         if (data) {
//
//             let fmtRoutes = formatRoutes(data);
//             router.addRoutes(fmtRoutes);
//             store.commit('initRoutes', fmtRoutes);
//             store.dispatch('connect');
//         }
//     })
// }
// //表单路由
// export const formatRoutes = (routes) => {
//     let fmRoutes = [];
//     routes.forEach(router => {
//         let {
//             path,
//             component,
//             name,
//             meta,
//             iconCls,
//             children
//         } = router;
//         if (children && children instanceof Array) {
//             children = formatRoutes(children);
//         }
//         let fmRouter = {
//             path: path,
//             name: name,
//             iconCls: iconCls,
//             meta: meta,
//             children: children,
//             //根据组件名拼接
//             component(resolve) {
//                 if (component.startsWith("Home")) {
//                     require(['../views/' + component + '.vue'], resolve);
//                 } else if (component.startsWith("Emp")) {
//                     require(['../views/emp/' + component + '.vue'], resolve);
//                 }  else if (component.startsWith("Sal")) {
//                     require(['../views/sal/' + component + '.vue'], resolve);
//                 } else if (component.startsWith("Sta")) {
//                     require(['../views/sta/' + component + '.vue'], resolve);
//                 } else if (component.startsWith("Sys")) {
//                     require(['../views/sys/' + component + '.vue'], resolve);
//                 }
//             }
//         }
//         fmRoutes.push(fmRouter);
//     })
//     return fmRoutes;
// }