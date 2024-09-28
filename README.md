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

## Setup all stuff from scratch

### 1. Golang

To install Golang, you can go to their official Website. Choose based on your Operating Systems. Available options:

- Windows.
- MacOS ARM.
- MacOS x86.
- Linux.
- Build from source.

### 2. Fiber

According to the Website:

> Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.

Based on that statement, we know that Fiber has gained traction due to its performance and developer-friendly features.

So why we choose Fiber?

- Express-like. You know Express JS? Nah, you will adapt quickly while using Fiber.
- It's fast.
- Rich middleware ecosystem. [CORS](https://docs.gofiber.io/api/middleware/cors/), [monitoring](https://docs.gofiber.io/api/middleware/monitor/), [helmet](https://docs.gofiber.io/api/middleware/helmet/), [Logger](https://docs.gofiber.io/api/middleware/logger/), [Caching](https://docs.gofiber.io/api/middleware/cache), and many more. See [here](https://docs.gofiber.io/category/-middleware) for more detail.
- Many built-in features like routing, template engines, and static file serving.

**How to install Fiber in a Go Project?**

```
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
