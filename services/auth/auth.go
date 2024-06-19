package auth

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/base"
	"github.com/dlankinl/bmstu-ppo-bl/pkg/logger"
)

type Service struct {
	authRepo domain.IAuthRepository
	crypto   base.IHashCrypto
	jwtKey   string
	logger   logger.ILogger
}

func NewService(
	repo domain.IAuthRepository,
	crypto base.IHashCrypto,
	jwtKey string,
	logger logger.ILogger,
) domain.IAuthService {
	return &Service{
		authRepo: repo,
		crypto:   crypto,
		jwtKey:   jwtKey,
		logger:   logger,
	}
}

func (s *Service) Register(authInfo *domain.UserAuth) (err error) {
	if authInfo.Username == "" {
		s.logger.Infof("должно быть указано имя пользователя")
		return fmt.Errorf("должно быть указано имя пользователя")
	}

	if authInfo.Password == "" {
		s.logger.Infof("должен быть указан пароль")
		return fmt.Errorf("должен быть указан пароль")
	}

	hashedPass, err := s.crypto.GenerateHashPass(authInfo.Password)
	if err != nil {
		s.logger.Infof("генерация хэша: %v", err)
		return fmt.Errorf("генерация хэша: %w", err)
	}

	authInfo.HashedPass = hashedPass

	ctx := context.Background()

	err = s.authRepo.Register(ctx, authInfo)
	if err != nil {
		s.logger.Infof("регистрация пользователя: %v", err)
		return fmt.Errorf("регистрация пользователя: %w", err)
	}

	return nil
}

func (s *Service) Login(authInfo *domain.UserAuth) (token string, err error) {
	if authInfo.Username == "" {
		s.logger.Infof("должно быть указано имя пользователя")
		return "", fmt.Errorf("должно быть указано имя пользователя")
	}

	if authInfo.Password == "" {
		s.logger.Infof("должен быть указан пароль")
		return "", fmt.Errorf("должен быть указан пароль")
	}

	ctx := context.Background()

	userAuth, err := s.authRepo.GetByUsername(ctx, authInfo.Username)
	if err != nil {
		s.logger.Infof("получение пользователя по username: %v", err)
		return "", fmt.Errorf("получение пользователя по username: %w", err)
	}

	if !s.crypto.CheckPasswordHash(authInfo.Password, userAuth.HashedPass) {
		s.logger.Infof("неверный пароль")
		return "", fmt.Errorf("неверный пароль")
	}

	token, err = base.GenerateAuthToken(authInfo.Username, s.jwtKey, userAuth.Role)
	if err != nil {
		s.logger.Infof("генерация токена: %v", err)
		return "", fmt.Errorf("генерация токена: %w", err)
	}

	return token, nil
}
