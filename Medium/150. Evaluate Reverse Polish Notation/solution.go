package evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
    nums := make([]int, 2)
    for _, token := range tokens {
        a, b := nums[len(nums) - 2], nums[len(nums) - 1]
        switch token {
        case "+": nums = append(nums[:len(nums) - 2], a + b)
        case "-": nums = append(nums[:len(nums) - 2], a - b)
        case "*": nums = append(nums[:len(nums) - 2], a * b)
        case "/": nums = append(nums[:len(nums) - 2], a / b)
        default : nums = append(nums, atoi(token))
        }
    }

    return nums[2]
}

func atoi(str string) int {
    num, _ := strconv.Atoi(str)
    return num
}