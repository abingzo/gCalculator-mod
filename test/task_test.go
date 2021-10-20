package test

import (
	"gCalculator-mod/alg/math"
	"gCalculator-mod/task"
	"strings"
	"testing"
)

type TaskCase map[string]string

var TPE TaskCase = map[string]string{
	"101|2|+|304|50|45|+|*|*":    "(101+2)*(304*(50+45))",
	"10|2|3|^|6|<<|+":            "10+2^3<<6",
	"10|2|3|6|<<|^|+|100|20|*|-": "10+2^(3<<6)-100*20",
	"10|2|^|sin(":                "sin((10^2)",
}

// 测试task相关的函数
func TestTask(t *testing.T) {
	for k, v := range TPE {
		if r := strings.Join(task.ToPostfixExp(v), "|"); r != k {
			t.Errorf("中缀表达式转后缀表达式失败:%s -> %s : result:%s", v, k, r)
		}
	}
	t.Log(strings.Join(task.ToPostfixExp("sin((10^2)"), "|"))
}

// 测试后缀表达式的计算
func TestCalculate(t *testing.T) {
	t.Log(task.Calculate(task.ToPostfixExp("2<<(2^6-1)")))
	t.Log(math.NewStep().Except("189945","4596"))
	t.Log(65 / 2)
	t.Log(len("1223991269129836461040288087039493387889969454293816708123394526087545181985875931563896746910243876364784265087469157369787673523135360084849536130430967233556314595717940459511592401185395176286057417655205010051147089483999646360761315463674773550919013488670539020453909349645827785758"))
}

// 测试封装的数学模块
func TestMathMode(t *testing.T) {
	fnd := new(math.HandlerDouble)
	fn := new(math.Handler)
	t.Log(fnd.Except("21","2"))
	t.Log(fn.Sqrt(fnd.Power("2","67")))
	t.Log(fn.Sqrt("64646.32765"))
}

// 对封装的math模块的一些性能测试
func BenchmarkMath(b *testing.B) {
	b.Run("Sqrt", func(b *testing.B) {
		fnd := new(math.HandlerDouble)
		fn := new(math.Handler)
		number := fnd.Power("2","2")
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fn.Sqrt(number)
		}
	})
	b.Run("Power", func(b *testing.B) {
		fnd := new(math.HandlerDouble)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fnd.Power("2","16")
		}
	})
}