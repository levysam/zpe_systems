package users

type Response struct {
	CurrentPage int
	TotalPages  int64
	Data        []DataResponse
}

type DataResponse struct {
	Id    string
	Name  string
	Email string
	Roles []string
}
