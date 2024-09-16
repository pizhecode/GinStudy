package controller

import (
	"web_App/logic"

	"github.com/gin-gonic/gin"
)

// 处理注册请求
func SignUpHandler(c *gin.Context) {
	//1 获取参数和参数校验
	//2 业务处理
	logic.SignUp()
	//3 返回相应
	c.JSON(200, "ok")
}
