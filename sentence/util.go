package sentence

import (
	"strings"
	"unicode"
)

func clean(text string) string {
	t := strings.ToUpper(text)
	t = strings.Trim(t, " ")
	t = RemoveAcentos(t)
	t = removeTrash(t)
	t = strings.ToUpper(t)

	return t
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func removeTrash(word string) string {
	w := strings.Replace(word, ".", " ", -1)

	w = strings.Replace(w, "&", "E", -1)
	w = strings.Replace(w, " DA ", " ", -1)
	w = strings.Replace(w, " DE ", " ", -1)
	w = strings.Replace(w, " DO ", " ", -1)

	w = strings.Replace(w, "- ME", "", -1)
	w = strings.Replace(w, "LTDA", "", -1)
	w = strings.Replace(w, "S/C", "", -1)
	w = strings.Replace(w, "S/S", "", -1)
	w = strings.Replace(w, "EPP", "", -1)
	w = strings.Replace(w, "EIRELI", "", -1)

	w = strings.Replace(w, "-", " ", -1)
	w = strings.Replace(w, "+", " ", -1)
	w = strings.Replace(w, `"`, "", -1)
	w = strings.Replace(w, `'`, "", -1)
	w = strings.Replace(w, `/`, "", -1)
	w = strings.Replace(w, `\`, "", -1)
	w = strings.Replace(w, `)`, "", -1)
	w = strings.Replace(w, `(`, "", -1)

	return w
}
