package Utils

import "math"

//Calculates the product of the ints in the slice
//@param slice -- slice to calc with
//@return product of slice values, 1 if empty
func ProductOfIntSlice(slice []int) int {
	product := 1
	for _, num := range slice {
		product *= num
	}
	return product
}

//Calculates the sum of the ints in the slice
//@param slice -- slice to calc with
//@return sum of slice values, 0 if slice is empty
func SumOfIntSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

//Finds the min int and its index in the slice
//@param slice -- slice to look in
//@return min, indexOfMin, index will be -1 if slice is empty
func MinOfIntSlice(slice []int) (int, int) {
	if len(slice) < 1 {
		return 0, -1
	}
	min := slice[0]
	minIndex := 0
	for i, num := range slice {
		if num < min {
			min = num
			minIndex = i
		}
	}
	return min, minIndex
}

//Finds the max int and its index in the slice
//@param slice -- slice to look in
//@return max, indexOfMax, index will be -1 if slice is empty
func MaxOfIntSlice(slice []int) (int, int) {
	if len(slice) < 1 {
		return 0, -1
	}
	max := slice[0]
	maxIndex := 0
	for i, num := range slice {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return max, maxIndex
}

//Calculates the mean of the int slice
//@param ints -- int slice to calc on
//@return mean of int slice
func Mean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum / len(ints)
}

//Calculates the mean of the int slice, rounds down
//@param ints -- int slice to calc on
//@return mean of int slice
func LowMean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return int(math.Floor(float64(sum) / float64(len(ints))))
}

//Calculates the mean of the int slice, rounds up
//@param ints -- int slice to calc on
//@return mean of int slice
func HighMean(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return int(math.Ceil(float64(sum) / float64(len(ints))))
}

//Finds the median of the integer slice
//@param ints -- sorted integer slice
//@return median of slice
func Median(ints []int) int {
	length := len(ints)
	middle := length / 2
	if length%2 != 0 {
		return ints[middle]
	}
	return (ints[middle-1] + ints[middle]) / 2
}
