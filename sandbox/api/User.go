package api

// UserRole is the role of the user
type UserRole int

const (
	// UserRoleAdmin is the admin role
	UserRoleAdmin UserRole = iota
	// UserRoleViewer is the viewer role
	UserRoleViewer
)

// User is the user that is logged in
type User struct {
	Id       string
	Email    string
	Username string
	Role     UserRole
}
