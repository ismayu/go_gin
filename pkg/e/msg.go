package e

var MsgFlags = map[int]string{
	SUCCESS : "OK",
	ERROR : "fail",
	INVALID_PARAMS: "request params error",
	ERROR_EXIST_TAG: "already exist tag",
	ERROR_NOT_EXIST_TAG: "not exist tag",
	ERROR_NOT_EXIST_ARTICLE: "this article not exist",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token permission denied",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token already timeout",
	ERROR_AUTH_TOKEN: "Token create fail",
	ERROR_AUTH: "Token error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR]
}