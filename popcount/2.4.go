package popcount



func PopCountr24(x uint64) int {
	var sum int = 0
	var i byte
	for  i = 0; i < 64; i++ {
		sum += int(byte(x&1))
		x = x>>1
	}
	return sum
}
