package models

// 数据库model示例

type UserInfo struct {
	Id              int    `orm:"column(id);auto"`
	UcGroupId       int    `orm:"column(uc_group_id)" description:"会员中心租户ID"`
	UcId            int64  `orm:"column(uc_id)" description:"会员中心的id"`
	Mobile          string `orm:"column(mobile);size(11)"`
	GroupId         uint   `orm:"column(group_id)" description:"集团ID"`
	LoginClient     int8   `orm:"column(login_client)" description:"登录设备 1 ios 2android "`
	RealName        string `orm:"column(real_name);size(50)" description:"真实姓名"`
	CertificateNo   string `orm:"column(certificate_no);size(50)" description:"证件号"`
	CertificateType int8   `orm:"column(certificate_type)" description:"证件类型"`
	ModifyTime      uint   `orm:"column(modify_time)" description:"最后修改时间"`
	AddTime         uint   `orm:"column(add_time)" description:"添加时间"`
}

func (t *UserInfo) TableName() string {
	return "user_info"
}
