package four_sum

import "slices"

func fourSum(nums []int, target int) [][]int {
    slices.Sort(nums)
    result := make([][]int, 0)
    ln := len(nums)

    for num1, i := nums[0] - 1, 0; i < ln - 3; i += 1 {
        if nums[i] == num1 {
            continue
        }

        num1 = nums[i]
        updatedTarget := target - num1

        if nums[i + 1] + nums[i + 2] + nums[i + 3] > updatedTarget {
            break
        }

        if nums[ln - 1] + nums[ln - 2] + nums[ln - 3] < updatedTarget {
            continue
        }

        for num2, j := nums[i + 1] - 1, i + 1; j < ln - 2; j += 1 {
            if nums[j] == num2 {
                continue
            }

            num2 = nums[j]
            currentTarget := updatedTarget - num2
            k, l := j + 1, ln - 1

            if nums[k] + nums[k + 1] > currentTarget {
                break
            }

            if nums[l] + nums[l - 1] < currentTarget {
                continue
            }
            
            for k < l {
                sum := nums[k] + nums[l]

                if sum == currentTarget {
                    result = append(result, []int {num1, num2, nums[k], nums[l]})
                    for k += 1; k < l && nums[k] == nums[k - 1]; k += 1 { }
                    for l -= 1; k < l && nums[l] == nums[l + 1]; l -= 1 { }
                } else if sum < currentTarget {
                    for k += 1; k < l && nums[k] == nums[k - 1]; k += 1 { }
                } else {
                    for l -= 1; k < l && nums[l] == nums[l + 1]; l -= 1 { }
                }
            }
        }
    }

    return result
}