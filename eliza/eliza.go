package eliza

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func init() {
}

func Greetings() string {
	return randChoice(Introductions)
}

func Goodbye() string {
	return randChoice(Goodbyes)
}

func preprocess(statement string) string {
	statement = strings.TrimRight(statement, "\n.!")
	statement = strings.ToLower(statement)
	return statement
}

func CompareStatement(statement string, compareArray []string, preprocessStatment bool) bool {
	if preprocessStatment {
		statement = preprocess(statement)
	}

	for _, compareStatement := range compareArray {
		if statement == compareStatement {
			return true
		}
	}

	return false
}

func ReplyTo(statement string) string {
	statement = preprocess(statement)

	for pattern, responses := range Psychobabble {
		fmt.Printf("Pattern %s 1\n", pattern)
		// Compile pattern and return regex to use
		re := regexp.MustCompile(pattern)
		// Find a match return match and capture
		matches := re.FindStringSubmatch(statement)

		fmt.Printf("Matches %d\n", len(matches))
		var fragment string
		if len(matches) > 1 {
			for _, match := range matches {
				fmt.Printf("Matches %s\n", match)
			}

			// Reflect the capture
			fragment = reflect(matches[1])

			fmt.Printf("Fragment %s", fragment)
		}

		if len(matches) > 0 {
			response := responses[0]

			// If there is response then sub the reflected fragment
			if strings.Contains(response, "%s") {
				response = fmt.Sprintf(response, fragment)
			}

			return response
		}
	}

	return randChoice(DefaultResponses)
}

func randChoice(list []string) string {
	randIndex := rand.Intn(len(list))
	return list[randIndex]
}

// Malay has word pairs that are reflected so
// need to cater for that as well. It has priority of
// singular for example: saya buny
func reflect(fragment string) string {
	words := strings.Split(fragment, " ")
	for i, word := range words {
		if reflectedWord, ok := ReflectedWords[word]; ok {
			words[i] = reflectedWord
		}
	}

	return strings.Join(words, " ")
}
