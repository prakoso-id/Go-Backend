# Go Backend API

A modern RESTful API built with Go, following Domain-Driven Design (DDD) principles. This project demonstrates a clean, modular architecture with proper separation of concerns.

## Features

- **Domain-Driven Design**: Clean separation of domain logic, application services, and infrastructure
- **Authentication**: JWT-based authentication with secure password hashing
- **CRUD Operations**: Full CRUD support for users and posts
- **Module Generator**: CLI tool for generating new DDD modules
- **API Documentation**: Comprehensive API documentation with examples
- **Testing**: Support for unit tests and mocks
- **Pagination**: Efficient data retrieval with pagination support
- **Modern Stack**: Go, Gin, GORM, PostgreSQL

## Project Structure

```
├── cmd/
│   ├── api/
│   │   └── main.go           # Application entry point
│   └── generator/            # Module generator tool
├── internal/
│   ├── infrastructure/       # Cross-cutting concerns
│   │   ├── database/        # Database connection and migrations
│   │   └── middleware/      # Global middleware (auth, logging)
│   ├── interfaces/
│   │   └── http/           # HTTP layer (router setup)
│   └── modules/            # Feature modules
│       ├── auth/           # Authentication module
│       ├── user/           # User management module
│       └── post/           # Post management module
│           ├── domain/     # Domain layer
│           │   ├── entity/
│           │   ├── repository/
│           │   └── service/
│           ├── handlers/   # HTTP handlers
│           ├── dto/       # Data transfer objects
│           ├── module.go  # Module setup
│           └── routes.go  # Route definitions
├── .env                    # Environment variables
├── go.mod                 # Go module file
└── Makefile              # Build automation
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository
2. Set up your environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Module Generation

Generate new DDD modules using our CLI tool:

```bash
# Basic module generation
go run cmd/generator/generate.go init -m <module-name>

# With tests and mocks
go run cmd/generator/generate.go init -m <module-name> -tests -mocks

# Skip confirmation
go run cmd/generator/generate.go init -m <module-name> -y

# Delete a module
go run cmd/generator/generate.go delete -m <module-name>
```

## API Endpoints

### Authentication
- POST `/api/v1/auth/login` - User login
- POST `/api/v1/auth/reset-password` - Reset password
- POST `/api/v1/auth/logout` - Logout (requires auth)

### Users
- POST `/api/v1/users` - Create user
- GET `/api/v1/users` - List users (paginated)
- GET `/api/v1/users/:id` - Get user by ID
- PUT `/api/v1/users/:id` - Update user (requires auth)
- DELETE `/api/v1/users/:id` - Delete user (requires auth)

### Posts
- POST `/api/v1/posts` - Create post (requires auth)
- GET `/api/v1/posts` - List posts (paginated)
- GET `/api/v1/posts/:id` - Get post by ID
- GET `/api/v1/posts/user/:user_id` - List user's posts
- PUT `/api/v1/posts/:id` - Update post (requires auth)
- DELETE `/api/v1/posts/:id` - Delete post (requires auth)

## Authentication

The API uses JWT (JSON Web Token) for authentication. To access protected endpoints:

1. Create a user account
2. Login to get your JWT token
3. Include the token in subsequent requests:
   ```
   Authorization: Bearer <your-jwt-token>
   ```

## Development

### Running Tests
```bash
go test ./... -v
```

### Code Generation
```bash
# Generate mocks
go generate ./...
```

### Database Migrations
The application handles migrations automatically on startup using GORM's AutoMigrate feature.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
