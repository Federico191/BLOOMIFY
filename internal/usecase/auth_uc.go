package usecase

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/jwt"
)

type AuthUseCaseItf interface {
	Register(ctx context.Context, req *model.UserRegister) (*model.UserResponse, error)
	Login(ctx context.Context, req *model.UserLogin) (string, error)
}

type AuthUseCase struct {
	userRepo repository.UserRepoItf
	token    jwt.JWTMakerItf
}

func NewAuthUseCase(userRepo repository.UserRepoItf, token jwt.JWTMakerItf) AuthUseCaseItf {
	return AuthUseCase{userRepo: userRepo, token: token}
}

func (a AuthUseCase) Register(ctx context.Context, req *model.UserRegister) (*model.UserResponse, error) {
	exist, _ := a.userRepo.GetByUsername(ctx, req.Username)

	if exist != nil {
		return nil, errors.New("user already exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		FullName: req.FullName,
		Avatar:   req.Avatar,
		Password: string(hashPassword),
	}

	err = a.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	userResponse := &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse, nil
}

func (a AuthUseCase) Login(ctx context.Context, req *model.UserLogin) (string, error) {
	user, err := a.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	createdToken, err := a.token.CreateToken(user.Username)
	if err != nil {
		return "", err
	}
	return createdToken, nil

}
