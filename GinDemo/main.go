package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 引擎实例
	r := gin.Default()

	// 定义一个 GET 请求的路由处理
	r.GET("/hello", func(c *gin.Context) {
		// 返回 HTTP 状态码 200 和 "ok" 字符串
		c.String(http.StatusOK, "ok")
	})

	// 启动 HTTP 服务器，默认监听在 :8080 端口
	r.Run()
}
