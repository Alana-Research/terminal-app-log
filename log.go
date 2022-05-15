package terminalapplog

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(a ...any) {
	c := color.New(color.FgCyan)

	c.Printf("INFO: ")
	fmt.Println(a...)
}

func BigInfo(a ...any) {
	c := color.New(color.FgCyan)
	c.Add(color.Bold)
	c.Add(color.BgHiWhite)

	c.Printf("INFO: ")
	fmt.Println(a...)
}

func Warn(a ...any) {
	c := color.New(color.FgYellow)

	c.Printf("WARNING: ")
	fmt.Println(a...)
}

func BigWarn(a ...any) {
	c := color.New(color.FgCyan)
	c.Add(color.Bold)
	c.Add(color.BgHiYellow)

	c.Printf("WARNING: ")
	fmt.Println(a...)
}

func Error(a ...any) {
	c := color.New(color.FgRed)

	c.Printf("ERROR: ")
	fmt.Println(a...)
}

func BigError(a ...any) {
	c := color.New(color.FgWhite)
	c.Add(color.Bold)
	c.Add(color.BgHiRed)

	c.Printf("ERROR: ")
	fmt.Println(a...)
}
