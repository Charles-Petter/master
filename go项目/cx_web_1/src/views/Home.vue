<template>
    <div>
        <el-container>
            <el-header class="homeHeader">
                <div class="title">人事管理系统</div>
                <div>
<!--                    <el-button icon="el-icon-bell" type="text" style="margin-right: 8px;color: #000000;" size="normal" @click="goChat"></el-button>-->
                    <el-dropdown class="userInfo" @command="commandHandler">
  <span class="el-dropdown-link">
<!--     <el-image :src="" class="eImage"></el-image>-->
    <span  :style="img_cx"> </span>
    <b>编号:{{userId}}<br />角色:{{userRole}}</b>
    <i></i>
  </span>

                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item command="userinfo">个人中心</el-dropdown-item>
<!--                            <el-dropdown-item command="setting">设置</el-dropdown-item>-->
                            <el-dropdown-item command="logout" divided>注销登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </div>
            </el-header>
            <el-container>
                <el-aside width="200px" hight="100%">
                    <el-menu
                        mode="vertical"
                        unique-opened
                        :default-active="$route.path"
                        background-color="#304156"
                        text-color="#fff"
                        active-text-color="#409EFF"
                        router unique-opened>
                        <el-submenu :index="index+''" v-for="(item,index) in $router.options.routes" v-if="!item.hidden" :key="index">
                            <template slot="title">
                                <i style="color: #162aef;margin-right: 5px" :class="item.iconCls"></i>

                                <span class="el-menu-font">{{item.name}}</span>
                            </template>
                            <el-menu-item :index="child.path" v-for="(child,indexj) in item.children" :key="indexj">
                              <span class="el-submenu-font">{{child.name}}</span>
                            </el-menu-item>
                        </el-submenu>
                    </el-menu>
                </el-aside>
                <el-main>
                    <el-breadcrumb separator-class="el-icon-arrow-right" v-if="this.$router.currentRoute.path!='/home'">
                        <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
                        <el-breadcrumb-item>{{this.$router.currentRoute.name}}</el-breadcrumb-item>
                    </el-breadcrumb>
                    <div class="homeWelcome" v-if="this.$router.currentRoute.path=='/home'">
                      人事管理系统！
                    </div>
                    <router-view class="homeRouterView"/>
                </el-main>
            </el-container>
        </el-container>
    </div>

</template>

<script>
    export default {
        name: "Home",
        data() {
            return {
                // user: JSON.parse(window.sessionStorage.getItem("user"))
              img_cx:{
                Image: "url(" + require("../assets/loginBackgroundImg.webp") + ")",//背景圖片

              },
            }

        },
        computed: {
            routes() {
              console.log("查找的routes = ", this.$store.state.routes)
                return this.$store.state.routes;
            },
            user() {
                return this.$store.state.currentHr;
            },
            userId() {
              // var currentId = localStorage.getItem('id');
              var currentId = localStorage.getItem("id");
              console.log("currentId = ", localStorage.getItem("id"));
              console.log("type = ", localStorage.getItem("type"));
              return currentId;
            },
            userRole() {
              var result = '';
              // console.log("登陆的员工角色：", role);
              // if (role === '总裁') {
              //   result = "总裁";
              // } else if (role === '主管') {
              //   result = "主管";
              // } else {
              //   result = "员工";
              // }
              return localStorage.getItem("role");
            }
        },
        methods: {
            goChat() {
                this.$router.push("/chat");
            },
            commandHandler(cmd) {
                if (cmd == 'logout') {
                    this.$confirm('此操作将注销登录, 是否继续?', '提示', {
                        confirmButtonText: '确定',
                        cancelButtonText: '取消',
                        type: 'warning'
                    }).then(() => {
                        // this.getRequest("/logout");
                        window.sessionStorage.removeItem("user")
                        this.$store.commit('initRoutes', []);
                        this.$router.replace("/");
                    }).catch(() => {
                        this.$message({
                            type: 'info',
                            message: '已取消操作'
                        });
                    });
                }else if (cmd == 'userinfo') {
                    this.$router.push('/hrinfo');
                }
            },
        }
    }
</script>

<style>
.sidebar-container {
  transition: width 0.28s;
  width: 180px !important;
  height: 100%;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  z-index: 1001;
}
.sidebar-container a {
  display: inline-block;
  width: 100%;
}
.sidebar-container .svg-icon {
  margin-right: 16px;
}
.sidebar-container .el-menu {
  border: none;
  width: 100%;
}
    .homeRouterView {
        margin-top: 10px;
    }

    .homeWelcome {
        text-align: center;
        font-size: 30px;
        font-family: 宋体;
        color: white;
        padding-top: 50px;
    }

    .homeHeader {
        background-color: #717a80;
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 0px 15px;
        box-sizing: border-box;
    }

    .homeHeader .title {
        font-size: 40px;
        font-family: 宋体;
        color: #f0f639
    }

    .homeHeader .userInfo {
        cursor: pointer;
    }

    .el-dropdown-link img {
        width: 48px;
        height: 48px;
        border-radius: 24px;
        margin-left: 8px;
    }

    .el-dropdown-link {
        display: flex;
        align-items: center;
    }
    .el-menu-font {
      font-family: 宋体;
    }
    .el-submenu-font {
      font-family: 宋体;
    }
.nest-menu .el-submenu > .el-submenu__title,.el-submenu .el-menu-item {min-width: 180px !important;background-color: #1f2d3d !important;}
.nest-menu .el-submenu > .el-submenu__title,.el-submenu .el-menu-item :hover {background-color: #001528 !important;}
.el-menu--collapse .el-menu .el-submenu {min-width: 180px !important;}
</style>