package cli

import "fmt"

func ErrorHandler(message string) {
	fmt.Println("Error Occur")
	fmt.Println(message)
}
