# Go Fiber Todo API

A RESTful API for managing todos built with Go Fiber framework, MongoDB, and JWT authentication.

## Features

- User registration and authentication
- JWT-based authorization
- Create, read, update, and delete todos
- User-specific todo management
- MongoDB database integration

## Tech Stack

- **Framework**: [Go Fiber v3](https://gofiber.io/)
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment Management**: godotenv

## Project Structure

```
go-fiber/
├── go.mod
├── server.go
└── src/
    ├── app.go
    ├── controller/
    │   ├── auth.controller.go
    │   └── todo.controller.go
    ├── db/
    │   └── db.go
    ├── middleware/
    │   └── auth.middleware.go
    ├── model/
    │   ├── todo.model.go
    │   └── user.model.go
    └── router/
        ├── auth.route.go
        └── todo.route.go
```

## Prerequisites

- Go 1.26.1 or higher
- MongoDB instance (local or cloud)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Frank2006x/Fibre.git
cd go-fiber
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory with the following variables:
```env
MONGO_URL=mongodb://localhost:27017
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret_key
```

4. Start MongoDB service if running locally.

5. Run the application:
```bash
go run server.go
```

The server will start on port 3000.

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login user
- `POST /auth/logout` - Logout user (requires authentication)

### Todos (All require authentication)

- `POST /todos` - Create a new todo
- `GET /todos` - Get all todos for the authenticated user
- `PUT /todos/:id` - Update a specific todo
- `DELETE /todos/:id` - Delete a specific todo

## Request/Response Examples

### Register User
```bash
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Login User
```bash
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Create Todo
```bash
POST /todos
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "Buy groceries",
  "description": "Milk, bread, eggs",
  "status": "incomplete"
}
```

### Get Todos
```bash
GET /todos
Authorization: Bearer <jwt_token>
```

## Data Models

### User
```json
{
  "id": "ObjectId",
  "username": "string",
  "email": "string",
  "password": "string (hashed)",
  "todo": ["ObjectId array"]
}
```

### Todo
```json
{
  "id": "ObjectId",
  "title": "string",
  "description": "string",
  "status": "complete" | "incomplete"
}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

This project is licensed under the MIT License.