package u

import (
	"fmt"
	"github.com/gookit/color"
)

// Println = fmt.Println
func Println(val ...interface{}) {
	fmt.Println(val...)
}

// Print = fmt.Print
func Print(val ...interface{}) {
	fmt.Print(val...)
}

// Printf = fmt.Printf
func Printf(format string, val ...interface{}) {
	fmt.Printf(format, val...)
}

// PrintRed 红色不换行
func PrintRed(val ...interface{}) {
	color.Red.Print(val...)
}

// PrintlnRed 红色换行
func PrintlnRed(val ...interface{}) {
	color.Red.Println(val...)
}

// PrintfRed 红色-格式化输出
func PrintfRed(format string, val ...interface{}) {
	color.Red.Printf(format, val...)
}

// PrintGreen 绿色不换行
func PrintGreen(val ...interface{}) {
	color.Green.Print(val...)
}

// PrintlnGreen 绿色换行
func PrintlnGreen(val ...interface{}) {
	color.Green.Println(val...)
}

// PrintfGreen 绿色-格式化输出
func PrintfGreen(format string, val ...interface{}) {
	color.Green.Printf(format, val...)
}

// PrintBlue 蓝色不换行
func PrintBlue(val ...interface{}) {
	color.Blue.Print(val...)
}

// PrintlnBlue 蓝色换行
func PrintlnBlue(val ...interface{}) {
	color.Blue.Println(val...)
}

// PrintfBlue 蓝色-格式化输出
func PrintfBlue(format string, val ...interface{}) {
	color.Blue.Printf(format, val...)
}

// PrintYellow 黄色不换行
func PrintYellow(val ...interface{}) {
	color.Yellow.Print(val...)
}

// PrintlnYellow 黄色换行
func PrintlnYellow(val ...interface{}) {
	color.Yellow.Println(val...)
}

// PrintfYellow 黄色-格式化输出
func PrintfYellow(format string, val ...interface{}) {
	color.Yellow.Printf(format, val...)
}

// PrintGray 灰色不换行
func PrintGray(val ...interface{}) {
	color.Gray.Print(val...)
}

// PrintlnGray 灰色换行
func PrintlnGray(val ...interface{}) {
	color.Gray.Println(val...)
}

// PrintfGray 灰色-格式化输出
func PrintfGray(format string, val ...interface{}) {
	color.Gray.Printf(format, val...)
}

// PrintCyan 青色不换行
func PrintCyan(val ...interface{}) {
	color.Cyan.Print(val...)
}

// PrintlnCyan 青色换行
func PrintlnCyan(val ...interface{}) {
	color.Cyan.Println(val...)
}

// PrintfCyan 青色-格式化输出
func PrintfCyan(format string, val ...interface{}) {
	color.Cyan.Printf(format, val...)
}
