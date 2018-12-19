// Package hello says hello to the user or world in different languages
package hello

import "fmt"

const (
	// French Option
	french = "French"
	// Spanish Option
	spanish = "Spanish"
	// Italian Option
	italian = "Italian"
	// English (default) prefix
	englishPrefix = "Hello, "
	// French Prefix
	frenchPrefix = "Bonjour, "
	// Spanish Prefix
	spanishPrefix = "Hola, "
	// Italian Prefix
	italianPrefix = "Ciao, "
)

func main() {
	fmt.Println(Hello("world", ""))
}

// Hello returns the string to print based on a name and a language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// greetingPrefix returns a pre-defined prefix string based on the provided language or default English
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case italian:
		prefix = italianPrefix
	default:
		prefix = englishPrefix
	}

	return
}
