/*
 * 文件用于接入数据库，所有数据库初始化在这个文件完成
 * 接入单个或多个数据库均需要在这里配置
 *
 * 注意：目前数据库初始化包括 gorm 库的初始化
 *
 *
 * 调用gorm的时候请采用 db.NewDb() 获取db，db会自动处理长链接断开的问题
 */

package db

import (
	"fmt"
	"sync"
	"time"

	"go-easy-frame/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	gormDb            *gorm.DB
	gormDbAutoCfg     *gorm.DB
	gormDbSelfDefined *gorm.DB
	IsInitDb          bool
	selfDns           string
	initLock          sync.Mutex

	err error
)

func init() {
	InitDb()
}

func New() *gorm.DB {
	initLock.Lock()
	defer initLock.Unlock()

	if !IsInitDb {
		InitDb()
	}
	if "" != selfDns {
		return gormDbSelfDefined
	} else {
		return gormDb
	}
}

func NewAutoCfg() *gorm.DB {
	if !IsInitDb {
		InitDb()
	}

	return gormDbAutoCfg
}

func InitSelfDb(dns string) {
	selfDns = dns
	gormDbSelfDefined = GormDb(selfDns)
	SetDbConfig(gormDbSelfDefined)

	IsInitDb = true
}

func InitDb() {
	fmt.Println("begin init mysql")

	// 初始化 gorm
	InitGorm()
	InitGormAutoCfg()

	IsInitDb = true

	fmt.Println("end int mysql")
}

// 初始化 gorm
func InitGorm() {
	gormDb = GormDb("sqlconn")
	SetDbConfig(gormDb)
}

func InitGormAutoCfg() {
	gormDbAutoCfg = GormDb("sqlconnAutoCfg")
	SetDbConfig(gormDbAutoCfg)
}

func SetDbConfig(ob *gorm.DB) {
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	ob.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	ob.DB().SetMaxOpenConns(1000)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	ob.DB().SetConnMaxLifetime(time.Hour)
}

// gorm 调用必须通过这个方法调用,函数处理了连接池和断开自动重连接的处理。
func GormDb(sqlDns string) *gorm.DB {
	dnsStr := config.GetEnv().MysqlDns
	gormDb, err := gorm.Open("mysql", dnsStr+"&parseTime=True&loc=Local")
	if nil != err {
		panic("gorm数据库连接错误！" + dnsStr)
	}

	fmt.Println(gormDb.DB().Ping())

	return gormDb
}
