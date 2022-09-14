package controller

import (
	"account_check/app/model"
	"account_check/bootstrap/driver"

	"github.com/gin-gonic/gin"
)

func Zzzz(ctx *gin.Context) {
	res := []model.User{{Username: "test"}}

	//添加
	// driver.GVA_DB.Create(&res)
	//查询
	// driver.GVA_DB.First(&res)
	driver.GVA_DB.Where("id = ?", 4).Find(&res)
	// ctx.JSONP(200, res)
	ctx.JSONP(200, res)

	// ctx.JSONP(http.StatusNotFound, gin.H{
	// 	"state": 200,
	// 	"msg":   "gin1",
	// 	"content": map[string]interface{}{
	// 		"time": "111",
	// 	},
	// })
}



