import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/axios'

import JsonExcel from "vue-json-excel";
Vue.component('downloadExcel', JsonExcel)
import ElementUI from 'element-ui';
Vue.use(ElementUI);

Vue.prototype.$ELEMENT = {size: 'small', zIndex: 3000};


import {postRequest} from "./utils/api";
import {postKeyValueRequest} from "./utils/api";
import {putRequest} from "./utils/api";
import {deleteRequest} from "./utils/api";
import {getRequest} from "./utils/api";
import {initMenu} from "./utils/menus";
import 'font-awesome/css/font-awesome.min.css'
import permission from "@/store/modules/permission";
import getters from "@/store/getters";

Vue.prototype.postRequest = postRequest;
Vue.prototype.postKeyValueRequest = postKeyValueRequest;
Vue.prototype.putRequest = putRequest;
Vue.prototype.deleteRequest = deleteRequest;
Vue.prototype.getRequest = getRequest;

Vue.config.productionTip = false


router.beforeEach(async (to, from, next) => {
    // if (to.path == '/') {
    //     next();
    // } else {
    //     if (window.sessionStorage.getItem("user")) {
    //         console.log("准备执行initMenu")
    //         initMenu(router, store);
    //         next();
    //     } else {
    //         next('/?redirect=' + to.path);
    //     }
    // }
    if (to.path !== '/') {
        const type = localStorage.getItem('type');
        // alert(type);
        let flag = 0;
        console.log("if执行前的type和flag：", type, flag)

        if ((type == '主管' || type == '员工') && flag == 0) {
            console.log(11111);

            const accessRoutes = await store.dispatch('permission/generateRoutes', ["id"]);
            // router.addRoutes(accessRoutes)
            router.options.routes = accessRoutes
            router.addRoutes(accessRoutes);
            flag++;
            console.log("FLAG = ", flag);
            // next();
            next({ ...to, replace: false });
            localStorage.setItem('type',null);
        } else {
            // if (whiteList.indexOf(to.path) !== -1) {
            //   // in the free login whitelist, go directly
            //   next();
            // } else {
            //   // other pages that do not have permission to access are redirected to the login page.
            //   next(`/login?redirect=${to.path}`);
            // }
            console.log("执行next");
            next();
        }
    } else {
        next();
    }

})

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
