# Golang REST API

- Basic REST API built on Golang using the Gin-Gonic Framework
- Connected to MongoDB
- Uses Go Modules
- Implements basic CRUD operations (for now)

## To run

- Clone repository
- Fill out .env file as given in env.example
- Run the main.go file

```sh
git clone https://github.com/rahulsai1999/go-rest.git
cd go-rest
go run ./main.go
```

## API

- Available on [Postman Docs](https://documenter.getpostman.com/view/5649815/SzYYzHzJ)

| Type of Request | Route     | Handler Package                                         |
| --------------- | --------- | ------------------------------------------------------- |
| GET             | /         | github.com/rahulsai1999/go-rest/service.ExtRouter       |
| GET             | /ping     | github.com/rahulsai1999/go-rest/service/api.Ping        |
| GET             | /blog     | github.com/rahulsai1999/go-rest/service/api.GetAllBlogs |
| GET             | /blog/:id | github.com/rahulsai1999/go-rest/service/api.GetBlogs    |
| POST            | /blog     | github.com/rahulsai1999/go-rest/service/api.InsertBlog  |
| PUT             | /blog/:id | github.com/rahulsai1999/go-rest/service/api.UpdateBlog  |
| DELETE          | /blog/:id | github.com/rahulsai1999/go-rest/service/api.DeleteBlog  |
