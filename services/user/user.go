package user

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
	"github.com/google/uuid"
	"strings"
)

type Service struct {
	userRepo     domain.IUserRepository
	companyRepo  domain.ICompanyRepository
	actFieldRepo domain.IActivityFieldRepository
	logger       logger.ILogger
}

func NewService(
	userRepo domain.IUserRepository,
	companyRepo domain.ICompanyRepository,
	actFieldRepo domain.IActivityFieldRepository,
	logger logger.ILogger,
) domain.IUserService {
	return &Service{
		userRepo:     userRepo,
		companyRepo:  companyRepo,
		actFieldRepo: actFieldRepo,
		logger:       logger,
	}
}

func (s *Service) Create(user *domain.User) (err error) {
	if user.Gender != "m" && user.Gender != "w" {
		s.logger.Infof("неизвестный пол")
		return fmt.Errorf("неизвестный пол")
	}

	if user.City == "" {
		s.logger.Infof("должно быть указано название города")
		return fmt.Errorf("должно быть указано название города")
	}

	if user.Birthday.IsZero() {
		s.logger.Infof("должна быть указана дата рождения")
		return fmt.Errorf("должна быть указана дата рождения")
	}

	if user.FullName == "" {
		s.logger.Infof("должны быть указаны ФИО")
		return fmt.Errorf("должны быть указаны ФИО")
	}

	if len(strings.Split(user.FullName, " ")) != 3 {
		s.logger.Infof("некорректное количество слов (должны быть фамилия, имя и отчество)")
		return fmt.Errorf("некорректное количество слов (должны быть фамилия, имя и отчество)")
	}

	ctx := context.Background()

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.Infof("создание пользователя: %v", err)
		return fmt.Errorf("создание пользователя: %w", err)
	}

	return nil
}

func (s *Service) GetByUsername(username string) (user *domain.User, err error) {
	ctx := context.Background()

	user, err = s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		s.logger.Infof("получение пользователя по username: %v", err)
		return nil, fmt.Errorf("получение пользователя по username: %w", err)
	}

	return user, nil
}

func (s *Service) GetById(userId uuid.UUID) (user *domain.User, err error) {
	ctx := context.Background()

	user, err = s.userRepo.GetById(ctx, userId)
	if err != nil {
		s.logger.Infof("получение пользователя по id: %v", err)
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	return user, nil
}

func (s *Service) GetAll(page int) (users []*domain.User, err error) {
	ctx := context.Background()

	users, err = s.userRepo.GetAll(ctx, page)
	if err != nil {
		s.logger.Infof("получение списка всех пользователей: %v", err)
		return nil, fmt.Errorf("получение списка всех пользователей: %w", err)
	}

	return users, nil
}

func (s *Service) Update(user *domain.User) (err error) {
	ctx := context.Background()

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		s.logger.Infof("обновление информации о пользователе: %v", err)
		return fmt.Errorf("обновление информации о пользователе: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(id uuid.UUID) (err error) {
	ctx := context.Background()

	err = s.userRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("удаление пользователя по id: %v", err)
		return fmt.Errorf("удаление пользователя по id: %w", err)
	}

	return nil
}
