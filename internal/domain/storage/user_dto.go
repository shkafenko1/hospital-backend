package storage

type UserDTO struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    string
}
