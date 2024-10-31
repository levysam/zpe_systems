package users

import "go-skeleton/internal/application/domain/users"

type Data struct {
	Id string
}

type Request struct {
	Data   *Data
	Domain *users.Users
	Err    error
}

func NewRequest(data *Data) Request {
	domain := &users.Users{}
	return Request{
		Data:   data,
		Domain: domain,
	}
}

func (r *Request) Validate() error {
	return nil
}
