package main

import (
	"fmt"
)

func main() {
	friends := []string{"Rahul", "Nitesh", "Sushil"}

	fmt.Printf("frineds: %v (type %T)\n", friends, friends)

	fmt.Println(len(friends))
	fmt.Println("------------")
	fmt.Println(friends[0])

	fmt.Println("------------")
	fmt.Println(friends[1:])

	fmt.Println("------------")
	for i, name := range friends {
		fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("------------")
	for _, name := range friends {
		fmt.Println(name)
	}
}
