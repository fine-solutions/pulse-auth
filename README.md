# JWT Authentication and Authorization Service in Go

This project is an HTTP service for user authentication and authorization, developed in the Go programming language. It utilizes JWT (JSON Web Tokens) for secure information transfer between the client and server.

## Used Libraries

- [Chi](https://github.com/go-chi/chi) - Lightweight and flexible HTTP router for Go.
- [Cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Simple and efficient parser for environment variables and configuration files in Go.
- [Zap](https://github.com/uber-go/zap) - Structured, powerful, and fast logger for Go.
- [Godotenv](https://github.com/joho/godotenv) - Loads environment variables from a `.env` file in the project.
- [SQLx](https://github.com/jmoiron/sqlx) - Enhanced database package for Go.
- [Squirrel](https://github.com/Masterminds/squirrel) - Simple and powerful SQL query builder for Go.

## Endpoints

- **POST** `/login` - User authentication. Provides a JWT token upon successful authentication.
- **POST** `/user/register` - Register a new user.
- **GET** `/user/{userID}` - Get user information by their identifier.
- **GET** `/user/search` - Search for a user by name and surname.

## Environment Variables

- `CONFIG_PATH` - Path to the configuration file.
- `DB_USER` - Database user name.
- `DB_PASSWORD` - Database user password.
- `DB_NAME` - Database name.
- `DB_PORT` - Database port.
- `PGADMIN_DEFAULT_EMAIL` - Email for accessing the PostgreSQL admin panel.
- `PGADMIN_DEFAULT_PASSWORD` - Password for accessing the PostgreSQL admin panel.
- `JWT_TOKEN_SALT` - Salt for signing JWT tokens.

## Running the Project

1. Install the necessary dependencies by running `go mod tidy`.
2. Create a `.env` configuration file and fill in the required environment variables.
3. Run the project by executing `go run main.go`.

