package rest

import (
	"github.com/nopsky/project-demo/di"
	"github.com/nopsky/project-demo/pkg/transports/http"
)

func Router(
	s *http.Server,
	d *di.UseCases,
) {
	u := s.Group("/user")
	{
		uc := NewUserController(d.UserUseCase)
		u.POST("/register", uc.RegisterUser)
	}
}
