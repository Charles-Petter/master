<template>
    <div>
        <el-container>

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
              var currentId = localStorage.getItem("id");
              console.log("currentId = ", localStorage.getItem("id"));
              console.log("type = ", localStorage.getItem("type"));
              return currentId;
            },
            userRole() {
              var result = '';
              console.log("登陆的员工角色：", role);
              if (role === '主管') {
                result = "主管";
              } else {
                result = "员工";
              }
              return localStorage.getItem("role");
            }
        },
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