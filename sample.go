// Onmyodo test sample — Go
package main

import "fmt"

func greet(name string) string {
	return fmt.Sprintf("Hello from Onmyodo, %s!", name)
}

func main() {
	fmt.Println(greet("world"))
}
