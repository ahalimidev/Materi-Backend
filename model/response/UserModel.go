package response

import "materi/helper"

type UserRespon struct {
	IdUser       helper.NullString `json:"id_user"`
	NamaLengkap  helper.NullString `json:"nama_lengkap"`
	JenisKelamin helper.NullString `json:"jenis_kelamin"`
	Alamat       helper.NullString `json:"alamat"`
	Foto         helper.NullString `json:"foto"`
	Username     helper.NullString `json:"-"`
	Password     helper.NullString `json:"-"`
	CreateDate   helper.NullString `json:"create_date"`
	CreateBy     helper.NullString `json:"create_by"`
	UpdateDate   helper.NullString `json:"update_date"`
	UpdateBy     helper.NullString `json:"upate_by"`
	Status       helper.NullString `json:"status"`
}
