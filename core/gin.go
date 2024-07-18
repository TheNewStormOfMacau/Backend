package core

import (
	v1 "backend/api/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Gin() {
	routers := Routers()
	err := routers.Run(":8888")
	if err != nil {
		panic(err)
	}
}

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))
	router := r.Group("/api")
	{
		router.GET("/v1/getuserinfo", v1.GetUserInfo)
		router.GET("/v1/getrewardinfo", v1.GetRewardInfo)
	}
	return r
}
