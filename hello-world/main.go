package main

import (
    "fmt"
    "os"
    "io"

    "github.com/svedjebratt/hello-world/hello"
    "github.com/svedjebratt/hello-world/world"
)

func displayGreetings(w io.Writer) {
    fmt.Fprintln(w, hello.Greet())
    fmt.Fprintln(w, world.Greet())
}

func main() {   
    displayGreetings(os.Stdout)
}
