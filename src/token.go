package src

type Token int
const (
	TOK_EOF Token = iota
	TOK_WORD
	TOK_PLUS
)

type Symbol struct {
	tokenType Token
	content string
}

