package service

import "testing"

func TestHelloService_GetHello(t *testing.T) {
	tests := []struct {
		name string
		s    *HelloService
		want string
	}{
		{
			name: "returns correct greeting message",
			s:    &HelloService{},
			want: "Hello, World!!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HelloService{}
			if got := s.GetHello(); got != tt.want {
				t.Errorf("HelloService.GetHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
