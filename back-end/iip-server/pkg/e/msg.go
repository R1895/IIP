package e

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	ErrorExist:                 "已存在",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorNotExist:              "不存在",
	ErrorNotUpdate:             "无效更新",
	ErrorExistUser:             "用户已存在",
	ErrorNotExistUser:          "用户不存在",
	ErrorFailEncryption:        "加密失败",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
