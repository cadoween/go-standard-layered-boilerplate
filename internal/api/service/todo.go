package service

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"go.uber.org/zap"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/pkg/errorsutil"
)

type TodoRepository interface {
	Create(ctx context.Context, params api.CreateParams) (api.Todo, error)
	List(ctx context.Context) ([]api.Todo, error)
	Find(ctx context.Context, id int64) (api.Todo, error)
	Update(ctx context.Context, id int64, params api.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type Todo struct {
	todoRepo TodoRepository
	log      *zap.Logger
}

func NewTodo(logger *zap.Logger, todoRepo TodoRepository) *Todo {
	return &Todo{
		todoRepo: todoRepo,
		log:      logger,
	}
}

func (t *Todo) List(ctx context.Context) ([]api.Todo, error) {
	todos, err := t.todoRepo.List(ctx)
	if err != nil {
		return todos, errorsutil.Wrapf(err, "Failed to get todos list", api.ErrCodeInternalDatabase)
	}

	return todos, nil
}

func (t *Todo) Create(ctx context.Context, params api.CreateParams) (api.Todo, error) {
	newTodo, err := t.todoRepo.Create(ctx, params)
	if err != nil {
		return api.Todo{}, errorsutil.Wrapf(err, "Failed to create new todo", api.ErrCodeInternalDatabase)
	}

	return newTodo, nil
}

func (t *Todo) Find(ctx context.Context, id int64) (api.Todo, error) {
	singleTodo, err := t.todoRepo.Find(ctx, id)
	if err != nil {
		if pgxscan.NotFound(err) {
			return singleTodo, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return singleTodo, errorsutil.Wrapf(err, "Failed to find a todo", api.ErrCodeInternalDatabase)
	}

	return singleTodo, nil
}

func (t *Todo) Update(ctx context.Context, id int64, params api.UpdateParams) (int64, error) {
	updatedID, err := t.todoRepo.Update(ctx, id, params)
	if err != nil {
		if pgxscan.NotFound(err) {
			return updatedID, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return updatedID, errorsutil.Wrapf(err, "Failed to update a todo", api.ErrCodeInternalDatabase)
	}

	return updatedID, nil
}

func (t *Todo) Delete(ctx context.Context, id int64) (int64, error) {
	deletedID, err := t.todoRepo.Delete(ctx, id)
	if err != nil {
		if pgxscan.NotFound(err) {
			return deletedID, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return deletedID, errorsutil.Wrapf(err, "Failed to delete a todo", api.ErrCodeInternalDatabase)
	}

	return deletedID, nil
}
