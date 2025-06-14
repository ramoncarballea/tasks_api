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
- Docker 20.10+ and Docker Compose 1.29+

## Prerequisites

- Go 1.18 or higher
- PostgreSQL 12 or higher
- Git
- Docker 20.10+ and Docker Compose 1.29+

```bash
git clone https://github.com/yourusername/task-management-api.git
cd task-management-api
```

### 2. Set up environment variables

Create a `.env` file in the root directory with the following variables:

```env
# Server Configuration
PORT=8080

# Database Configuration
CONNECTION_STRING=postgres://username:password@localhost:5432/tasks_db?sslmode=disable

# Auto Migration
AUTO_MIGRATE=false
```

#### Configuration Details:

- `PORT`: The port on which the application will run (default: 8080)
- `CONNECTION_STRING`: PostgreSQL connection string in the format: 
  ```
  postgres://username:password@host:port/database_name?sslmode=disable
  ```
- `AUTO_MIGRATE`: Set to `true` to automatically run database migrations on startup (default: false)

For production, make sure to:
1. Use proper database credentials
2. Set `AUTO_MIGRATE` to `false` after initial setup
3. Consider enabling SSL in the connection string for production use

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the application

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## Docker Compose Files

- `docker-compose.yml`: Base configuration with common services
- `docker-compose.dev.yml`: Development environment with hot-reloading
- `docker-compose.prod.yml`: Production-optimized configuration

### Development Commands

- Start development environment:
  ```bash
  docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build
  ```

- View logs:
  ```bash
  docker-compose logs -f
  ```

- Stop services:
  ```bash
  docker-compose down
  ```

### Production Commands

- Start production environment:
  ```bash
  docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
  ```

- View logs:
  ```bash
  docker-compose logs -f
  ```

- Stop services:
  ```bash
  docker-compose down
  ```

- Update the application:
  ```bash
  git pull
  docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d --build
  ```

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