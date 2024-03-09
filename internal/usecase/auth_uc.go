package usecase

import (
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/email"
	"projectIntern/pkg/jwt"
)

type AuthUseCaseItf interface {
	Register(req *model.UserRegister) (*model.UserResponse, error)
	Login(req *model.UserLogin) (string, error)
}
type AuthUseCase struct {
	userRepo repository.UserRepoItf
	token    jwt.JWTMakerItf
	email    email.EmailItf
}

func NewAuthUseCase(userRepo repository.UserRepoItf, token jwt.JWTMakerItf, email email.EmailItf) AuthUseCaseItf {
	return AuthUseCase{userRepo: userRepo, token: token, email: email}
}

func (a AuthUseCase) Register(req *model.UserRegister) (*model.UserResponse, error) {
	exist, _ := a.userRepo.GetByEmail(req.Email)

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
		Password: string(hashPassword),
	}

	err = a.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	verificationUrl := randstr.Hex(20)

	err = a.email.SendEmailVerification(user, verificationUrl)
	if err != nil {
		return nil, err
	}

	user.IsVerified = true

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

func (a AuthUseCase) Login(req *model.UserLogin) (string, error) {
	user, err := a.userRepo.GetByEmail(req.Email)
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
