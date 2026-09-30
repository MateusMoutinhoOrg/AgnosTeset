package api

type UserRole int

const (
	UserRoleAdmin UserRole = iota
	UserRoleViewer
)

// User is the user that is logged in
type User struct {
	Id       string
	Email    string
	Username string
	Role     UserRole
}
