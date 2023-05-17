// GithubActionsTest
package main

import (
	"log"
)

func main() {
	log.Println("Start")

	log.Printf("2 + 2 = %d\n", AddTwoNumbers(2, 2))

	log.Println("End")
}

func AddTwoNumbers(n1, n2 int) int {
	return n1 + n2
}
