## Installation Steps

1. Clone the repository

```bash
git clone https://github.com/fajarihsan21/go-backend.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Run the app

```bash
go run *.go serve
# or
go run main.go serve
```

You are all set!


## API Documentation
- GET [`/users/`](https://my-go-backend.herokuapp.com/users/) (get all users)
- GET [`/vehicles/`](https://my-go-backend.herokuapp.com/vehicles/) (get all vehicles)

## Built with

-   [Golang](https://go.dev/)
-   [gorilla/mux](https://github.com/gorilla/mux): for handle http request
-   [Postgres](https://www.postgresql.org/): for DBMS

<hr>

