package v1

import (
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	addr, success := c.GetQuery("chain_addr")
	if success != true {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
	}

	var user *model.User
	user = service.GetUserByAddr(addr)
	if user == nil {
		c.JSON(202, gin.H{
			"message": "user not found",
		})
	} else {
		_ = c.ShouldBind(&user)
	}
}

func GetRewardInfo(c *gin.Context) {
	addr, success := c.GetQuery("chain_addr")
	if success != true {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
	}

	var records *[]model.Record
	records = service.GetRecordsByAddr(addr)
	if records == nil {
		c.JSON(202, gin.H{
			"message": "user not found",
		})
	} else {
		c.JSON(http.StatusOK, records)
	}
}
