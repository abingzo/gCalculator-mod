package test

import (
	"gCalculator-mod/task"
	"strings"
	"testing"
)

type TaskCase map[string]string

// 测试task相关的函数
func TestTask(t *testing.T) {
	var TPE TaskCase = map[string]string{
		"101|2|+|304|50|45|+|*|*":    "(101+2)*(304*(50+45))",
		"10|2|3|^|6|<<|+":            "10+2^3<<6",
		"10|2|3|6|<<|^|+|100|20|*|-": "10+2^(3<<6)-100*20",
		"10|2|^|sin(":                "sin((10^2)",
	}
	for k, v := range TPE {
		if r := strings.Join(task.ToPostfixExp(v), "|"); r != k {
			t.Errorf("中缀表达式转后缀表达式失败:%s -> %s : result:%s", v, k, r)
		}
	}
	t.Log(strings.Join(task.ToPostfixExp("sin((10^2)"), "|"))
}
