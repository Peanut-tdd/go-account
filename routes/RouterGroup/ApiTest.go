package RouterGroup

import (
	"account_check/app/Http/Middlewares"
	"account_check/app/Http/controller"
	"account_check/app/http/controller/bill"
	"account_check/app/model"
	"account_check/bootstrap/driver"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiTest(route *gin.Engine) {
	api := route.Group("/api/", Middlewares.HttpCorsApi)

	gen := api.Group("/gen1/")
	gen.GET("1", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 200,
			"msg":   "gin1",
			"content": map[string]interface{}{
				"time": "111",
			},
		})
	})

	gen.GET("test", controller.Zzzz)

	gen.GET("wx", bill.WxBill)

	//form-data
	gen.POST("form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "1")
		username := c.PostForm("username")
		password := c.PostForm("password")

		fmt.Println(types, username, password)

	})

	gen.POST("json", func(c *gin.Context) {
		data,_:=c.GetRawData()

		var body map[string]interface{}
		_ = json.Unmarshal(data, &body)

		//获取json中的key，注意使用["key"]获取


		fmt.Println(body)

		//fmt.Println(string(data))
		//bodyByts, err := ioutil.ReadAll(c.Request.Body)
		//
		//if err != nil {
		//	// 返回错误信息
		//	c.String(http.StatusBadRequest, err.Error())
		//	// 执行退出
		//	c.Abort()
		//}
		//
		//
		//fmt.Println(string(bodyByts))
		//// 返回的 code 和 对应的参数星系
		//c.String(http.StatusOK, "%s \n", string(bodyByts))

	})
	//
	//utils.RedisSet("name", "aaaaa", time.Second*60)
	//value := utils.RedisGet("name")
	//
	//fmt.Println(value)
	//driver.GVA_DB.AutoMigrate(&model.UserTest{})
	//user := model.UserTest{Username: "Jinzhu"}
	//
	//driver.GVA_DB.Create(&user) // 通过数据的指针来创

	//1
	a := model.OrderBill{}
	driver.GVA_DB.Where("id=?", 1).First(&a)
	//for _, item := range a {
	//	fmt.Println(item.Username)
	//}
	fmt.Println(a)

	//2
	//result := make([]map[string]interface{},0)
	//driver.GVA_DB.Model(&model.UserTest{}).Find(&result)
	//for _, item := range result {
	//	v1, _ := json.Marshal(&item)
	//	fmt.Printf("v1:%s\n", v1)
	//}

}
