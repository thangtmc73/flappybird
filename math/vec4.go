package math

import "fmt"

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (vec *Vec4) Transform(mat Mat4) *Vec4 {
	result := &Vec4{}
	result.X = (vec.X * mat.C11) + (vec.Y * mat.C21) + (vec.Z * mat.C31) + (vec.W * mat.C41)
	result.Y = (vec.X * mat.C12) + (vec.Y * mat.C22) + (vec.Z * mat.C32) + (vec.W * mat.C42)
	result.Z = (vec.X * mat.C13) + (vec.Y * mat.C23) + (vec.Z * mat.C33) + (vec.W * mat.C43)
	result.W = (vec.X * mat.C14) + (vec.Y * mat.C24) + (vec.Z * mat.C34) + (vec.W * mat.C44)
	fmt.Println(result.X, result.Y, result.Z, mat.C22)
	return result
}
