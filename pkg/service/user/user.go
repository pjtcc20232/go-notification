package user

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type UserServiceInterface interface {
	//GetAll(ctx context.Context, filters model.FilterUser, limit, page int64) (*model.Paginate, error)
	GetAll(ctx context.Context) (*model.UsertList, error)
	GetByID(ctx context.Context, ID string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (user *model.User, err error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, ID string, userToChange *model.User, previousUserData *model.User, userRequest *model.User) (bool, error)
	Delete(ctx context.Context, ID string, currentUserData *model.User, userRequest *model.User) (bool, error)
	ChangePassword(ctx context.Context, currentPassword, newPassword string, userRequest *model.User) error
}

type userservice struct {
	dbp pgsql.DatabaseInterface
}

func NewUserService(database_pool pgsql.DatabaseInterface) *userservice {
	return &userservice{
		dbp: database_pool,
	}
}

func (us *userservice) GetAll(ctx context.Context) (*model.UsertList, error) {
	rows, err := us.dbp.GetDB().QueryContext(ctx, "SELECT id, user_name, email, passwrod, hashpass, enable, locked, created_at, updated_at  FROM usuarios_user ")
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)
		return &model.UsertList{}, err

	}

	defer rows.Close()

	usr_list := &model.UsertList{}

	for rows.Next() {
		usr := model.User{}
		if err := rows.Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.HashedPassword, &usr.Enable, &usr.IsLocked, &usr.CreatedAt, &usr.UpdatedAt); err != nil {
			logger.Error("Erro to list users:"+err.Error(), err)
		} else {
			usr_list.List = append(usr_list.List, &usr)
		}
	}

	return usr_list, nil
}

func (us *userservice) GetByID(ctx context.Context, ID int64) (*model.User, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, user_name, email, passwrod, hashpass, enable, locked, created_at, updated_at  FROM usuarios_user WHERE id = $1")
	usr := model.User{}
	if err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &usr, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.HashedPassword, &usr.Enable, &usr.IsLocked, &usr.CreatedAt, &usr.UpdatedAt); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &usr, err
	}

	return &usr, nil
}

func (us *userservice) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, user_name, email, passwrod, hashpass, enable, locked, created_at, updated_at  FROM usuarios_user WHERE email = $1")
	usr := model.User{}
	if err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &usr, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, email).Scan(&usr.ID, &usr.Name, &usr.Email, &usr.Password, &usr.HashedPassword, &usr.Enable, &usr.IsLocked, &usr.CreatedAt, &usr.UpdatedAt); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &usr, err
	}

	return &usr, nil
}

func (us *userservice) Create(ctx context.Context, user *model.User) (*model.User, error) {

	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	usr := model.User{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		return &usr, err
	}

	query := "INSERT INTO users (name, email, password, enable, is_locked) VALUES ($1, $2, $3, $4, $5);"

	_, err = tx.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Enable, user.IsLocked)
	if err != nil {
		logger.Error("erro to Exec SQL Query:"+err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit TX:"+err.Error(), err)
		return &usr, err
	} else {
		logger.Info("Insert Transaction committed")

	}

	return &usr, nil
}
