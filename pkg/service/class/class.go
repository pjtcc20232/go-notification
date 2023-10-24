package class

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type ClassServiceInterface interface {
	GetAll(ctx context.Context) (*model.ClasstList, error)
	GetByID(ctx context.Context, ID string) (*model.Class, error)
	GetByName(ctx context.Context, name string) (*model.Class, error)
	Create(ctx context.Context, cls *model.Class) (*model.Class, error)
	Update(ctx context.Context, ID int, clsToChange *model.Class) (bool, error)
	Delete(ctx context.Context, ID int) (bool, error)
}

type classservice struct {
	dbp pgsql.DatabaseInterface
}

func NewClassService(database_pool pgsql.DatabaseInterface) *classservice {
	return &classservice{
		dbp: database_pool,
	}
}

func (cl *classservice) GetAll(ctx context.Context) (*model.ClasstList, error) {
	rows, err := cl.dbp.GetDB().QueryContext(ctx, "SELECT id, horario, curso_id FROM turma ")
	if err != nil {
		logger.Error("Erro to list all:"+err.Error(), err)
		return &model.ClasstList{}, err

	}

	defer rows.Close()

	cl_list := &model.ClasstList{}

	for rows.Next() {
		class := model.Class{}
		if err := rows.Scan(&class.ID, &class.Schedules, &class.Tbl_course_id); err != nil {
			logger.Error("Erro to list class:"+err.Error(), err)
		} else {
			cl_list.List = append(cl_list.List, &class)
		}
	}

	return cl_list, nil
}

func (us *classservice) GetByID(ctx context.Context, ID int64) (*model.Class, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, horario, curso_id FROM turma WHERE id = $1")
	cl := model.Class{}
	if err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &cl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&cl.ID, &cl.Schedules, &cl.Tbl_course_id); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &cl, err
	}

	return &cl, nil
}

func (us *classservice) GetByName(ctx context.Context, name string) (*model.Class, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, horario, curso_id from turma WHERE name = $1")
	cl := model.Class{}
	if err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &cl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, name).Scan(&cl.ID, &cl.Schedules, &cl.Tbl_course_id); err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &cl, err
	}

	return &cl, nil
}

func (us *classservice) Create(ctx context.Context, cls *model.Class) (*model.Class, error) {

	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	cl := model.Class{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		return &cl, err
	}

	query := "INSERT INTO turmas (horario, course_id) VALUES ($1, $2);"

	_, err = tx.ExecContext(ctx, query, cls.Schedules, cls.Tbl_course_id)
	if err != nil {
		logger.Error("erro to Exec SQL Query:"+err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit TX:"+err.Error(), err)
		return &cl, err
	} else {
		logger.Info("Insert Transaction committed")

	}

	return &cl, nil
}

func (us *classservice) Update(ctx context.Context, ID int, clsToChange *model.Class) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()

	query := "UPDATE turma SET horario = $1, curso_id = $2 WHERE id = $3"

	// Execute a consulta de atualização dentro da transação
	_, err = tx.ExecContext(ctx, query, clsToChange.Schedules, clsToChange.Tbl_course_id, ID)
	if err != nil {
		logger.Error("Erro ao executar a consulta de atualização:", err)
		return false, err
	}

	// Commit da transação se tudo estiver correto
	err = tx.Commit()
	if err != nil {
		logger.Error("Erro ao confirmar a transação:", err)
		return false, err
	}

	logger.Info("Atualização da classe bem-sucedida")
	return true, nil
}

func (us *classservice) Delete(ctx context.Context, ID int) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()
	query := "DELETE FROM turma WHERE id = $1"

	_, err = tx.ExecContext(ctx, query, ID)
	if err != nil {
		logger.Error("Erro ao executar a consulta de exclusão:", err)
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Erro ao confirmar a transação:", err)
		return false, err
	}

	logger.Info("Exclusão da classe bem-sucedida")
	return true, nil
}
