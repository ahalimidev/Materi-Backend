package service

import (
	"context"
	"fmt"
	"materi/helper"
	"materi/model/request"
	"materi/model/response"
	"materi/repository"
)

type UserService interface {
	All(ctx context.Context) ([]response.UserRespon, error)
	FindByID(ctx context.Context, IdUser string) (response.UserRespon, error)
	Create(ctx context.Context, input request.UserCreate) (response.UserRespon, error)
	Update(ctx context.Context, input request.UserUpdate) error
	Delete(ctx context.Context, IdUser string) error
	Upload(ctx context.Context, input request.UserProfile) error
}

type userService struct {
	userRepository repository.UserRepository
}

// All implements UserService
func (s *userService) All(ctx context.Context) ([]response.UserRespon, error) {
	result, err := s.userRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Create implements UserService
func (s *userService) Create(ctx context.Context, input request.UserCreate) (response.UserRespon, error) {
	User := response.UserRespon{}
	User.NamaLengkap.String = input.NamaLengkap
	User.JenisKelamin.String = input.JenisKelamin
	User.Alamat.String = input.Alamat
	User.Username.String = input.Username
	User.Password.String = helper.Base64Encode(input.Password)
	result, err := s.userRepository.Create(ctx, User)
	if err != nil {
		return result, err
	}
	fmt.Print(result)
	return result, nil
}

// Delete implements UserService
func (s *userService) Delete(ctx context.Context, IdUser string) error {
	return s.userRepository.Delete(ctx, IdUser)
}

// FindByID implements UserService
func (s *userService) FindByID(ctx context.Context, IdUser string) (response.UserRespon, error) {
	result, err := s.userRepository.FindByID(ctx, IdUser)
	if err != nil {
		return result, err
	}
	if !result.IdUser.Valid {
		return result, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Update implements UserService
func (s *userService) Update(ctx context.Context, input request.UserUpdate) error {
	result, err := s.userRepository.FindByID(ctx, input.IdUser)
	if err != nil {
		return err
	}

	if !result.IdUser.Valid {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.NamaLengkap.String = input.NamaLengkap
	result.JenisKelamin.String = input.JenisKelamin
	result.Alamat.String = input.Alamat
	result.Status.String = input.Status
	return s.userRepository.Update(ctx, result)

}

// Upload implements UserService
func (s *userService) Upload(ctx context.Context, input request.UserProfile) error {
	result, err := s.userRepository.FindByID(ctx, input.IdUser)
	if err != nil {
		return err
	}

	if !result.IdUser.Valid {
		return helper.Error("Data Tidak Ditemukan")
	}
	result.Foto.String = input.Foto
	return s.userRepository.Upload(ctx, result)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
