package postgresql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/db"
)

type Todo struct {
	conn *pgxpool.Pool
	q    *db.Queries
}

func NewTodo(conn *pgxpool.Pool) *Todo {
	return &Todo{
		conn: conn,
		q:    db.New(conn),
	}
}

func (t *Todo) Create(ctx context.Context, params api.CreateParams) (api.Todo, error) {
	result, err := t.q.CreateTodo(ctx, db.CreateTodoParams{
		Task:   params.Task,
		IsDone: params.IsDone,
	})
	if err != nil {
		return api.Todo{}, err
	}

	return api.Todo{
		ID:   result.ID,
		Task: result.Task,
	}, nil
}

func (t *Todo) Delete(ctx context.Context, id int64) (int64, error) {
	deletedID, err := t.q.DeleteTodo(ctx, id)
	if err != nil {
		return deletedID, err
	}

	return deletedID, nil
}

func (t *Todo) Find(ctx context.Context, id int64) (api.Todo, error) {
	result, err := t.q.FindTodo(ctx, id)
	if err != nil {
		return api.Todo{}, err
	}

	return api.Todo{
		ID:     result.ID,
		Task:   result.Task,
		IsDone: result.IsDone,
	}, nil
}

func (t *Todo) List(ctx context.Context) ([]api.Todo, error) {
	var todos []api.Todo

	sql, _, err := sq.Select("id", "task", "is_done").From("todos").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return todos, err
	}

	if err := pgxscan.Select(ctx, t.conn, &todos, sql); err != nil {
		return todos, err
	}

	return todos, nil
}

func (t *Todo) Update(ctx context.Context, id int64, params api.UpdateParams) (int64, error) {
	updatedID, err := t.q.UpdateTodo(ctx, db.UpdateTodoParams{
		ID:     id,
		Task:   params.Task,
		IsDone: params.IsDone,
	})
	if err != nil {
		return updatedID, err
	}

	return updatedID, nil
}
