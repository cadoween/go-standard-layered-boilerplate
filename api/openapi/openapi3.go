package openapi

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

//go:generate go run ../../cmd/openapi-gen/main.go -path .

// Swagger tags list.
const (
	tagTodo = "todo"
)

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Go Standard Modular API",
			Description: "REST APIs example for standard boilerplate",
			Version:     "1.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				Name: "Krishna",
				URL:  "https://github.com/KrisCatDog",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://localhost:1234",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Todo": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewInt64Schema()).
				WithProperty("task", openapi3.NewStringSchema()).
				WithProperty("is_done", openapi3.NewBoolSchema())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateTodoRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a todo.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("task", openapi3.NewStringSchema()).
					WithProperty("is_done", openapi3.NewBoolSchema())),
		},
		"UpdateTodoRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a todo.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("task", openapi3.NewStringSchema()).
					WithProperty("is_done", openapi3.NewBoolSchema())),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateTodoResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating tasks.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Todo",
					}))),
		},
		"ReadTodoResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one task.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Todo",
					}))),
		},
	}

	swagger.Tags = openapi3.Tags{
		&openapi3.Tag{
			Name: tagTodo,
		},
	}

	swagger.Paths = openapi3.Paths{
		"/todos": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateTodo",
				Tags:        []string{tagTodo},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateTodoRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateTodoResponse",
					},
				},
			},
			Get: &openapi3.Operation{
				OperationID: "TodosList",
				Tags:        []string{tagTodo},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadTodoResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
		"/todos/{id}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "ReadTodo",
				Tags:        []string{tagTodo},
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("id").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadTodoResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateTodo",
				Tags:        []string{tagTodo},
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("id").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateTodoRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task updated"),
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Delete: &openapi3.Operation{
				OperationID: "DeleteTodo",
				Tags:        []string{tagTodo},
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("id").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task updated"),
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Task not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
	}

	return swagger
}

// RegisterSpecifications update and serve OpenAPI specification files.
func RegisterSpecifications(r *gin.Engine) {
	swagger := NewOpenAPI3()

	r.GET("/openapi3.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, swagger)
	})

	r.GET("/openapi3.yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, swagger)
	})
}
