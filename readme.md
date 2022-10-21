# Template Fiber app

Stack:
`go`, `fiber`, `postgres`, `gorm` 

##### Install dependencies
```bash
> go get
```

##### Run project
```bash
 > air
 ```

#### Build

Create build file
```bash
> go build
```

Run build file
```bash
> ./todo-fiber
```
*todo-fiber - build name

#### Errors and fix

1. cannot find package "github.com/gofiber/fiber/v2" in any of
<br />
Fix: `go get` or `go mod tidy`
2. go: modules disabled by GO111MODULE=off; see 'go help modules'
<br />
Fix: `export GO111MODULE="on"`
