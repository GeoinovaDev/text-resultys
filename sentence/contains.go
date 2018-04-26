package sentence

import (
	"strings"
)

// Contains verifica se um a sentenca contem um texto
// Return bool
func (s *Sentence) Contains(text string, total int) bool {
	sentence := clean(s.text)
	text = clean(text)

	if len(sentence) == 0 || len(text) == 0 {
		return false
	}

	contains := 0
	p := strings.Split(text, " ")
	for i := 0; i < len(p); i++ {
		if strings.Index(sentence, p[i]) > -1 {
			contains++
		}
	}

	if contains >= total {
		return true
	}

	return false
}
