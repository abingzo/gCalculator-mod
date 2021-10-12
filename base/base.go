// Package base 用于保存一些全局变量和一些公共数据的地方
package base

import (
	"gCalculator-mod/bus"
	"os"
)

var (
	StdIn    *bus.StdIn
	StdOut   *bus.StdOut
	OsStdIn  = os.Stdin
	OsStdOut = os.Stdout
)

// ShowTopLevel 用于展示运算符的优先级
// -1 表示没有级别，比如0-9的数字
var ShowTopLevel = map[int][]string{
	-1: {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	4:  {"(", ")"},
	3:  {"^", "!", "ln", "lg", "<<", ">>", "π", "e", "sin(", "cos(", "tan("},
	2:  {"*", "/", "%"},
	1:  {"+", "-"},
}

// TopLevel 从展示运算符优先级的map中构造用于匹配的map
var TopLevel = func() map[string]int {
	tl := make(map[string]int)
	for k, v := range ShowTopLevel {
		for _, vv := range v {
			tl[vv] = k
		}
	}
	return tl
}()
