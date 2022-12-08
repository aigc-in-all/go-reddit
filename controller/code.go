package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "invalid param",
	CodeUserExist:       "user exist",
	CodeUserNotExist:    "user not exist",
	CodeInvalidPassword: "invalid password",
	CodeServerBusy:      "server busy",

	CodeNeedLogin:    "need login",
	CodeInvalidToken: "invalid token",
}

func (r ResCode) Msg() string {
	msg, ok := codeMsgMap[r]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
