
# Task Manager API

A RESTful API built with **Go (Golang)** for managing tasks. The API enables users to register, log in, and manage their tasks (create, read, update, delete) with secure authentication.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Execution](#execution)
- [API Endpoints](#api-endpoints)
- [Testing with Postman](#testing-with-postman)
- [Database Choice Rationale](#database-choice-rationale)
- [Conclusion](#conclusion)

## Features

- User registration and login (JWT Authentication)
- Task management (CRUD operations)
- User-specific task access
- Input validation
- Concurrent task update handling
- Logging and error handling middleware

## Technologies Used

- **Go (Golang)**: Backend language
- **Gin**: HTTP web framework for Go
- **GORM**: ORM for database interaction
- **PostgreSQL**: Relational database
- **bcrypt**: Password hashing for secure user authentication

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/gerry1818/task-manager-api.git
cd task-manager-api
```

### 2. Install Dependencies

Ensure you have Go installed, then run:

```bash
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file in the root directory with the following content:

```plaintext
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=postgres
DB_TIMEZONE=Asia/Kolkata

```

Replace `username` and `password` with your PostgreSQL credentials.

### 4. Set Up the Database

Ensure PostgreSQL is running. To create the required database, run:

```sql
CREATE DATABASE postgres;
```

## Execution

To run the application:

```bash
go run main.go
```

The API will be accessible at `http://localhost:8080`.

## API Endpoints

### 1. User Registration

- **POST** `/register`
  
  **Request Body**:
  ```json
  {
      "username": "your_username",
      "email": "your_email@example.com",
      "password": "your_password"
  }
  ```

### 2. User Login

- **POST** `/login`
  
  **Request Body**:
  ```json
  {
      "email": "your_email@example.com",
      "password": "your_password"
  }
  ```

### 3. Task Management

- **GET** `/tasks`: Get all tasks for the logged-in user.
- **POST** `/tasks`: Create a new task.
  
  **Request Body**:
  ```json
  {
      "title": "Task title",
      "description": "Task description"
  }
  ```

- **PUT** `/tasks/:id`: Update an existing task.
  
  **Request Body**:
  ```json
  {
      "title": "Updated title",
      "description": "Updated description"
  }
  ```

- **DELETE** `/tasks/:id`: Delete a task.

## Testing with Postman

1. Open Postman and import the provided `Postman_collection.json`.
2. Make requests to the API using the pre-configured requests.
3. Ensure to first register and log in to retrieve a JWT token for authenticated requests.

## Database Choice Rationale

**PostgreSQL** was chosen due to its strong feature set and robustness for relational data models. It supports:

- Concurrency control with **ACID** compliance
- Advanced querying capabilities, including full-text search and JSONB support
- Seamless integration with **GORM** for migrations and schema management
- Future-proof scalability

**Alternatives considered**:
- **MySQL**: A viable choice, but PostgreSQL provides more features and better concurrency handling.
- **SQLite**: Not suitable for production environments with concurrent write requirements.
- **MongoDB**: Suitable for unstructured data but adds complexity for relational models.

## Conclusion

This `README.md` covers the key steps to set up, run, and test the **Task Manager API**. Contributions are welcome—feel free to open an issue or pull request.