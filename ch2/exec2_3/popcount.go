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
		c += pc[byte(x>>(i*8))]
	}
	return int(c)
}
