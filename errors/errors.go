// Package errors 集中定义错误的包
package errors

import "errors"

var (
	StdInPipeNotExist  = errors.New("stdIn pipe not exist")
	StdOutPipeNotExist = errors.New("stdOut pipe not exist")
)
