<div align="center">
  <h1>Go GraphQL Simulation Project</h1>
  <p>This repo is for research and experiment purpose</p>
</div>

## Technologies

- Golang
- Fiber
- Gorm
- GraphQL with `gqlgen`
- PostgreSQL
- Redis
- Air

## Setup all stuff from scratch

### 1. Golang

To install Golang, you can go to their [Official Website](https://go.dev/dl/). Choose based on your Operating Systems. Available options:

- [Windows](https://go.dev/dl/go1.23.1.windows-amd64.msi).
- [MacOS ARM](https://go.dev/dl/go1.23.1.darwin-arm64.pkg).
- [MacOS x86](https://go.dev/dl/go1.23.1.darwin-amd64.pkg).
- [Linux](https://go.dev/dl/go1.23.1.linux-amd64.tar.gz).
- [Build from source](https://go.dev/dl/go1.23.1.src.tar.gz).

Getting started with Golang:

- First, create your project folder. In Linux, you can go to the terminal and type `mkdir <new_folder>`.
- Go to the project folder, and initialize new Golang project with command `go mod init <your_module_name>`.

### 2. Fiber

According to the Website:

> Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.

Based on that statement, we know that Fiber has gained traction due to its performance and developer-friendly features.

**So why we choose Fiber?**

- Express-like. You know Express JS? Nah, you will adapt quickly while using Fiber.
- It's fast.
- Rich middleware ecosystem. [CORS](https://docs.gofiber.io/api/middleware/cors/), [monitoring](https://docs.gofiber.io/api/middleware/monitor/), [helmet](https://docs.gofiber.io/api/middleware/helmet/), [Logger](https://docs.gofiber.io/api/middleware/logger/), [Caching](https://docs.gofiber.io/api/middleware/cache), and many more. See [here](https://docs.gofiber.io/category/-middleware) for more detail.
- Many built-in features like routing, template engines, and static file serving.

**How to install Fiber in a Go Project?**

```sh
go get github.com/gofiber/fiber/v2
```

**Very basic usage example in a `main.go` file:**

```go
package main

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
    Message string
}

func main() {
    app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
            Message: "Sukses mengakses API!",
        })
	})
}
```

### 3. Gorm

Gorm is one of the popular ORM (Object Relation Mapping) in Golang that simplifies database interactions.

**So, why we choose Gorm?**

- **Ease of use**.
- **Feature-Rich**, such as Auto Migrations, Association Handling, Eager Loading, and Transactions Support.
- **Support multiple databases**, one of them is PostgreSQL.
- **Well-supported Project**.

**How to install Gorm in a Golang project?**

```sh
go get -u gorm.io/gorm
```

### 4. GraphQL with `gqlgen`

We can implement GraphQL in Golang using `gqlgen`. `gqlgen` is a popular Go library for building GraphQL servers in Golang. It automatically generates much of the boilerplate code required for creating a fully functional GraphQL API, making it easier to work with GraphQL in Golang. Visit their [Official Website](https://gqlgen.com/) for more detail.

**How to setup and install `gqlgen` in your Golang Project?**

```sh
# Create tools.go file in the root project and fill it with needed packages
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

# Clean Up packages and optimize
go mod tidy

# Initialize gqlgen config and models
go run github.com/99designs/gqlgen init

# Clean Up packages and optimize
go mod tidy

# Start the GraphQL server
go run server.go # or just using air

# Generate new schemas based on your changes
go run github.com/99designs/gqlgen generate
```

The generated schemas and resolvers will looks like this:

```
.
├── generated.go
├── model
│   └── models_gen.go
├── resolver.go
├── schema.graphqls
└── schema.resolvers.go
```

If you want to modify and adjust GraphQL folder structure based on your needs, you can edit `gqlgen.yml` file that located in your root project.

### 5. PostgreSQL

To install PostgreSQL, you can download the packages and installers in their [Official Website](https://www.postgresql.org/download/). Choose based on your Operating Systems. Available options:

- [Linux](https://www.postgresql.org/download/linux/).
- [Mac OS](https://www.postgresql.org/download/macosx/).
- [Windows](https://www.postgresql.org/download/windows/).
- [BSD](https://www.postgresql.org/download/bsd/).
- [Solaris](https://www.postgresql.org/download/solaris/).

**How to install and setup PostgreSQL in your Golang Project?**

Because we are using Gorm as our ORM, we don't need to install any additional packages. We can use [built-in PostgreSQL driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL) that provided by Gorm.

**Connect to Database example:**

```go
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

dsn := "host=<your_host, localhost/127.0.0.1/0.0.0.0> user=<your_db_username> password=<your_db_password> dbname=<your_db_name> port=<your_port, normally in 5432> sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
	log.LogFatalf("failed to connect database: %v", err)
}
```

See [server.go](https://github.com/dpalhz/gogql-simulation-proj/blob/main/backend/cmd/server/server.go) file for more detail.

But if you don't use Gorm, you can install and implement PostgreSQL to your Go project using [pgx](https://github.com/jackc/pgx).

```sh
go get github.com/jackc/pgx/v5
```

**Basic example using pgx**

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)
}
```

### 6. Redis

To install Redis, you can see the detail instructions in their [Official Website](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/). Available options:

- [Windows](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-windows/).
- [Linux](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-linux/).
- [MacOS](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-mac-os/).
- [Build from source](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-from-source/).
- [With Redis Stack](https://redis.io/docs/latest/operate/oss_and_stack/install/install-stack/).

**How to install and setup Redis in your Golang Project?**

You can install and implement Redis to your Go project using [go-redis](https://github.com/redis/go-redis).

```sh
go get github.com/redis/go-redis/v9
```

Initialize Redis

```go
func InitRedis(addr string, password string) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Test the connection
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Reset database
	err = rdb.FlushDB(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to reset Redis database: %v", err)
	}

	return rdb, nil
}
```

See [session.go](https://github.com/dpalhz/gogql-simulation-proj/blob/main/backend/utils/session.go) file for more detail.

### 7. Air

Air is a live reloading tool designed to streamline the development process of your Golang project. Air automatically watch your codebase changes and reloads the applications without manual stop.

**How to install Air?**

There are some options to install Air.

- Using `go install`

  ```sh
  go install github.com/cosmtrek/air@latest
  ```

- Via install.sh

  ```sh
  curl -fLo ./air https://raw.githubusercontent.com/cosmtrek/air/master/bin/install
  chmod +x ./air
  ```

**Usage**

Work in progress.
