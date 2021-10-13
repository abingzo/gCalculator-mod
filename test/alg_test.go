package test

import (
	"gCalculator-mod/alg"
	"gCalculator-mod/alg/math"
	bmath "math/big"
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

// Result|num|num
type Case map[string]map[string]string

// 辅助函数的测试用例类型
type AuxCase map[bool]map[string]string

func TestMath(t *testing.T) {
	// 加法的测试用例
	var AddCase Case = map[string]map[string]string{
		"100":                           {"10": "90", "1": "99", "200": "-100"},
		"-200":                          {"-100": "-100", "-199.5": "-0.5"},
		"99.55":                         {"10": "89.55", "9": "90.55", "99.1": "0.45"},
		"0":                             {"-199": "199"},
		"18070406285660840000000000000": {"9.903520314283042e+27": "9.903520314283042e+27"},
	}
	// 减法的测试用例
	var SubCase Case = map[string]map[string]string{
		"-100": {"100": "200", "-200": "-100"},
		"999":  {"-999": "-1998", "1998": "999"},
		"50.5": {"100": "49.5", "170.9": "120.4"},
	}
	// 乘法的测试用例
	var RideCase Case = map[string]map[string]string{
		"100":   {"10": "10", "1": "100", "20": "5", "-20": "-5"},
		"-120":  {"-20": "6", "20": "-6", "-240": "0.5"},
		"16.25": {"2.5": "6.5", "4": "4.0625"},
	}
	// 除法的测试用例
	var ExceptCase Case = map[string]map[string]string{
		"100":        {"10000": "100", "1": "0.01", "10": "0.1"},
		"10086":      {"90774": "9"},
		"1086":       {"9774": "9"},
		"4568000950": {"15988003325": "3.5"},
		"400":        {"0.89": "0.002225"},
		"107800":     {"970200": "9"},
		"0.0475":     {"9.5": "200"},
		"0.05":       {"1": "20"},
		"16":         {"32": "2","64":"4"},
	}
	var x, y math.BigNum
	// 测试加法
	for k, v := range AddCase {
		for kk, vv := range v {
			if r := x.Add(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("加法测试失败: %s+%s=%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("加法测试成功")
	// 测试减法
	for k, v := range SubCase {
		for kk, vv := range v {
			if r := x.Sub(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("减法测试失败: %s-%s=%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("减法测试成功")
	// 测试乘法
	for k, v := range RideCase {
		for kk, vv := range v {
			if r := x.Ride(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("乘法测试失败: %s*%s=%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("乘法测试成功")
	// 测试除法
	for k, v := range ExceptCase {
		for kk, vv := range v {
			if r := x.Except(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("除法测试失败: %s/%s=%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("除法测试成功")
}

// 测试辅助函数的正确性
func TestBigNumAuxFunction(t *testing.T) {
	// 比较大的测试用例
	var Max Case = map[string]map[string]string{
		"0.1835": {"0.18304": "0.1835"},
	}
	// 比较小的测试用例
	var Min Case = map[string]map[string]string{
		"0.18304": {"0.18304": "0.18305"},
	}
	// 布尔类型的比较测试用例
	// 比较是否大于
	var GT AuxCase = map[bool]map[string]string{
		true:  {"9": "8"},
		false: {"8": "9", "9": "9", "10.505": "10.50506"},
	}
	// 比较是否小于
	var LT AuxCase = map[bool]map[string]string{
		true:  {"8": "9", "0.80001": "0.800019"},
		false: {"0.800019": "0.80001"},
	}
	// 比较是否等于
	var EQ AuxCase = map[bool]map[string]string{
		true:  {"0.1": "0.10"},
		false: {"0.183040000": "0.18304001"},
	}
	// 比较是否小于等于
	var LE AuxCase = map[bool]map[string]string{
		true:  {"0.2": "0.5"},
		false: {"456789.02": "456789.0000001"},
	}
	// 比较是否大于等于
	var GE AuxCase = map[bool]map[string]string{
		true:  {"76": "75.9", "76.2": "76.08", "76.002": "76.0020000"},
		false: {"76.0": "76.0000001"},
	}
	var x, y math.BigNum
	// 测试max
	for k, v := range Max {
		for kk, vv := range v {
			if r := x.Max(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("找出最大的数测试失败: %s和%s中最大的数为:%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("Max测试成功")
	// 测试min
	for k, v := range Min {
		for kk, vv := range v {
			if r := x.Min(x.FromString(kk), y.FromString(vv)); !x.EQ(r, x.FromString(k)) {
				t.Errorf("找出最小的数测试失败: %s和%s中最小的数为:%s", kk, vv, r)
				break
			}
		}
	}
	t.Log("Min测试成功")
	// 测试GT
	for k, v := range GT {
		for kk, vv := range v {
			if r := x.GT(x.FromString(kk), y.FromString(vv)); !(r == k) {
				t.Errorf("GT测试失败: %s>%s的结果为%v", kk, vv, r)
				break
			}
		}
	}
	t.Log("GT测试成功")
	// 测试LT
	for k, v := range LT {
		for kk, vv := range v {
			if r := x.LT(x.FromString(kk), y.FromString(vv)); !(r == k) {
				t.Errorf("LT测试失败: %s<%s的结果为%v", kk, vv, r)
				break
			}
		}
	}
	t.Log("LT测试成功")
	// 测试EQ
	for k, v := range EQ {
		for kk, vv := range v {
			if r := x.EQ(x.FromString(kk), y.FromString(vv)); !(r == k) {
				t.Errorf("EQ测试失败: %s==%s的结果为%v", kk, vv, r)
				break
			}
		}
	}
	t.Log("EQ测试成功")
	// 测试GE
	for k, v := range GE {
		for kk, vv := range v {
			if r := x.GE(x.FromString(kk), y.FromString(vv)); !(r == k) {
				t.Errorf("GE测试失败: %s>=%s的结果为%v", kk, vv, r)
				break
			}
		}
	}
	t.Log("GE测试成功")
	// 测试LE
	for k, v := range LE {
		for kk, vv := range v {
			if r := x.LE(x.FromString(kk), y.FromString(vv)); !(r == k) {
				t.Errorf("LE测试失败: %s<=%s的结果为%v", kk, vv, r)
				break
			}
		}
	}
	t.Log("LE测试成功")
}

// 对大数操作的一些性能测试
func BenchmarkBigNum(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		var x, y math.BigNum
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x.Add(x.FromString("13696391639698649264985275947507057027074290650924860609265757858587592645969465946"),
				y.FromString("89365965965964956925649976592465982659826557859629465984265984264596259642956924659645"))
		}
	})
	b.Run("Sub", func(b *testing.B) {
		var x, y math.BigNum
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x.Sub(x.FromString("13696391639698649264985275947507057027074290650924860609265757858587592645969465946273658256283523"),
				y.FromString("89365965965964956925649976592465982659826557859629465984265984264596259642956924659645"))
		}
	})
	b.Run("Ride", func(b *testing.B) {
		var x, y math.BigNum
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.Log(x.Ride(x.FromString("1369639163969864926498527594750705702707429065092486060926575785858759264596946594686259265926592659269169469164926942689164914164916491649162"),
				y.FromString("893659659659649569256499765924659826598265578596294659842659842645962596429569246596452465942659265926956295639286592659265965926956956926526574659")).String())
		}
	})
	b.Run("Big", func(b *testing.B) {
		var x, y bmath.Int
		xx, _ := x.SetString("1369639163969864926498527594750705702707429065092486060926575785858759264596946594686259265926592659269169469164926942689164914164916491649162", 10)
		yy, _ := y.SetString("893659659659649569256499765924659826598265578596294659842659842645962596429569246596452465942659265926956295639286592659265965926956956926526574659", 10)
		x.Mul(xx, yy)
	})
	b.Run("Except", func(b *testing.B) {
		var x, y math.BigNum
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			x.Except(x.FromString("13696391639698649264985275947507057027074290650924860609265757858587592645969465946"),
				y.FromString("89365965965964956925649976592465982659826557859629465984265984264596259642956924659645"))
		}
	})
}
