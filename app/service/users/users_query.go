package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Query(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "用户列表", 
	})
}