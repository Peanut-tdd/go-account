package bill

import (
	"github.com/gin-gonic/gin"
)

func KsBill(c *gin.Context) {
	var request = make(map[string]string)
	request["app_id"] = "ks695806146341101215"
	request["start_date"] = "20220819000000"
	request["end_date"] = "20220820000000"
	request["bill_type"] = "PAY"
	//kuaishou.GetBills(request)

	c.JSONP(200, "success")
}
