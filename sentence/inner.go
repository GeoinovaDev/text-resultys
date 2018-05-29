package sentence

import (
	"strings"
)

// IsInner ...
func (s *Sentence) IsInner(text string) bool {
	text = clean(text)

	if len(s.raw) == 0 || len(text) == 0 {
		return false
	}

	p := strings.Split(s.raw, " ")
	count := 0
	for i := 0; i < len(p); i++ {
		if !isValidWord(p[i]) {
			continue
		}

		if strings.Index(text, p[i]) > -1 {
			count++
		}
	}

	if float32(count)/float32(len(p)) > s.degree {
		return true
	}

	return false
}

func isValidWord(word string) bool {
	commons := []string{"DE", "DO", "DA", "EM", "NO", "NA", "PARA"}
	for i := 0; i < len(commons); i++ {
		if commons[i] == word {
			return false
		}
	}

	if len(word) == 1 {
		return false
	}

	return true
}
