package text

import (
	"strings"

	"git.resultys.com.br/lib/text/extract"
)

// OnlyNumbers remove tudo que nao for numero do texto
func OnlyNumbers(text string) string {
	arr := extract.New(text).Regex(`[0-9]+`).ToArray()
	return strings.Join(arr, "")
}
