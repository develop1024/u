package u

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"sort"
	"strings"
)

// GetErrorExit 返回错误并退出
func GetErrorExit(err error, r *ghttp.Request) {
	if err != nil {
		ResponseError(r, err)
	}
}

// GetErrorAndSkip 返回错误并退出
func GetErrorAndSkip(err error, r *ghttp.Request) {
	if err != nil {
		ResponseErrorAndSkip(r, err)
	}
}

// GetStructAndValid 获取结构体中的数据并验证
func GetStructAndValid(params interface{}, r *ghttp.Request) {
	err := r.GetStruct(params)
	GetErrorExit(err, r)

	if err := gvalid.CheckStruct(r.Context(), params, nil); err != nil {
		ResponseFailWithMsg(err.FirstString(), r)
	}
}

// GenerateParamsSign 参数签名
func GenerateParamsSign(params map[string]interface{}) string {
	// 参数签名，保证参数不被篡改
	paramSecret := g.Cfg().GetString("web.paramSecret")

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var stringList []string
	for _, k := range keys {
		key, val := k, params[k]
		stringList = append(stringList, fmt.Sprintf("%v=%v", key, val))
	}
	stringA := strings.Join(stringList, "&")

	stringSignTemp := fmt.Sprintf("%s&paramsSecret=%s", stringA, paramSecret)
	sign := strings.ToUpper(MD5(stringSignTemp))

	return sign
}

// Paginate 分页组件
func Paginate(total int, r *ghttp.Request, function func(page, pageSize int) (gdb.Result, error)) {
	page := r.GetInt("page", 1)
	pageSize := r.GetInt("pageSize", 100)

	totalPage := total / pageSize
	if total % pageSize !=0 {
		totalPage += 1
	}

	result, err := function(page, pageSize)
	GetErrorExit(err, r)

	hasNextPage := true
	if page >= totalPage {
		hasNextPage = false
	}

	ResponseSuccessWithData(g.Map{
		"total": total,
		"totalPage": totalPage,
		"currentPage": page,
		"data": result,
		"hasNextPage": hasNextPage,
	}, r)
}

// LayPage layui分页
func LayPage(r *ghttp.Request, tableName string, where interface{})  {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)

	data, err := g.DB().Model(tableName).Where(where).Offset((page-1)*limit).Limit(limit).FindAll()
	GetErrorExit(err, r)

	count, err := g.DB().Model(tableName).Where(where).Count()
	GetErrorExit(err, r)

	_ = r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg": "ok",
		"count": count,
		"data": data,
	})
}

// LayPageCallback layui分页支持回调函数
func LayPageCallback(r *ghttp.Request, tableName string, where interface{}, callback func(data *gdb.Result))  {
	page := r.GetInt("page", 1)
	limit := r.GetInt("limit", 10)

	data, err := g.DB().Model(tableName).Where(where).Offset((page-1)*limit).Limit(limit).FindAll()
	GetErrorExit(err, r)

	// 调用回调函数
	callback(&data)

	count, err := g.DB().Model(tableName).Where(where).Count()
	GetErrorExit(err, r)

	_ = r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg": "ok",
		"count": count,
		"data": data,
	})
}
