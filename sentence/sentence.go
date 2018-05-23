package sentence

import (
	"git.resultys.com.br/lib/text"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Sentence estrutura para comparação entre sentencas
type Sentence struct {
	text   string
	degree float32
}

// New cria uma sentenca de analise
func New(text string) *Sentence {
	return &Sentence{text, 0.66}
}

// SetDegree ...
func (s *Sentence) SetDegree(d float32) *Sentence {
	s.degree = d

	return s
}

// RemoveAcentos remove acentos da string
func RemoveAcentos(text string) string {
	// Retorna string
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	n, _, _ := transform.String(t, text)

	return n
}

// RemoveAcentos remove acentos da string
// Retorna string
func (s *Sentence) RemoveAcentos() string {
	return RemoveAcentos(s.text)
}

// OnlyNumber extrai todos os numeros do texto
// Retorna string
func (s *Sentence) OnlyNumber() string {
	return text.OnlyNumbers(s.text)
}
