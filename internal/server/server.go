package server

import (
	"embed"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"

	"github.com/KrisCatDog/go-standard-layered-boilerplate/api/openapi"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api/http/rest"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api/repository/postgresql"
	"github.com/KrisCatDog/go-standard-layered-boilerplate/internal/api/service"
)

type Config struct {
	Address string
	DB      *pgxpool.Pool
	Logger  *zap.Logger
	Static  embed.FS
}

func New(cfg Config) (*http.Server, error) {
	// Construct new gin with default options.
	r := gin.Default()

	// Set gin mode that depends on ENV.
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Implement cors middleware to gin.
	r.Use(cors.Default())

	// Register repositories used for dependency injection to the services.
	todoRepo := postgresql.NewTodo(cfg.DB)

	// Register services used for dependency injection to the handlers.
	todoSvc := service.NewTodo(cfg.Logger, todoRepo)

	// Serve swagger specification files.
	openapi.RegisterSpecifications(r)

	// Serve embedded static files that live on a rest-server/static folder.
	r.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(c.Request.URL.Path, http.FS(cfg.Static))
	})

	// Register all REST HTTP handlers.
	rest.NewTodoHandler(todoSvc).Register(r)

	return &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}, nil
}
