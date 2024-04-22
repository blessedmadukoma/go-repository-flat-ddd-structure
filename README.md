# Golang Repository Pattern with a Flat Domain-Driven Structure v1

A simple Golang repository pattern backend service providing RESTful APIs for managing users, and authentication through a flat structure approach.

Tools/Libraries involved:
- Golang
- PostgreSQL Database
- [pgx](https://github.com/jackc/pgx)
- [Sqlc](https://sqlc.dev)

## Features

- User registration
- User authentication
- Token-based authentication using JWT
- User management (CRUD operations)

## API Endpoints

### Authentication

#### Register User

- `POST /v1/auth/register`
  - Register a new user
  - Request Body:
    ```json
    {
        "firstname": "John",
        "lastname": "Doe",
        "email": "john@example.com",
        "password": "password123"
    }
    ```
  - Response:
    ```json
    {
    "status": true,
    "data": {
        "id": 1,
        "firstname": "John",
        "lastname": "Doe",
        "email": "john@example.com",
        "token": "<JWT_TOKEN>",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z"
    },
    "error": [],
    "message": "operation was successful"
    }
    ```

#### Login User

- `POST /v1/auth/login`
  - Login with existing credentials
  - Request Body:
    ```json
    {
        "email": "john@example.com",
        "password": "password123"
    }
    ```
  - Response:
    ```json
    {
        "token": "<JWT_TOKEN>"
    }
    ```

### User Management

#### Get User Profile

- `GET /v1/users/:id`
  - Get user profile by ID
  - Response:
    ```json
    {
        "id": 1,
        "firstname": "John",
        "lastname": "Doe",
        "email": "john@example.com"
    }
    ```

#### Update User Profile

- `PUT /v1/users/:id`
  - Update user profile by ID
  - Request Body:
    ```json
    {
        "firstname": "John",
        "lastname": "Doe",
        "email": "john@example.com"
    }
    ```
  - Response:
    ```json
    {
        "id": 1,
        "firstname": "John",
        "lastname": "Doe",
        "email": "john@example.com"
    }
    ```

#### Delete User

- `DELETE /v1/users/:id`
  - Delete user by ID
  - No request body
  - No response body, status code 204 on success

## Setup

### Dependencies

- Go 1.16+
- PostgreSQL

### Configuration

1. Copy `.env.example` to `.env` and update with your configuration.

### Build and Run

1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Build the binary:
   ```bash
   go build -o goRepositoryPattern
   ```
3. Run the server:
   ```bash
   ./goRepositoryPattern
   ```

## Environment Variables

- `DB_DRIVER`: Database driver (e.g., postgres)
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_HOST`: Database host
- `DB_PORT`: Database port
- `JWT_SECRET`: Secret key for JWT

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE)