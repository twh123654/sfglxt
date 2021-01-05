package main

import (
	"sfglxt/app/router"
	model "sfglxt/app/model"
)


func main() {
	model.InitTable()
	r:= router.InitRouter()
	r.Run(":8080")
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context){
	// 	// c.JSON(200, gin.H{
	// 	// 	"Blog":"www.flysnow.org",
	// 	// 	"wechat":"flysnow_org",
	// 	// })
	// 	// id := c.Param("id")
	// 	// c.String(200, "this id is %s", id)

	// 	param := c.Query("wechat")
	// 	c.String(200, "the wechat is %s", param)
	// })
	// r.Run(":8080")
}