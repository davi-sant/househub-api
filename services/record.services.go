package services

import (
	"context"

	"github.com/davi-sant/househub-go/models"
	"github.com/davi-sant/househub-go/repositories"
)

type RecordService struct {
	RecordRepository *repositories.RecordRepository
}

func NewRecordService(rr *repositories.RecordRepository) *RecordService {
	return &RecordService{RecordRepository: rr}
}

func (s *RecordService) Create(ctx context.Context, rc models.RecordCreate) (*models.Record, error) {
	record, err := s.RecordRepository.Create(ctx, rc)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *RecordService) FindAll(ctx context.Context) ([]models.Record, error) {
	records, err := s.RecordRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *RecordService) FindById(ctx context.Context, fr models.FindRecordById) (*models.Record, error) {
	record, err := s.RecordRepository.FindById(ctx, fr)

	if err != nil {
		return nil, err
	}

	return record, nil
}
