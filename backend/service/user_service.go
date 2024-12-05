package service

import (
	"context"

	"github.com/prnndk/final-project-golang-pbkk/dto"
	"github.com/prnndk/final-project-golang-pbkk/entity"
	"github.com/prnndk/final-project-golang-pbkk/repository"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
		GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error)
		Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error)
		Delete(ctx context.Context, userId string) error
		}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

const (
	LOCAL_URL          = "http://localhost:3000"
	VERIFY_EMAIL_ROUTE = "register/verify_email"
)

func (s *userService) Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:    userReg.ID.String(),
		Name:  userReg.Name,
		Email: userReg.Email,
	}, nil
}

func (s *userService) GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := s.userRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserPaginationResponse{}, err
	}

	var datas []dto.UserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.UserResponse{
			ID:          user.ID.String(),
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (dto.UserResponse, error) {
	emails, err := s.userRepo.GetUserByEmail(ctx, nil, email)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserByEmail
	}

	return dto.UserResponse{
		ID:          emails.ID.String(),
		Name:        emails.Name,
		PhoneNumber: emails.PhoneNumber,
		Email:       emails.Email,
	}, nil
}

func (s *userService) Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error) {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUserNotFound
	}

	data := entity.User{
		ID:          user.ID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	userUpdate, err := s.userRepo.UpdateUser(ctx, nil, data)
	if err != nil {
		return dto.UserUpdateResponse{}, dto.ErrUpdateUser
	}

	return dto.UserUpdateResponse{
		ID:          userUpdate.ID.String(),
		Name:        userUpdate.Name,
		PhoneNumber: userUpdate.PhoneNumber,
		Email:       userUpdate.Email,
	}, nil
}

func (s *userService) Delete(ctx context.Context, userId string) error {
	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.ErrUserNotFound
	}

	err = s.userRepo.DeleteUser(ctx, nil, user.ID.String())
	if err != nil {
		return dto.ErrDeleteUser
	}

	return nil
}
