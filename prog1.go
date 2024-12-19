package main

import (
	"fmt"
	"os"
	"time"
)

var value int = 2

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	fmt.Printf("Adresse von 'value': %p\n", &value)
	fmt.Println("Variable 'value' wurde auf 2 gesetzt. Warte auf Speicherzugriff...")
	for {
		time.Sleep(time.Second)
	}
}
