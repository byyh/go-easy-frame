package ctrAdminV1_1

type ReqUserTest struct {
	Verify string `form:"verify"`
	Mobile string `form:"mobile"`
	Code   string `form:"code"`
}
