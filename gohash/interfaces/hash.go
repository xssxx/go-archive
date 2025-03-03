package interfaces

type PasswordHasher interface {
	HashPassword(password string) string
}
