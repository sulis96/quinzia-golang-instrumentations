package service

import (
	"context"

	"github.com/sulis96/quinzia-golang-instrumentations/internal/model"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/repository"
)

type (
	service struct {
		repo repository.IRepository
	}

	IService interface {
		CreateMember(ctx context.Context, data model.Member) error
		ReadMember(ctx context.Context) (data []model.Member, err error)
	}
)

func NewService(r repository.IRepository) IService {
	return &service{
		repo: r,
	}
}

func (s *service) CreateMember(ctx context.Context, data model.Member) error {
	return s.repo.InsertMember(ctx, data)
}

func (s *service) ReadMember(ctx context.Context) (data []model.Member, err error) {
	return s.repo.ReadMember(ctx)
}
