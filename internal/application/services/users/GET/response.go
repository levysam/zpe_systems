package users

type Response struct {
	Data UserResponse
}

type UserResponse struct {
	Id    string
	Name  string
	Email string
	Roles []string
}
