//Package iteration is used to explore go itterations
package iteration

//Repeat returns a string of a provided character repeated a provided number of times
func Repeat(character string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
