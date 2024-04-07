package xcode

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/text/language"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWriter interface {
	error
	Code() ErrorCode
}

func NewResponse(code ErrorCode, tag language.Tag) error {
	return &responseWriterImpl{tag: tag, code: code}
}

type responseWriterImpl struct {
	tag  language.Tag
	code ErrorCode
}

func (r *responseWriterImpl) Code() ErrorCode {
	return r.code
}

func (r *responseWriterImpl) Error() string {
	return responseMsg[r.code][r.tag]
}

func SuccessResponse(data interface{}, tag language.Tag) *Response {
	return &Response{
		Code:    int(Success),
		Message: responseMsg[Success][tag],
		Data:    data,
	}
}

func ErrorResponse(code ErrorCode, msg string) *Response {
	return &Response{
		Code:    int(code),
		Message: msg,
		Data:    nil,
	}
}

func HttpResponse(r *http.Request, w http.ResponseWriter, resp interface{}, err error, lang language.Tag) {
	// 请求成功
	if err == nil {
		httpx.WriteJson(w, http.StatusOK, resp)
		return
	}
	// 记录日志
	loggerError(r.Context(), err)
	if e, ok := errors.Cause(err).(ResponseWriter); ok {
		writeErrorResponse(w, http.StatusOK, e.Code(), e.Error())
		return
	}
	writeErrorResponse(
		w,
		http.StatusInternalServerError,
		StatusInternalServerError,
		responseMsg[StatusInternalServerError][lang],
	)
}

func loggerError(ctx context.Context, err error) {
	logx.WithContext(ctx).Errorw(
		"api_error",
		logx.Field("err", err),
	)
}

func writeErrorResponse(w http.ResponseWriter, status int, errCode ErrorCode, errMsg string) {
	httpx.WriteJson(w, status, ErrorResponse(errCode, errMsg))
}
