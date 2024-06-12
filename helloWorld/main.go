package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	} else if name == "anmol" {
		return "Hello, " + name + "! notnoob", nil
	}

	return "Hello, " + name + "! noob", nil
}
