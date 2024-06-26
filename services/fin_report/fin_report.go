package fin_report

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	finRepo domain.IFinancialReportRepository
	logger  logger.ILogger
}

func NewService(
	finRepo domain.IFinancialReportRepository,
	logger logger.ILogger,
) domain.IFinancialReportService {
	return &Service{
		finRepo: finRepo,
		logger:  logger,
	}
}

func (s *Service) Create(finReport *domain.FinancialReport) (err error) {
	if finReport.Revenue < 0 {
		s.logger.Infof("выручка не может быть отрицательной")
		return fmt.Errorf("выручка не может быть отрицательной")
	}

	if finReport.Costs < 0 {
		s.logger.Infof("расходы не могут быть отрицательными")
		return fmt.Errorf("расходы не могут быть отрицательными")
	}

	if finReport.Quarter > 4 || finReport.Quarter < 1 {
		s.logger.Infof("значение квартала должно находиться в отрезке от 1 до 4")
		return fmt.Errorf("значение квартала должно находиться в отрезке от 1 до 4")
	}

	now := time.Now()
	if finReport.Year > now.Year() {
		s.logger.Infof("значение года не может быть больше текущего года")
		return fmt.Errorf("значение года не может быть больше текущего года")
	}

	if finReport.Year == now.Year() && finReport.Quarter > (int(now.Month()-1)/3) {
		s.logger.Infof("нельзя добавить отчет за квартал, который еще не закончился")
		return fmt.Errorf("нельзя добавить отчет за квартал, который еще не закончился")
	}

	ctx := context.Background()

	err = s.finRepo.Create(ctx, finReport)
	if err != nil {
		s.logger.Infof("добавление финансового отчета: %v", err)
		return fmt.Errorf("добавление финансового отчета: %w", err)
	}

	return nil
}

func (s *Service) CreateByPeriod(finReportByPeriod *domain.FinancialReportByPeriod) (err error) {
	for _, report := range finReportByPeriod.Reports {
		err = s.Create(&report)
		if err != nil {
			s.logger.Infof("добавление отчетов за период: %v", err)
			return fmt.Errorf("добавление отчетов за период: %w", err)
		}
	}

	return nil
}

func (s *Service) GetById(id uuid.UUID) (finReport *domain.FinancialReport, err error) {
	ctx := context.Background()

	finReport, err = s.finRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Infof("получение финансового отчета по id: %v", err)
		return nil, fmt.Errorf("получение финансового отчета по id: %w", err)
	}

	return finReport, nil
}

func (s *Service) GetByCompany(companyId uuid.UUID, period *domain.Period) (
	finReport *domain.FinancialReportByPeriod, err error) {
	if period.StartYear > period.EndYear ||
		(period.StartYear == period.EndYear && period.StartQuarter > period.EndQuarter) {
		s.logger.Infof("дата конца периода должна быть позже даты начала")
		return nil, fmt.Errorf("дата конца периода должна быть позже даты начала")
	}

	ctx := context.Background()

	finReport, err = s.finRepo.GetByCompany(ctx, companyId, period)
	if err != nil {
		s.logger.Infof("получение финансового отчета по id компании: %v", err)
		return nil, fmt.Errorf("получение финансового отчета по id компании: %w", err)
	}

	return finReport, nil
}

func (s *Service) Update(finReport *domain.FinancialReport) (err error) {
	ctx := context.Background()

	err = s.finRepo.Update(ctx, finReport)
	if err != nil {
		s.logger.Infof("обновление отчета: %v", err)
		return fmt.Errorf("обновление отчета: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(id uuid.UUID) (err error) {
	ctx := context.Background()

	err = s.finRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("удаление отчета по id: %v", err)
		return fmt.Errorf("удаление отчета по id: %w", err)
	}

	return nil
}
