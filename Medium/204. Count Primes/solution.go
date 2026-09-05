package count_primes

import "math"

func countPrimes(n int) int {
    if n < 5 {
        return max(n - 2, 0)
    }

    cnt, primes := 1, make([]bool, (n - (n & 1) - 1) >> 1)
	sqrtn := (int(math.Floor(math.Sqrt(float64(n - (n & 1) - 1)))) - 3) >> 1

    for i := 0; i <= sqrtn; i += 1 {
        if !primes[i] {
			num, j := (i << 1) + 3, (((i << 1) + 3) * ((i << 1) + 3) - 3) >> 1
            for cnt += 1; j < len(primes); primes[j], j = true, j + num { }
        }
    }

	for i := sqrtn + 1; i < len(primes); i += 1 {
		if !primes[i] {
			cnt += 1
		}
	}

    return cnt
}