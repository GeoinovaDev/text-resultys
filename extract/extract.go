package extract

import (
	"regexp"
	"strings"

	"git.resultys.com.br/lib/lower/convert/decode"
)

// Extract estrutura para clip de texto
type Extract struct {
	list []string
}

var regex map[string]*regexp.Regexp
var b = false

func init() {
	if b {
		return
	}

	regex = make(map[string]*regexp.Regexp)
	b = true
}

// Clone clona o clip atual
func (e *Extract) Clone() *Extract {
	nw := &Extract{}

	for i := 0; i < len(e.list); i++ {
		nw.list = append(nw.list, e.list[i])
	}

	return nw
}

// New cria a estrutura de recorte
func New(content string) *Extract {
	return &Extract{list: []string{content}}
}

// First retorna o primeiro item recortado
func (e *Extract) First() string {
	if len(e.list) == 0 {
		return ""
	}

	return e.list[0]
}

// ToArray converte para array
func (e *Extract) ToArray() []string {
	return e.list
}

func compile(pattern string) *regexp.Regexp {
	if val, ok := regex[pattern]; ok {
		return val
	} else {
		c, _ := regexp.Compile(pattern)
		regex[pattern] = c

		return c
	}
}

// Regex executa uma ER sobre as list clipada
func (e *Extract) Regex(pattern string) *Extract {
	list := e.list
	e.list = []string{}

	for i := 0; i < len(list); i++ {
		c := compile(pattern)
		arr := c.FindAllString(list[i], -1)
		for j := 0; j < len(arr); j++ {
			e.list = append(e.list, arr[j])
		}
	}

	return e
}

// Clips recorta fragmentos dentro de conteudo
func (e *Extract) Clips(parts ...string) *Extract {
	contents := e.list
	e.list = []string{}
	if len(parts) < 2 {
		panic("str.IndexOf: numero de parametros incorreto")
	}

	for j := 0; j < len(contents); j++ {
		index := 0
		content := contents[j]

		if parts[0] == "IGNORE" {
			content = " IGNORE" + content
		}

	loop:
		for {
			if index >= len(content) {
				break
			}

			for i := 0; i < len(parts)-1; i++ {
				part := parts[i]
				_index := strings.Index(string(content[index:]), part)
				if _index == -1 {
					break loop
				}
				index += _index + len(part)
			}

			if index == 0 {
				break
			}

			f := strings.Index(string(content[index:]), parts[len(parts)-1])
			if f == -1 {
				break
			}
			f += index

			str := string(content[index:f])
			str = decode.HTML(str)
			str = strings.Trim(str, " ")

			e.list = append(e.list, str)
			index = f
		}
	}

	return e
}
