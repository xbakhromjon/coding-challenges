package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
)

func main() {
	chi.RegisterMethod("/git-contribution-visualisation-tool")
	fmt.Println("Git contribution visualisation")
}
