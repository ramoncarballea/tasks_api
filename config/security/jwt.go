package security

type jwtProvider struct {
	Provider
	secret          string
	issuer          string
	audience        string
	expirationHours int
}

func NewJWTProvider(secret string, issuer string, audience string, expirationHours int) Provider {
	return &jwtProvider{
		secret:          secret,
		issuer:          issuer,
		audience:        audience,
		expirationHours: expirationHours,
	}
}

func (p *jwtProvider) CreateSignedKey(userId string) (string, error) {
	return "", nil
}

func (p *jwtProvider) ValidateSignedKey(signedKey string) bool {
	return false
}
