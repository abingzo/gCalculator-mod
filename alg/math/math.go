// 存储一些支持的运算符的运算方法
package math

import (
	"math/big"
	"strings"
)

type Step interface {
	// 加
	Add(a,b string) string
	// 减
	Sub(a,b string) string
	// 乘
	Ride(a,b string) string
	// 除
	Except(a,b string) string
}

const (
	// 小数精度
	DECIMAL_PRECISION = 14
	// 整数的精度
	INTERGER_PRECISION = 20
)

// 数字的类型
type numType int

// 存储数字的结构体
type Base struct {
	// 左数
	leftNum string
	// 右数
	rightNum string
	// 左数的类型
	leftType numType
	// 右数的类型
	rightType numType
	// 小数点的位置
	pointPtr int
	// 存储数字的底层类型
	bottomType *big.Int
}

// 函数入参绑定类型
type Handler func(a,b string) string

// 确定数字的类型
// 并将结果保存到结构体的字段中
func (b2 *Base) checkNumType()  {
	// 包含小数点则为小数
	if result := strings.Index(b2.leftNum,"."); result > 0 {
		b2.pointPtr = len(b2.leftNum) - result
	}
	// 第二个操作数为小数且大于原有的小数的位置
	if result := strings.Index(b2.rightNum,"."); result > 0 && len(b2.rightNum) - result > b2.pointPtr {

	}
}

func (b2 *Base) Add(a, b string) string {
	panic("implement me")
}

func (b2 *Base) Sub(a, b string) string {
	panic("implement me")
}

func (b2 *Base) Ride(a, b string) string {
	panic("implement me")
}

func (b2 *Base) Except(a, b string) string {
	panic("implement me")
}
