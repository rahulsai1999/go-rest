# Golang REST API

[![Build Status](https://travis-ci.com/rahulsai1999/go-rest.svg?branch=master)](https://travis-ci.com/rahulsai1999/go-rest)

- Available on [Postman Docs](https://documenter.getpostman.com/view/5649815/SzYYzHzJ)
- Basic REST API built on Golang using the Gin-Gonic Framework
- Connected to MongoDB
- Uses Go Modules
- Implements basic CRUD operations (for now)
- Implements Authentication from scratch using Bcrypt and JWT

## To run

### Direct (Golang installed)

- Clone repository
- Fill out .env file as given in env.example
- Run the main.go file

```sh
git clone https://github.com/rahulsai1999/go-rest.git
cd go-rest
go run ./main.go
```

### Using Docker

```sh
docker pull noneuser2183/go-rest
docker run -p 5000:5000 noneuser2183/go-rest
```

(or)

- Clone the Repository
- Fill out the .env file
- Build the Docker image

```sh
git clone https://github.com/rahulsai1999/go-rest.git
cd go-rest
docker build -t go-rest .
docker run -p 5000:5000 go-rest
```

## API

| Type of Request | Route     | Handler Package                                         |
| --------------- | --------- | ------------------------------------------------------- |
| GET             | /         | github.com/rahulsai1999/go-rest/service.ExtRouter       |
| GET             | /ping     | github.com/rahulsai1999/go-rest/service/api.Ping        |
| GET             | /blog     | github.com/rahulsai1999/go-rest/service/api.GetAllBlogs |
| GET             | /blog/:id | github.com/rahulsai1999/go-rest/service/api.GetBlogs    |
| POST            | /blog     | github.com/rahulsai1999/go-rest/service/api.InsertBlog  |
| PUT             | /blog/:id | github.com/rahulsai1999/go-rest/service/api.UpdateBlog  |
| DELETE          | /blog/:id | github.com/rahulsai1999/go-rest/service/api.DeleteBlog  |
