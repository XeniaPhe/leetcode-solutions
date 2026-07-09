package maximum_product_subarray

import "math"

func maxProduct(nums []int) int {
    neg, negProd, prod, maxProd := false, 1, 1, math.MinInt
    for i := 0; i < len(nums); i += 1 {
        if nums[i] == 0 {
            neg, negProd, prod, maxProd = false, 1, 1, max(maxProd, 0)
        } else {
            negProd, prod = negProd * nums[i], prod * nums[i]
            if maxProd = max(maxProd, negProd, prod); !neg && prod < 0 {
                neg, negProd = true, 1
            }
        }
    }

    return maxProd
}