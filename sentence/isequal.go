package sentence

// IsEqual retorna se duas palavras são iguais
// Retorna boolean
func (s *Sentence) IsEqual(text string) bool {
	sentence2 := clean(text)

	if len(s.raw) == 0 || len(sentence2) == 0 {
		return false
	}

	return s.raw == sentence2
}
