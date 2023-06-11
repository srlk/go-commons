package maths

func Max[K int | float32 | float64](a, b K) K {
	if a > b {
		return a
	}
	return b
}

func Min[K int | float32 | float64](a, b K) K {
	if a < b {
		return a
	}
	return b
}
