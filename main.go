package main

import (
	_ "api/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "api/models"
    _ "github.com/go-sql-driver/mysql"
)

func init(){
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/radius?charset=utf8", 30)

    orm.RegisterModel(new(models.Radcheck))
    orm.RunSyncdb("default", false, true)
}

func main() {
    orm.Debug = true
    //ok, err := regexp.MatchString("/topic/edit/[0-9]+", "/topic/edit/123")
    //beego.Debug(ok, err)
	beego.Run()
}
