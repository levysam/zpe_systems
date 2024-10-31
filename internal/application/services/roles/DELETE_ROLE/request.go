package roles

type Data struct {
	RoleName string
}

type Request struct {
	Data *Data
	Err  error
}

func NewRequest(data *Data) Request {
	return Request{
		Data: data,
	}
}

func (r *Request) Validate() error {
	if err := r.rolesCreateRule(); err != nil {
		return err
	}

	return nil
}

func (r *Request) rolesCreateRule() error {
	// Add validation...
	return nil
}
