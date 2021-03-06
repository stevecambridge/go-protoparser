package lexer

import "github.com/yoheimuta/go-protoparser/internal/lexer/scanner"

// ReadFullIdent reads a fullIdent.
// fullIdent = ident { "." ident }
func (lex *Lexer) ReadFullIdent() (string, error) {
	lex.Next()
	if lex.Token != scanner.TIDENT {
		return "", lex.unexpected(lex.Text, "TIDENT")
	}
	fullIdent := lex.Text
	lex.Next()

	for !lex.IsEOF() {
		if lex.Token != scanner.TDOT {
			lex.UnNext()
			break
		}

		lex.Next()
		if lex.Token != scanner.TIDENT {
			return "", lex.unexpected(lex.Text, "TIDENT")
		}
		fullIdent += "." + lex.Text
		lex.Next()
	}
	return fullIdent, nil
}
