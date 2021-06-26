package exec2_5

func PopCount(x uint64) int {
	c := 0
	for ; x > 0; c++ {
		x &= (x - 1)
	}
	return c
}
