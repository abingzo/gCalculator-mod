// Package base 用于保存一些全局变量和一些公共数据的地方
package base

import (
	"gCalculator-mod/bus"
	"os"
)

var (
	StdIn *bus.StdIn
	StdOut *bus.StdOut
	OsStdIn = os.Stdin
	OsStdOut = os.Stdout
)
