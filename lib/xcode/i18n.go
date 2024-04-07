package xcode

import (
	"golang.org/x/text/language"
)

type ErrorCode int

const (
	Success                   ErrorCode = 0
	ErrInvalidParam           ErrorCode = 400
	StatusInternalServerError ErrorCode = 500
	//ErrInvalidOldPassword 旧密码错误
	ErrInvalidOldPassword ErrorCode = 1001
	ErrUserLoginFailed    ErrorCode = 1002
	ErrUserIsDisabled     ErrorCode = 1003
	//ErrThirdPartyPluginError 第三方插件错误
	ErrThirdPartyPluginError ErrorCode = 3001

	ErrUpdateDataFailed ErrorCode = iota + 4998
	// ErrDataNotFoundFailed 数据未找到
	ErrDataNotFoundFailed
	ErrDataInsertFailed
	ErrDataDeleteFailed
)

var (
	responseMsg = map[ErrorCode]map[language.Tag]string{
		Success: {
			language.Chinese: "成功",
			language.English: "Success",
		},
		StatusInternalServerError: {
			language.Chinese: "服务器内部错误",
			language.English: "Internal Server Error",
		},
		ErrUserLoginFailed: {
			language.Chinese: "用户登录失败",
			language.English: "User login failed",
		},
		ErrInvalidParam: {
			language.Chinese: "无效参数",
			language.English: "Invalid parameter",
		},
		ErrInvalidOldPassword: {
			language.Chinese: "密码错误",
			language.English: "Invalid password",
		},
		ErrThirdPartyPluginError: {
			language.Chinese: "第三方插件错误",
			language.English: "third party plugin error",
		},
		ErrUpdateDataFailed: {
			language.Chinese: "更新数据失败",
			language.English: "Update data failed",
		},
		ErrDataNotFoundFailed: {
			language.Chinese: "数据未找到",
			language.English: "Data not found",
		},
		ErrDataInsertFailed: {
			language.Chinese: "数据插入失败",
			language.English: "Data insert failed",
		},
		ErrDataDeleteFailed: {
			language.Chinese: "数据删除失败",
			language.English: "Data delete failed",
		},
	}
)
