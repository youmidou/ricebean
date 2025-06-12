package sys_base

import (
	"fmt"
	"math"
)

type Axis int

/*
	T向量2

added by yh @ 2024/04/19 15:48
注意:
*/

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

// 加法
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{X: v.X + other.X, Y: v.Y + other.Y}
}

// 减法
func (v Vector2) Subtract(other Vector2) Vector2 {
	return Vector2{X: v.X - other.X, Y: v.Y - other.Y}
}

// 标量乘法
func (v Vector2) ScalarMultiply(scalar float64) Vector2 {
	return Vector2{X: v.X * scalar, Y: v.Y * scalar}
}

// 计算模长
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v1 := Vector2{X: 3, Y: 4}
	v2 := Vector2{X: 1, Y: 2}

	// 加法
	fmt.Println("Addition:", v1.Add(v2))

	// 减法
	fmt.Println("Subtraction:", v1.Subtract(v2))

	// 标量乘法
	fmt.Println("Scalar Multiplication:", v1.ScalarMultiply(2))

	// 计算模长
	fmt.Println("Magnitude:", v1.Magnitude())
}
