package httprouter_server

import (
	"github.com/develop1024/u"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

type Router struct {
	*httprouter.Router
}

var _server *Router

// Server 获取server对象
func Server() *Router {
	_server = &Router{
		httprouter.New(),
	}
	//http.Handle("/", _server)
	return _server
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
}

// Handle server handle
func Handle(path string, method string, f func(c *Context)) {
	_server.Handle(method, path, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := new(Context)
		ctx.Response = writer
		ctx.Request = request
		ctx.Params = params
		f(ctx)
	})
}

// GET method GET
func (receiver *Router) GET(path string, f func(c *Context)) {
	Handle(path, http.MethodGet, f)
}

// POST method POST
func (receiver *Router) POST(path string, f func(c *Context)) {
	Handle(path, http.MethodPost, f)
}

// PUT method PUT
func (receiver *Router) PUT(path string, f func(c *Context)) {
	Handle(path, http.MethodPut, f)
}

// DELETE method DELETE
func (receiver *Router) DELETE(path string, f func(c *Context)) {
	Handle(path, http.MethodDelete, f)
}

// HEAD method HEAD
func (receiver *Router) HEAD(path string, f func(c *Context)) {
	Handle(path, http.MethodHead, f)
}

// OPTIONS method options
func (receiver *Router) OPTIONS(path string, f func(c *Context)) {
	Handle(path, http.MethodOptions, f)
}

// Run 运行
func (receiver *Router) Run(addr string) error {
	u.Yellow.Println("####################################")
	u.Blue.Printf("uWeb listen %s\n", addr)
	u.Yellow.Println("####################################")
	return http.ListenAndServe(addr, _server)
}

// JSON json响应
func (receiver *Context) JSON(data interface{}) {
	receiver.Response.WriteHeader(200)
	receiver.Response.Header().Set("Content-Type", "application/json")
	_, _ = receiver.Response.Write(u.ToBytes(u.Json(data)))
}

// String 字符串响应
func (receiver *Context) String(data string) {
	receiver.Response.WriteHeader(200)
	_, _ = receiver.Response.Write([]byte(data))
}

// GetParam 获取路由参数
func (receiver *Context) GetParam(key string) string {
	return receiver.Params.ByName(key)
}

// GetVar 获取查询参数
func (receiver *Context) GetVar(key string) string {
	return receiver.Request.URL.Query().Get(key)
}

// Bind 绑定参数
func (receiver *Context) Bind(obj interface{}) {
	dataMap := receiver.Request.URL.Query()

	v := reflect.ValueOf(obj)
	for i:=0;i<v.Elem().NumField();i++ {
		t := v.Type()
		fieldValue := dataMap.Get(t.Elem().Field(i).Tag.Get("field"))

		switch v.Elem().Field(i).Kind() {
		case reflect.Int,reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v.Elem().Field(i).SetInt(u.ToInt64(fieldValue))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v.Elem().Field(i).SetUint(u.ToUint64(fieldValue))
		case reflect.Float32, reflect.Float64:
			v.Elem().Field(i).SetFloat(u.ToFloat64(fieldValue))
		case reflect.Bool:
			v.Elem().Field(i).SetBool(u.ToBool(fieldValue))
		case reflect.String:
			v.Elem().Field(i).SetString(fieldValue)
		default:
			v.Elem().Field(i).SetString(fieldValue)
		}
	}
}
