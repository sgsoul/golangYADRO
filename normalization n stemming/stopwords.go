package main

// оставлено 2 версии для удобного добавления новых слов
// (при копировании из новой библиотеки будет удобно вставить в подходящий блок)

var Stopwords = map[string]bool{
	"i":      true,
	"me":     true,
	"my":     true,
	"you":    true,
	"your":   true,
	"he":     true,
	"she":    true,
	"him":    true,
	"his":    true,
	"her":    true,
	"it":     true,
	"we":     true,
	"us":     true,
	"our":    true,
	"they":   true,
	"them":   true,
	"their":  true,
	"this":   true,
	"that":   true,
	"these":  true,
	"those":  true,
	"am":     true,
	"is":     true,
	"are":    true,
	"was":    true,
	"were":   true,
	"be":     true,
	"been":   true,
	"being":  true,
	"have":   true,
	"has":    true,
	"had":    true,
	"do":     true,
	"does":   true,
	"did":    true,
	"shall":  true,
	"will":   true,
	"should": true,
	"would":  true,
	"may":    true,
	"might":  true,
	"must":   true,
	"can":    true,
	"could":  true,
	"of":     true,
	"a":      true,
	"an":     true,
	"the":    true,
	"as":     true,
}

func IsStopWord(word string) bool {
	switch word {
	case "a", "about", "above", "after", "again", "against", "all", "am", "an",
		"and", "any", "are", "as", "at", "be", "because", "been", "before",
		"being", "below", "between", "both", "but", "by", "can", "did", "do",
		"does", "doing", "don", "down", "during", "each", "few", "for", "from",
		"further", "had", "has", "have", "having", "he", "her", "here", "hers",
		"herself", "him", "himself", "his", "how", "i", "if", "in", "into", "is",
		"it", "its", "itself", "just", "me", "more", "most", "my", "myself",
		"no", "nor", "not", "now", "of", "off", "on", "oh", "once", "only", "or",
		"other", "our", "ours", "ourselves", "out", "over", "own", "s", "same",
		"she", "should", "so", "some", "such", "t", "than", "that", "the", "their",
		"theirs", "them", "themselves", "then", "there", "these", "they",
		"this", "those", "through", "to", "too", "under", "until", "up",
		"very", "was", "we", "were", "what", "when", "where", "which", "while",
		"who", "whom", "why", "will", "with", "you", "your", "yours", "yourself",
		"yourselves":
		return true
	}
	return false
}
