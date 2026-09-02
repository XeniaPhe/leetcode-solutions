package rotate_array

import "slices"
import "math/rand"

var variations = [4]func([]int, int) {rotate1, rotate2, rotate3, rotate4}

func rotate(nums []int, k int) {
    if k = k % len(nums); k != 0 {
		variations[rand.Intn(4)](nums, k)
    }
}

func rotate1(nums []int, k int) {
    copy(nums, append(nums[len(nums) - k:], nums[:len(nums) - k]...))
}

func rotate2(nums []int, k int) {
	s, e, d := 0, len(nums), 1
	if reverse := len(nums) - k; reverse < k {
		k, s, e, d = reverse, e - 1, s - 1, -1
	}

	for prev, i, j := 0, 0, 0; i < k; i += 1 {
		for prev, nums[s], j = nums[s], nums[e - d], s + d; j != e; j += d {
			prev, nums[j] = nums[j], prev
		}
	}
}

func rotate3(nums []int, k int) {
    slices.Reverse(nums)
    slices.Reverse(nums[:k])
    slices.Reverse(nums[k:])
}

func rotate4(nums []int, k int) {
	for len(nums) > k && k > 0 {
		cycles, rem, start, end, iter, dir := (len(nums) - k) / k, (len(nums) - k) % k, len(nums) - k, len(nums), 0, 1
		if cycles == 0 {
			k, cycles, rem, start, end, iter, dir = start, k / start, k % start, start - 1, -1, end - 1, -1
		}

		for i := 0; i < cycles; i += 1 {
			for j := start; j != end; iter, j = iter + dir, j + dir {
				nums[iter], nums[j] = nums[j], nums[iter]
			}
		}

		if dir == 1 {
			nums = nums[len(nums) - k - rem:]
		} else {
			k, nums = rem, nums[:k + rem]
		}
	}
}