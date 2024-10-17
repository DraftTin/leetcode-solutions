package main

import "math"

// math: we can convert (2^(p - 1) - 1) numbers to '1', and (2^(p - 1) - 1) nums to '(2^p - 2)' by swapping bits
func minNonZeroProduct(p int) int {
	mod := uint64(1e9 + 7)
	val := uint64(math.Pow(2, float64(p)))
	ans := expn(val-2, val/2-1, mod)
	return int(ans * ((val - 1) % mod) % mod)
}

func expn(val, k, mod uint64) uint64 {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return val % mod
	}
	if k%2 == 0 {
		tmp := expn(val, k/2, mod)
		return (tmp * tmp) % mod
	} else {
		tmp := expn(val, k/2, mod)
		tmp = (tmp * tmp) % mod
		return (tmp * (val % mod)) % mod
	}

}
