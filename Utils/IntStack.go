package Utils

import "fmt"

type IntStack struct {
	Nums []int
}

func (is *IntStack) Push(i int) {
	is.Nums = append([]int{i}, is.Nums...)
}

func (is *IntStack) Pop() (int, error) {
	if len(is.Nums) > 0 {
		toReturn := is.Nums[len(is.Nums) - 1]
		is.Nums = is.Nums[:len(is.Nums) - 1]
		return toReturn, nil
	}

	return 0, fmt.Errorf("stack is empty, cannot pop value")
}

func (is IntStack) Max() (int, error) {
	if len(is.Nums) < 1 { return 0, fmt.Errorf("stack is empty, cannot calc max")}
	max := is.Nums[0]
	for _, i := range is.Nums {
		if i > max { max = i }
	}
	return max, nil
}

func (is IntStack) Min() (int, error) {
	if len(is.Nums) < 1 { return 0, fmt.Errorf("stack is empty, cannot calc max")}
	min := is.Nums[0]
	for _, i := range is.Nums {
		if i < min { min = i }
	}
	return min, nil
}

func (is *IntStack) Limit(numElements int) {
	for len(is.Nums) > numElements {
		is.Pop()
	}
}

func (is IntStack) Sum() int {
	sum := 0
	for _, num := range is.Nums {
		sum += num
	}
	return sum
}

func (is IntStack) Product() int {
	prod := 1
	for _, num := range is.Nums {
		prod *= num
	}
	return prod
}

func (is IntStack) String() string {
	toReturn := "["
	for _, num := range is.Nums {
		toReturn += IntToString(num) + " "
	}
	toReturn = toReturn[:len(toReturn) -1]
	toReturn += "]"
	return toReturn
}