package ctrAdminV1_1

import (
	"go-easy-frame/services/atom"
)

type ReqUser struct {
	UcId int64 `json:"uc_id"`
}

type RespUser struct {
	UcId        int64  `json:"uc_id"`
	Mobile      string `json:"mobile"`
	GroupId     uint   `json:"group_id"`
	LoginClient int8   `json:"login_client"`
	RealName    string `json:"real_name"`
	Good        uint   `json:"good"`
	Fans        uint   `json:"fans"`
	Friends     uint   `json:"friends"`
	Star        uint   `json:"star"`
}

func (r *ReqUser) ParamsValidator() {
	if 0 == r.UcId {
		panic(atom.NewMyError("ucid不能为空", atom.ErrCodeInput))
	}
}
