# Task Management API

A robust and scalable Task Management API built with Go, Gin, and PostgreSQL. This API provides endpoints for managing tasks with features like CRUD operations, pagination, and more.

## Features

- **Task Management**: Create, read, update, and delete tasks
- **RESTful API**: Follows REST principles
- **Pagination**: Supports pagination for task listings
- **Structured Logging**: Uses Zap for efficient logging
- **Dependency Injection**: Built with Uber's FX for dependency management
- **Configuration Management**: Environment-based configuration
- **Database**: PostgreSQL with connection pooling
- **Caching**: In-memory caching support

## Prerequisites

- Go 1.18 or higher
- PostgreSQL 12 or higher
- Git

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/task-management-api.git
cd task-management-api
```

### 2. Set up environment variables

Create a `.env` file in the root directory with the following variables:

```env
# Server
PORT=8080

# Database
CONNECTION_STRING=postgres://username:password@localhost:5432/tasks_db?sslmode=disable
AUTO_MIGRATE=false
```

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Tasks

- `GET /api/v1/task` - List all tasks (with pagination)
- `POST /api/v1/task` - Create a new task
- `GET /api/v1/task/:id` - Get task details
- `PUT /api/v1/task/:id` - Update a task
- `DELETE /api/v1/task/:id` - Delete a task

## Project Structure

```
.
├── app/                  # Application setup and configuration
├── cmd/                  # Main application entry points
├── config/               # Configuration packages (cache, database, etc.)
├── modules/             # Feature modules
│   └── task/            # Task management module
│       ├── domain/      # Domain models and interfaces
│       ├── dto/         # Data transfer objects
│       ├── handlers/    # HTTP handlers
│       ├── repositories/# Data access layer
│       └── services/    # Business logic
├── .env                 # Environment variables
├── go.mod              # Go module definition
└── go.sum              # Go module checksums
```

## Development

### Running Tests

```bash
go test ./...
```

### Code Formatting

```bash
gofmt -s -w .
```

### Building the Application

```bash
go build -o bin/task-api cmd/main.go
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Uber FX](https://github.com/uber-go/fx)
- [Zap Logger](https://github.com/uber-go/zap)
- [Go PostgreSQL Driver](https://github.com/lib/pq)