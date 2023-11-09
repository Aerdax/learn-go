package di

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
