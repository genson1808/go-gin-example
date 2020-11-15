package e

var MsgFlags = map[int]string {
	SUCCESS: "ok",
	ERROR: "fail",
	INVALID_PARAMS: "Request parameter error",
	ERROR_EXIST_TAG: "The tag name already exists",
	ERROR_NOT_EXIST_TAG: "The tag does not exist",
	ERROR_NOT_EXIST_ARTICLE: "The article does not exist",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "Token authentication failed",
	ERROR_AUTH_TOKEN: "Token generation failed",
	ERROR_AUTH: "Token Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}