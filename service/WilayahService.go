package service

import (
	"context"
	"materi/helper"
	"materi/model"
	"materi/repository"
)

type WilayahService interface {
	Provinsi(ctx context.Context) ([]model.WilayahProvinsi, error)
	Kabupaten(ctx context.Context, IdProvinsi string) ([]model.WilayahKabupaten, error)
	Kecamatan(ctx context.Context, IdKabupaten string) ([]model.WilayahKecamatan, error)
	Desa(ctx context.Context, IdKecamatan string) ([]model.WilayahDesa, error)
}

type wilayahService struct {
	wilayahRepository repository.WilayahRepository
}

// Desa implements WilayahService
func (s *wilayahService) Desa(ctx context.Context, IdKecamatan string) ([]model.WilayahDesa, error) {
	result, err := s.wilayahRepository.Desa(ctx, IdKecamatan)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Kabupaten implements WilayahService
func (s *wilayahService) Kabupaten(ctx context.Context, IdProvinsi string) ([]model.WilayahKabupaten, error) {
	result, err := s.wilayahRepository.Kabupaten(ctx, IdProvinsi)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Kecamatan implements WilayahService
func (s *wilayahService) Kecamatan(ctx context.Context, IdKabupaten string) ([]model.WilayahKecamatan, error) {
	result, err := s.wilayahRepository.Kecamatan(ctx, IdKabupaten)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

// Provinsi implements WilayahService
func (s *wilayahService) Provinsi(ctx context.Context) ([]model.WilayahProvinsi, error) {
	result, err := s.wilayahRepository.Provinsi(ctx)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, helper.Error("Data Tidak Ditemukan")
	}
	return result, nil
}

func NewWilayahService(wilayahRepository repository.WilayahRepository) WilayahService {
	return &wilayahService{
		wilayahRepository: wilayahRepository,
	}
}
