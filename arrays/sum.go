//Package arrays is used to explore go arrays
package arrays

//Sum returns the integer value from adding everything in the provided array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//SumAll adds up all the provided slices and returns a slice with all the answers of their Sum
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
