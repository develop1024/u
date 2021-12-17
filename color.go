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
	color.Redp(val...)
}

// PrintlnRed 红色换行
func PrintlnRed(val ...interface{}) {
	color.Redln(val...)
}

// PrintGreen 绿色不换行
func PrintGreen(val ...interface{}) {
	color.Greenp(val...)
}

// PrintlnGreen 绿色换行
func PrintlnGreen(val ...interface{}) {
	color.Greenln(val...)
}

// PrintBlue 蓝色不换行
func PrintBlue(val ...interface{}) {
	color.Bluep(val...)
}

// PrintlnBlue 蓝色换行
func PrintlnBlue(val ...interface{}) {
	color.Blueln(val)
}

// PrintYellow 黄色不换行
func PrintYellow(val ...interface{}) {
	color.Yellowp(val...)
}

// PrintlnYellow 黄色换行
func PrintlnYellow(val ...interface{}) {
	color.Yellowln(val...)
}

// PrintGray 灰色不换行
func PrintGray(val ...interface{}) {
	color.Grayp(val...)
}

// PrintlnGray 灰色换行
func PrintlnGray(val ...interface{}) {
	color.Grayln(val...)
}

// PrintCyan 青色不换行
func PrintCyan(val ...interface{}) {
	color.Cyanp(val...)
}

// PrintlnCyan 青色换行
func PrintlnCyan(val ...interface{}) {
	color.Cyanln(val...)
}
