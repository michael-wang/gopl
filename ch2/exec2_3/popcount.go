package exec2_3

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var c byte
	for i := 0; i < 8; i++ {
		// notice >> and * has same precedence but >> will be associated first.
		// (see p. 52 of the book).
		// so (i*8) is required, or x>>i first then *8 which is NOT what we want.
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}
