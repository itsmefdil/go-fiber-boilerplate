## Golang Fiber REST API with GORM and PostgreSQL

This is a simple REST API written in Golang using the Fiber framework, GORM and PostgreSQL.

### Tools

1. Golang 1.21.1
2. Fiber v2.49.2

### Prerequisites

1. Install [Golang](https://golang.org/doc/install)
2. Install [PostgreSQL](https://www.postgresql.org/download/)
3. Install [Fiber](https://gofiber.io/)
4. Install [GORM](https://gorm.io/docs/#installation)
5. Install [GORM PostgreSQL Driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)

### Setup

1. Clone this repository
2. Create a database in PostgreSQL
3. Copy `.env.example` to `.env` and change some values
4. Run `go get` to install all dependencies
5. Run `go run main.go` to start the server
6. Open your browser and go to `http://localhost:3000/api/note` to see the list of notes

### API Endpoints

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | /api/note     | Get all notes     |
| GET    | /api/note/:id | Get a single note |
| POST   | /api/note     | Create a new note |
| PUT    | /api/note/:id | Update a note     |
| DELETE | /api/note/:id | Delete a note     |

Example Create a new note

    ```json
    {
    "title": "Test",
    "subtitle": "Testsss",
    "text": "Test Juga"
    }
    ```

### Swagger Documentation

1. Open your browser and go to `http://localhost:3000/docs` to see the swagger documentation

### References

1. [Golang](https://golang.org/)
2. [Fiber](https://gofiber.io/)
3. [GORM](https://gorm.io/)
4. [GORM PostgreSQL Driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
5. [PostgreSQL](https://www.postgresql.org/)
6. [Tutorial](https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34)
7. [Swagger](https://swagger.io/)
