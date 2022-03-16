package hello

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	english            = "English"
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
)

func hello() string {
	return englishHelloPrefix + "world!"
}

func helloGreeting(user, language string) string {
	if user == "" {
		user = "world"
	}
	prefix := greetingPrefix(language)
	return fmt.Sprintf("%s%s!", prefix, user)
}

func greetingPrefix(language string) string {
	var prefix string
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
		break
	case french:
		prefix = frenchHelloPrefix
		break
	case english:
		fallthrough
	default:
		prefix = englishHelloPrefix
		break
	}
	return prefix
}
