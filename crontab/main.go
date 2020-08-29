package main

/**
 * 定时任务
 * 定时任务不依赖linux的crontab
 */
import (
	"log"

	"github.com/byyh/go-easy-frame/crontab/user"

	"github.com/robfig/cron"
)

func main() {
	log.Println("begin")
	c := cron.New() // 新建一个定时任务对象
	// 秒 分 时 日 月 周（0代表周日）
	c.AddFunc("*/5 * * * * *", new(cronUser.RaskUserLvl).Run)      // 定时任务示例1，多个定时任务参考示例添加多个
	c.AddFunc("*/5 * * * * *", new(cronUser.RaskUserQuantity).Run) // 定时任务示例1，多个定时任务参考示例添加多个

	c.Start()

	log.Println("start...")
	select {}

	log.Println("end...")
}
