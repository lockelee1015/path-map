package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {

	router := gin.Default()

	store := sessions.NewCookieStore([]byte("path-map"))

	router.Use(sessions.Sessions("path-map-session", store))

	router.Use(CORSMiddleware())

	/**
	获取属于我的项目列表
	 */
	router.GET("/projects")

	/**
	获取项目中的成员
	 */
	router.GET("/project/members")

	/**
	获取项目路线图
	 */
	router.GET("/map")

	/**
	获取项目测试用例
	 */
	router.GET("/case")

	/**
	保存项目路线图
	 */
	router.POST("/map")

	/**
	执行测试
	 */
	router.POST("/test")

	/**
	逐条保存测试用例结果
	 */
	router.POST("/test/case/save")

	router.Run()
}
