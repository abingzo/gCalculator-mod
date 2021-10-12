package task

import (
	"bytes"
	"gCalculator-mod/alg"
	"gCalculator-mod/base"
)

/*
	parse.go 负责把中缀表达式转换为后缀表达式
*/

// ToPostfixExp 从中缀表达式到后缀表达式
func ToPostfixExp(exp string) []string {
	// 没有括号的加上左括号和右括号
	if exp[0] != '(' || exp[len(exp)-1] != ')' {
		exp = string(append([]byte{'('}, append([]byte(exp), ')')...))
	}
	// 将所有支持的运算符展开
	tmpSupport := make([]string, 0)
	for k := range base.ShowTopLevel {
		tmpSupport = append(tmpSupport, base.ShowTopLevel[k]...)
	}
	support := make(map[string]int, 0)
	for _, v := range tmpSupport {
		support[v] = 1
	}
	// 分割算式
	// 分割后的算式
	var split []string
	buf := bytes.Buffer{}
	for i := 0; i < len(exp); i++ {
		buf.WriteString(string(exp[i]))
		// 运算符存在的情况
		if support[buf.String()] == 1 {
			split = append(split, buf.String())
			buf.Reset()
		}
	}
	tmp, _ := toPostfixExp(split, 0, alg.NewStack())
	return tmp
}

func toPostfixExp(split []string, i int, stack *alg.Stack) ([]string, int) {
	// 需要用到的栈
	// 统一入栈类型为string
	result := make([]string, 0, len(split))
	// TODO:使用有限状态机改写
	// 分割按照合法的运算符分割，并没有将数字合法化，所以数字需要处理
	for ; i < len(split); i++ {
		switch r := split[i]; base.TopLevel[r] {
		case -1:
			// -1表示没有级别，即为数字
			var bts []byte
			bts = append(bts, []byte(r)...)
			for i != len(split)-1 && base.TopLevel[split[i+1]] == -1 {
				i++
				bts = append(bts, []byte(split[i])...)
			}
			result = append(result, string(bts))
		case 4:
			// 4表示需要右括号的级别
			if r != ")" {
				stack.Push(r)
			}
			if split[i] != ")" {
				i++
				var tmp []string
				tmp, i = toPostfixExp(split, i, stack)
				// 统计总长
				result = append(result, tmp...)
				continue
			}
			if split[i] == ")" {
				for !stack.IsEmpty() && base.TopLevel[stack.Peek().(string)] != 0 {
					if tr := stack.Pop().(string); base.TopLevel[tr] == 4 {
						break
					} else {
						result = append(result, tr)
					}
				}
				return result, i
			}
		case 0:
			// 0 为不存在或不支持的运算符/数字
			continue
		default:
			// 其他则根据定义好的运算符优先级进行操作
			// 比栈顶的运算符优先级高则入栈
			if stack.IsEmpty() || base.TopLevel[stack.Peek().(string)] < base.TopLevel[r] || base.TopLevel[stack.Peek().(string)] == 4 {
				stack.Push(r)
			} else {
				// 否则出栈
				for !stack.IsEmpty() {
					if tr := stack.Peek().(string); base.TopLevel[tr] >= base.TopLevel[r] && base.TopLevel[tr] != 4 {
						result = append(result, stack.Pop().(string))
					} else {
						stack.Push(r)
						break
					}
				}
			}
		}
	}
	// 全部出栈
	for !stack.IsEmpty() {
		if r := stack.Pop().(string); base.TopLevel[r] != 4 {
			result = append(result, r)
		}
	}
	return result, i
}
