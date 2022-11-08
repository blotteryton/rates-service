package token

import "golang.org/x/exp/slices"

type Token struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewToken(id uint32, name string, code string) *Token {
	token := Token{}
	token.ID = id
	token.Name = name
	token.Code = code
	return &token
}

func Fake() []Token {
	return []Token{
		*NewToken(11419, "TON", "toncoin"),
	}
}

func Find(code string) *Token {
	tokens := Fake()
	idx := slices.IndexFunc(tokens, func(t Token) bool { return t.Code == code })
	if idx == -1 {
		return nil
	}
	return &tokens[idx]
}
