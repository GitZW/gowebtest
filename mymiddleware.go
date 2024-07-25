package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func middleware1(c *gin.Context) {
	//中间件逻辑
	c.Request.Header.Set("abc", "123")
	fmt.Println("start middleware1")
	c.Next()
	fmt.Println("end middleware1")

}

func middleware2(c *gin.Context) {
	//中间件逻辑

	fmt.Println(c.Request.Header.Get("abc"))
	fmt.Println("start middleware2")
	c.Next()
	fmt.Println("end middleware2")
}

func middleware3(c *gin.Context) {
	fmt.Println("start middleware3")
	c.Next()
	fmt.Println("end middleware3")
}

// 执行顺序
// a ->b >c -> handler ->c->b>a
func main() {
	r := gin.New()
	r.Use(middleware1)
	r.Use(middleware2)
	r.Use(middleware3)

	r.GET("/test", func(c *gin.Context) {

		// 打印："12345"
		log.Println("example")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
