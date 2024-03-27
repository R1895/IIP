package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorExist           = 10001
	ErrorExistUser       = 10002
	ErrorNotExistUser    = 10003
	ErrorExistRole       = 10004
	ErrorFailEncryption  = 10006
	ErrorNotCompare      = 10007
	ErrorExistPermission = 10008
	ErrorExistWorkshop   = 10009

	ErrorNotExist              = 20001
	ErrorNotUpdate             = 20002
	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001
)
