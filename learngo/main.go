package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, favFood}
	nico2 := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
	fmt.Println(nico2)
}
