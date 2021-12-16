package u

import (
	"math/rand"
	"strings"
)

// RandNum 生成随机数
func RandNum(num int) int {
	return rand.Intn(num)
}

// RandRangeNum 生成指定范围的随机数
func RandRangeNum(min, max int) int {
	return RandNum(max - min + 1) + 1
}

// RandStr 生成随机字符串
func RandStr(length int) string {
	strs := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strSlice := strings.Split(strs, "")
	strList := make([]string, length)
	for i:=0;i<length;i++ {
		num := RandNum(len(strs))
		strList = append(strList, strSlice[num])
	}
	return strings.Join(strList, "")
}
