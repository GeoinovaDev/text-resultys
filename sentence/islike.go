package sentence

// IsLike retorna se duas palavras são semelhantes
// Retorna boolean
func (s *Sentence) IsLike(text string) bool {
	return similar(s.text, text) >= s.degree
}
