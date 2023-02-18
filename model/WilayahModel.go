package model

import "materi/helper"

type WilayahProvinsi struct {
	IdProvinsi helper.NullString `json:"provinsi_id"`
	Provinsi   helper.NullString `json:"provinsi"`
}

type WilayahKabupaten struct {
	IdKabupaten helper.NullString `json:"kabupaten_id"`
	Kabupaten   helper.NullString `json:"kabupaten"`
}

type WilayahKecamatan struct {
	IdKecamatan helper.NullString `json:"kecamatan_id"`
	Kecamatan   helper.NullString `json:"kecamatan"`
}

type WilayahDesa struct {
	IdDesa helper.NullString `json:"desa_id"`
	Desa   helper.NullString `json:"desa"`
}
