package math

import (
	"gCalculator-mod/alg"
	"strconv"
	"strings"
)

const (
	FLOAT numType = iota
	INTERGER
	// 正数
	PN int = 11
	// 负数
	NN int = 10
)

var table = map[byte]int8{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'-': int8(NN),
	'+': int8(PN),
}

// 序列化的反表
var reverseTable = map[int8]byte{
	0:        '0',
	1:        '1',
	2:        '2',
	3:        '3',
	4:        '4',
	5:        '5',
	6:        '6',
	7:        '7',
	8:        '8',
	9:        '9',
	int8(NN): '-',
	int8(PN): '+',
}

// BigNum 可伸缩大数字类型
type BigNum struct {
	// 数字的类型
	_type numType
	// 存储底层数据的类型
	// S|Integer|Float
	// 符号位|整数位|小数位
	data []int8
	// 小数数据
	// Integer,没有符号位
	pointData []int8
}

// 重置所有数据
func (b *BigNum) Reset() {
	b._type = 0
	b.data = nil
	b.pointData = nil
}

// 科学计数法的格式化
func (b *BigNum) sNotation(s string, ptr int) *BigNum {
	tmpN := s[:ptr]
	tmpS := s[ptr+1:]
	// 格式化的数量
	formatN := -1
	for i := 0; i < len(tmpN); i++ {
		if tmpN[i] == '.' {
			continue
		}
		b.data = append(b.data, table[tmpN[i]])
		formatN++
	}
	eN, err := strconv.Atoi(tmpS)
	if err != nil {
		panic(err)
	}
	eN = eN - formatN
	for i := 0; i <= eN; i++ {
		b.data = append(b.data, 0)
	}
	// 科学计数法无小数
	b._type = INTERGER
	return b
}

// 从字符串格式化为大数
func (b *BigNum) FromString(s string) *BigNum {
	b.Reset()
	if r := strings.Index(s, "e"); r > 0 {
		return b.sNotation(s, r)
	} else if r := strings.Index(s, "E"); r > 0 {
		return b.sNotation(s, r)
	}
	// 普通十进制格式化
	b.data = make([]int8, 1, len(s)+1)
	// 遍历字符串的初始偏移量
	offset := 0
	// 判断正负数
	if strings.HasPrefix(s, "-") {
		b.data[0] = int8(NN)
		offset += 1
	} else if strings.HasPrefix(s,"+") {
		offset += 1
		b.data[0] = int8(PN)
	} else {
		b.data[0] = int8(PN)
	}
	for i := offset; i < len(s); i++ {
		if s[i] == '.' {
			for j := i + 1; j < len(s); j++ {
				b.pointData = append(b.pointData, table[s[j]])
			}
			break
		}
		b.data = append(b.data, table[s[i]])
	}
	// 判断类型
	if len(b.pointData) > 0 {
		b._type = FLOAT
	} else {
		b._type = INTERGER
	}
	return b
}

// 格式化为字符串
func (b *BigNum) String() string {
	if b._type == FLOAT {
		return b.floatString()
	} else if b._type == INTERGER {
		return b.integerString()
	} else {
		return ""
	}
}

// 整数格式化
func (b *BigNum) integerString() string {
	lens := len(b.data)
	result := make([]byte, lens)
	for i := 0; i < lens; i++ {
		result[i] = reverseTable[b.data[i]]
	}
	return checkString(result)
}

// 小数格式化
func (b *BigNum) floatString() string {
	lens := len(b.data) + len(b.pointData) + 1
	result := make([]byte, lens)
	for k, v := range b.data {
		result[k] = reverseTable[v]
	}
	// 码上小数点
	result[len(b.data)] = '.'
	for i := len(b.data) + 1; i < lens; i++ {
		result[i] = reverseTable[b.pointData[i-(len(b.data)+1)]]
	}
	return checkString(result)
}

// 格式化验证
func checkString(p []byte) string {
	if p[0] == '+' {
		return string(p[1:])
	} else {
		return string(p)
	}
}

// 三元表达式的简单实现
func hit(bl bool, a, b interface{}) interface{} {
	if bl {
		return a
	} else {
		return b
	}
}

// 比较得出最大的数
// a,c 为nil时返回nil
func (b *BigNum) Max(a,c *BigNum) *BigNum {
	if a == nil && c == nil {
		return nil
	} else if a != nil && c == nil {
		return a
	} else if c != nil && a == nil {
		return c
	}
	if len(a.data) > len(c.data) {
		return a
	}  else if len(a.data) < len(c.data){
		return c
	} else if len(a.data) == len(c.data) {
		// 长度相同的情况下
		// 从大到小遍历
		for i := 0; i < len(a.data);i++ {
			if a.data[i] > c.data[i] {
				return a
			} else if a.data[i] < c.data[i] {
				return c
			}
		}
		// 如果整数位相等且任意一数为小数
		if c._type == FLOAT {
			return c
		} else if a._type == FLOAT {
			return a
		} else if a._type == c._type && a._type == INTERGER{
			// 都为整数且相等则返回第一个数
			return a
		}
		// 整数位相等且都为小数则比较小数位
		a2 := b.copy(a)
		c2 := b.copy(c)
		toData := func(p *BigNum) {
			p._type = INTERGER
			dst := []int8{int8(PN)}
			p.data = append(dst,p.pointData...)
		}
		toData(a2)
		toData(c2)
		// 递归比较
		tmp := b.Max(a2,c2)
		if tmp == a2 {
			return a
		} else {
			return c
		}
	}
	return nil
}

// 比较得出最小的数
func (b *BigNum) Min(a,c *BigNum) *BigNum {
	tmp := b.Max(a,c)
	if tmp == a {
		return c
	} else {
		return a
	}
}

// 绝对值
func (b *BigNum) Abs() *BigNum {
	c := b.copy(b)
	c.data[0] = int8(PN)
	return c
}

// 拷贝
func (b *BigNum) copy(p *BigNum) *BigNum {
	return &BigNum{
		_type:     p._type,
		data: func() []int8 {
			dst := make([]int8,len(p.data))
			copy(dst,p.data)
			return dst
		}(),
		pointData: func() []int8 {
			if p.pointData == nil {
				return nil
			}
			dst := make([]int8,len(p.pointData))
			copy(dst,p.pointData)
			return dst
		}(),
	}
}

// 加法递归运算结果
// 栈填充的内容统一为int
// TODO:解决与负数相加的BUG
func (b *BigNum) Add(a, c *BigNum) *BigNum {
	// 用于存储整数加法溢出的符号位
	tmpFlowI := alg.NewStack()
	// 用于存储小数加法溢出的符号位
	tmpFlowP := alg.NewStack()
	// 存储整数计算结果的栈
	integerResult := alg.NewStack()
	// 存储小数计算结果的栈
	pointResult := alg.NewStack()
	// 判断符号位
	// 正负相加
	if a.data[0] == int8(PN) && c.data[0] == int8(NN) {
		return b.Sub(a,c)
	} else if a.data[0] == c.data[0] && a.data[0] == int8(NN) {
		// 负负相加
		a2 := b.copy(a)
		c2 := b.copy(c)
		// 调整符号位
		a2.data[0] = int8(PN)
		c2.data[0] = int8(PN)
		nResult := b.Add(a2,c2)
		nResult.data[0] = int8(NN)
		return nResult
	} else if a.data[0] == int8(NN) && c.data[0] == int8(PN) {
		// 负正相加
		return b.Sub(c,a)
	}
	// 不同的类型采用不同的策略
	// 正正相加
	loop:
	if a._type == c._type && a._type == INTERGER {
		// 找出最大的数对其计算
		lenMax := hit(len(a.data) > len(c.data), len(a.data), len(c.data)).(int)
		var maxBigNum, minBigNum *BigNum
		if lenMax == len(c.data) {
			maxBigNum = c
			minBigNum = a
		} else {
			maxBigNum = a
			minBigNum = c
		}
		// 符号位不计算
		for i := 0; i < len(minBigNum.data) - 1; i++ {
			var r int8
			if !tmpFlowI.IsEmpty() {
				flow := tmpFlowI.Pop().(int)
				r = minBigNum.data[len(minBigNum.data) - 1 - i] + maxBigNum.data[len(maxBigNum.data) - 1 - i] + int8(flow)
			} else {
				r = minBigNum.data[len(minBigNum.data) - 1 - i] + maxBigNum.data[len(maxBigNum.data) - 1 - i]
			}
			if r >= 10 {
				tmpFlowI.Push(int(r / 10))
			}
			integerResult.Push(int(r % 10))
		}
		// 接下来计算没有对齐的数字
		for i := len(maxBigNum.data) - len(minBigNum.data); i > 0; i-- {
			var r int8
			if !tmpFlowI.IsEmpty() {
				flow := tmpFlowI.Pop().(int)
				r = maxBigNum.data[i] + int8(flow)
			} else {
				r = maxBigNum.data[i]
			}
			if r >= 10 {
				tmpFlowI.Push(int(r / 10))
			}
			integerResult.Push(int(r % 10))
		}
		// 如果结果还有溢出则需要进位
		if !tmpFlowI.IsEmpty() {
			integerResult.Push(tmpFlowI.Pop())
		}
		// 补上符号位
		integerResult.Push(PN)
	} else if a._type == c._type && a._type == FLOAT {
		// 找出最大长度的小数，方便后续验证溢出
		fLenMax := hit(len(a.pointData) > len(c.pointData),a,c).(*BigNum)
		fLenMin := hit(len(a.pointData) > len(c.pointData),c,a).(*BigNum)
		// 小数对齐计算
		if l := len(fLenMax.pointData) - len(fLenMin.pointData); l > 0 {
			for i := 0; i < l; i++ {
				fLenMin.pointData = append(fLenMin.pointData,0)
			}
		}
		var str1,str2 []byte
		for _,v := range fLenMax.pointData {
			str1 = append(str1,reverseTable[v])
		}
		for _,v := range fLenMin.pointData {
			str2 = append(str2,reverseTable[v])
		}
		var x,y BigNum
		fElement := b.Add(x.FromString(string(str1)),y.FromString(string(str2)))
		// 长度大于原来最长的数则代表有溢出
		// 忽略符号位
		fResult := fElement.data[1:]
		if len(fResult) > len(fLenMax.pointData) {
			// 将溢出结果压栈
			tmpFlowP.Push(int(fResult[0]))
			// 裁切结果
			fResult = fResult[1:]
		}
		// 把得到的结果压入小数栈
		for i := len(fResult) - 1; i >= 0; i-- {
			pointResult.Push(int(fResult[i]))
		}
		// 计算整数
		d1 := make([]int8,len(a.data))
		d2 := make([]int8,len(c.data))
		copy(d1,a.data)
		copy(d2,c.data)
		iResult := b.Add(&BigNum{_type: INTERGER,data: d1},&BigNum{_type: INTERGER,data: d2}).data
		// 将得到的整数结果压入栈
		for i := len(iResult) - 1; i >= 0; i-- {
			integerResult.Push(int(iResult[i]))
		}
	} else if a._type != c._type {
		// 找到其中为小数的类型
		if a._type == FLOAT {
			// 临时修改，以满足跳转条件
			c._type = FLOAT
			defer func(bn *BigNum) {
				bn._type = INTERGER
			}(c)
			c.pointData = append(c.pointData,0)
		} else if c._type == FLOAT {
			// 临时修改，以满足跳转条件
			a._type = FLOAT
			defer func(bn *BigNum) {
				bn._type = INTERGER
			}(a)
			a.pointData = append(a.pointData,0)
		}
		// 将小数对其并运算
		// 设置好结果并运算
		goto loop
	}
	// 拼接成BigNum类型
	bytes := make([]byte, 0)
	for !integerResult.IsEmpty() {
		r := integerResult.Pop().(int)
		bytes = append(bytes, reverseTable[int8(r)])
	}
	// 如果小数有溢出
	element := &BigNum{}
	if !tmpFlowP.IsEmpty() {
		var x, y BigNum
		element = b.Add(x.FromString(string(bytes)), y.FromString(string(reverseTable[int8(tmpFlowP.Pop().(int))])))
	} else {
		x := BigNum{}
		element = x.FromString(string(bytes))
	}
	// 取出小数的值,并修改数的类型
	for !pointResult.IsEmpty() {
		element._type = FLOAT
		element.pointData = append(element.pointData, int8(pointResult.Pop().(int)))
	}
	return element
}

// 减法
func (b *BigNum) Sub(a,c *BigNum) *BigNum {
	// 判断符号位，以便优化
	// 负正相减，变换符号位
	if a.data[0] != c.data[0] && a.data[0] == int8(NN) {
		a.data[0] = int8(PN)
		defer func(p *BigNum) {
			p.data[0] = int8(NN)
		}(a)
		tmp := b.Add(a,c)
		tmp.data[0] = int8(NN)
		return tmp
	} else if a.data[0] != c.data[0] && a.data[0] == int8(PN) {
		// 正减负
		c.data[0] = int8(PN)
		defer func(p *BigNum) {
			p.data[0] = int8(NN)
		}(c)
		return b.Add(a,c)
	}
	// 存储整数计算结果的栈
	integerResult := alg.NewStack()
	// 正整数减法
	if a._type == c._type && a._type == INTERGER && a.data[0] == c.data[0] && a.data[0] == int8(PN) {
		// 找出最大的数
		max := b.Max(a,c)
		var min *BigNum
		if max == c {
			min = a
		} else {
			min = c
		}
		minuend := b.copy(max)
		for i := 1; i < len(min.data); i++{
			// 小于被减数则借位
			if r := minuend.data[len(minuend.data) - i]; r < min.data[len(min.data) - i] {
				minuend.data[len(minuend.data) - (i+1)] -= 1
				r += 10
				integerResult.Push(int(r - min.data[len(min.data) - i]))
			}  else {
				integerResult.Push(int(r - min.data[len(min.data) - i]))
			}
		}
		// 长度不相等则补上剩余的数字
		// 忽略符号位
		// 扫描到小于0的数字要借位
		for i := len(minuend.data) - len(min.data); i > 0; i-- {
			if r := int(minuend.data[i]); r < 0 {
				minuend.data[i - 1] -= 1
				integerResult.Push(r + 10)
			} else {
				integerResult.Push(r)
			}
		}
		// 去除首位零及补充符号位
		for !integerResult.IsEmpty() && integerResult.Peek().(int) == 0 {
			integerResult.Pop()
		}
		// 根据被减数和减数大小关系补充符号位
		if min ==  a {
			integerResult.Push(NN)
		} else {
			integerResult.Push(PN)
		}
	} else if a._type == c._type && a._type == FLOAT && a.data[0] == c.data[0] && a.data[0] == int8(PN) {
		// 正小数减法
		// 分解步骤
		/*
			小数部分的运算 --> 整数部分的运算
			运算之后的整数 + 运算之后的小数
		*/
		a2 := b.copy(a)
		c2 := b.copy(c)
		iResult := b.Sub(&BigNum{_type: INTERGER,data: a2.data},&BigNum{_type: INTERGER,data: c2.data})
		// 补上符号位
		// 小数对齐
		maxFloatLen := a2
		if len(c2.pointData) > len(a2.pointData) {
			maxFloatLen = c2
		}
		fResult := b.Sub(&BigNum{
			_type:     INTERGER,
			data:      append([]int8{int8(PN)},append(a2.pointData,make([]int8,len(maxFloatLen.pointData) - len(a2.pointData))...)...),
		},&BigNum{
			_type:     INTERGER,
			data:      append([]int8{int8(PN)},append(c2.pointData,make([]int8,len(maxFloatLen.pointData) - len(c2.pointData))...)...),
		})
		// 恢复为浮点数
		fResult.pointData = append(fResult.pointData,fResult.data[1:]...)
		fResult.data = append(fResult.data[:1],0)
		fResult._type = FLOAT
		// add
		return b.Add(iResult,fResult)
	} else if a._type != c._type {
		// 类型不相等时
		// 对齐对应的浮点类型
		// 涉及到修改时不破坏原始类型
		a2 := b.copy(a)
		c2 := b.copy(c)
		if a2._type == FLOAT {
			c2._type = FLOAT
			c2.pointData = make([]int8,len(a2.pointData))
		} else if c._type == FLOAT {
			a2._type = FLOAT
			a2.pointData = make([]int8,len(c2.pointData))
		}
		return b.Sub(a2,c2)
	}
	// 序列化
	// 拼接成BigNum类型
	bytes := make([]byte, 0)
	for !integerResult.IsEmpty() {
		r := integerResult.Pop().(int)
		bytes = append(bytes, reverseTable[int8(r)])
	}
	var x BigNum
	return x.FromString(string(bytes))
}

