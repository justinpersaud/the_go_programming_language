// Rewrite PopCount to use a loop instead of a single expression. Compare the performance of the two versions. (Refer to section 11.4 on how to compare performance differences)

package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	// return int(pc[byte(x>>(0*8))] +
	// 	pc[byte(x>>(1*8))] +
	// 	pc[byte(x>>(2*8))] +
	// 	pc[byte(x>>(3*8))] +
	// 	pc[byte(x>>(4*8))] +
	// 	pc[byte(x>>(5*8))] +
	// 	pc[byte(x>>(6*8))] +
	// 	pc[byte(x>>(7*8))])
	var total byte
	for i := uint64(0); i < 8; i++ {
		total += pc[byte(x>>(i*8))]
	}
	return int(total)
}
