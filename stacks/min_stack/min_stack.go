package stack

type MinStack struct {
	arr []int
	min int
}


func Constructor() MinStack {
	return MinStack {
			arr: []int{},
			min: 1<<63-1,
	}
}


func (this *MinStack) Push(val int)  {
	this.arr = append(this.arr, val)
	if val < this.min {
			this.min = val
	}
}


func (this *MinStack) Pop()  {
	// getting the top element
	topElement := this.arr[len(this.arr)-1]
	// removing the top element
	this.arr = this.arr[:len(this.arr)-1]

	if topElement == this.min && len(this.arr) > 0 {
			newMinVal := this.arr[0]

			for _, item := range this.arr {
					if item < newMinVal {
							newMinVal = item
					}
			}

			this.min = newMinVal
	}

	if len(this.arr) == 0 {
			this.min = 1<<63-1   
	}
}


func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
	return this.min
}
// write test case for MinStack