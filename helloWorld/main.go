package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, language string) (string, string, error) {
	if name == "" && language == "" {
		name = "world"
		return "Hello, " + name + "! noob", "english", nil
	} else if name == "anmol" {
		return "Hello, " + name + "! notnoob", "english", nil
	}

	return "Hello, " + name + "! noob", "english", nil
}
