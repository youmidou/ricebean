package sys_base

import (
	"fmt"
	"math"
)

/*
	T向量3

added by yh @ 2024/04/19 15:48
注意:
*/
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// 加法
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// 减法
func (v Vector3) Subtract(other Vector3) Vector3 {
	return Vector3{X: v.X - other.X, Y: v.Y - other.Y, Z: v.Z - other.Z}
}

// 计算模长
func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// 计算点乘
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// 计算叉乘
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func main1() {
	v1 := Vector3{X: 1, Y: 2, Z: 3}
	v2 := Vector3{X: 4, Y: 5, Z: 6}

	// 加法
	fmt.Println("Addition:", v1.Add(v2))

	// 减法
	fmt.Println("Subtraction:", v1.Subtract(v2))

	// 计算模长
	fmt.Println("Magnitude of v1:", v1.Magnitude())

	// 计算点乘
	fmt.Println("Dot product:", v1.Dot(v2))

	// 计算叉乘
	fmt.Println("Cross product:", v1.Cross(v2))
}
