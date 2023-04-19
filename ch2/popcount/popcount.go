package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountByLoop(x uint64) int {
	var c byte
	for i := 0; i < 8; i++ {
		// notice >> and * has same precedence but >> will be associated first.
		// (see p. 52 of the book).
		// so (i*8) is required, or x>>i first then *8 which is NOT what we want.
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}

func PopCountByShifting(x uint64) int {
	var c byte
	for i := 0; i < 64; i++ {
		c += byte(x & 1)
		x >>= 1
	}
	return int(c)
}

func PopCountByRightMostOne(x uint64) int {
	c := 0
	for ; x > 0; c++ {
		x &= (x - 1)
	}
	return c
}
