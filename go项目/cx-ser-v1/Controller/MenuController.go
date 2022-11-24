package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"cx/Global"
	"cx/Model"
	"net/http"
)

func AllMenu(context *gin.Context)  {
	fmt.Println("进入AllMenu")
	var menu []Model.Menu
	var err error
	//err = Global.Db.Raw("select m1.id as id1,m1.name as name1,m2.id as id2,m2.name as name2,m3.id as id3,m3.name as name3 from menus m1,menus m2,menus m3 where m1.id=m2.parent_id and m2.id=m3.parent_id and m3.enabled=1 order by m1.id,m2.id,m3.id").Scan(&menu).Error
	//err = Global.Db.Raw("select m.*,r.id as rid,r.name as rname,r.name_zh as rname_zh from menus m,menu_roles mr,roles r where m.id=mr.mid and mr.rid=r.id order by m.id").Scan(&menu).Error
	err = Global.Db.Raw("select distinct m1.*,m2.id as id2,m2.component as component2,m2.enabled as enabled2,m2.icon_cls as icon_cls2,m2.keep_alive as keep_alive2,m2.name as name2,m2.parent_id as parent_id2,m2.require_auth as require_auth2,m2.path as path2 from menus m1,menus m2,hr_roles hrr,menu_roles mr where m1.id=m2.parent_id and hrr.hrid = ? and hrr.rid=mr.rid and mr.mid=m2.id and m2.enabled=1 order by m1.id,m2.id", Global.Id).Scan(&menu).Error
	if err != nil {
		fmt.Println("查询菜单栏出错")
		return
	}
	//select m1.`id` as id1,m1.`name` as name1,m2.`id` as id2,m2.`name` as name2,m3.`id` as id3,m3.`name` as name3 from menu m1,menu m2,menu m3 where m1.`id`=m2.`parentId` and m2.`id`=m3.`parentId` and m3.`enabled`=true order by m1.`id`,m2.`id`,m3.`id`
	//select m.*,r.`id` as rid,r.`name` as rname,r.`nameZh` as rnameZh from menu m,menu_role mr,role r where m.`id`=mr.`mid` and mr.`rid`=r.`id` order by m.`id`
	//select distinct m1.*,m2.`id` as id2,m2.`component` as component2,m2.`enabled` as enabled2,m2.`iconCls` as iconCls2,m2.`keepAlive` as keepAlive2,m2.`name` as name2,m2.`parentId` as parentId2,m2.`requireAuth` as requireAuth2,m2.`path` as path2 from menu m1,menu m2,hr_role hrr,menu_role mr where m1.`id`=m2.`parentId` and hrr.`hrid`=#{hrid} and hrr.`rid`=mr.`rid` and mr.`mid`=m2.`id` and m2.`enabled`=true order by m1.`id`,m2.`id`
	fmt.Println("菜单栏：", menu)
	context.JSON(http.StatusOK, menu)
}
