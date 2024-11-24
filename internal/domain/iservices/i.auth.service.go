package services

type AuthService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) bool
}
