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
	t.Log("浮点数与整数相加:" + z2.String())
	t.Log("浮点数:" + f1.String())
	// 负数相加
	z2 = f1.Add(f1.FromString("-20"),f2.FromString("-40"))
	t.Log("负负相加:" + z2.String())
	// 负正相加
	z2 = f1.Add(f1.FromString("-20.7456356"),f2.FromString("40.54363846"))
	t.Log("负正相加:" + z2.String())
	// 正负相加
	z2 = f1.Add(f1.FromString("7954.874297492"),f2.FromString("-987593275.5827592"))
	t.Log("正负相加:" + z2.String())
	// 精度小数相加
	z2 = f1.Add(f1.FromString("0.30000004"),f2.FromString("0.3004400005"))
	t.Log("精度小数相加:" + z2.String())
	// 大整数减法
	z3 := i1.Sub(&i1,&i2)
	t.Log(z3.String())
	// 整数减法
	z3 = i1.Sub(i1.FromString("10000"),i2.FromString("999"))
	t.Log(z3.String())
	// 被减数小于减数
	z3 = i1.Sub(i1.FromString("10"),i2.FromString("999"))
	t.Log(z3.String())
	// 负数相减
	z3 = i1.Sub(i1.FromString("-10.589572"),i2.FromString("67.58796796794"))
	t.Log(z3.String())
	// 正负相减
	z3 = i1.Sub(i1.FromString("10.589572"),i2.FromString("-67.58796796794"))
	t.Log(z3.String())
	// 正小数相减
	z3 = i1.Sub(i1.FromString("0.663"),i2.FromString("0.7"))
	t.Log(z3.String())
	// 正整数单乘多
	z3 = i1.Ride(i1.FromString("9"),i2.FromString("9223"))
	t.Log(z3.String())
	// 正整数多乘多
	z3 = i1.Ride(i1.FromString(strconv.FormatUint(2 << 62,10)),i2.FromString(strconv.FormatUint(2 << 62,10)))
	t.Log(z3.String())
	// 小数乘与小数
	z3 = i1.Ride(i1.FromString("999.9559"),i2.FromString("92.9223"))
	t.Log(z3.String())
	// 小数乘与整数
	z3 = i1.Ride(i1.FromString("999"),i2.FromString("92.9223"))
	t.Log(z3.String())
	// 负负相乘
	z3 = i1.Ride(i1.FromString("-999"),i2.FromString("-92.9223"))
	t.Log(z3.String())
	// 正负相乘
	z3 = i1.Ride(i1.FromString("9"),i2.FromString("-92.9223"))
	t.Log(z3.String())
	// 负正相乘
	z3 = i1.Ride(i1.FromString("-9"),i2.FromString("92.9223"))
	t.Log(z3.String())
}
