package test

import (
	"gCalculator-mod/alg"
	"gCalculator-mod/alg/math"
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	stack := alg.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Peek())
	t.Log(stack.IsEmpty())
}

func TestMath(t *testing.T) {
	var x,y math.BigNum
	x.FromString("+" + strconv.FormatUint(2 << 62,10) + "223372036854775808")
	y.FromString("-" + strconv.FormatUint(2 << 62,10) + "." + strconv.FormatInt(2 << 61,10))
	t.Log(x.String())
	t.Log(y.String())
	var i1,i2 math.BigNum
	i1.FromString("223372036854775808")
	i2.FromString("999")
	z := i1.Add(&i1,&i2)
	t.Log(z.String())
	// 浮点数相加
	var f1,f2 math.BigNum
	f1.FromString("0.333")
	f2.FromString("0.33")
	z2 := f1.Add(&f1,&f2)
	t.Log(z2.String())
	// 浮点数与整数相加
	f1.FromString("37332686")
	f2.FromString("726578685.76786875779")
	z2 = f1.Add(&f1,&f2)
	t.Log(z2.String())
	t.Log(f1.String())
}
