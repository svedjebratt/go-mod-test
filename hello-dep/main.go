package main

import (
    "fmt"
    "os"
    "io"

    "github.com/svedjebratt/go-mod-test/greetings/hello"
    "github.com/svedjebratt/go-mod-test/greetings/world"
)

func displayGreetings(w io.Writer) {
    fmt.Fprintln(w, hello.Greet())
    fmt.Fprintln(w, world.Greet())
}

func main() {   
    displayGreetings(os.Stdout)
}
