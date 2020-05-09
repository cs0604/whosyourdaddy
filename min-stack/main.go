package min_stack

type MinStack struct {
	ele []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{}
}

func (this *MinStack) Push(x int) {

	curMin := x
	if len(this.min) > 0 && x > this.min[len(this.min)-1] {
		curMin = this.min[len(this.min)-1]
	}
	this.min = append(this.min, curMin)
	this.ele = append(this.ele, x)

}

func (this *MinStack) Pop() {

	if len(this.ele) == 0 {
		return
	}
	this.min = this.min[:len(this.min)-1]
	this.ele = this.ele[:len(this.ele)-1]
}

func (this *MinStack) Top() int {

	if len(this.ele) == 0 {
		return 0
	}
	return this.ele[len(this.ele)-1]

}

func (this *MinStack) GetMin() int {

	if len(this.min) == 0 {
		return 0
	}

	return this.min[len(this.min)-1]
}
