package ctrAdminV1_1

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-easy-frame/controllers"
	_ "go-easy-frame/models"
	"go-easy-frame/services/atom"
	"go-easy-frame/services/cache"
	"go-easy-frame/services/db"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Exam struct {
	ctr.BaseController
}

// 测试db
func (this *Exam) Db(ctx *gin.Context) {
	// 获取上传参数
	var inputParam ReqUser
	if nil != ctx.ShouldBindJSON(&inputParam) {
		panic(atom.NewMyErrorByCode(atom.ErrCodeInput))
	}
	inputParam.ParamsValidator()

	//
	var syncW sync.WaitGroup

	ob := db.New()

	// 获取企业id

	var userRes RespUser
	var err, err2 error
	//var userInfo models.UserInfo
	//var userStatis models.UserStatis

	syncW.Add(2)

	go func() {
		err = ob.Table("user_info").
			Where("id=? ", inputParam.UcId).
			First(&userRes).Error
		syncW.Done()
		// 这个里面不能直接panic，会导致程序崩溃
	}()

	go func() {
		err2 = ob.Table("user_statis").
			Where("id=? ", inputParam.UcId).
			First(&userRes).Error

		syncW.Done()
	}()

	syncW.Wait()

	if nil != err {
		if "record not found" == err.Error() {
			panic(atom.NewMyError("user_info不存在", atom.ErrCodeDb))
		} else {
			panic(errors.New(fmt.Sprintf("获取数据库user_info失败:%v", err)))
		}
	}
	if nil != err2 {
		if "record not found" != err2.Error() {
			panic(errors.New(fmt.Sprintf("获取数据库user_statis失败:%v", err2)))
		}
	}

	// 写入redis
	ocah := cache.New()
	key := "K1_TABLE_USER-" + strconv.Itoa(int(inputParam.UcId))
	val, err := json.Marshal(&userRes)
	if nil != err {
		panic(atom.NewMyError("json转换失败", atom.ErrCodeJson))
	}

	tm := 16 * 3600 * time.Second // 16小时
	valStr := string(val)
	fmt.Println(key, valStr, tm)
	err = ocah.Set(key, valStr, tm).Err()
	if err != nil {
		panic(err)
	}

	this.Success(ctx, map[string]interface{}{
		"out":   userRes,
		"input": inputParam,
	})
}

func (this *Exam) Redis(ctx *gin.Context) {
	var inputParam ReqUser
	var er error
	inputParam.UcId, er = strconv.ParseInt(ctx.Query("uc_id"), 10, 64)
	if nil != er {
		log.Println(er)
		panic(atom.NewMyErrorByCode(atom.ErrCodeInput))
	}
	inputParam.ParamsValidator()

	key := "K1_TABLE_USER-" + strconv.Itoa(int(inputParam.UcId))

	ocah := cache.New()

	val, err := ocah.Get(key).Result()
	if err == redis.Nil {
		//fmt.Println("key2 does not exist")
	} else if err != nil {
		// 错误
		panic(atom.NewMyError("获取缓存失败", atom.ErrCodeJson))
	}

	var user RespUser
	fmt.Println(val)
	if err = json.Unmarshal([]byte(val), &user); nil != err {
		fmt.Println(err)
		panic(atom.NewMyError("json转换失败", atom.ErrCodeJson))
	}

	this.Success(ctx, map[string]interface{}{
		"user": user,
	})

}
