package main

import "fmt"

type emptyIface interface{}

func main() {
	var i emptyIface
	i = 23
	fmt.Printf("%T, %v\n", i, i)

	var j emptyIface
	fmt.Printf("%T, %v\n", j, j)

	empIface(i, j)
}

func empIface(ifaces ...emptyIface) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}
