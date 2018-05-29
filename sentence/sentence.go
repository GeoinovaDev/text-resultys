package sentence

import (
	"strings"

	"git.resultys.com.br/lib/text"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Sentence estrutura para comparação entre sentencas
type Sentence struct {
	raw    string
	text   string
	degree float32
}

// New cria uma sentenca de analise
func New(text string) *Sentence {
	return &Sentence{
		raw:    clean(text),
		text:   text,
		degree: 0.66,
	}
}

// SetDegree ...
func (s *Sentence) SetDegree(d float32) *Sentence {
	s.degree = d

	return s
}

// Normaliza ...
func Normaliza(text string) string {
	// Retorna string
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	n, _, _ := transform.String(t, text)

	return n
}

// RemoveAcentos remove acentos da string
func RemoveAcentos(text string) string {
	withAccents := []string{
		"Á", "À", "á", "à",
		"É", "È", "é", "è",
		"Í", "Ì", "í", "ì",
		"Ó", "Ò", "ó", "ò",
		"Ú", "Ù", "ú", "ù",
		"Â", "â",
		"Ê", "ê",
		"Î", "î",
		"Ô", "ô",
		"Ç", "ç",
	}

	withoutAccents := []string{
		"A", "A", "a", "a",
		"E", "E", "e", "e",
		"I", "I", "i", "i",
		"O", "O", "o", "o",
		"U", "U", "u", "u",
		"A", "a",
		"E", "e",
		"I", "i",
		"O", "o",
		"C", "c",
	}

	for i := 0; i < len(withAccents); i++ {
		text = strings.Replace(text, withAccents[i], withoutAccents[i], -1)
	}

	return text
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
