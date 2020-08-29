package atom

const (
	Success             = 0
	ErrCodeNoToken      = 1
	ErrCodeTokenInvalid = 1
	ErrCodeTokenTimeout = 2

	// ...
)

var (
	RetCodeMap map[int]string
)

func init() {
	RetCodeMap = make(map[int]string)

	RetCodeMap[Success] = "成功" // 不等于0表示失败。
	RetCodeMap[ErrCodeNoToken] = "token无效"
	RetCodeMap[ErrCodeTokenInvalid] = "token无效"
	RetCodeMap[ErrCodeTokenTimeout] = "登陆超时"
}

func GetRetMsgByCode(code int) string {
	if v, ok := RetCodeMap[code]; ok {
		return v
	}

	return "失败"
}
