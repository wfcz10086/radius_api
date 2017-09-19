package controllers

import (
	"strings"
	"api/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type NewController struct {
        beego.Controller
}
func (this *NewController) Get() {
	text := this.Input().Get("text")
	rt1 := models.Base64_de(text)
	user:=strings.Split(rt1, "_")[0]
	pass:=strings.Split(rt1, "_")[1]
	var text1 string = "Cleartext-Password"
	var text2 string = ":="
	var text3 string = "Expiration"
	rac1 := models.Radcheck{Username: user, Attribute: text1, Op: text2, Value: pass}
	models.InsertPassword(&rac1)
	var text4 string = models.Expiration()
	rac2 := models.Radcheck{Username: user, Attribute: text3, Op: text2, Value: text4}
	models.InsertExpiration(&rac2)
	models.UpdatePassword(&rac1)
	models.UpdateExpiration(&rac2)

    	this.Data["json"] = &rac2
    	this.ServeJSON()

	
}
