package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/pkg/errorsutil"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/pkg/resputil"
)

type TodoService interface {
	Create(ctx context.Context, params api.CreateParams) (api.Todo, error)
	List(ctx context.Context) ([]api.Todo, error)
	Find(ctx context.Context, id int64) (api.Todo, error)
	Update(ctx context.Context, id int64, params api.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type TodoHandler struct {
	todoSvc TodoService
}

func NewTodoHandler(todoSvc TodoService) *TodoHandler {
	return &TodoHandler{
		todoSvc: todoSvc,
	}
}

func (h *TodoHandler) Register(r *gin.Engine) {
	r.GET("/todos", h.list)
	r.POST("/todos", h.create)
	r.GET("/todos/:id", h.find)
	r.PUT("/todos/:id", h.update)
	r.DELETE("/todos/:id", h.delete)
}

type Todo struct {
	ID     int64  `json:"id,omitempty"`
	Task   string `json:"task,omitempty"`
	IsDone bool   `json:"is_done,omitempty"`
}

type CreateTodoRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) create(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Invalid type of request body", api.ErrCodeBadRequest))

		return
	}

	newTodo, err := h.todoSvc.Create(c, api.CreateParams{
		Task:   req.Task,
		IsDone: false,
	})
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendJSON(c, http.StatusCreated, "Todo successfully created", &CreateTodoResponse{
		Todo: Todo{
			ID:     newTodo.ID,
			Task:   newTodo.Task,
			IsDone: newTodo.IsDone,
		},
	})
}

type DeleteTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type DeleteTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) delete(c *gin.Context) {
	var reqURI DeleteTodoRequest
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	deletedID, err := h.todoSvc.Delete(c, reqURI.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendJSON(c, http.StatusOK, "Todo successfully deleted", &DeleteTodoResponse{
		Todo: Todo{
			ID: deletedID,
		},
	})
}

type FindTodoRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type FindTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) find(c *gin.Context) {
	var req FindTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	singleTodo, err := h.todoSvc.Find(c, req.ID)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendJSON(c, http.StatusOK, "Successfully got todo details", &FindTodoResponse{
		Todo: Todo{
			ID:     singleTodo.ID,
			Task:   singleTodo.Task,
			IsDone: singleTodo.IsDone,
		},
	})
}

type ListTodosResponse struct {
	Todos []Todo `json:"todos"`
}

func (h *TodoHandler) list(c *gin.Context) {
	items, err := h.todoSvc.List(c)
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	todos := make([]Todo, len(items))

	for i, item := range items {
		todos[i].ID = item.ID
		todos[i].Task = item.Task
		todos[i].IsDone = item.IsDone
	}

	resputil.SendJSON(c, http.StatusOK, "Successfully got todos list", &ListTodosResponse{
		Todos: todos,
	})
}

type UpdateTodoRequestURI struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateTodoRequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type UpdateTodoResponse struct {
	Todo Todo `json:"todo"`
}

func (h *TodoHandler) update(c *gin.Context) {
	var reqURI UpdateTodoRequestURI
	if err := c.ShouldBindUri(&reqURI); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Failed to bind request URI parameters", api.ErrCodeBadRequest))

		return
	}

	var reqBody UpdateTodoRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		resputil.SendError(c, errorsutil.Wrapf(err, "Invalid type of request body", api.ErrCodeBadRequest))

		return
	}

	updatedID, err := h.todoSvc.Update(c, reqURI.ID, api.UpdateParams{
		Task:   reqBody.Task,
		IsDone: reqBody.IsDone,
	})
	if err != nil {
		resputil.SendError(c, err)

		return
	}

	resputil.SendJSON(c, http.StatusOK, "Todo successfully updated", &UpdateTodoResponse{
		Todo: Todo{
			ID: updatedID,
		},
	})
}
