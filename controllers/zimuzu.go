package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zhaoyao/showtime/zimuzu"
)

var zimuzuCtx *zimuzu.Ctx

func init() {
	account := beego.AppConfig.String("zimuzu.account")
	pwd := beego.AppConfig.String("zimuzu.password")
	zimuzuCtx = zimuzu.New(account, pwd)
}
