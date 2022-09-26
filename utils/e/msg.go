package e

var MsgFlags = map[int]string{
	SUCCESS:        "",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	VERSION:        "header版本号错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

func GetStatus(code int) string {
	if code == SUCCESS {
		return "success"
	} else {
		return "failure"
	}
}
