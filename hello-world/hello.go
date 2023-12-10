package helloworld

import "fmt"

const (
	// Languages
	spanish = "spanish"
	french  = "french"

	// Prefixes
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func Greet(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(lang)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
