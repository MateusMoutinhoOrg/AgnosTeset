package api

type UserRole int

const (
	UserRoleAdmin UserRole = iota
	UserRoleViewer
)

type User struct {
	Id       string
	Email    string
	Username string
	Role     UserRole
}
