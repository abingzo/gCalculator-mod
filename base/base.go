// Package base 用于保存一些全局变量和一些公共数据的地方
package base

import (
	"gCalculator-mod/alg/math"
	"os"
)

var (
	OsStdIn  = os.Stdin
	OsStdOut = os.Stdout
)

// ShowTopLevel 用于展示运算符的优先级
// -1 表示没有级别，比如0-9的数字
var ShowTopLevel = map[int][]string{
	-1: {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9","."},
	4:  {"(", ")"},
	3:  {"^", "!", "ln", "lg", "<<", ">>", "π", "e", "sin(", "cos(", "tan("},
	2:  {"*", "/", "%"},
	1:  {"+", "-"},
}

type ElementBindDouble map[string]math.HandlerDouble

type ElementBind map[string]math.Handler

// ElementDouble 对符号提供计算支持的函数,双参数
var ElementDouble = ElementBindDouble{
	"+":  math.NewStep().Add,
	"-":  math.NewStep().Sub,
	"*":  math.NewStep().Ride,
	"/":  math.NewStep().Except,
	"%":  math.NewStep().Mod,
	"^":  math.NewStep().Power,
	"<<": math.NewStep().LeftShift,
	">>": math.NewStep().RightShift,
}

// Element 对符号提供计算支持的函数，单参数
var Element = ElementBind {
	"sin(": math.NewStep().Sin,
	"cos(": math.NewStep().Cos,
	"tan(": math.NewStep().Tan,
	"π":    math.NewStep().Pi,
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
