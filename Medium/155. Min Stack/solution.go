package min_stack

type MinStack struct {
    stack [][2]int
}

func Constructor() MinStack {
    return MinStack{nil}
}

func (this *MinStack) Push(value int) {
    currMin := value
    if len(this.stack) > 0 {
        currMin = min(currMin, this.GetMin())
    }

    this.stack = append(this.stack, [2]int{value, currMin})
}

func (this *MinStack) Pop() {
    this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1][0]
}

func (this *MinStack) GetMin() int {
    return this.stack[len(this.stack) - 1][1]
}