package ctrAdminV1_1Login

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type ReqEmployeeIn struct {
	Verify string `form:"verify"`
	Mobile string `form:"mobile"`
	Code   string `form:"code"`
}

// 获取request json参数
func (this *ReqEmployeeIn) GetReq(ctx *gin.Context) (res ReqEmployeeIn) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		panic("参数缺失")
	}

	log.Println("shuru :", string(data))

	if err := json.Unmarshal(data, &res); nil != err {
		log.Println(err)
		panic("参数错误")
	}

	this.Validator()

	return
}

// 验证参数合法性
func (this *ReqEmployeeIn) Validator() {
	// ...
}
