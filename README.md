# Simple Api with Go
It is a just simple RESTful API with Go using:
1. **gorilla/mux**
2. **gorm** 

## Installation & Run
```bash
# Download this project
$ go get github.com/dedidot/simple-api-golang

# Download Mux Router
$ go get github.com/gorilla/mux

# Download GORM
$ go get github.com/jinzhu/gorm

# Download stringer / generate random id
$ go get github.com/dedidot/generate/stringer
```

Setting DB in config/database.go
```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "testinger",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
cd simple-api-golang
go build
./simple-api-golang

# API Endpoint : http://127.0.0.1:8000
```

## Structure
```
├── app
│   ├── routes.go                   // Router
│   ├── controllers                 // Our API core handlers
│   │   ├── ActionPostBook.go       
│   │   ├── ActionViewBook.go       
│   └── models
│   |   └── book.go     // APi Model
|   └── utils
|       └── utils.go // generate id
├── config
│   └── database.go        // Configuration
├── migrate
|   └── book.go // Models for our application
└── main.go
```

## API

#### /book
* `GET` : Get all book
* `POST` : Create a new book

#### /book/:codes
* `GET` : Get a book
* `PUT` : Update a book
* `DELETE` : Delete a book

#Post Params
```
{
	"author": "Op Super John Doe Bilw",
	"name": "Implementation Golang",
	"category": "Knowledge"
}
```

```bash
# Inspired By
github.com/mingrammer/go-todo-rest-api-example
```