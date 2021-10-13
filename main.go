package main

import (
	"gCalculator-mod/base"
	"gCalculator-mod/stdio"
	"gCalculator-mod/task"
	"os"
)

func main() {
	base.OsStdOut = os.Stdout
	base.OsStdIn = os.Stdin
	// 错误恢复
	defer func(stdOut *os.File) {
		if err := recover(); err != nil {
			_, _ = stdOut.WriteString(err.(error).Error())
		}
	}(base.OsStdOut)
	// 创建一个goroutine检测信号
	// 检测输入和输出
	for {
		// 阻塞检测输入
		bytes, err := stdio.ReadStdIn()
		if err != nil {
			panic(err)
		}
		// 执行计算任务
		// 检测输出
		_ = stdio.WriteStdOut([]byte(task.NewCalculationTask(string(bytes))))
	}
}
