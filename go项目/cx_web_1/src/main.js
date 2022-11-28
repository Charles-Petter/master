import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/axios'

import JsonExcel from "vue-json-excel";
Vue.component('downloadExcel', JsonExcel)
// import {
//     Button,
//     Input,
//     Table,
//     TableColumn,
//     Dialog,
//     Card,
//     Container,
//     Icon,
//     Select,
//     Form,
//     Tag,
//     Tree,
//     Pagination,
//     Badge,
//     Loading,
//     Message,
//     MessageBox,
//     Menu,
//     Tabs,
//     TabPane,
//     Breadcrumb,
//     BreadcrumbItem,
//     Dropdown,
//     Steps,
//     Step,
//     Tooltip,
//     Popover,
//     Collapse,
//     FormItem,
//     Checkbox,
//     Header,
//     DropdownMenu,
//     DropdownItem,
//     Aside,
//     Main,
//     MenuItem,
//     Submenu,
//     Option,
//     Col,
//     Row,
//     Upload,
//     Radio,
//     DatePicker,
//     RadioGroup,
//     CollapseItem,
//     Switch
// } from 'element-ui';
// import 'element-ui/lib/theme-chalk/index.css';

import ElementUI from 'element-ui';
Vue.use(ElementUI);

Vue.prototype.$ELEMENT = {size: 'small', zIndex: 3000};
// Vue.use(Switch);
// Vue.use(CollapseItem);
// Vue.use(Radio);
// Vue.use(RadioGroup);
// Vue.use(DatePicker);
// Vue.use(Upload);
// Vue.use(Row);
// Vue.use(Col);
// Vue.use(Option);
// Vue.use(Submenu);
// Vue.use(MenuItem);
// Vue.use(Header);
// Vue.use(DropdownMenu);
// Vue.use(DropdownItem);
// Vue.use(Aside);
// Vue.use(Main);
// Vue.use(Checkbox);
// Vue.use(FormItem);
// Vue.use(Collapse);
// Vue.use(Popover);
// Vue.use(Menu);
// Vue.use(Tabs);
// Vue.use(TabPane);
// Vue.use(Breadcrumb);
// Vue.use(BreadcrumbItem);
// Vue.use(Dropdown);
// Vue.use(Steps);
// Vue.use(Step);
// Vue.use(Tooltip);
// Vue.use(Tree);
// Vue.use(Pagination);
// Vue.use(Badge);
// Vue.use(Loading);
// Vue.use(Button);
// Vue.use(Input);
// Vue.use(Table);
// Vue.use(TableColumn);
// Vue.use(Dialog);
// Vue.use(Card);
// Vue.use(Container);
// Vue.use(Icon);
// Vue.use(Select);
// Vue.use(Form);
// Vue.use(Tag);
// Vue.prototype.$alert = MessageBox.alert
// Vue.prototype.$confirm = MessageBox.confirm

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

        if ((type == '总裁' || type == '主管' || type == '员工') && flag == 0) {
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
