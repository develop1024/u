package u

import (
	"fmt"
	"github.com/go-basic/uuid"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"log"
	"strings"
)

type Map map[string]interface{}

// Author contact author
func Author() {
	fmt.Println("author: 王哈哈")
	fmt.Println("contact: 31931727@qq.com")
}

// ErgodicParentChild 遍历父子组合关系
func ErgodicParentChild(tableName string,
	childName string,
	parentField string,
	childParentField string,
	isParentCondition interface{}) *gdb.Result {
	parentData, _ := g.DB().Model(tableName).Where(isParentCondition).FindAll()

	for _, parent := range parentData {
		childData, _ := g.DB().Model(tableName).Where(g.Map{
			childParentField: parent[parentField],
		}).FindAll()
		parent[childName] = g.NewVar(childData)
	}
	return &parentData
}

// ErgodicParentChild1 遍历父子组合关系
func ErgodicParentChild1(tableName string,
	childName string,
	parentField string,
	childParentField string) *gdb.Result {
	return ErgodicParentChild(tableName,
		childName,
		parentField,
		childParentField,
		childParentField+` is null or `+childParentField+` =''`,
	)
}

// ErgodicParentChild2 遍历父子组合关系
func ErgodicParentChild2(tableName string,
	childName string,
	parentField string,
	childParentField string) *gdb.Result {
	return ErgodicParentChild(tableName,
		childName,
		parentField,
		childParentField,
		parentField+`=0`,
	)
}

// HasExists 判断元素是否存在与某列表
func HasExists(dataList interface{}, data interface{}) bool {
	result := false

	switch dataList.(type) {
	case []string:
		for _, item := range dataList.([]string) {
			if item == data {
				result = true
			}
		}
	case []int:
		for _, item := range dataList.([]int) {
			if item == data {
				result = true
			}
		}
	case []int8:
		for _, item := range dataList.([]int8) {
			if item == data {
				result = true
			}
		}
	case []int16:
		for _, item := range dataList.([]int16) {
			if item == data {
				result = true
			}
		}
	case []int32:
		for _, item := range dataList.([]int32) {
			if item == data {
				result = true
			}
		}
	case []int64:
		for _, item := range dataList.([]int64) {
			if item == data {
				result = true
			}
		}
	case []uint:
		for _, item := range dataList.([]int) {
			if item == data {
				result = true
			}
		}
	case []uint8:
		for _, item := range dataList.([]int8) {
			if item == data {
				result = true
			}
		}
	case []uint16:
		for _, item := range dataList.([]int16) {
			if item == data {
				result = true
			}
		}
	case []uint32:
		for _, item := range dataList.([]int32) {
			if item == data {
				result = true
			}
		}
	case []uint64:
		for _, item := range dataList.([]int64) {
			if item == data {
				result = true
			}
		}

	case []float32:
		for _, item := range dataList.([]float32) {
			if item == data {
				result = true
			}
		}
	case []float64:
		for _, item := range dataList.([]float64) {
			if item == data {
				result = true
			}
		}
	case []gdb.Record:
		for _, item := range dataList.([]gdb.Record) {
			if item.Json() == data.(gdb.Record).Json() {
				result = true
			}
		}
	}

	return result
}

// UUID 获取UUID
func UUID() string {
	s := uuid.New()
	return s
}

// UUIDStr 获取UUID字符串, 不含"-"
func UUIDStr() string {
	s := uuid.New()
	return strings.ReplaceAll(s, "-", "")
}

// UUIDStrUpper 获取UUID字符串, 不含"-", 转大写
func UUIDStrUpper() string {
	s := uuid.New()
	return strings.ToUpper(strings.ReplaceAll(s, "-", ""))
}

// Foreach 循环执行
func Foreach(n int, callback func()) {
	for i := 0; i < n; i++ {
		callback()
	}
}

// separator 分隔符
const separator = "#############################################"

// Debug 调试输出
func Debug(data ...interface{}) {
	log.Println(separator)
	log.Println(data...)
	log.Println(separator)
}

// DebugYellow 调试输出-黄色
func DebugYellow(data ...interface{}) {
	PrintlnYellow(Time().DateTime(), separator)
	PrintYellow(Time().DateTime(), " ")
	PrintlnYellow(data...)
	PrintlnYellow(Time().DateTime(), separator)
}

// DebugBlue 调试输出-蓝色
func DebugBlue(data ...interface{}) {
	PrintlnBlue(Time().DateTime(), separator)
	PrintBlue(Time().DateTime(), " ")
	PrintlnBlue(data...)
	PrintlnBlue(Time().DateTime(), separator)
}

// DebugGreen 调试输出-绿色
func DebugGreen(data ...interface{}) {
	PrintlnGreen(Time().DateTime(), separator)
	PrintGreen(Time().DateTime(), " ")
	PrintlnGreen(data...)
	PrintlnGreen(Time().DateTime(), separator)
}

// DebugGray 调试输出-灰色
func DebugGray(data ...interface{}) {
	PrintlnGray(Time().DateTime(), separator)
	PrintGray(Time().DateTime(), " ")
	PrintlnGray(data...)
	PrintlnGray(Time().DateTime(), separator)
}

// DebugCyan 调试输出-青色
func DebugCyan(data ...interface{}) {
	PrintlnCyan(Time().DateTime(), separator)
	PrintCyan(Time().DateTime(), " ")
	PrintlnCyan(data...)
	PrintlnCyan(Time().DateTime(), separator)
}

// Version 输出版本
func Version() {
	fmt.Println("v0.0.18")
	fmt.Println("Last update time: 2021-12-19 14:45:00")
}
