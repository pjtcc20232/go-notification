package teacher

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type TeacherServiceInterface interface {
	GetAll(ctx context.Context) (*model.TeacherList, error)
	GetByID(ctx context.Context, ID int) (*model.Teacher, error)
	GetByName(ctx context.Context, name string) (*model.Teacher, error)
	Create(ctx context.Context, cls *model.Student) (*model.Teacher, error)
	Update(ctx context.Context, ID int, tsToChange *model.Teacher) (bool, error)
	Delete(ctx context.Context, ID int) (bool, error)
}

type teacherservice struct {
	dbp pgsql.DatabaseInterface
}

func NewStudentService(database_pool pgsql.DatabaseInterface) *teacherservice {
	return &teacherservice{
		dbp: database_pool,
	}
}

func (tc *teacherservice) GetAll(ctx context.Context) (*model.TeacherList, error) {
	rows, err := tc.dbp.GetDB().QueryContext(ctx, "SELECT id, nome, cadeira, usuario_id FROM alunos ")
	if err != nil {
		logger.Error("Erro to list all:"+err.Error(), err)
		return &model.TeacherList{}, err

	}

	defer rows.Close()

	tl_list := &model.TeacherList{}

	for rows.Next() {
		teacher := model.Teacher{}
		if err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Chair, &teacher.Tbl_usr_id); err != nil {
			logger.Error("Erro to list class:"+err.Error(), err)
		} else {
			tl_list.List = append(tl_list.List, &teacher)
		}
	}

	return tl_list, nil
}

func (us *teacherservice) GetByID(ctx context.Context, ID int64) (*model.Teacher, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome, cadeira, usuario_id from professores WHERE id = $1")
	tl := model.Teacher{}
	if err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &tl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&tl.ID, &tl.Name, &tl.Chair, &tl.Tbl_usr_id); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &tl, err
	}

	return &tl, nil
}

func (us *teacherservice) GetByName(ctx context.Context, name string) (*model.Teacher, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome, cadeira, usuario_id from professores WHERE nome = $1")
	tl := model.Teacher{}
	if err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &tl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, name).Scan(&tl.ID, &tl.Name, &tl.Chair, &tl.Tbl_usr_id); err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &tl, err
	}

	return &tl, nil
}

func (us *teacherservice) Create(ctx context.Context, std *model.Teacher) (*model.Teacher, error) {

	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	tl := model.Teacher{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		return &tl, err
	}

	query := "INSERT INTO professores (nome, cadeira, usuario_id) VALUES ($1, $2, $3);"

	_, err = tx.ExecContext(ctx, query, std.Name, std.Chair, std.Tbl_usr_id)
	if err != nil {
		logger.Error("erro to Exec SQL Query:"+err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit TX:"+err.Error(), err)
		return &tl, err
	} else {
		logger.Info("Insert Transaction committed")

	}

	return &tl, nil
}

func (us *teacherservice) Update(ctx context.Context, ID int, lsToChange *model.Teacher) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()

	query := "UPDATE professores SET nome=$1, cadeira=$2, usuario_id=$3 WHERE id = $4"

	// Execute a consulta de atualização dentro da transação
	_, err = tx.ExecContext(ctx, query, lsToChange.Name, lsToChange.Chair, lsToChange.Tbl_usr_id, ID)
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

func (us *teacherservice) Delete(ctx context.Context, ID int) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()
	query := "DELETE FROM professores WHERE id = $1"

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

	logger.Info("Exclusão do aluno bem-sucedida")
	return true, nil
}
