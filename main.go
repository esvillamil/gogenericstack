// Ejercicio de Programacion Generica de un stack
package main

import (
"fmt"

"github.com/esvillamil/gogenericstack/genericstack"
)


func main() {
	testStack := genericstack.EsvStack[string]()
	testStack.Push("1")
	fmt.Println("Push: 1")
	testStack.Push("2")
	fmt.Println("Push: 2")
	testStack.Push("3")
	fmt.Println("Push: 3")
	fmt.Println("Pop: ",testStack.Pop())
	testStack.Push("4")
	fmt.Println("Push: 4")
	fmt.Println("Pop: ",testStack.Pop())
	fmt.Println("Pop: ",testStack.Pop())
	fmt.Println("Pop: ",testStack.Pop())
	fmt.Println("Pop: ",testStack.Pop())
}
