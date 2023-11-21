package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqBody struct{
	Username 	string 	`json:"username"`
	Password 	string	`json:"password"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/get",func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"Hey there we are testing things out")
	})

	r.POST("/post",func(ctx *gin.Context) {
		var reqBody ReqBody
		if err:= ctx.BindJSON(&reqBody);err!=nil{
			ctx.Abort()
		}
		ctx.String(http.StatusOK,fmt.Sprintf("%s %s",reqBody.Username,reqBody.Password))
	})

	r.Run(":3001")
}