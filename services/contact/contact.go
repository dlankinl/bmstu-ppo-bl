package contact

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
	"github.com/google/uuid"
)

type Service struct {
	contactRepo domain.IContactsRepository
	logger      logger.ILogger
}

func NewService(
	conRepo domain.IContactsRepository,
	logger logger.ILogger,
) domain.IContactsService {
	return &Service{
		contactRepo: conRepo,
		logger:      logger,
	}
}

func (s *Service) Create(contact *domain.Contact) (err error) {
	if contact.Name == "" {
		s.logger.Infof("должно быть указано название средства связи")
		return fmt.Errorf("должно быть указано название средства связи")
	}

	if contact.Value == "" {
		s.logger.Infof("должно быть указано значение средства связи")
		return fmt.Errorf("должно быть указано значение средства связи")
	}

	ctx := context.Background()

	err = s.contactRepo.Create(ctx, contact)
	if err != nil {
		s.logger.Infof("добавление средства связи: %v", err)
		return fmt.Errorf("добавление средства связи: %w", err)
	}

	return nil
}

func (s *Service) GetById(id uuid.UUID) (contact *domain.Contact, err error) {
	ctx := context.Background()

	contact, err = s.contactRepo.GetById(ctx, id)
	if err != nil {
		s.logger.Infof("получение средства связи по id: %v", err)
		return nil, fmt.Errorf("получение средства связи по id: %w", err)
	}

	return contact, nil
}

func (s *Service) GetByOwnerId(id uuid.UUID, page int) (contacts []*domain.Contact, err error) {
	ctx := context.Background()

	contacts, err = s.contactRepo.GetByOwnerId(ctx, id, page)
	if err != nil {
		s.logger.Infof("получение всех средств связи по id владельца: %v", err)
		return nil, fmt.Errorf("получение всех средств связи по id владельца: %w", err)
	}

	return contacts, nil
}

func (s *Service) Update(contact *domain.Contact) (err error) {
	ctx := context.Background()

	err = s.contactRepo.Update(ctx, contact)
	if err != nil {
		s.logger.Infof("обновление информации о средстве связи: %v", err)
		return fmt.Errorf("обновление информации о средстве связи: %w", err)
	}

	return nil
}

func (s *Service) DeleteById(id uuid.UUID) (err error) {
	ctx := context.Background()

	err = s.contactRepo.DeleteById(ctx, id)
	if err != nil {
		s.logger.Infof("удаление средства связи по id: %v", err)
		return fmt.Errorf("удаление средства связи по id: %w", err)
	}

	return nil
}
