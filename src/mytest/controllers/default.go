package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "hezll@qq.com"
	c.Ctx.WriteString("hello world")
	v:= c.GetSession("asta")
	if v==nil{
		c.SetSession("asta",int(1))
		c.Data["num"] = 0
	} else {
		c.SetSeesion("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	} 



}
