package request

import (
	"github.com/go-playground/validator/v10"
)

type UserCreate struct {
	NamaLengkap  string `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required,eq=L|eq=P"`
	Alamat       string `json:"alamat" form:"alamat" validate:"required"`
	Username     string `json:"username" form:"username" validate:"required,min=6,max=32"`
	Password     string `json:"password" form:"password" validate:"required,min=6,max=32"`
	// CreateDate   string `json:"create_date" form:"create_date" binding:"required"`
	// CreateBy     string `json:"create_by" form:"create_by" binding:"required"`
}

type UserUpdate struct {
	IdUser       string `json:"id_user"`
	NamaLengkap  string `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" validate:"required"`
	Alamat       string `json:"alamat" form:"alamat" validate:"required"`
	// UdateDate    string `json:"update_date" form:"update_date" validate:"required"`
	// UdateBy      string `json:"update_by" form:"update_by" validate:"required"`
	Status string `json:"status" form:"status" validate:"required"`
}

type UserProfile struct {
	IdUser string `json:"id_user"`
	Foto   string
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(validasi interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(validasi)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
