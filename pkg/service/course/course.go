package course

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type CourseServiceInterface interface {
	GetAll(ctx context.Context) (*model.CourseList, error)
	GetByID(ctx context.Context, ID int) (*model.CourseList, error)
	GetByName(ctx context.Context, name string) (*model.Courses, error)
	Create(ctx context.Context, cls *model.Class) (*model.Courses, error)
	Update(ctx context.Context, ID int, corsToChange *model.Courses) (bool, error)
	Delete(ctx context.Context, ID int) (bool, error)
}

type courseService struct {
	dbp pgsql.DatabaseInterface
}

func NewCourseService(database_pool pgsql.DatabaseInterface) *courseService {
	return &courseService{
		dbp: database_pool,
	}
}

func (cl *courseService) GetAll(ctx context.Context) (*model.CourseList, error) {
	rows, err := cl.dbp.GetDB().QueryContext(ctx, "SELECT id, nome_curso FROM cursos ")
	if err != nil {
		logger.Error("Erro to list all:"+err.Error(), err)
		return &model.CourseList{}, err

	}

	defer rows.Close()

	cl_list := &model.CourseList{}

	for rows.Next() {
		course := model.Courses{}
		if err := rows.Scan(&course.ID, &course.Name); err != nil {
			logger.Error("Erro to list class:"+err.Error(), err)
		} else {
			cl_list.List = append(cl_list.List, &course)
		}
	}

	return cl_list, nil
}

func (us *courseService) GetByID(ctx context.Context, ID int64) (*model.Courses, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome_curso FROM cursos WHERE id = $1")
	cl := model.Courses{}
	if err != nil {
		logger.Error("Erro to list cursos:"+err.Error(), err)
		return &cl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&cl.ID, &cl.Name); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &cl, err
	}

	return &cl, nil
}

func (us *courseService) GetByName(ctx context.Context, name string) (*model.Courses, error) {
	stmt, err := us.dbp.GetDB().PrepareContext(ctx, "SELECT id, nome_curso from turma WHERE nome_curso = $1")
	cl := model.Courses{}
	if err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &cl, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, name).Scan(&cl.ID, &cl.Name); err != nil {
		logger.Error("Erro to list course:"+err.Error(), err)
		return &cl, err
	}

	return &cl, nil
}

func (us *courseService) Create(ctx context.Context, cls *model.Courses) (*model.Courses, error) {

	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	cl := model.Courses{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		return &cl, err
	}

	query := "INSERT INTO cursos (nome_curso) VALUES ($1);"

	_, err = tx.ExecContext(ctx, query, cls.Name)
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

func (us *courseService) Update(ctx context.Context, ID int, clsToChange *model.Courses) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()

	query := "UPDATE cursos SET nome_curso = $1  WHERE id = $2"

	// Execute a consulta de atualização dentro da transação
	_, err = tx.ExecContext(ctx, query, clsToChange.Name, ID)
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

func (us *courseService) Delete(ctx context.Context, ID int) (bool, error) {
	// Comece uma transação
	tx, err := us.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()
	query := "DELETE FROM cursos WHERE id = $1"

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
