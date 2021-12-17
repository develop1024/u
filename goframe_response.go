package u

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"log"
)

// 指定返回的状态
const (
	Success     = 0
	ServerError = 50001
	LogicError  = 50002
	AuthenticationError = 401
)

// ResponseError 返回错误
func ResponseError(r *ghttp.Request, err error) {
	log.Println(err)
	g.Log().Error(err.Error())

	ResponseJsonExit(Map{
		"code": ServerError,
		"msg":  "服务器开小差了, 请联系管理员~",
	}, r)
}

// ResponseJson response json
func ResponseJson(content interface{}, r *ghttp.Request) {
	_ = r.Response.WriteJson(content)
}

// ResponseJsonExit response json and exit
func ResponseJsonExit(content interface{}, r *ghttp.Request) {
	_ = r.Response.WriteJsonExit(content)
}

// ResponseErrorAndSkip 返回错误
func ResponseErrorAndSkip(r *ghttp.Request, err error) {
	log.Println(err)

	ResponseJsonExit(Map{
		"code": ServerError,
		"msg":  "服务器开小差了, 请联系管理员~",
	}, r)
}

// ResponseErrorWithMsg 返回错误并指定消息
func ResponseErrorWithMsg(msg string, err error, r *ghttp.Request) {
	log.Println(err)
	g.Log().Error(err.Error())

	ResponseJsonExit(Map{
		"code": ServerError,
		"msg":  msg,
	}, r)
}

// ResponseFail 返回失败
func ResponseFail(r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": LogicError,
		"msg":  "操作失败",
	}, r)
}

// ResponseFailWithMsg 返回失败并指定消息
func ResponseFailWithMsg(msg string, r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": LogicError,
		"msg":  msg,
	}, r)
}

// ResponseSuccess 返回成功
func ResponseSuccess(r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": Success,
		"msg":  "操作成功",
	}, r)
}

// ResponseSuccessWithMsg 返回成功并指定消息
func ResponseSuccessWithMsg(msg string, r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": Success,
		"msg":  msg,
	}, r)
}

// ResponseSuccessWithData 返回成功并指定数据
func ResponseSuccessWithData(data interface{}, r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": Success,
		"msg":  "操作成功",
		"data": data,
	}, r)
}

// ResponseSuccessWithCustomData 返回成功并指定数据
func ResponseSuccessWithCustomData(data map[string]interface{}, r *ghttp.Request) {
	data["code"] = Success
	data["msg"] = "操作成功"

	ResponseJsonExit(data, r)
}

// ResponseSuccessWithDataMsg 返回成功并指定数据和消息
func ResponseSuccessWithDataMsg(data map[string]interface{}, msg string, r *ghttp.Request) {
	ResponseJsonExit(Map{
		"code": Success,
		"msg":  msg,
		"data": data,
	}, r)
}
