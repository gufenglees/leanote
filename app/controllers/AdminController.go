package controllers

import (
	"github.com/revel/revel"
)

// admin 首页

type Admin struct {
	AdminBaseController
}

// admin 主页
func (c Admin) Index() revel.Result {
	c.SetUserInfo()
	
	c.RenderArgs["title"] = "leanote"
	c.RenderArgs["openRegister"] = openRegister
	c.SetLocale()
	
	return c.RenderTemplate("admin/index.html");
}