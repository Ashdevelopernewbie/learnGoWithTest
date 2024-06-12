package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name, language string) (string, string, error) {
	if language == "" {
		language = "english"
	}

	switch language {
	case "spanish":
		if name == "" {
			name = "Mundo"
		}
		const spanishPrefix = "¡Hola, "
		const spanishSuffix = "! novato"
		return spanishPrefix + name + spanishSuffix, language, nil
	case "english":
		if name == "" {
			name = "world"
		}
		const englishPrefix = "Hello, "
		const englishSuffix = "! noob"
		if name == "anmol" {
			return englishPrefix + name + "! notnoob", language, nil
		}
		return englishPrefix + name + englishSuffix, language, nil
	default:
		return "", "", fmt.Errorf("unsupported language: %v", language)
	}
}
