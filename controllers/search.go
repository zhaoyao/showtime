package controllers

import (
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	q := c.GetString("q", "")
	ret, err := zimuzuCtx.Search(q)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 1,
			"msg":  err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"code": 0,
			"data": ret,
		}
	}

	c.ServeJSON()
}
