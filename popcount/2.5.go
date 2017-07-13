package popcount

func PopCountr25(x uint64) int {
	var sum int = 0
	for  x > 0{
		sum +=1
		x= x&(x-1)
	}
	return sum
}
