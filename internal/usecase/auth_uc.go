package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"log"
	"projectIntern/internal/entity"
	"projectIntern/internal/model"
	"projectIntern/internal/repository"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/email"
	"projectIntern/pkg/encode"
	"projectIntern/pkg/jwt"
)

type AuthUseCaseItf interface {
	Register(req *model.UserRegister) (*model.UserResponse, error)
	Login(req *model.UserLogin) (string, error)
	GetByVerificationCode(code string) (*entity.User, error)
	VerifyEmail(id uuid.UUID) error
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	rawCode := randstr.Hex(20)

	code := encode.Encode(rawCode)

	user := &entity.User{
		ID:               uuid.New(),
		Email:            req.Email,
		FullName:         req.FullName,
		Password:         string(hashPassword),
		VerificationCode: rawCode,
	}

	err = a.userRepo.Create(user)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailAlreadyExists) {

			return nil, err
		}
		return nil, err
	}

	err = a.email.SendEmailVerification(user, code)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := &model.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Avatar:    user.PhotoLink,
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

	if !user.IsVerified {
		return "", customerrors.ErrNotVerified
	}

	createdToken, err := a.token.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return createdToken, nil

}

func (a AuthUseCase) VerifyEmail(id uuid.UUID) error {
	user, err := a.userRepo.GetById(id)
	if err != nil {
		return err
	}

	err = a.userRepo.Update(user, model.UserUpdate{IsVerified: true})
	if err != nil {
		return err
	}

	return nil
}

func (a AuthUseCase) GetByVerificationCode(code string) (*entity.User, error) {
	user, err := a.userRepo.GetByVerificationCode(code)
	if err != nil {
		return nil, err
	}

	return user, nil
}
