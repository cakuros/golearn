package main

import "fmt"
import "io"
import "os"

func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main(){
	Greet(os.Stdout, "Elodie")
}