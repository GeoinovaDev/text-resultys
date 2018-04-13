package text

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Word estrutura para comparação entre palavras
type Word struct {
	text   string
	degree float32
}

// New cria uma palavra de analise
func New(word string) *Word {
	return &Word{word, 0.66}
}

// IsLike retorna se duas palavras são semelhantes
func (w *Word) IsLike(text string) bool {
	return similar(w.text, text) >= w.degree
}

func similar(word1 string, word2 string) float32 {
	word1 = clean(word1)
	word2 = clean(word2)

	if len(word1) == 0 || len(word2) == 0 {
		return 0
	}

	arr1 := strings.Split(word1, " ")
	arr2 := strings.Split(word2, " ")
	swap := arr2

	if len(arr1) > len(arr2) {
		arr2 = arr1
		arr1 = swap
	}

	counter := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				counter++
				break
			}
		}
	}

	return float32(counter / len(arr1))
}

func clean(text string) string {
	t := strings.ToUpper(text)
	t = strings.Trim(text, " ")
	t = RemoveAcentos(text)
	t = removeTrash(text)
	t = strings.ToUpper(text)

	return t
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func removeTrash(word string) string {
	w := strings.Replace(word, ".", "", -1)

	w = strings.Replace(w, " DA ", "", -1)
	w = strings.Replace(w, " DE ", "", -1)
	w = strings.Replace(w, " DO ", "", -1)

	w = strings.Replace(w, "- ME", "", -1)
	w = strings.Replace(w, "LTDA", "", -1)
	w = strings.Replace(w, "S/C", "", -1)
	w = strings.Replace(w, "S/S", "", -1)
	w = strings.Replace(w, "EPP", "", -1)
	w = strings.Replace(w, "EIRELI", "", -1)

	w = strings.Replace(w, "-", "", -1)
	w = strings.Replace(w, `"`, "", -1)
	w = strings.Replace(w, `'`, "", -1)
	w = strings.Replace(w, `/`, "", -1)
	w = strings.Replace(w, `\`, "", -1)
	w = strings.Replace(w, `)`, "", -1)
	w = strings.Replace(w, `(`, "", -1)

	return w
}

// RemoveAcentos remove acentos da string
func RemoveAcentos(text string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	n, _, _ := transform.String(t, text)

	return n
}

// RemoveAcentos remove acentos da string
func (w *Word) RemoveAcentos() string {
	return RemoveAcentos(w.text)
}
