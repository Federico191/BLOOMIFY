package usecase

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
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
	exist, _ := a.userRepo.GetByEmail(ctx, req.Email)

	if exist != nil {
		return nil, customerrors.ErrEmailAlreadyExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:       uuid.New(),
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
		Email:     user.Email,
		FullName:  user.FullName,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return userResponse, nil
}

func (a AuthUseCase) Login(ctx context.Context, req *model.UserLogin) (string, error) {
	user, err := a.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if user == nil {
			return "", customerrors.ErrEmailInvalid
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", customerrors.ErrPasswordInvalid
	}

	createdToken, err := a.token.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return createdToken, nil

}
