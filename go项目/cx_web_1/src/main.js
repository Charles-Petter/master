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

//注册一个全局前置守卫
//在路由跳转之前做判断，符合条件就放行
router.beforeEach(async (to, from, next) => {
    if (to.path !== '/') {
        const type = localStorage.getItem('type');
        // alert(type);
        let flag = 0;
        console.log("if执行前的type和flag：", type, flag)

        if (( type === '主管' || type === '员工') && flag === 0) {
            console.log(11111);

            const accessRoutes = await store.dispatch('permission/generateRoutes', ["id"]);
            router.options.routes = accessRoutes;
            router.addRoutes(accessRoutes);
            flag++;
            console.log("FLAG = ", flag);
            next({ ...to, replace: false });
            localStorage.setItem('type',null);
        } else {
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
