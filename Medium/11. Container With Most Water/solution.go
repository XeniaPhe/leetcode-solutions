package container_with_most_water

func maxArea(height []int) int {
    maxArea := 0
    for left, right := 0, len(height) - 1; right > left; {
        area := mint2(height[left], height[right]) * (right - left)
        
        if area > maxArea {
            maxArea = area
        }

        if height[left] <= height[right] {
            left += 1
        } else {
            right -= 1
        }
    }

    return maxArea
}

func mint2(a int, b int) int {
    if a <= b {
        return a
    }

    return b
}