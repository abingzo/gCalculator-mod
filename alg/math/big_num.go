package math

import (
	"gCalculator-mod/alg"
	"strconv"
	"strings"
)

const (
	// FLOAT 小数
	FLOAT NumType = iota
	// INTEGER 整数
	INTEGER
	// PN 正数
	PN int = 11
	// NN 负数
	NN int = 10
	// DivisionAccuracy 除法精度
	DivisionAccuracy int = 20
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
	_type NumType
	// 存储底层数据的类型
	// S|Integer|Float
	// 符号位|整数位|小数位
	data []int8
	// 小数数据
	// Integer,没有符号位
	pointData []int8
}

// Reset 重置所有数据
func (b *BigNum) Reset() {
	b._type = 0
	b.data = nil
	b.pointData = nil
}

// Type 返回对应的类型
func (b *BigNum) Type() NumType {
	return b._type
}

/*
	手动伸缩类型
*/
// 伸缩为Int意味将抛弃所有的小数位
func (b *BigNum) ToInteger() *BigNum {
	b._type = INTEGER
	b.pointData = nil
	return b
}

// ToFloat 伸缩为Float意味将当前数设置为值相等的小数
// 本身为FLOAT则返回本身
func (b *BigNum) ToFloat() *BigNum {
	if b._type == FLOAT {
		return b
	}
	b._type = FLOAT
	b.pointData = []int8{0}
	return b
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
	b._type = INTEGER
	return b
}

// FromString 从字符串格式化为大数
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
	} else if strings.HasPrefix(s, "+") {
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
		b._type = INTEGER
	}
	return b
}

// 格式化为字符串
func (b *BigNum) String() string {
	if b._type == FLOAT {
		b.data = sliceDeleteLeftZero(b.data)
		b.pointData = sliceDeleteRightZero(b.pointData)
		return b.floatString()
	} else if b._type == INTEGER {
		b.data = sliceDeleteLeftZero(b.data)
		return b.integerString()
	}
	return ""
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
// 小数的值为零则自动伸缩为整数类型,如:2.00 -> 2
func (b *BigNum) floatString() string {
	if b.EQ(b,&BigNum{_type: INTEGER,data: append(b.data)}) {
		b.ToInteger()
		return b.integerString()
	}
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

// 找出BigNum中的非nil
// 都为nil则返回nil
// 都不为nil则返回nil
func checkNil(a, c *BigNum) *BigNum {
	if a == nil && c == nil {
		return nil
	} else if a != nil && c == nil {
		return a
	} else if c != nil && a == nil {
		return c
	} else {
		return nil
	}
}

// 均为整数?
// 判断是否均为整数切片
func isInteger(a, c []int8) bool {
	return a[0] == int8(PN) && c[0] == int8(PN)
}

// 均为小数？
func isFloat(a, c []int8) bool {
	return a[0] <= 9 && c[0] <= 9
}

// 忽略符号位
// 值相等的前提下，删除无意义的0值
// 去除左边的零
// 只有一个零的时候不去除
func sliceDeleteLeftZero(a []int8) []int8 {
	aDst := make([]int8, 0, len(a))
	if (len(a) == 2 || len(a) == 1) && a[len(a)-1] == 0{
		return append(aDst,a...)
	}
	// 有符号位的情况
	ptr := 0
	if a[0] == int8(PN) || a[0] == int8(NN) {
		aDst = append(aDst, a[0])
		ptr = 1
	} else {
		// 无符号位
		ptr = 0
	}
	for i := 1; i < len(a); i++ {
		if r := a[i]; r != 0 {
			break
		} else if r == 0 {
			ptr++
		}
	}
	aDst = append(aDst, a[ptr:]...)
	return aDst
}

// 忽略符号位
// 值相等的前提下，删除无意义的0值
// 去除右边的零
func sliceDeleteRightZero(a []int8) []int8 {
	aDst := make([]int8, 0, len(a))
	ptr := len(a)
	for i := len(a) - 1; i >= 0; i-- {
		if r := a[i]; r != 0 {
			break
		} else if r == 0 {
			ptr--
		}
	}
	// 保证至少留下一个零
	if ptr == 0 && len(a) > 0 {
		ptr++
	}
	aDst = append(aDst, a[:ptr]...)
	return aDst
}

// 比较切片的数值大小 a > c,并返回true或者false
func sliceValueGT(a, c []int8) bool {
	if c == nil && a != nil {
		return true
	}
	if a == nil && c != nil {
		return false
	}
	// 判断为整数还是小数
	if isInteger(a, c) {
		a, c = sliceDeleteLeftZero(a), sliceDeleteLeftZero(c)
		// 整数比较则对齐左零
		maxLen := hit(len(a) > len(c), len(a), len(c)).(int)
		a = append(a[:1], append(make([]int8, maxLen-len(a)), a[1:]...)...)
		c = append(c[:1], append(make([]int8, maxLen-len(c)), c[1:]...)...)
	} else if isFloat(a, c) {
		a, c = sliceDeleteRightZero(a), sliceDeleteRightZero(c)
		// 在小数比较值大小前对齐长度，防止越界
		// 小数则对齐右零
		maxLen := hit(len(a) > len(c), len(a), len(c)).(int)
		a = append(a, make([]int8, maxLen-len(a))...)
		c = append(c, make([]int8, maxLen-len(c))...)
	}
	// TODO:不能取巧判断长度来判断值得大小，因为可能有02和000009这类小数值
	// 从大到小遍历
	for i := 0; i < len(a); i++ {
		if a[i] > c[i] {
			return true
		} else if a[i] < c[i] {
			return false
		}
	}
	return false
}

// 比较切片的数值是否相等
// TODO:会忽略整数开始的零和小数后面的零
func sliceValueEQ(a, c []int8) bool {
	// 判断为整数还是小数
	if isInteger(a, c) {
		a, c = sliceDeleteLeftZero(a), sliceDeleteLeftZero(c)
	} else if isFloat(a, c) {
		a, c = sliceDeleteRightZero(a), sliceDeleteRightZero(c)
	}
	if len(a) != len(c) {
		return false
	}
	for k, v := range a {
		if v != c[k] {
			return false
		}
	}
	return true
}

// Max 比较得出最大的数
// a,c 为nil时返回nil
// 相等时返回第一个数
func (b *BigNum) Max(a, c *BigNum) *BigNum {
	if a == nil || c == nil {
		return checkNil(a, c)
	}
	if b.EQ(a, c) {
		return a
	}
	return hit(b.GT(a, c), a, c).(*BigNum)
}

// Min 比较得出最小的数
func (b *BigNum) Min(a, c *BigNum) *BigNum {
	tmp := b.Max(a, c)
	if tmp == a {
		return c
	} else {
		return a
	}
}

/*
	比较运算符
*/
// 值等于/大小等于
// NOTE: 8.0 == 8 (9 == 9.00000) ( 17.35 == 17.3500)
func (b *BigNum) EQ(a, c *BigNum) bool {
	if a._type == c._type && a._type == INTEGER {
		return sliceValueEQ(a.data, c.data)
	}
	if a._type == c._type && a._type == FLOAT {
		return sliceValueEQ(a.data, c.data) && sliceValueEQ(a.pointData, c.pointData)
	} else if a._type != c._type {
		// 一方为小数时比较大小
		if !sliceValueEQ(a.data, c.data) {
			return false
		}
		if a._type == FLOAT {
			c = b.copy(c)
			c._type = FLOAT
			c.pointData = make([]int8, len(a.pointData))
		} else if c._type == FLOAT {
			a = b.copy(a)
			a._type = FLOAT
			a.pointData = make([]int8, len(c.pointData))
		}
		return hit(sliceValueEQ(a.pointData, c.pointData), true, false).(bool)
	}
	return false
}

// NE 不等于
func (b *BigNum) NE(a, c *BigNum) bool {
	return !b.EQ(a, c)
}

// GT 大于
func (b *BigNum) GT(a, c *BigNum) bool {
	if a._type == c._type && a._type == INTEGER {
		return sliceValueGT(a.data, c.data)
	} else if a._type == FLOAT || c._type == FLOAT {
		// 小数的比较大小的bool值要取反
		return hit(sliceValueEQ(a.data, c.data), sliceValueGT(a.pointData, c.pointData), sliceValueGT(a.data, c.data)).(bool)
	} else if a._type != c._type {
		// 一方为小数时比较大小
		if !sliceValueEQ(a.data, c.data) {
			return sliceValueGT(a.data, c.data)
		}
		if a._type == FLOAT {
			c = b.copy(c)
			c._type = FLOAT
			c.pointData = make([]int8, len(a.pointData))
		} else if c._type == FLOAT {
			a = b.copy(a)
			a._type = FLOAT
			a.pointData = make([]int8, len(c.pointData))
		}
		return hit(sliceValueEQ(a.pointData, c.pointData), false, sliceValueGT(a.pointData, c.pointData)).(bool)
	}
	return false
}

// LT 小于
func (b *BigNum) LT(a, c *BigNum) bool {
	// 防止等于
	if b.EQ(a, c) {
		return false
	}
	return !b.GT(a, c)
}

// GE 大于等于
func (b *BigNum) GE(a, c *BigNum) bool {
	return b.GT(a, c) || b.EQ(a, c)
}

// LE 小于等于
func (b *BigNum) LE(a, c *BigNum) bool {
	return b.LT(a, c) || b.EQ(a, c)
}

// 数字存储的真实长度
// 不包括符号位和小数点
func (b *BigNum) len() int {
	if b._type == INTEGER {
		return len(b.data) - 1
	} else if b._type == FLOAT {
		return len(b.data) - 1 + len(b.pointData)
	} else {
		return 0
	}
}

// Abs 绝对值
func (b *BigNum) Abs() *BigNum {
	c := b.copy(b)
	c.data[0] = int8(PN)
	return c
}

// 拷贝
func (b *BigNum) copy(p *BigNum) *BigNum {
	return &BigNum{
		_type: p._type,
		data: func() []int8 {
			dst := make([]int8, len(p.data))
			copy(dst, p.data)
			return dst
		}(),
		pointData: func() []int8 {
			if p.pointData == nil {
				return nil
			}
			dst := make([]int8, len(p.pointData))
			copy(dst, p.pointData)
			return dst
		}(),
	}
}

// Add 加法递归运算结果
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
		c2 := b.copy(c)
		c2.data[0] = int8(PN)
		return b.Sub(a, c2)
	} else if a.data[0] == c.data[0] && a.data[0] == int8(NN) {
		// 负负相加
		a2 := b.copy(a)
		c2 := b.copy(c)
		// 调整符号位
		a2.data[0] = int8(PN)
		c2.data[0] = int8(PN)
		nResult := b.Add(a2, c2)
		nResult.data[0] = int8(NN)
		return nResult
	} else if a.data[0] == int8(NN) && c.data[0] == int8(PN) {
		// 负正相加
		a2 := b.copy(a)
		a2.data[0] = int8(PN)
		return b.Sub(c, a2)
	}
	// 不同的类型采用不同的策略
	// 正正相加
loop:
	if a._type == c._type && a._type == INTEGER {
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
		for i := 0; i < len(minBigNum.data)-1; i++ {
			var r int8
			if !tmpFlowI.IsEmpty() {
				flow := tmpFlowI.Pop().(int)
				r = minBigNum.data[len(minBigNum.data)-1-i] + maxBigNum.data[len(maxBigNum.data)-1-i] + int8(flow)
			} else {
				r = minBigNum.data[len(minBigNum.data)-1-i] + maxBigNum.data[len(maxBigNum.data)-1-i]
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
		fLenMax := hit(len(a.pointData) > len(c.pointData), a, c).(*BigNum)
		fLenMin := hit(len(a.pointData) > len(c.pointData), c, a).(*BigNum)
		// 小数对齐计算
		if l := len(fLenMax.pointData) - len(fLenMin.pointData); l > 0 {
			for i := 0; i < l; i++ {
				fLenMin.pointData = append(fLenMin.pointData, 0)
			}
		}
		var str1, str2 []byte
		for _, v := range fLenMax.pointData {
			str1 = append(str1, reverseTable[v])
		}
		for _, v := range fLenMin.pointData {
			str2 = append(str2, reverseTable[v])
		}
		var x, y BigNum
		fElement := b.Add(x.FromString(string(str1)), y.FromString(string(str2)))
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
		d1 := make([]int8, len(a.data))
		d2 := make([]int8, len(c.data))
		copy(d1, a.data)
		copy(d2, c.data)
		iResult := b.Add(&BigNum{_type: INTEGER, data: d1}, &BigNum{_type: INTEGER, data: d2}).data
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
				bn._type = INTEGER
			}(c)
			c.pointData = append(c.pointData, 0)
		} else if c._type == FLOAT {
			// 临时修改，以满足跳转条件
			a._type = FLOAT
			defer func(bn *BigNum) {
				bn._type = INTEGER
			}(a)
			a.pointData = append(a.pointData, 0)
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

// Sub 减法
func (b *BigNum) Sub(a, c *BigNum) *BigNum {
	// 判断符号位，以便优化
	// 负正相减，变换符号位
	if a.data[0] != c.data[0] && a.data[0] == int8(NN) {
		a.data[0] = int8(PN)
		defer func(p *BigNum) {
			p.data[0] = int8(NN)
		}(a)
		tmp := b.Add(a, c)
		tmp.data[0] = int8(NN)
		return tmp
	} else if a.data[0] != c.data[0] && a.data[0] == int8(PN) {
		// 正减负
		c.data[0] = int8(PN)
		defer func(p *BigNum) {
			p.data[0] = int8(NN)
		}(c)
		return b.Add(a, c)
	} else if a.data[0] == c.data[0] && a.data[0] == int8(NN) {
		// 负负相减
		c2 := b.copy(c)
		c2.data[0] = int8(PN)
		return b.Add(a, c2)
	}
	// 存储整数计算结果的栈
	integerResult := alg.NewStack()
	// 存储小数计算结果的栈
	pointResult := alg.NewStack()
	// 正整数减法
	if a._type == c._type && a._type == INTEGER && a.data[0] == c.data[0] && a.data[0] == int8(PN) {
		// 找出最大的数
		max := b.Max(a, c)
		var min *BigNum
		if max == c {
			min = a
		} else {
			min = c
		}
		minuend := b.copy(max)
		for i := 1; i < len(min.data); i++ {
			// 小于被减数则借位
			if r := minuend.data[len(minuend.data)-i]; r < min.data[len(min.data)-i] {
				minuend.data[len(minuend.data)-(i+1)] -= 1
				r += 10
				integerResult.Push(int(r - min.data[len(min.data)-i]))
			} else {
				integerResult.Push(int(r - min.data[len(min.data)-i]))
			}
		}
		// 长度不相等则补上剩余的数字
		// 忽略符号位
		// 扫描到小于0的数字要借位
		for i := len(minuend.data) - len(min.data); i > 0; i-- {
			if r := int(minuend.data[i]); r < 0 {
				minuend.data[i-1] -= 1
				integerResult.Push(r + 10)
			} else {
				integerResult.Push(r)
			}
		}
		// 长度大于1时去除首位零及补充符号位
		for !integerResult.IsEmpty() && integerResult.Peek().(int) == 0 && integerResult.Len() > 1 {
			integerResult.Pop()
		}
		// 根据被减数和减数大小关系补充符号位
		if min == a {
			integerResult.Push(NN)
		} else {
			integerResult.Push(PN)
		}
	} else if a._type == c._type && a._type == FLOAT && a.data[0] == c.data[0] && a.data[0] == int8(PN) {
		// 正小数减法
		// 分解步骤
		//TODO:不依赖加法接口
		/*
			强制被减数大于减数
			小数部分的运算 --> 整数部分的运算
			运算之后的整数 + 运算之后的小数
		*/
		// 找出最大的数，以便可以借位
		max := b.Max(a, c)
		min := b.Min(a, c)
		// 被减数
		minuend := b.copy(max)
		sub := min
		// 小数要对齐长度
		if len(minuend.pointData) > len(sub.pointData) {
			sub.pointData = append(sub.pointData, make([]int8, len(minuend.pointData)-len(sub.pointData))...)
		} else if len(minuend.pointData) < len(sub.pointData) {
			minuend.pointData = append(minuend.pointData, make([]int8, len(sub.pointData)-len(minuend.pointData))...)
		}
		// 遍历相减
		for i := len(minuend.pointData) - 1; i >= 0; i-- {
			if r := minuend.pointData[i] - sub.pointData[i]; r < 0 {
				// 借位,小数位没位置则向整数位借
				if i-1 < 0 {
					minuend.data[len(minuend.data)-1] -= 1
				} else {
					minuend.pointData[i-1] -= 1
				}
				pointResult.Push(int(r + 10))
			} else {
				pointResult.Push(int(r))
			}
		}
		// 整数相减
		iResult := b.Sub(&BigNum{_type: INTEGER, data: minuend.data}, &BigNum{_type: INTEGER, data: min.data})
		// 根据最小数字确认符号位
		if a == min {
			iResult.data[0] = int8(NN)
		} else {
			iResult.data[0] = int8(PN)
		}
		// 将结果压栈
		for i := len(iResult.data) - 1; i >= 0; i-- {
			integerResult.Push(int(iResult.data[i]))
		}
	} else if a._type != c._type {
		// 类型不相等时
		// 对齐对应的浮点类型
		// 涉及到修改时不破坏原始类型
		a2 := b.copy(a)
		c2 := b.copy(c)
		if a2._type == FLOAT {
			c2._type = FLOAT
			c2.pointData = make([]int8, len(a2.pointData))
		} else if c._type == FLOAT {
			a2._type = FLOAT
			a2.pointData = make([]int8, len(c2.pointData))
		}
		return b.Sub(a2, c2)
	}
	// 序列化
	// 拼接成BigNum类型
	element := &BigNum{}
	if !integerResult.IsEmpty() {
		element._type = INTEGER
	}
	if !pointResult.IsEmpty() {
		element._type = FLOAT
	}
	for !integerResult.IsEmpty() {
		element.data = append(element.data, int8(integerResult.Pop().(int)))
	}
	for !pointResult.IsEmpty() {
		element.pointData = append(element.pointData, int8(pointResult.Pop().(int)))
	}
	return element
}

// Ride 乘法
/*
	单乘多 -> 多乘多 -> 小数多乘多
*/
func (b *BigNum) Ride(a, c *BigNum) *BigNum {
	// 存储整数计算结果的栈
	integerResult := alg.NewStack()
	// 存储小数计算结果的栈
	pointResult := alg.NewStack()
	// 正整数单乘多
	if a._type == c._type && a._type == INTEGER && (a.len() == 1 || c.len() == 1) {
		var multiplicand *BigNum
		var multiplier int8
		if a.len() == 1 {
			multiplicand = c
			multiplier = a.data[len(a.data)-1]
		} else if c.len() == 1 {
			multiplier = c.data[len(c.data)-1]
			multiplicand = a
		}
		// 结果集,用于竖式对齐
		resultSet := alg.NewStack()
		for i := len(multiplicand.data) - 1; i > 0; i-- {
			// 相乘
			r := multiplicand.data[i] * multiplier
			var resultData []int8
			if r > 9 {
				resultData = []int8{r / 10, r % 10}
			} else {
				resultData = []int8{r}
			}
			// 补零
			resultData = append(resultData, make([]int8, multiplicand.len()-i)...)
			resultSet.Push(&BigNum{_type: INTEGER, data: append([]int8{int8(PN)}, resultData...)})
		}
		// 相加并把结果压栈
		// 取出来并逐个相加
		addResult := &BigNum{}
		// 创建哨兵条件，避免多次if导致的流水线
		if !resultSet.IsEmpty() {
			addResult = resultSet.Pop().(*BigNum)
		}
		for !resultSet.IsEmpty() {
			r := resultSet.Pop().(*BigNum)
			addResult = b.Add(addResult, r)
		}
		// 将最后的结果压栈
		// 忽略符号位
		for i := len(addResult.data) - 1; i > 0; i-- {
			integerResult.Push(int(addResult.data[i]))
		}
	} else if a._type == c._type && a._type == INTEGER && a.len() > 1 && c.len() > 1 {
		// 整数多乘多
		multiplicand := b.Max(a, c)
		multiplier := b.Min(a, c)
		// 结果集
		resultSet := alg.NewStack()
		// 逐个相乘并补上零之后放入结果集
		// 符号位不计算
		for i := len(multiplier.data) - 1; i > 0; i-- {
			r := b.Ride(&BigNum{_type: INTEGER, data: []int8{int8(PN), multiplier.data[i]}}, multiplicand)
			// 补零
			r.data = append(r.data, make([]int8, multiplier.len()-i)...)
			// 压栈
			resultSet.Push(r)
		}
		// 取出来并逐个相加
		addResult := &BigNum{}
		// 创建哨兵条件，避免多次if导致的流水线
		if !resultSet.IsEmpty() {
			addResult = resultSet.Pop().(*BigNum)
		}
		for !resultSet.IsEmpty() {
			r := resultSet.Pop().(*BigNum)
			addResult = b.Add(addResult, r)
		}
		// 将最后的结果压栈
		// 忽略符号位
		for i := len(addResult.data) - 1; i > 0; i-- {
			integerResult.Push(int(addResult.data[i]))
		}
	} else if a._type == c._type && a._type == FLOAT {
		// 均为小数相乘的情况
		// 将小数化为整数
		a2 := b.copy(a)
		c2 := b.copy(c)
		// 记录化整的偏移量
		offset := len(a2.pointData) + len(c2.pointData)
		result := b.Ride(&BigNum{_type: INTEGER, data: append(a2.data, a2.pointData...)},
			&BigNum{_type: INTEGER, data: append(c2.data, c2.pointData...)})
		// 根据偏移量恢复小数
		// 将小数结果压栈
		for i := len(result.data) - 1; len(result.data)-offset <= i; i-- {
			pointResult.Push(int(result.data[i]))
		}
		// 将整数结果压栈
		for i := len(result.data) - 1 - offset; i > 0; i-- {
			integerResult.Push(int(result.data[i]))
		}
	} else if a._type != c._type && (a._type == FLOAT || c._type == FLOAT) {
		// 一方为小数相乘的情况
		// 注意非小数*小数与小数*小数有很大区别
		// 找出为小数的一方
		if a._type == FLOAT {
			c2 := b.copy(c)
			c2._type = FLOAT
			c2.pointData = []int8{0}
			return b.Ride(a, c2)
		} else if c._type == FLOAT {
			a2 := b.copy(a)
			a2._type = FLOAT
			a2.pointData = []int8{0}
			return b.Ride(a2, c)
		}
	}
	// 序列化
	// 拼接成BigNum类型
	element := &BigNum{}
	if !integerResult.IsEmpty() {
		element._type = INTEGER
	}
	if !pointResult.IsEmpty() {
		element._type = FLOAT
	}
	// 负数的符号位确定比较简单，所以可以在后面确定
	// 负负为正|有一负均为负|正正为正
	if a.data[0] == c.data[0] && a.data[0] == int8(NN) {
		element.data = append(element.data, int8(PN))
	} else if a.data[0] == c.data[0] && a.data[0] == int8(PN) {
		element.data = append(element.data, int8(PN))
	} else if a.data[0] == int8(NN) || c.data[0] == int8(NN) {
		element.data = append(element.data, int8(NN))
	}
	for !integerResult.IsEmpty() {
		element.data = append(element.data, int8(integerResult.Pop().(int)))
	}
	for !pointResult.IsEmpty() {
		element.pointData = append(element.pointData, int8(pointResult.Pop().(int)))
	}
	return element
}

// Except 除法
/*
	整数除法 -> 小数除法
*/
func (b *BigNum) Except(a, c *BigNum) *BigNum {
	// 处理负数，有任一数为负数则结果为负数
	// 负负为正
	if a.data[0] == c.data[0] && a.data[0] == int8(NN) {
		a2, c2 := b.copy(a), b.copy(c)
		a2.data[0] = int8(PN)
		c2.data[0] = int8(PN)
		return b.Except(a2, c2)
	} else if a.data[0] == int8(NN) || c.data[0] == int8(NN) {
		a2, c2 := b.copy(a), b.copy(c)
		a2.data[0] = int8(PN)
		c2.data[0] = int8(PN)
		tmp := b.Except(a2, c2)
		tmp.data[0] = int8(NN)
		return tmp
	}
	// 存储整数计算结果的栈
	integerResult := alg.NewStack()
	// 存储小数计算结果的栈
	pointResult := alg.NewStack()

	if a._type == c._type && a._type == INTEGER {
		// 临时整数结果存储
		tmpIntResult := make([]int, 0)
		tmpFloatResult := make([]int, 0)
		// 创建哨兵条件
		// 找出一个比被除数要大的数
		// 被除数比除数大则借位
		dividend := b.copy(a)
		divisor := b.copy(c)
		offset := dividend.len()
		// 记录匹配点，用于补零
		// 初始为1，防止首位小于的情况
		point := []int{1}
		// 匹配点的向量,从第一个匹配的数开始
		pointIV := 0
		// 从左向右除法未补零的原生长度
		nativeLen := 0
		// 一些函数
		appendData := func(d int8) {
			// 偏移量大于原始的被除数则填入小数结果集
			if offset > a.len() {
				tmpFloatResult = append(tmpFloatResult, int(d))
			} else {
				tmpIntResult = append(tmpIntResult, int(d))
			}
		}
		// 校验匹配点
		checkPoint := func() {
			// 比较两次的匹配点
			// 匹配点有间隔的数则补上缺失的零
			// 对齐匹配点
			if p, j := point[len(point)-1], point[len(point)-2]; p-j >= 2 {
				for j := p - j - 1; j > 0; j-- {
					// 解决残留的问题，比如21/2导致的补零之后的结果为1.05
					if offset - a.len() == 1 && point[len(point)-1] - point[len(point)-2] >= 2 {
						tmpIntResult = append(tmpIntResult,0)
					} else {
						appendData(0)
					}
				}
			}
		}
		// 初始除数大于被除数的情况
		for i:= 0;b.GT(divisor, dividend);i++ {
			//if len(tmpIntResult) == 0 {
			//	offset++
			//	pointIV--
			//	tmpIntResult = append(tmpIntResult, 0)
			//}
			//if r := divisor.len() - dividend.len() - 1 + i; r > 0 {
			//	offset += r
			//	pointIV -= r
			//	tmpFloatResult = append(tmpFloatResult,make([]int,r)...)
			//}
			dividend.data = append(dividend.data, make([]int8, divisor.len()-dividend.len()+i)...)
		}
		// 根据补零之后的被除数和原来的被除数比较得出补零结果
		for i := dividend.len() - a.len(); i > 0; i-- {
			if len(tmpIntResult) == 0 {
				tmpIntResult = append(tmpIntResult,0)
				offset++
				continue
			}
			tmpFloatResult = append(tmpFloatResult,make([]int,i)...)
			break
		}
		cDividend := b.copy(dividend)
		// 记录截取除数的指针
		divPtr := len(divisor.data)
		for len(tmpIntResult)+len(tmpFloatResult) <= DivisionAccuracy {
			// 运算中被除数小于除数的情况
			if b.LT(dividend, divisor) && !b.EQ(dividend, divisor) {
				// 被除数首位小于等于除数则还需要补一位
				flow := len(divisor.data) - len(dividend.data)
				if dividend.data[1] <= divisor.data[1] {
					flow++
				}
				offset++
				// 借位大于1的话则补上对应的零
				if flow > 1 {
					tmpFloatResult = append(tmpFloatResult, make([]int, flow-1)...)
				}
				divPtr = len(divisor.data)
				// 匹配点的向量减去借位的位数
				pointIV += flow
				dividend.data = append(dividend.data, make([]int8, flow)...)
			}
			// 截取与被除数相等的长度进行比较,包括符号位
			tmp := BigNum{_type: INTEGER, data: dividend.data[:divPtr]}
			if !b.GT(&tmp, divisor) && b.NE(&tmp, divisor) {
				divPtr++
				pointIV++
				continue
			}
			// 添加实际长度的匹配点，忽略符号位
			// 确定point
			point = append(point,cDividend.len() - dividend.len() + 1)
			point[len(point)-1] += pointIV
			pointIV = 0
			// 预测除数与被除数相差多少倍
			// 截取被除数4位于除数3位，利用语言原生提供的除法预测
			for i := 2; i <= 10; i++ {
				if r := b.Ride(divisor, &BigNum{_type: INTEGER, data: []int8{int8(PN), int8(i)}}); b.GT(r, &tmp) {
					checkPoint()
					appendData(int8(i - 1))
					r = b.Sub(r, divisor)
					// 将r补上零并减去
					// 结果是整数位时才需要补零
					if offset == a.len() {
						nativeLen = len(dividend.data) - len(r.data)
						if dividend.data[1] < r.data[1] {
							r.data = append(r.data, make([]int8, len(dividend.data)-len(r.data)-1)...)
						} else {
							r.data = append(r.data, make([]int8, len(dividend.data)-len(r.data))...)
						}
					}
					dividend = b.Sub(dividend, r)
					break
				}
			}
			// TODO:优化补零算法
			// 如果减法之后的结果为0，则直接输出
			// 需要补上剩余的零，补零数量 = dividend.zeroLen - divisor.zeroLen
			if b.EQ(dividend, &BigNum{_type: INTEGER, data: []int8{int8(PN), 0}}) {
				for i := nativeLen; i > 0; i-- {
					appendData(0)
				}
				break
			}
			divPtr = len(divisor.data)
		}
		// 将结果去除并打上符号位
		for i := len(tmpIntResult) - 1; i >= 0; i-- {
			integerResult.Push(tmpIntResult[i])
		}
		for i := len(tmpFloatResult) - 1; i >= 0; i-- {
			pointResult.Push(tmpFloatResult[i])
		}
		// 负数不在此处理
		integerResult.Push(PN)
	} else if a._type == c._type && a._type == FLOAT {
		// 小数对齐计算
		a2 := b.copy(a)
		c2 := b.copy(c)
		// 确定偏移量
		offset := hit(len(a2.pointData) > len(c2.pointData), len(a2.pointData), len(c2.pointData)).(int)
		// 去除左边的无意义零，保证计算结果的准确性
		a2.data = sliceDeleteLeftZero(append(a2.data, append(a2.pointData, make([]int8, offset-len(a2.pointData))...)...))
		c2.data = sliceDeleteLeftZero(append(c2.data, append(c2.pointData, make([]int8, offset-len(c2.pointData))...)...))
		a2._type = INTEGER
		c2._type = INTEGER
		// 清空pointData
		a2.pointData = nil
		c2.pointData = nil
		return b.Except(a2, c2)
	} else if a._type != c._type {
		// 将不为小数的一方装换为小数进行运算
		integer := b.copy(hit(a._type == INTEGER, a, c).(*BigNum))
		float := b.copy(hit(a._type == FLOAT, a, c).(*BigNum))
		integer._type = FLOAT
		integer.pointData = append(integer.pointData, 0)
		dividend := hit(b.EQ(integer, a), integer, float).(*BigNum)
		divisor := hit(b.EQ(float, c), float, integer).(*BigNum)
		return b.Except(dividend, divisor)
	}

	// 序列化
	// 拼接成BigNum类型
	element := &BigNum{}
	if !integerResult.IsEmpty() {
		element._type = INTEGER
	}
	if !pointResult.IsEmpty() {
		element._type = FLOAT
	}
	for !integerResult.IsEmpty() {
		element.data = append(element.data, int8(integerResult.Pop().(int)))
	}
	for !pointResult.IsEmpty() {
		element.pointData = append(element.pointData, int8(pointResult.Pop().(int)))
	}
	return element
}

func (b *BigNum) Mod(a,c *BigNum) *BigNum {
	result := b.Except(a,c)
	dec := &BigNum{_type: FLOAT,data: []int8{result.data[0],0},pointData: append(result.pointData)}
	return b.Ride(dec,c)
}