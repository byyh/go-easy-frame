package middleware

// 登陆验证，已登陆的用户用户信息参数通过Keys存储到上下文中，后面的控制器可以获取

import (
	"github.com/gin-gonic/gin"
)

func LoginVerify(c *gin.Context) {
	// 验证登陆

	// 当前用户的登陆user_id和user_info存储到上下文的keys中
	c.Set("cur_user_id", "abci48df8")
	c.Set("cur_user", map[string]interface{}{
		"id":     "abci48df8",
		"name":   "张三",
		"mobile": "15022336644",
	})
}
