package repository

import (
	"context"
	"database/sql"
	"materi/model"
)

type WilayahRepository interface {
	Provinsi(ctx context.Context) ([]model.WilayahProvinsi, error)
	Kabupaten(ctx context.Context, IdProvinsi string) ([]model.WilayahKabupaten, error)
	Kecamatan(ctx context.Context, IdKabupaten string) ([]model.WilayahKecamatan, error)
	Desa(ctx context.Context, IdKecamatan string) ([]model.WilayahDesa, error)
}

type WilayahConnetion struct {
	tx *sql.DB
}

// Desa implements WilayahRepository
func (r *WilayahConnetion) Desa(ctx context.Context, IdKecamatan string) ([]model.WilayahDesa, error) {
	SQL := "SELECT desa_id,desa FROM ref_desa where kecamatan_id = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdKecamatan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	desa_ := []model.WilayahDesa{}
	for rows.Next() {
		desa := model.WilayahDesa{}
		err := rows.Scan(&desa.IdDesa, &desa.Desa)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		desa_ = append(desa_, desa)
	}
	return desa_, nil
}

// Kabupaten implements WilayahRepository
func (r *WilayahConnetion) Kabupaten(ctx context.Context, IdProvinsi string) ([]model.WilayahKabupaten, error) {
	SQL := "SELECT kabupaten_id,kabupaten FROM ref_kabupaten where provinsi_id  = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdProvinsi)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kabupaten_ := []model.WilayahKabupaten{}
	for rows.Next() {
		kabupaten := model.WilayahKabupaten{}
		err := rows.Scan(&kabupaten.IdKabupaten, &kabupaten.Kabupaten)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		kabupaten_ = append(kabupaten_, kabupaten)
	}
	return kabupaten_, nil
}

// Kecamatan implements WilayahRepository
func (r *WilayahConnetion) Kecamatan(ctx context.Context, IdKabupaten string) ([]model.WilayahKecamatan, error) {
	SQL := "SELECT kecamatan_id,kecamatan FROM ref_kecamatan where kabupaten_id  = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdKabupaten)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kecamatan_ := []model.WilayahKecamatan{}
	for rows.Next() {
		kecamatan := model.WilayahKecamatan{}
		err := rows.Scan(&kecamatan.IdKecamatan, &kecamatan.Kecamatan)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		kecamatan_ = append(kecamatan_, kecamatan)
	}
	return kecamatan_, nil
}

// Provinsi implements WilayahRepository
func (r *WilayahConnetion) Provinsi(ctx context.Context) ([]model.WilayahProvinsi, error) {
	SQL := "SELECT provinsi_id,provinsi FROM ref_provinsi"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	provinsi_ := []model.WilayahProvinsi{}
	for rows.Next() {
		provinsi := model.WilayahProvinsi{}
		err := rows.Scan(&provinsi.IdProvinsi, &provinsi.Provinsi)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		provinsi_ = append(provinsi_, provinsi)
	}
	return provinsi_, nil
}

func NewWilayahRepository(DB *sql.DB) WilayahRepository {
	return &WilayahConnetion{
		tx: DB,
	}
}
