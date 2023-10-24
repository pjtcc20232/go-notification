package student

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type StudentServiceInterface interface {
	GetAll(ctx context.Context) (*model.StudenttList, error)
	GetByID(ctx context.Context, ID int) (*model.Student, error)
	GetByName(ctx context.Context, name string) (*model.Student, error)
	Create(ctx context.Context, cls *model.Student) (*model.Student, error)
	Update(ctx context.Context, ID int, stsToChange *model.Class) (bool, error)
	Delete(ctx context.Context, ID int) (bool, error)
}

type studentservice struct {
	dbp pgsql.DatabaseInterface
}

func NewStudentService(database_pool pgsql.DatabaseInterface) *studentservice {
	return &studentservice{
		dbp: database_pool,
	}
}

func (cl *studentservice) GetAll(ctx context.Context) (*model.StudenttList, error) {
	rows, err := cl.dbp.GetDB().QueryContext(ctx, "SELECT id, nome, periodo, matricula, turma_id, usuario_id FROM alunos ")
	if err != nil {
		logger.Error("Erro to list all:"+err.Error(), err)
		return &model.StudenttList{}, err

	}

	defer rows.Close()

	sl_list := &model.StudenttList{}

	for rows.Next() {
		student := model.Student{}
		if err := rows.Scan(&student.ID, &student.Name, &student.Registration, &student.Tbl_class_id, &student.Tbl_usr_id); err != nil {
			logger.Error("Erro to list class:"+err.Error(), err)
		} else {
			sl_list.List = append(sl_list.List, &student)
		}
	}

	return sl_list, nil
}

func (us *studentservice) GetByID(ctx context.Context, ID int64) (*model.Student, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome, periodo, matricula, turma_id, usuario_id from alunos WHERE id = $1")
	sl := model.Student{}
	if err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &sl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&sl.ID, &sl.Name, &sl.Registration, &sl.Tbl_class_id, &sl.Tbl_usr_id); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &sl, err
	}

	return &sl, nil
}

func (us *studentservice) GetByName(ctx context.Context, name string) (*model.Student, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome, periodo, matricula, turma_id, usuario_id from aluno WHERE nome = $1")
	sl := model.Student{}
	if err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &sl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, name).Scan(&sl.ID, &sl.Name, &sl.Registration, &sl.Tbl_class_id, &sl.Tbl_usr_id); err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &sl, err
	}

	return &sl, nil
}

func (us *studentservice) Create(ctx context.Context, std *model.Student) (*model.Student, error) {

	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	sl := model.Student{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		return &sl, err
	}

	query := "INSERT INTO alunos (nome, periodo, matricula, turma_id, usuario_id) VALUES ($1, $2, $3, $4, $5);"

	_, err = tx.ExecContext(ctx, query, std.Name, std.Period, std.Registration, std.Tbl_class_id, std.Tbl_usr_id)
	if err != nil {
		logger.Error("erro to Exec SQL Query:"+err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit TX:"+err.Error(), err)
		return &sl, err
	} else {
		logger.Info("Insert Transaction committed")

	}

	return &sl, nil
}

func (us *studentservice) Update(ctx context.Context, ID int, lsToChange *model.Student) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()

	query := "UPDATE alunos SET nome=$1, periodo=$2, matricula=$3, turma_id=$4, usuario_id=$5 WHERE id = $6"

	// Execute a consulta de atualização dentro da transação
	_, err = tx.ExecContext(ctx, query, lsToChange.Name, lsToChange.Period, lsToChange.Registration, lsToChange.Tbl_class_id, lsToChange.Tbl_usr_id, ID)
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

func (us *studentservice) Delete(ctx context.Context, ID int) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()
	query := "DELETE FROM alunos WHERE id = $1"

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
