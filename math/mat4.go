package math

type Mat4 struct {
	C11, C12, C13, C14,
	C21, C22, C23, C24,
	C31, C32, C33, C34,
	C41, C42, C43, C44 float64
}

func Mat4Identity() *Mat4 {
	mat := &Mat4{
		C11: 1, C12: 0, C13: 0, C14: 0,
		C21: 0, C22: 1, C23: 0, C24: 0,
		C31: 0, C32: 0, C33: 1, C34: 0,
		C41: 0, C42: 0, C43: 0, C44: 1,
	}
	return mat
}

