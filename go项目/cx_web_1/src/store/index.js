// import Vue from 'vue'
// import Vuex from 'vuex'
// import { Notification } from 'element-ui';
// import {getRequest} from "../utils/api";
// import '../utils/stomp'
// import '../utils/sockjs'
//
// Vue.use(Vuex)
//
// const now = new Date();
//
// // const myData : {
// //     [
// //         name : '员工资料',
// //         children : [
// //             namee : '基本资料'
// //         ],
// //     ]
// //
// // },
//
// const store = new Vuex.Store({
//     state: {
//         routes: [],//？？？
//         sessions: {},
//         hrs: [],
//         currentSession: null,
//         currentHr: JSON.parse(window.sessionStorage.getItem("user")),
//         filterKey: '',
//         stomp: null,
//         isDot: {}
//     },
//     mutations: {
//         INIT_CURRENTHR(state, hr) {
//             state.currentHr = hr;
//         },
//         initRoutes(state, data) {
//             //后端根据不同的权限传来不同的data导航栏
//             console.log("router的data：", data);
//             state.routes = data;
//             // let name, path, component, children, meta, iconCls
//             // state.routes = [
//             //     name = "员工资料",
//             //     component = "Home",
//             //     path = "/home",
//             //     children = [
//             //
//             //     ],
//             // ]
//         },
//         changeCurrentSession(state, currentSession) {
//             Vue.set(state.isDot, state.currentHr.username + '#' + currentSession.username, false);
//             state.currentSession = currentSession;
//         },
//         addMessage(state, msg) {
//             let mss = state.sessions[state.currentHr.username + '#' + msg.to];
//             if (!mss) {
//                 // state.sessions[state.currentHr.username+'#'+msg.to] = [];
//                 Vue.set(state.sessions, state.currentHr.username + '#' + msg.to, []);
//             }
//             state.sessions[state.currentHr.username + '#' + msg.to].push({
//                 content: msg.content,
//                 date: new Date(),
//                 self: !msg.notSelf
//             })
//         },
//         INIT_DATA(state) {
//             //浏览器本地的历史聊天记录可以在这里完成
//             let data = localStorage.getItem('vue-chat-session');
//             if (data) {
//                 state.sessions = JSON.parse(data);
//             }
//         },
//         INIT_HR(state, data) {
//             state.hrs = data;
//         }
//     },
//     actions: {
//         connect(context) {
//             context.state.stomp = Stomp.over(new SockJS('/ws/ep'));
//             context.state.stomp.connect({}, success => {
//                 context.state.stomp.subscribe('/user/queue/chat', msg => {
//                     let receiveMsg = JSON.parse(msg.body);
//                     if (!context.state.currentSession || receiveMsg.from != context.state.currentSession.username) {
//                         Notification.info({
//                             title: '【' + receiveMsg.fromNickname + '】发来一条消息',
//                             message: receiveMsg.content.length > 10 ? receiveMsg.content.substr(0, 10) : receiveMsg.content,
//                             position: 'bottom-right'
//                         })
//                         Vue.set(context.state.isDot, context.state.currentHr.username + '#' + receiveMsg.from, true);
//                     }
//                     receiveMsg.notSelf = true;
//                     receiveMsg.to = receiveMsg.from;
//                     context.commit('addMessage', receiveMsg);
//                 })
//             }, error => {
//
//             })
//         },
//         initData(context) {
//             context.commit('INIT_DATA')
//             getRequest("/chat/hrs").then(resp => {
//                 if (resp) {
//                     context.commit('INIT_HR', resp);
//                 }
//             })
//         }
//     }
// })
//
// store.watch(function (state) {
//     return state.sessions
// }, function (val) {
//     localStorage.setItem('vue-chat-session', JSON.stringify(val));
// }, {
//     deep: true/*这个貌似是开启watch监测的判断,官方说明也比较模糊*/
// })
//
//
// export default store;


import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'

Vue.use(Vuex)

// https://webpack.js.org/guides/dependency-management/#requirecontext
const modulesFiles = require.context('./modules', true, /\.js$/)

// you do not need `import app from './modules/app'`
// it will auto require all vuex module from modules file
const modules = modulesFiles.keys().reduce((modules, modulePath) => {
    // set './app.js' => 'app'
    const moduleName = modulePath.replace(/^\.\/(.*)\.\w+$/, '$1')
    const value = modulesFiles(modulePath)
    modules[moduleName] = value.default
    return modules
}, {})

const store = new Vuex.Store({
    state: {
        currentHr: JSON.parse(window.sessionStorage.getItem("id")),
    },
    modules,
    getters
})

export default store
