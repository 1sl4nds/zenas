package address

import (
	"context"
)

func NewService(r *Repository) *Service {
	return &Service{r}
}

type Service struct {
	r *Repository
}

func (s *Service) Create() {

}

func (s *Service) Delete() {

}

func (s *Service) Get(ctx context.Context) (*Address, error) {
	return s.r.SelectByID(ctx, "9db0a034-dcfd-428d-b121-ee1034bd4568")
}

func (s *Service) List() {

}
