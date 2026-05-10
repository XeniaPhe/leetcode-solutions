package jump_game

func canJump(nums []int) bool {
    for target, max, currMax, i := len(nums) - 1, -1, 0, 0; i < len(nums); i += 1 {
        if maxi := i + nums[i]; maxi > max {
            if max = maxi; max >= target {
                return true
            }
        }

        if temp := currMax; i == currMax {
            if currMax = max; temp == max {
                break
            }
        }
    }

    return false
}