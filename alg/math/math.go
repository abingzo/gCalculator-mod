// Package math 存储一些支持的运算符的运算方法
package math

type Step interface {
	// Add 加
	Add(a, b string) string
	// Sub 减
	Sub(a, b string) string
	// Ride 乘
	Ride(a, b string) string
	// Except 除
	Except(a, b string) string
	// Power 幂
	Power(a, b string) string
	// Mod 求余
	Mod(a, b string) string
	// LeftShift 模仿二进制的左移
	LeftShift(a, b string) string
	// RightShift 模仿二进制的右移
	RightShift(a, b string) string
	// Sin 三角函数
	Sin(a string) string
	Cos(a string) string
	Tan(a string) string
	// Pi π
	Pi(a string) string
}

const PI string = "3.141592653589793"

type step struct {
	Handler
	HandlerDouble
}

func NewStep() Step {
	return Step(&step{})
}

func newFormString(a string) *BigNum {
	num := BigNum{}
	return num.FromString(a)
}

// NumType 数字的类型
type NumType int

// HandlerDouble 函数入参绑定类型,双操作数
type HandlerDouble func(string,string) string

// Handler 函数入参绑定类型,单操作数
type Handler func(string) string


// Add 加
func (hd HandlerDouble) Add(a, b string) string {
	var x,y BigNum
	return x.Add(x.FromString(a),y.FromString(b)).String()
}

// Sub 减
func (hd HandlerDouble) Sub(a, b string) string {
	var x,y BigNum
	return x.Sub(x.FromString(a),y.FromString(b)).String()
}

// Ride 乘
func (hd HandlerDouble) Ride(a, b string) string {
	var x,y BigNum
	return x.Ride(x.FromString(a),y.FromString(b)).String()
}

// Except 除
func (hd HandlerDouble) Except(a, b string) string {
	var x,y BigNum
	return x.Except(x.FromString(a),y.FromString(b)).String()
}

// LeftShift 模仿二进制左移
func (hd HandlerDouble) LeftShift(a, b string) string {
	var x,y BigNum
	return x.Ride(newFormString(a),y.FromString(hd.Power("2",b))).String()
}

// RightShift 模仿二进制右移
func (hd HandlerDouble) RightShift(a, b string) string {
	var x,y BigNum
	return x.Except(newFormString(a),y.FromString(hd.Power("2",b))).String()
}

// Power 幂
func (hd HandlerDouble) Power(a, b string) string {
	var x,y BigNum
	if r := x.FromString(b);x.GT(r,y.FromString("0")) {
		return power(a,b)
	} else {
		return x.Except(x.FromString("1"),y.FromString(power(a,r.Abs().String()))).String()
	}
}

func power(a,b string) string {
	var x,y BigNum
	ans := newFormString("1")
	n := newFormString(b)
	// 贡献的初始值为 x
	xContribute := newFormString(a)
	for x.GT(n,newFormString("0")) {
		if x.EQ(x.Mod(n,y.FromString("2")).ToInteger(),y.FromString("1")) {
			// n的二进制位最低为1则需要计入贡献
			ans = x.Ride(ans,xContribute)
		}
		xContribute = x.Ride(xContribute,xContribute)
		n = x.Except(n,y.FromString("2")).ToInteger()
	}
	return ans.String()
}

// Mod 求余
func (hd HandlerDouble) Mod(a, b string) string {
	var x,y BigNum
	return x.Mod(x.FromString(a),y.FromString(b)).String()
}

// Sin sin
func (h Handler) Sin(a string) string {
	panic("sin")
}

func (h Handler) Cos(a string) string {
	panic("cos")
}

// Tan tan
func (h Handler) Tan(a string) string {
	panic("tan")
}

// Pi n*pi
func (h Handler) Pi(a string) string {
	var x,y BigNum
	return x.Ride(x.FromString(a),y.FromString(PI)).String()
}
