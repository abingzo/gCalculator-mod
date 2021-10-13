package task

/*
	对生成的后缀表达式进行计算
*/

import (
	"gCalculator-mod/alg"
	"gCalculator-mod/base"
)

func Call(handler string,num ...string) string {
	switch checkHandler(handler) {
	case 1:
		return base.Element[handler](num[0])
	case 2:
		return base.ElementDouble[handler](num[0],num[1])
	default:
		panic("handler is not defined")
	}
}

func checkHandler(handler string) int {
	if _,ok := base.ElementDouble[handler]; ok {
		return 2
	} else if _,ok := base.Element[handler]; ok {
		return 1
	} else {
		panic("handler is not defined")
	}
}

// 验证分割后的字符串对应的优先级类型
func checkTopLevelType(splitString string) int {
	if l := len(splitString); l == 1 {
		return base.TopLevel[splitString]
	} else {
		if r := base.TopLevel[string(splitString[len(splitString)-1])]; r == -1 {
			return r
		} else {
			return base.TopLevel[splitString]
		}
	}
}

func Calculate(split []string) string {
	stack := alg.NewStack()
	// 遍历处理好的逆波兰式
	for i := 0; i < len(split); i++ {
		switch r := split[i] ;checkTopLevelType(r) {
		case -1:
			// 数字则压栈
			stack.Push(split[i])
		case 1, 2, 3:
			// 表达式则计算
			if j := checkHandler(r); j == 1 {
				stack.Push(Call(r,stack.Pop().(string)))
			} else if j == 2 {
				n1 := stack.Pop().(string)
				stack.Push(Call(r,stack.Pop().(string),n1))
			}
		}
	}
	return stack.Pop().(string)
}