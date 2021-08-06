package eliza

var Introductions = []string{
	"Hello, Apa Khabar",
	"Selamat Datang",
}

var Goodbyes = []string{
	"Selamat Tinggal",
	"Bye",
}

var Psychobabble = map[string][]string{
	`i need (.*)`: {
		"Why do you need %s?",
		"Would it really help you to get %s?",
		"Are you sure you need %s?",
	},
}

// This is a table to reflect words in question fragments inside the response.
// For example, the phrase "your jacket" in "I want your jacket" should be
// reflected to "my jacket" in the response.
var ReflectedWords = map[string]string{
	"am":     "are",
	"was":    "were",
	"i":      "you",
	"i'd":    "you would",
	"i've":   "you have",
	"i'll":   "you will",
	"my":     "your",
	"are":    "am",
	"you've": "I have",
	"you'll": "I will",
	"your":   "my",
	"yours":  "mine",
	"you":    "me",
	"me":     "you",
}
