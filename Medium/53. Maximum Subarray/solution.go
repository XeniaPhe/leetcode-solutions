package maximum_subarray

import "math"

/*
	This is what I was able to come up with after staring at the problem for an hour or two
	without looking up the Kadane's algorithm. As far as I can see, it uses the same idea
	but it keeps extra unnecessary state and explicitly handles each edge case
	that didn't really need to be explicitly handled. Basically it implements an extra
	state machine on top of Kadane's simple algorithm. So overall I have an ugly version of
	Kadane's algorithm with a bigger constant factor and it's also probably prone to more
	cache misses due to so many branches. But it's spiritually and mathematicaly equivalent to Kadane!
*/

func maxSubArray(nums []int) int {
    currSum, prevSum, bestSum := 0, 0, math.MinInt
    for temp, num := range nums {
        if currSum >= 0 {
            if num < 0 {
                prevSum, currSum = currSum, 0
            }

            currSum += num
        } else if num >= 0 {
            temp, currSum = currSum, num
            if merged := temp + prevSum; merged > 0 {
                currSum += merged
            }
        } else {
            if num > bestSum {
                bestSum = num
            }

            currSum += num
            continue
        }

        if currSum > bestSum {
            bestSum = currSum
        }
    }

    return bestSum
}