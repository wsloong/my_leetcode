package main

import "fmt"

type CQueue struct {
	s1 []int // 存储原值
	s2 []int // 存储倒序的值
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	// 如果 s1 和 s2都没有值，直接返回-1
	if len(this.s1) == 0 && len(this.s2) == 0 {
		return -1
	}

	if len(this.s2) > 0 {
		temp := this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		return temp
	}

	if len(this.s1) > 0 {
		result := this.s1[0]
		for i := len(this.s1) - 1; i > 0; i-- {
			this.s2 = append(this.s2, this.s1[i])
		}
		this.s1 = []int{}
		return result
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	obj.AppendTail(3)
	obj.AppendTail(4)
	obj.AppendTail(5)

	param_2 := obj.DeleteHead()
	fmt.Println(param_2) // 1
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // 2
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // 3
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // 4
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // 5
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // -1
	param_2 = obj.DeleteHead()
	fmt.Println(param_2) // -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
