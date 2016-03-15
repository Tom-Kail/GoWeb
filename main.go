package main

import (
	_ "WebScan/controllers"
	_ "WebScan/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	//链接数据库
	beego.Info("Connecting database...")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@tcp(192.168.56.128:3306)/test?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "root:@/webscan?charset=utf8")
	beego.Info("Connected")
}
func main() {

	beego.Run()
}
