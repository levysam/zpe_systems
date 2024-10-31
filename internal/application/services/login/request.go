package login

import "go-skeleton/internal/application/domain/users"

type Data struct {
	Email    string
	Password string
}

type Request struct {
	Data *Data
	User *users.Users
	Err  error
}

func NewRequest(data *Data) Request {
	domain := &users.Users{}
	return Request{
		Data: data,
		User: domain,
	}
}

func (r *Request) Validate() error {
	if err := r.loginCreateRule(); err != nil {
		return err
	}

	return nil
}

func (r *Request) loginCreateRule() error {
	// Add validation...
	return nil
}
