package controllers

import (
	"github.com/astaxie/beego"
)

type ResourceController struct {
	beego.Controller
}

func (c *ResourceController) Get() {
	id := c.Ctx.Input.Param(":id")
	files, err := zimuzuCtx.GetResource(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"id":   id,
			"data": map[string]interface{}{
				"Episodes": files,
			},
		}
	}

	c.ServeJSON()
}
