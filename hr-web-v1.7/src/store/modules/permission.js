import { startLogin,bossLogin,directorLogin, employeeLogin } from '@/router'

/**
 * Use meta.role to determine if the current user has permission
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role))
  } else {
    return true
  }
}

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
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
      // if (roles.includes('admin')) {
      //   accessedRoutes = asyncRoutes || []
      // } else {
      //   accessedRoutes = filterAsyncRoutes(adminLogin, roles)
      // }
      const type = localStorage.getItem('type');
      // alert("aaaa"+type);
      // type:1是老板 2是主管 3是员工
      if(type == '总裁'){
        accessedRoutes = filterAsyncRoutes(bossLogin, roles)
      } else if (type == '主管'){
        accessedRoutes = filterAsyncRoutes(directorLogin, roles)
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
