package src

import (
	"fmt"
	"strings"
	"unicode"
)

var delims []byte = []byte{'(', ')', '+', '-', '*', '\'', '"', ' ', '&'}
var digits []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
func is_digit(char byte) bool {
	for i := range len(digits) {
		digit := digits[i]
		if char == digit {
			return true
		}
	}
	return false
}
func is_delimeter(char byte) bool {
	for i := range len(delims) {
		delim := delims[i]
		if char == delim {
			return true
		}
	}
	return false
}
func is_alpha(char byte) bool {
	return unicode.IsLetter(rune(char))
}
func is_whitespace(char byte) bool {
	return unicode.IsSpace(rune(char))
}

// CharType
// `CHAR_NONE` 
type CharType int
const (
	CHAR_NONE CharType = iota
	CHAR_LETTER
	CHAR_WHITESPACE
	CHAR_SPECIAL
)
const (
	WORD_UNINITIALIZED = -1
)  

type Token int
const (
	TOK_WORD Token = iota
)
type Symbol struct {
	Type Token
	Content string
}

type Word struct {
	Start int
	End   int
	Type  CharType
}

func WordInitialized(word Word) bool {
	return word.Start != -1
}
func GetCharacterTypeOfByte(b byte) CharType {
	if is_alpha(b) { 
		return CHAR_LETTER 
	} else if is_whitespace(b) { 
		return CHAR_WHITESPACE 
	} else { 
		return CHAR_SPECIAL 
	}
}
func Substring(str string, start, end int) string {
  return strings.TrimSpace(str[start:end])
}

func Tokenize(source string) []string {
	var symbols []Word
	quote := false
	word_type := CHAR_LETTER
	word_start := 0
	word_end := 0
	for i := range len(source) {
		char := source[i]
		char_type := GetCharacterTypeOfByte(char)
		if !quote && char_type != word_type || word_type == CHAR_SPECIAL {
			word_end = i
			word := Word{
				Start: word_start,
				End: word_end,
				Type: word_type,
			}
			symbols = append(symbols, word)
			word_start = i
			word_type = GetCharacterTypeOfByte(char)
		}
		if char == '"' {
			quote = !quote
			if quote {
				word_start = i
			} else {
				word_end = i
				word := Word{
					Start: word_start,
					End: word_end,
					Type: word_type,
				}
				symbols = append(symbols, word)
				word_start = i
			}
		} 
	}
	symbols = append(symbols, Word{Start: word_end, End: len(source), Type: CHAR_LETTER})

	// Go through the slices

	var tokens []string
	for i := range len(symbols) {
		symbol := symbols[i]
		token := Substring(source, symbol.Start, symbol.End)
		tokens = append(tokens, token)
	}
	for i := range len(tokens) {
		fmt.Printf("tokens[i]: %v\n", tokens[i])
	}
	return tokens
}

// func tokenize(source string) {
// 	var symbols []string
// 	var word string
// 	var quote bool = false
// 	for i := range len(source) {
// 		c := source[i]
// 		// token_character_dbg(c)
// 		if c == '"' {
// 			quote = !quote
// 		}
// 		if quote {
// 			word = word + string(c)
// 			continue
// 		}
// 		if is_whitespace(c) || is_delimeter(c) && c != '"' {
// 			symbols = append(symbols, word)
// 			if !is_whitespace(c) {
// 				symbols = append(symbols, string(c))
// 			}
// 			word = ""
// 		} else  {
// 			word = word + string(c)
// 		}
// 	}
// 	if strings.Trim(word, "") != "" {
// 		symbols = append(symbols, word)
// 	}
	
// 	// TEST BS DELETE LATER
// 	for i := range len(symbols) {
// 		fmt.Printf("[%q]\n", symbols[i])
// 	}
// }
