package security

type Provider interface {
	CreateSignedKey(userId string) (string, error)
	ValidateSignedKey(signedKey string) bool
}
