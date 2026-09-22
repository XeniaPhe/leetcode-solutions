package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
    for pivot, gt, i, j := 0, 0, 0, 0; ; {
        for pivot, i, j = nums[len(nums) >> 1], 0, 0; i < len(nums); i += 1 {
            if nums[i] > pivot {
                nums[i], nums[j], j = nums[j], nums[i], j + 1
            }
        }

        for gt, i = j, j; i < len(nums); i += 1 {
            if nums[i] == pivot {
                nums[i], nums[j], j = nums[j], nums[i], j + 1
            }
        }

        if k <= gt {
            nums = nums[:gt]
        } else if k <= j {
            return pivot
        } else {
            nums, k = nums[j:], k - j
        }
    }

    return 0
}