package atom

const (
	Success               = 0
	ErrCodeNoToken        = 1
	ErrCodeTokenInvalid   = 1
	ErrCodeTokenTimeout   = 2
	ErrCodePgTokenInvalid = 5

	ErrCodeCallHigh = 1005

	ErrCodeVerifyErr     = 1320
	ErrCodeVerifyTimeout = 1321

	ErrCodeBlockUser         = 1011
	ErrCodeMobileVerifyErr   = 1020
	ErrCodeMobileVerifyTimes = 1021
	ErrCodePasswordErr       = 1030
	ErrCodeUrlErr            = 1040
	ErrCodeNoAuth            = 1120
	ErrCodeNoCall            = 1110
	ErrCodeUploadFile        = 1200
	ErrCodeFileLarge         = 1205
	ErrCodeNoShopId          = 1055

	ErrCodeNoPage = 2400
	ErrCodeInput  = 2500
	ErrCodeParam  = 2600

	// 大于2000 为后端错误，小于3000为输入错误
	ErrCodeAutoConfig = 4020

	ErrCodeEnv = 4100

	ErrCodeDb    = 3010
	ErrCodeRedis = 3020
	ErrCodeLog   = 3040
	ErrCodeExcel = 3050

	ErrCodeHandle    = 5010
	ErrCodeArgsTrans = 5020
	ErrCodeFuncCall  = 5030

	ErrCodeCommon      = 5000 // 通用一般性错误
	ErrCodeLogicHandle = 5001 // 通用逻辑处理错误
	ErrCodeException   = 5005
	ErrCodeJson        = 5007
)

var (
	RetCodeMap map[int]string
)

func init() {
	RetCodeMap = make(map[int]string)

	RetCodeMap[Success] = "成功" // 不等于0表示失败。
	RetCodeMap[ErrCodeNoToken] = "token无效"

	RetCodeMap[ErrCodeNoPage] = "页面不存在"
	RetCodeMap[ErrCodeNoShopId] = "shopId缺失"

	RetCodeMap[ErrCodeInput] = "参数不合法"
	RetCodeMap[ErrCodeParam] = "参数不合法"
	RetCodeMap[ErrCodeCallHigh] = "高频访问"
	RetCodeMap[ErrCodeBlockUser] = "黑名单"
	RetCodeMap[ErrCodeMobileVerifyErr] = "手机验证码错误"
	RetCodeMap[ErrCodeMobileVerifyTimes] = "手机验证码获取次数太多"
	RetCodeMap[ErrCodePasswordErr] = "用户或密码错误"
	RetCodeMap[ErrCodeUrlErr] = "url错误"
	RetCodeMap[ErrCodeNoCall] = "禁止访问"
	RetCodeMap[ErrCodeNoAuth] = "没有权限"
	RetCodeMap[ErrCodeUploadFile] = "上传文件失败"
	RetCodeMap[ErrCodeFileLarge] = "上传文件太大"

	RetCodeMap[ErrCodeVerifyErr] = "验证码错误"
	RetCodeMap[ErrCodeVerifyTimeout] = "验证码过期"

	RetCodeMap[ErrCodeEnv] = "环境配置错误"

	RetCodeMap[ErrCodeAutoConfig] = "at配置错误"
	RetCodeMap[ErrCodeDb] = "数据库错误"
	RetCodeMap[ErrCodeRedis] = "缓存错误"
	RetCodeMap[ErrCodeLog] = "日志错误"
	RetCodeMap[ErrCodeExcel] = "excel错误"
	RetCodeMap[ErrCodeHandle] = "执行错误"
	RetCodeMap[ErrCodeArgsTrans] = "参数转换错误"
	RetCodeMap[ErrCodeFuncCall] = "函数调用错误"

	RetCodeMap[ErrCodeCommon] = "处理失败"        // 通用
	RetCodeMap[ErrCodeLogicHandle] = "处理逻辑失败" // 通用
	RetCodeMap[ErrCodeException] = "处理异常失败"   // 通用
	RetCodeMap[ErrCodeJson] = "json处理失败"
}

func GetRetMsgByCode(code int) string {
	if v, ok := RetCodeMap[code]; ok {
		return v
	}

	return "失败"
}
