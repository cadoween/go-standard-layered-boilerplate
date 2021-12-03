# go-standard-layered-boilerplate

Go Standard Layered Boilerplate

## Technologies

The various packages used are as follows:

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP handler framework
- [PGX](https://github.com/jackc/pgx) - PostgreSQL database driver
- [Squirrel](https://github.com/Masterminds/squirrel) - SQL query builder
- [Scany](https://github.com/georgysavva/scany) - Database scanner helper from data to structs
- [Kin OpenAPI](https://github.com/getkin/kin-openapi) - OpenAPI 3.0 / Swagger specifications generator from code
- [CORS](https://github.com/gin-contrib/cors) - Gin Gonic CORS middleware
- [Go Dot ENV](https://github.com/joho/godotenv) - Dot ENV vars loader
- [Zap](https://github.com/uber-go/zap) - Logger package

## Installation

This boilerplate requires [Golang](https://golang.org/) v1.17+ to run.

Clone the repository then change your current directory to the repository

```sh
cd go-standard-layered-boilerplate
```

Copy the example env file and make the required configuration changes in the .env file

```sh
cp .env.example .env
```

Start the REST server.

```sh
make rest
```

## Server

Open the development server by navigating to your server address in your preferred browser.

http://localhost:1234

## Tools as Dependencies

Go Standard Layered Boilerplate is implementing tools as dependencies with go modules.
Instructions on how to use them in the application are linked below.

| Tools   | Repository                                |
| ------- | ----------------------------------------- |
| Migrate | https://github.com/golang-migrate/migrate |
| SQLC    | https://github.com/kyleconroy/sqlc        |

## Database Migrations & Seeding

Run database migrations up.

```sh
make migrate-up
```

Run database migrations down.

```sh
make migrate-down
```
