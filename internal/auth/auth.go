package auth

type AuthRepository interface {
	// GetUserByID retrieves a user by their ID.
	GetUserByID(id string) (*User, error)
	// GetUserByEmail retrieves a user by their email
}