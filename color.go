package u

import (
	"fmt"
	"github.com/gookit/color"
)

const (
	Red     = color.Red
	Cyan    = color.Cyan
	Gray    = color.Gray
	Blue    = color.Blue
	Black   = color.Black
	Green   = color.Green
	White   = color.White
	Yellow  = color.Yellow
	Magenta = color.Magenta

	FgDarkGray =color.FgDarkGray
	FgLightRed = color.FgLightRed
	FgLightGreen = color.FgLightGreen
	FgLightYellow = color.FgLightYellow
	FgLightBlue = color.FgLightBlue
	FgLightMagenta = color.FgLightMagenta
	FgLightCyan = color.FgLightCyan
	FgLightWhite = color.FgLightWhite
	FgGray = color.FgGray

	BgBlack = color.BgBlack
	BgRed = color.BgRed
	BgGreen = color.BgGreen
	BgYellow = color.BgYellow
	BgBlue = color.BgBlue
	BgMagenta = color.BgMagenta
	BgCyan = color.BgCyan
	BgWhite = color.BgWhite
	BgDefault = color.BgDefault

	BgDarkGray = color.BgDarkGray
	BgLightRed = color.BgLightRed
	BgLightGreen = color.BgLightGreen
	BgLightYellow = color.BgLightYellow
	BgLightBlue = color.BgLightBlue
	BgLightMagenta = color.BgLightMagenta
	BgLightCyan = color.BgLightCyan
	BgLightWhite = color.BgLightWhite
	BgGray = color.BgGray

	LightRed     = FgLightRed
	LightCyan    = FgLightCyan
	LightBlue    = FgLightBlue
	LightGreen   = FgLightGreen
	LightWhite   = FgLightWhite
	LightYellow  = FgLightYellow
	LightMagenta = FgLightMagenta

	HiRed     = FgLightRed
	HiCyan    = FgLightCyan
	HiBlue    = FgLightBlue
	HiGreen   = FgLightGreen
	HiWhite   = FgLightWhite
	HiYellow  = FgLightYellow
	HiMagenta = FgLightMagenta

	BgHiRed     = BgLightRed
	BgHiCyan    = BgLightCyan
	BgHiBlue    = BgLightBlue
	BgHiGreen   = BgLightGreen
	BgHiWhite   = BgLightWhite
	BgHiYellow  = BgLightYellow
	BgHiMagenta = BgLightMagenta
)

type _color struct {
	Red     color.Color
	Cyan    color.Color
	Gray    color.Color
	Blue    color.Color
	Black   color.Color
	Green   color.Color
	White   color.Color
	Yellow  color.Color
	Magenta color.Color

	FgDarkGray color.Color
	FgLightRed color.Color
	FgLightGreen color.Color
	FgLightYellow color.Color
	FgLightBlue color.Color
	FgLightMagenta color.Color
	FgLightCyan color.Color
	FgLightWhite color.Color
	FgGray color.Color

	BgBlack color.Color
	BgRed color.Color
	BgGreen color.Color
	BgYellow color.Color
	BgBlue color.Color
	BgMagenta color.Color
	BgCyan color.Color
	BgWhite color.Color
	BgDefault color.Color

	BgDarkGray color.Color
	BgLightRed color.Color
	BgLightGreen  color.Color
	BgLightYellow  color.Color
	BgLightBlue  color.Color
	BgLightMagenta  color.Color
	BgLightCyan  color.Color
	BgLightWhite  color.Color
	BgGray  color.Color
}

func Color() *_color {
	return &_color{
		Red:     color.Red,
		Cyan:    color.Cyan,
		Gray:    color.Gray,
		Blue:    color.Blue,
		Black:   color.Black,
		Green:   color.Green,
		White:   color.White,
		Yellow:  color.Yellow,
		Magenta: color.Magenta,

		FgDarkGray :color.FgDarkGray,
		FgLightRed : color.FgDarkGray,
		FgLightGreen : color.FgLightGreen,
		FgLightYellow : color.FgLightYellow,
		FgLightBlue : color.FgLightBlue,
		FgLightMagenta : color.FgLightMagenta,
		FgLightCyan : color.FgLightCyan,
		FgLightWhite : color.FgLightWhite,
		FgGray : color.FgGray,

		BgBlack : color.BgBlack,
		BgRed : color.BgRed,
		BgGreen : color.BgGreen,
		BgYellow : color.BgYellow,
		BgBlue : color.BgBlue,
		BgMagenta : color.BgMagenta,
		BgCyan : color.BgCyan,
		BgWhite : color.BgWhite,
		BgDefault : color.BgDefault,

		BgDarkGray : color.BgDarkGray,
		BgLightRed : color.BgLightRed,
		BgLightGreen : color.BgLightGreen,
		BgLightYellow : color.BgLightYellow,
		BgLightBlue : color.BgLightBlue,
		BgLightMagenta : color.BgLightMagenta,
		BgLightCyan : color.BgLightCyan,
		BgLightWhite : color.BgLightWhite,
		BgGray : color.BgGray,
	}
}

// NewColor create style
func NewColor(c ...color.Color) color.Style {
	return color.New(c...)
}

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

// PrintRed ???????????????
func PrintRed(val ...interface{}) {
	color.Red.Print(val...)
}

// PrintlnRed ????????????
func PrintlnRed(val ...interface{}) {
	color.Red.Println(val...)
}

// PrintfRed ??????-???????????????
func PrintfRed(format string, val ...interface{}) {
	color.Red.Printf(format, val...)
}

// PrintGreen ???????????????
func PrintGreen(val ...interface{}) {
	color.Green.Print(val...)
}

// PrintlnGreen ????????????
func PrintlnGreen(val ...interface{}) {
	color.Green.Println(val...)
}

// PrintfGreen ??????-???????????????
func PrintfGreen(format string, val ...interface{}) {
	color.Green.Printf(format, val...)
}

// PrintBlue ???????????????
func PrintBlue(val ...interface{}) {
	color.Blue.Print(val...)
}

// PrintlnBlue ????????????
func PrintlnBlue(val ...interface{}) {
	color.Blue.Println(val...)
}

// PrintfBlue ??????-???????????????
func PrintfBlue(format string, val ...interface{}) {
	color.Blue.Printf(format, val...)
}

// PrintYellow ???????????????
func PrintYellow(val ...interface{}) {
	color.Yellow.Print(val...)
}

// PrintlnYellow ????????????
func PrintlnYellow(val ...interface{}) {
	color.Yellow.Println(val...)
}

// PrintfYellow ??????-???????????????
func PrintfYellow(format string, val ...interface{}) {
	color.Yellow.Printf(format, val...)
}

// PrintGray ???????????????
func PrintGray(val ...interface{}) {
	color.Gray.Print(val...)
}

// PrintlnGray ????????????
func PrintlnGray(val ...interface{}) {
	color.Gray.Println(val...)
}

// PrintfGray ??????-???????????????
func PrintfGray(format string, val ...interface{}) {
	color.Gray.Printf(format, val...)
}

// PrintCyan ???????????????
func PrintCyan(val ...interface{}) {
	color.Cyan.Print(val...)
}

// PrintlnCyan ????????????
func PrintlnCyan(val ...interface{}) {
	color.Cyan.Println(val...)
}

// PrintfCyan ??????-???????????????
func PrintfCyan(format string, val ...interface{}) {
	color.Cyan.Printf(format, val...)
}
