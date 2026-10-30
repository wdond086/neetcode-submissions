type MinStack struct {
	stack 	   []int
	orderStack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, orderStack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	stackLength := len(this.orderStack)
	if stackLength == 0 {
		this.orderStack = append(this.orderStack, val)
	} else {
		currentMin := this.orderStack[stackLength - 1]
		if val < currentMin {
			this.orderStack = append(this.orderStack, val)
		} else {
			this.orderStack = append(this.orderStack, currentMin)
		}
	}
}

func (this *MinStack) Pop() {
	stackLength := len(this.stack)
	this.stack = this.stack[:stackLength - 1]
	this.orderStack = this.orderStack[:stackLength - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}

func (this *MinStack) GetMin() int {
	return this.orderStack[len(this.orderStack) - 1]
}
