package cronUser

import (
	"log"
	"time"
)

// 示例：用户等级调整
type RaskUserQuantity struct {
}

func (this *RaskUserQuantity) Run() {
	log.Println("begin RaskUserQuantity run", time.Now().String())

}
