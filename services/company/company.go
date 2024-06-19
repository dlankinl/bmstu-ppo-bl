package company

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
	"github.com/google/uuid"
)

type Service struct {
	companyRepo domain.ICompanyRepository
	logger      logger.ILogger
}

func NewService(
	companyRepo domain.ICompanyRepository,
	logger logger.ILogger,
) domain.ICompanyService {
	return &Service{
		companyRepo: companyRepo,
		logger:      logger,
	}
}

func (s *Service) Create(company *domain.Company) (err error) {
	if company.Name == "" {
		s.logger.Infof("должно быть указано название компании")
		return fmt.Errorf("должно быть указано название компании")
	}

	if company.City == "" {
		s.logger.Infof("должно быть указано название города")
		return fmt.Errorf("должно быть указано название города")
	}

	ctx := context.Background()

	err = s.companyRepo.Create(ctx, company)
	if err != nil {
		s.logger.Infof("добавление компании: %v", err)
		return fmt.Errorf("добавление компании: %w", err)
	}

	return nil
}

func (s *Service) GetById(id uuid.UUID) (company *domain.Company, err error) {
	ctx := context.Background()

	company, err = s.companyRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Infof("получение компании по id: %v", err)
		return nil, fmt.Errorf("получение компании по id: %w", err)
	}

	return company, nil
}

func (s *Service) GetByOwnerId(id uuid.UUID, page int) (companies []*domain.Company, err error) {
	ctx := context.Background()

	companies, err = s.companyRepo.GetByOwnerId(ctx, id, page)
	if err != nil {
		s.logger.Infof("получение списка компаний по id владельца: %v", err)
		return nil, fmt.Errorf("получение списка компаний по id владельца: %w", err)
	}

	return companies, nil
}

func (s *Service) GetAll(page int) (companies []*domain.Company, err error) {
	ctx := context.Background()

	companies, err = s.companyRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Infof("получение списка всех компаний: %v", err)
		return nil, fmt.Errorf("получение списка всех компаний: %w", err)
	}

	return companies, nil
}

func (s *Service) Update(company *domain.Company) (err error) {
	ctx := context.Background()

	err = s.companyRepo.Update(ctx, company)
	if err != nil {
		s.logger.Infof("обновление информации о компании: %v", err)
		return fmt.Errorf("обновление информации о компании: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(id uuid.UUID) (err error) {
	ctx := context.Background()

	err = s.companyRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("удаление компании по id: %v", err)
		return fmt.Errorf("удаление компании по id: %w", err)
	}

	return nil
}
