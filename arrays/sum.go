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
