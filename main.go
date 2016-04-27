package main

import (
	"fmt"
	// "os"
)
// func main() {
// 	if len(os.Args) != 2 {
// 		os.Exit(1)
// 	}
// 	fmt.Println("It's over", os.Args[1])

// 	// power := 9000
// 	// fmt.Printf("It's over %d\n", power)
// 	power := getPower()
// }

// func getPower() int {
// 	return 9001
// }


// Can only declar variable once example: so you must use = instead of := in the second indication of power

// func main() {
// 	power := 9000
// 	fmt.Printf("It's over %d\n", power)

// 	power = 9001
// 	fmt.Printf("It's also over %d\n", power)
// }


// you can declare two variables with := on the left hand side

// func main() {
// 	name, power := "Goku", 9000
// 	fmt.Printf("%s's power is over %d\n", name, power)
// }

// You can redeclare a variable with := as long as their is one undeclared variable on the left hand side

// func main() {
// 	power := 1000
// 	fmt.Printf("default power is %d\n", power)

// 	name, power := "Goku", 9000
// 	fmt.Printf("%s's power is over %d\n", name, power)
// }

// Go doesn't allow unused variables

// func main() {
// 	name, power := "Goku", 1000
// 	fmt.Printf("default power is %d\n", power)
// }