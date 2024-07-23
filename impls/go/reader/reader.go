package reader

import (
	"regexp"
)

type Reader struct {
	tokens []string
	pos    int
}

func (r *Reader) haveNext() bool {
	return r.pos < len(r.tokens)
}

func (r *Reader) next() string {
	t := r.tokens[r.pos]
	r.pos += 1
	return t
}

func (r *Reader) peek() string {
	return r.tokens[r.pos]
}

func ReadStr(s string) string {
	tokens := tokenize(s)
	reader := &Reader{tokens: tokens}
	return ReadForm(reader)
}

func ReadForm(r *Reader) string {
	result := ""
	for r.haveNext() {
		if result != "" {
			result += " "
		}
		result += r.next()
	}
	return result
}

func ReadList() {

}

func ReadAtom() {

}

func tokenize(s string) []string {
	reg := regexp.MustCompile(`[\s,]*(~@|[\[\]{}()'` + "`" + `~^@]|"(?:\\.|[^\\"])*"?|;.*|[^\s\[\]{}('"` + "`" + `,;)]*)`)
	subs := reg.FindAllStringSubmatch(s, -1)
	result := make([]string, 0, len(subs))
	for _, v := range subs {
		if v[1] == "" {
			continue
		}
		result = append(result, v[1])
	}
	return result
}
