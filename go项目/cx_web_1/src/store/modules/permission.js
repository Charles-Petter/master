import { startLogin,directorLogin, employeeLogin } from '@/router'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
//判断是否有权限
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes 异步路由
 * @param roles  作用
 */
//通过递归过滤异步路由表
export function filterAsyncRoutes(routes, roles) {
  const res = []
  console.log('routes + roles')
  console.log(routes)
  console.log(roles)
  routes.forEach(route => {
    const tmp = { ...route }
    if (hasPermission(roles, tmp)) {
      if (tmp.children) {
        tmp.children = filterAsyncRoutes(tmp.children, roles)
      }
      res.push(tmp)
    }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = startLogin.concat(routes)
  }
}

const actions = {
  generateRoutes({ commit }, roles) {
    return new Promise(resolve => {
      let accessedRoutes
//根据按钮类型来filter跳转不同权限的页面
      const type = localStorage.getItem('type');
     if (type === '主管'){
        accessedRoutes = filterAsyncRoutes(directorLogin, roles)//主管登录
      } else {
        accessedRoutes = filterAsyncRoutes(employeeLogin, roles)
      }
      console.log(accessedRoutes);
      console.log('accessedRoutes');
      commit('SET_ROUTES', accessedRoutes)
      resolve(accessedRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
