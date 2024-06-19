package skill

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
	"github.com/google/uuid"
)

type Service struct {
	skillRepo domain.ISkillRepository
	logger    logger.ILogger
}

func NewService(
	skillRepo domain.ISkillRepository,
	logger logger.ILogger,
) domain.ISkillService {
	return &Service{
		skillRepo: skillRepo,
		logger:    logger,
	}
}

func (s *Service) Create(skill *domain.Skill) (err error) {
	if skill.Name == "" {
		s.logger.Infof("должно быть указано название навыка")
		return fmt.Errorf("должно быть указано название навыка")
	}

	if skill.Description == "" {
		s.logger.Infof("должно быть указано описание навыка")
		return fmt.Errorf("должно быть указано описание навыка")
	}

	ctx := context.Background()

	err = s.skillRepo.Create(ctx, skill)
	if err != nil {
		s.logger.Infof("добавление навыка: %v", err)
		return fmt.Errorf("добавление навыка: %w", err)
	}

	return nil
}

func (s *Service) GetById(id uuid.UUID) (skill *domain.Skill, err error) {
	ctx := context.Background()

	skill, err = s.skillRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Infof("получение навыка по id: %v", err)
		return nil, fmt.Errorf("получение навыка по id: %w", err)
	}

	return skill, nil
}

func (s *Service) GetAll(page int) (skills []*domain.Skill, err error) {
	ctx := context.Background()

	skills, err = s.skillRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Infof("получение списка всех навыков: %v", err)
		return nil, fmt.Errorf("получение списка всех навыков: %w", err)
	}

	return skills, nil
}

func (s *Service) Update(skill *domain.Skill) (err error) {
	ctx := context.Background()

	err = s.skillRepo.Update(ctx, skill)
	if err != nil {
		s.logger.Infof("обновление информации о навыке: %v", err)
		return fmt.Errorf("обновление информации о навыке: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(id uuid.UUID) (err error) {
	ctx := context.Background()

	err = s.skillRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("удаление навыка по id: %v", err)
		return fmt.Errorf("удаление навыка по id: %w", err)
	}

	return nil
}
