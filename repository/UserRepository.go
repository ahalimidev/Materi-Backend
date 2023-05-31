package repository

import (
	"context"
	"database/sql"
	"fmt"
	"materi/model/response"
	"strconv"
)

type UserRepository interface {
	All(ctx context.Context) ([]response.UserRespon, error)
	FindByID(ctx context.Context, IdUser string) (response.UserRespon, error)
	Create(ctx context.Context, input response.UserRespon) (response.UserRespon, error)
	Update(ctx context.Context, input response.UserRespon) error
	Delete(ctx context.Context, IdUser string) error
	Upload(ctx context.Context, input response.UserRespon) error
}

type UserConnetion struct {
	tx *sql.DB
}

// All implements UserRepository
func (r *UserConnetion) All(ctx context.Context) ([]response.UserRespon, error) {
	SQL := "SELECT * FROM user"
	rows, err := r.tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user_ := []response.UserRespon{}
	for rows.Next() {
		user := response.UserRespon{}
		err := rows.Scan(
			&user.IdUser,
			&user.NamaLengkap,
			&user.JenisKelamin,
			&user.Alamat,
			&user.Foto,
			&user.Username,
			&user.Password,
			&user.CreateBy,
			&user.CreateDate,
			&user.UpdateBy,
			&user.UpdateDate,
			&user.Status)
		if err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		user_ = append(user_, user)
	}
	return user_, nil
}

// Create implements UserRepository
func (r *UserConnetion) Create(ctx context.Context, input response.UserRespon) (response.UserRespon, error) {
	SQL := "INSERT INTO user (nama_lengkap,jenis_kelamin,alamat,username,password) values (?, ?, ?, ?, ?)"
	result, err := r.tx.ExecContext(ctx, SQL,
		input.NamaLengkap.String,
		input.JenisKelamin.String,
		input.Alamat.String,
		input.Username.String,
		input.Password.String)
	if err != nil {
		return response.UserRespon{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return response.UserRespon{}, err
	}
	input.IdUser.String = strconv.Itoa(int(id))
	return input, err

}

// Delete implements UserRepository
func (r *UserConnetion) Delete(ctx context.Context, IdUser string) error {
	SQL := "DELETE FROM user WHERE id_user = ?"
	_, err := r.tx.ExecContext(ctx, SQL, IdUser)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements UserRepository
func (r *UserConnetion) FindByID(ctx context.Context, IdUser string) (response.UserRespon, error) {
	SQL := "SELECT * FROM user WHERE id_user = ?"
	rows, err := r.tx.QueryContext(ctx, SQL, IdUser)
	if err != nil {
		return response.UserRespon{}, err
	}
	defer rows.Close()

	user := response.UserRespon{}
	if rows.Next() {
		err := rows.Scan(
			&user.IdUser,
			&user.NamaLengkap,
			&user.JenisKelamin,
			&user.Alamat,
			&user.Foto,
			&user.Username,
			&user.Password,
			&user.CreateBy,
			&user.CreateDate,
			&user.UpdateBy,
			&user.UpdateDate,
			&user.Status)
		if err != nil {
			return user, err
		}
		return user, nil
	} else {
		return user, err
	}
}

// Update implements UserRepository
func (r *UserConnetion) Update(ctx context.Context, input response.UserRespon) error {
	fmt.Print(input)
	SQL := "UPDATE user SET nama_lengkap = ?, jenis_kelamin = ?, alamat = ?, status = ? WHERE id_user = ?"
	_, err := r.tx.QueryContext(ctx, SQL,
		input.NamaLengkap.String,
		input.JenisKelamin.String,
		input.Alamat.String,
		input.Status.String,
		input.IdUser.String)
	if err != nil {
		return err
	}
	return nil
}

// Upload implements UserRepository
func (r *UserConnetion) Upload(ctx context.Context, input response.UserRespon) error {
	SQL := "UPDATE user SET foto = ? WHERE id_user = ?"
	_, err := r.tx.QueryContext(ctx, SQL, input.Foto.String, input.IdUser.String)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserConnetion{
		tx: DB,
	}
}
