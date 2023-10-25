package event

import (
	"context"

	"github.com/notification/back-end/internal/config/logger"
	"github.com/notification/back-end/pkg/adapter/pgsql"
	"github.com/notification/back-end/pkg/model"
)

type EventServiceInterface interface {
	GetAll(ctx context.Context) (*model.EventList, error)
	GetByID(ctx context.Context, ID int) (*model.Event, error)
	GetByName(ctx context.Context, name string) (*model.Event, error)
	Create(ctx context.Context, cls *model.Event) (*model.Event, error)
	Update(ctx context.Context, ID int, corsToChange *model.Event) (bool, error)
	Delete(ctx context.Context, ID int) (bool, error)
}

type eventService struct {
	dbp pgsql.DatabaseInterface
}

func NewEventService(database_pool pgsql.DatabaseInterface) *eventService {
	return &eventService{
		dbp: database_pool,
	}
}

func (ev *eventService) GetAll(ctx context.Context) (*model.EventList, error) {
	rows, err := ev.dbp.GetDB().QueryContext(ctx, "SELECT id, data_evento, descricao, turma_id, professor_id, data_criacao, data_atualizacao, status_evento FROM eventos")
	if err != nil {
		logger.Error("Erro to list all:"+err.Error(), err)
		return &model.EventList{}, err

	}

	defer rows.Close()

	el_list := &model.EventList{}

	for rows.Next() {
		event := model.Event{}
		if err := rows.Scan(&event.ID, &event.EventDate, &event.Description, &event.Tbl_class_id, &event.Tbl_class_teacher, &event.CreatedAt, &event.UpdatedAt, &event.StatusEvent); err != nil {
			logger.Error("Erro to list class:"+err.Error(), err)
		} else {
			el_list.List = append(el_list.List, &event)
		}
	}

	return el_list, nil
}

func (ev *eventService) GetByID(ctx context.Context, ID int64) (*model.Event, error) {
	stmt, err := ev.dbp.GetDB().PrepareContext(ctx, "SELECT id, data_evento, descricao, turma_id, professor_id, data_criacao, data_atualizacao, status_evento FROM eventos WHERE id = $1")
	event := model.Event{}
	if err != nil {
		logger.Error("Erro to list cursos:"+err.Error(), err)
		return &event, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, ID).Scan(&event.ID, &event.EventDate, &event.Description, &event.Tbl_class_id, &event.Tbl_class_teacher, &event.CreatedAt, &event.UpdatedAt, &event.StatusEvent); err != nil {
		logger.Error("Erro to list users:"+err.Error(), err)
		return &event, err
	}

	return &event, nil
}

func (ev *eventService) GetByName(ctx context.Context, name string) (*model.Event, error) {
	stmt, err := ev.dbp.GetDB().PrepareContext(ctx, "SELECT id, data_evento, descricao, turma_id, professor_id, data_criacao, data_atualizacao, status_eventofrom eventos descricao LIKE '%' || $1 || '%';")
	event := model.Event{}
	if err != nil {
		logger.Error("Erro to list class:"+err.Error(), err)
		return &event, err
	}

	defer stmt.Close()

	if err := stmt.QueryRowContext(ctx, name).Scan(&event.ID, &event.EventDate, &event.Description, &event.Tbl_class_id, &event.Tbl_class_teacher, &event.CreatedAt, &event.UpdatedAt, &event.StatusEvent); err != nil {
		logger.Error("Erro to list course:"+err.Error(), err)
		return &event, err
	}

	return &event, nil
}

func (ev *eventService) Create(ctx context.Context, evt *model.Event) (*model.Event, error) {

	tx, err := ev.dbp.GetDB().BeginTx(ctx, nil)
	event := model.Event{}
	if err != nil {
		logger.Error("Erro to create:"+err.Error(), err)

		event := model.Event{}
		return &event, err
	}

	query := "INSERT INTO eventos (data_evento, descricao, turma_id, professor_id, data_criacao, data_atualizacao, status_evento) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	_, err = tx.ExecContext(ctx, query, evt.EventDate, evt.Description, evt.Tbl_class_id, evt.Tbl_class_teacher, evt.CreatedAt, evt.UpdatedAt, evt.StatusEvent)
	if err != nil {
		logger.Error("erro to Exec SQL Query:"+err.Error(), err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("erro to Commit TX:"+err.Error(), err)
		return &event, err
	} else {
		logger.Info("Insert Transaction committed")

	}

	return &event, nil
}

func (ev *eventService) Update(ctx context.Context, ID int, evtToChange *model.Event) (bool, error) {
	// Comece uma transação
	tx, err := ev.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()

	query := "UPDATE eventos SET data_evento = $1, descricao = $2, turma_id = $3, professor_id = $4, data_atualizacao = $5, status_evento = $6 WHERE id = $7;"

	// Execute a consulta de atualização dentro da transação
	_, err = tx.ExecContext(ctx, query, evtToChange.EventDate, evtToChange.Description, evtToChange.Tbl_class_id, evtToChange.Tbl_class_teacher, evtToChange.CreatedAt, evtToChange.UpdatedAt, evtToChange.StatusEvent, ID)
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

func (ev *eventService) Delete(ctx context.Context, ID int) (bool, error) {
	// Comece uma transação
	tx, err := ev.dbp.GetDB().BeginTx(ctx, nil)
	if err != nil {
		logger.Error("Erro ao iniciar a transação:", err)
		return false, err
	}
	defer tx.Rollback()
	query := "DELETE FROM eventos WHERE id = $1"

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
