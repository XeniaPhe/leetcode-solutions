package two_sum_II__input_array_is_sorted

func twoSum(numbers []int, target int) (indices []int) {
    for i, j := 0, len(numbers) - 1; i < j; {
        for ; numbers[i] + numbers[j] < target && i < j; i += 1 { }
        for ; numbers[i] + numbers[j] > target && i < j; j -= 1 { }
        if numbers[i] + numbers[j] == target && i < j {
            indices = []int{i + 1, j + 1}
            break
        }
    }

    return indices
}