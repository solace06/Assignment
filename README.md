# Task Management API

A minimal REST API built in **Go** for managing tasks.

The API supports:
- Creating a new task
- Listing all existing tasks

Tasks are stored **in memory** and no database is required.

---

## Tech Stack

- Go (1.25)
- Gin Web Framework
- In-memory storage

---

## How to Run the Service

### 1. Clone the repository

```bash
git clone <repository-url>
cd ANRA_Assignment
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Start the server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## API Endpoints

### Create Task

**POST** `/api/v1/tasks`

Example request:

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Learn Go","status":"todo"}'
```

Example response:

```json
{
  "data": {
    "id": "uuid",
    "title": "Learn Go",
    "status": "todo"
  },
  "message": "task created successfully"
}
```

---

### List Tasks

**GET** `/api/v1/tasks`

Example request:

```bash
curl http://localhost:8080/api/v1/tasks
```

Example response:

```json
{
  "data": [
    {
      "id": "uuid",
      "title": "Learn Go",
      "status": "todo"
    }
  ],
  "message": "tasks fetched successfully"
}
```

---

## Validation Rules

- `title` is **required**
- `title` must be **≤ 200 characters**
- `status` allowed values:
  - `todo`
  - `in_progress`
  - `done`
- default status is **todo**

---

## Running Tests

Run all tests:

```bash
go test -v ./...
```

---

## Project Structure

```
api/
  router.go

internal/
  controller.go
  tasks_test.go
  models.go
  repo.go
  service.go
  utils.go

main.go
README.md
```

Architecture:

```
Router → Controller → Service → Repository
```