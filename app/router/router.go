package router

import (
	"github.com/gin-gonic/gin"
	"sfglxt/app/service/users"
)

func InitRouter() *gin.Engine{
	router := gin.Default()
	router.GET("/users/query", users.Query)
	return router;
}