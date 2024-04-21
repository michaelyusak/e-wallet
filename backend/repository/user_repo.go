package repository

import (
	"context"
	"database/sql"
	"e-wallet/entity"
	"fmt"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserById(ctx context.Context, userId int) (*entity.User, error)
	PostOneUser(ctx context.Context, userReq entity.User) (*entity.User, error)
	ResetPassword(ctx context.Context, userId int, hashPwd []byte) error
	UpdateProfilePicture(ctx context.Context, userId int, imagePath string) error
	UpdateUserEmail(ctx context.Context, userReq entity.User) error
	UpdateUserName(ctx context.Context, userReq entity.User) error
	WithTx(ctx context.Context, fn func(repo UserRepository) (*entity.User, error)) (*entity.User, error)
}

type userRepositoryPostgres struct {
	db   *sql.DB
	dbtx DBTX
}

func NewUserRepositoryPostgres(db *sql.DB) userRepositoryPostgres {
	return userRepositoryPostgres{
		db:   db,
		dbtx: db,
	}
}

func (r *userRepositoryPostgres) injectDBTX(dbtx DBTX) {
	r.dbtx = dbtx
}

func (r *userRepositoryPostgres) UpdateProfilePicture(ctx context.Context, userId int, imagePath string) error {
	query := `
		UPDATE users
			SET user_profile_picture_name = $1
			WHERE user_id = $2
	`

	_, err := r.dbtx.ExecContext(ctx, query, imagePath, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) PostOneUser(ctx context.Context, userReq entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (user_email, user_name, user_password, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING user_id
	`

	err := r.dbtx.QueryRowContext(ctx, query, userReq.Email, userReq.Name, userReq.HashPassword).Scan(&userReq.Id)
	if err != nil {
		return nil, err
	}

	return &userReq, nil
}

func (r *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT user_id, user_name, user_password
		FROM users
		WHERE user_email = $1
		AND deleted_at
			IS NULL
	`

	var user entity.User

	err := r.dbtx.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Name, &user.HashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	user.Email = email

	return &user, nil
}

func (r *userRepositoryPostgres) GetUserById(ctx context.Context, userId int) (*entity.User, error) {
	query := `
		SELECT user_email, user_name, user_profile_picture_name, user_password
		FROM users
		WHERE user_id = $1
	`

	var user entity.User
	var imageName *string

	err := r.dbtx.QueryRowContext(ctx, query, userId).Scan(&user.Email, &user.Name, &imageName, &user.HashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if imageName != nil {
		user.ProfilePictureName = *imageName
	}

	user.Id = userId

	return &user, nil
}

func (r *userRepositoryPostgres) UpdateUserEmail(ctx context.Context, userReq entity.User) error {
	query := `
		UPDATE users
		SET user_email = $1
		WHERE user_id = $2
	`

	_, err := r.dbtx.ExecContext(ctx, query, userReq.Email, userReq.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) UpdateUserName(ctx context.Context, userReq entity.User) error {
	query := `
		UPDATE users
		SET user_name = $1
		WHERE user_id = $2
	`

	_, err := r.dbtx.ExecContext(ctx, query, userReq.Name, userReq.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) ResetPassword(ctx context.Context, userId int, hashPwd []byte) error {
	query := `
		UPDATE users
		SET user_password = $1
		WHERE user_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, hashPwd, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryPostgres) WithTx(ctx context.Context, fn func(repo UserRepository) (*entity.User, error)) (*entity.User, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	repo := NewUserRepositoryPostgres(r.db)
	repo.injectDBTX(tx)

	user, err := fn(&repo)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf(err.Error(), rbErr.Error())
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return user, nil
}
