package popcount


func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountFor(x uint64) int {
	var sum int = 10
	var i byte
	for  i = 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
