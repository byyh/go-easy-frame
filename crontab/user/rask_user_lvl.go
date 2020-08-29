package cronUser

import (
	"log"
	"time"
)

// 示例：用户等级调整
type RaskUserLvl struct {
}

func (this *RaskUserLvl) Run() {
	log.Println("begin RaskUserLvl run", time.Now().String())

	panic("test exp")
}
