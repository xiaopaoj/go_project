package service

import (
	"context"
	"go_project/internal/model"
	"go_project/internal/repository"
)

type TestService interface {
	Find(ctx context.Context) (t model.TTable, err error)
}

type testService struct {
	repo repository.Repository
}

var _ TestService = (*testService)(nil)

func NewTest(svc *service) *testService {
	return &testService{repo: svc.repo}
}

func (s *testService) Find(ctx context.Context) (t model.TTable, err error) {
	t, _ = s.repo.Find()
	return t, nil
}
