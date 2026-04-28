package usecase

type Hello struct {
}

func NewHello() Hello {
	return Hello{}
}

func (h Hello) Say() string {
	return "Hello!"
}
