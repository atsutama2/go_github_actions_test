package service

type HelloService struct{}

// GetHello...
func (s *HelloService) GetHello() string {
	return "Hello, World!!!!"
}
