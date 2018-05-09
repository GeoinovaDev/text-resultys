package sentence

// IsEqual retorna se duas palavras são iguais
// Retorna boolean
func (s *Sentence) IsEqual(text string) bool {
	sentence1 := clean(s.text)
	sentence2 := clean(text)

	if len(sentence1) == 0 || len(sentence2) == 0 {
		return false
	}

	return sentence1 == sentence2
}
