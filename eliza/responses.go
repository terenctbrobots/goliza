package eliza

var Introductions = []string{
	"Apa Khabar",
	"Selamat Datang",
}

var Goodbyes = []string{
	"Selamat Tinggal",
	"Bye",
}

var Psychobabble = map[string][]string{
	`saya mahu (.*)`: {
		"Kenapa awak mahu %s?",
		"Adakah ia benar-benar membantu anda untuk mendapatkan %s?",
		"Adakah anda pasti memerlukan %s?",
	},
}

// If ELIZA doesn't understand the question, then it will reply with one of
// these default responses
var DefaultResponses = []string{
	"Bagi beritahu saya lagi.",
	"Mari ubah fokus sedikit... Beritahu saya tentang keluarga anda.",
	"Bolehkah anda menghuraikannya?",
	"Saya faham.",
	"Sugguh menarik.",
	"Saya faham. Dan apa yang memberitahu anda?",
	"Bgaimana itu membuat anda merasa?",
	"Bagaimana perasaan anda apabila mengatakannya",
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
