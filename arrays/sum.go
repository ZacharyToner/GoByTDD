//Package arrays is used to explore go arrays
package arrays

//Sum returns the integer value from adding everything in the provided array
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
