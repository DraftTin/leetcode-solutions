import "math"

func kthFactor(n int, k int) int {
	s := int(math.Sqrt(float64(n)))

	count := 0

	for i := 1; i <= s; i++ {
		if n%i == 0 {
			count++
		}

		if count == k {
			return i
		}
	}

	total := count * 2

	if s*s == n {
		total--
		count--
	}

	if k > total {
		return -1
	}

	for i := s; i >= 1; i-- {
		if n%i == 0 {
			count++
		}

		if count == k {
			return n / i
		}
	}

	return -1
}
