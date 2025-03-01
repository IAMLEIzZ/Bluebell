package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvildToken
)

const (
	PostOrderByTime = "time"
	PostOrderByScore = "score"
)

const (
	CodeUrlQueryID = "id"
	ContextUserIDKey = "userID"
	CodeUrlQueryPage = "page"
	CodeUrlQuerySize = "size"
)

var codeMsgMap = map[ResCode]string {
	CodeSuccess: "success",
	CodeInvalidParam: "请求参数错误",
	CodeUserExist: "用户名已存在",
	CodeUserNotExist: "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy: "服务繁忙",
	CodeNeedLogin: "用户未登录",
	CodeInvildToken: "token不合法",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}