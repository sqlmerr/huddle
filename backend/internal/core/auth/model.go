package core_auth

type Token struct {
	AccessToken string
}

func NewToken(accessToken string) Token {
	return Token{AccessToken: accessToken}
}
