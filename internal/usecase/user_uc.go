package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"projectIntern/internal/entity"
	"projectIntern/internal/repository"
	"projectIntern/model"
	"projectIntern/pkg/customerrors"
	"projectIntern/pkg/email"
	"projectIntern/pkg/encode"
	"projectIntern/pkg/jwt"
	"projectIntern/pkg/supabase"
)

type UserUCItf interface {
	Register(req *model.UserRegister) (interface{}, error)
	Login(req *model.UserLogin) (string, error)
	GetByVerificationCode(code string) (*entity.User, error)
	VerifyEmail(id uuid.UUID) error
	GetById(id uuid.UUID) (interface{}, error)
	UpdatePhoto(ctx *gin.Context, param model.UserUploadPhoto) error
}

type UserUC struct {
	userRepo repository.UserRepoItf
	token    jwt.JWTMakerItf
	email    email.EmailItf
	supabase supabase.SupabaseStorageItf
}

func (u UserUC) UpdatePhoto(ctx *gin.Context, param model.UserUploadPhoto) error {
	user, err := u.token.GetLoginUser(ctx)
	if err != nil {
		return err
	}

	if user.PhotoLink != "" {
		err = u.supabase.Delete(user.PhotoLink)
	}

	link, err := u.supabase.Upload(param.Photo)
	if err != nil {
		return err
	}

	err = u.userRepo.Update(user, model.UserUpdate{PhotoLink: link})
	if err != nil {
		return err
	}

	return nil
}

func NewUseUC(repo repository.UserRepoItf, token jwt.JWTMakerItf, email email.EmailItf, supabase supabase.SupabaseStorageItf) UserUCItf {
	return &UserUC{userRepo: repo, token: token, email: email, supabase: supabase}
}

func (u UserUC) GetById(id uuid.UUID) (interface{}, error) {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	userResponse := &model.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Avatar:   user.PhotoLink,
	}

	return userResponse, nil
}

func (u UserUC) Register(req *model.UserRegister) (interface{}, error) {
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

	err = u.userRepo.Create(user)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailAlreadyExists) || errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, err
		}
		return nil, err
	}

	err = u.email.SendEmailVerification(user, code)
	if err != nil {
		return nil, err
	}

	userResponse := &model.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Avatar:   user.PhotoLink,
	}
	return userResponse, nil
}

func (u UserUC) Login(req *model.UserLogin) (string, error) {
	user, err := u.userRepo.GetByEmail(req.Email)

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

	createdToken, err := u.token.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return createdToken, nil

}

func (u UserUC) VerifyEmail(id uuid.UUID) error {
	user, err := u.userRepo.GetById(id)
	if err != nil {
		return err
	}

	err = u.userRepo.Update(user, model.UserUpdate{IsVerified: true})
	if err != nil {
		return err
	}

	return nil
}

func (u UserUC) GetByVerificationCode(code string) (*entity.User, error) {
	user, err := u.userRepo.GetByVerificationCode(code)
	if err != nil {
		return nil, err
	}

	return user, nil
}
