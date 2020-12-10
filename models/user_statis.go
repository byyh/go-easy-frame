package models

// 数据库model示例

type UserStatis struct {
	Id         int   `orm:"column(id);auto"`
	UcId       int64 `orm:"column(uc_id)" description:"会员中心的id"`
	Good       uint  `orm:"column(good);size(11)"`
	Fans       uint  `orm:"column(fans)" description:""`
	Friends    uint  `orm:"column(friends)" description:""`
	Star       uint  `orm:"column(star);" description:""`
	ModifyTime uint  `orm:"column(modify_time)" description:"最后修改时间"`
	AddTime    uint  `orm:"column(add_time)" description:"添加时间"`
}

func (t *UserStatis) TableName() string {
	return "user_statis"
}
