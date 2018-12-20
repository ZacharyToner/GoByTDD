//Package main greets a user and writes it to the provided writer
//Used to explore dependency injection within Go
package main

import (
	"fmt"
	"io"
	"net/http"
)

//Greet construscts a greeting using a name and writes it to a Writer interface
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
