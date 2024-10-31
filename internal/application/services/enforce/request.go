package enforce

type Data struct {
	UserId   string
	Resource string
	Method   string
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
	if err := r.enforceCreateRule(); err != nil {
		return err
	}

	return nil
}

func (r *Request) enforceCreateRule() error {
	// Add validation...
	return nil
}
