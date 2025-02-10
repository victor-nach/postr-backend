# **User Management API**

This is a Go-based API that leverages SQLite as its primary data store, designed with simplicity and ease of deployment in mind. The project embraces a Domain-Driven Design (DDD) approach by clearly separating concerns into distinct layers—domain, repository, service, and handler—which promotes maintainability and scalability. Dependency injection (DI) is at the heart of the architecture, ensuring that components remain loosely coupled and highly testable.

Key features include:

- **SQLite Integration:** A lightweight yet powerful database engine that keeps the API simple and efficient.
- **Flexible Migrations:** Supports both manual migrations and automatic migrations on startup, ensuring that the database schema is always up-to-date.
- **Optional Database Seeding:** Easily populate the database with initial data, which is especially useful during development and testing.
- **DDD & DI:** Leverages a Domain-Driven Design approach along with dependency injection to promote clean architecture, making the system easier to test, maintain, and scale.
- **Ease of Deployment:** With minimal configuration required, the API is ready to deploy quickly, making it ideal for rapid prototyping or production use.

---

## **Table of Contents**

1. [Project structure](#project-structure)
2. [Installation and Setup](#installation-and-setup)
3. [Makefile Commands](#makefile-commands)
4. [Database Migrations and Seeding](#database-migrations-and-seeding)
5. [API Documentation](#api-documentation)
   - Entities
   - Endpoints
   - Errors
   - Postman Collection

---

## 1. **Project Structure**

```
POSTR-BACKEND
├── cmd
│   ├── app
│   │   └── main.go
│   └── migration_runner
│       └── main.go
├── data
│   └── app.db
├── internal
│   ├── config
│   ├── domain
│   │   ├── domain.go
│   │   ├── errors.go
│   │   └── models.go
│   ├── handlers
│   │   ├── posts.go
│   │   ├── request.go
│   │   ├── response.go
|   |   ├── users.go
│   │   └── handler_test.go
│   ├── infrastructure
│   │   └── db
│   │       └── db.go
│   ├── repositories
│   │   ├── posts.go
|   |   |── posts_test.go
│   │   |── users.go
|   |   └── users_test.go
│   └── services
│       ├── postsservice
│       │   |── posts.go
|       |   └── posts_test.go
│       └── usersservice
│           |── users.go
|           └── users_test.go
├── migrations
│   ├── 0001_init_tables.down.sql
│   └── 0001_init_tables.up.sql
├── pkg
│   ├── logger
│   │   └── logger.go
│   └── migrator
│       └── migrator.go
├── seeds
│   ├── posts.json
│   └── users.json
├── .gitignore
├── .postman_collection.json
├── go.mod
├── go.sum
├── makefile
└── README.md
```

- **cmd/**: Contains the entry points for the application.

  - **app/**: Contains the main application (`main.go`).
  - **migration_runner/**: Contains the migration runner (`main.go`).

- **data/**: Contains database files, such as `app.db`.

- **internal/**: Contains internal application logic and modules.

  - **config/**: Configuration files and utilities.
  - **domain/**: Domain logic including models, errors, and domain-specific functions.
  - **handlers/**: HTTP handlers (e.g., for posts and users).
  - **infrastructure/**: Infrastructure code such as database connections.
  - **repositories/**: Code that interacts with the database.
  - **services/**: Business logic divided into services for posts and users.

- **migrations/**: SQL migration files for setting up and tearing down database schemas.

- **pkg/**: External or reusable packages.

  - **logger/**: Logging utilities.
  - **migrator/**: Migration management utilities.

- **seeds/**: JSON files containing seed data for posts and users.

- Additional files like `.gitignore`, `go.mod`, `go.sum`, `makefile`, and `README.md` are included in the root directory.

---

## 2. **Installation and Setup**

### Prerequisites:

- **Go** (version 1.18+)
- **Git**

### Steps:

1. Clone the repository:

   ```sh
   git clone git@github.com:victor-nach/postr-backend.git
   cd postr-backend
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Build and run the application:
   ```sh
   make build-run
   ```

The application will start on `http://localhost:8080`.

You can specify an alternative port `PORT` via a .env file in the project root

---

## **Makefile Commands**

The following `Makefile` commands are available to simplify development and management:

| Command             | Description                               |
| ------------------- | ----------------------------------------- |
| `make build-run`    | Build and run the application.            |
| `make migrate`      | Run the migrations (without seeding).     |
| `make migrate-seed` | Run the migrations and seed the database. |
| `make run`          | Start the application using `go run`.     |
| `make test`         | Run tests `go run`.                       |

---

## **Database Migrations and Seeding**

### Automatic Migrations

The application will apply the latest migrations **automatically on startup**. This ensures that the database schema is always up-to-date.

### Optional Migrations and Seeding

You can run migrations and optionally seed the database manually:

- **Run migrations only**:

  ```sh
  make migrate
  ```

- **Run migrations and seed the database**:
  ```sh
  make migrate-seed
  ```

---

## **API Documentation**

### **Entities**

### **User**

| **Field**    | **Type**   | **Description**                       |
| ------------ | ---------- | --------------------------------------|
| `id`         | `string`   | Unique identifier for the user (UUID) |
| `firstname`  | `string`   | User's first name                     |
| `lastname`   | `string`   | User's last name                      |
| `email`      | `string`   | User's email (must be unique)         |
| `street`     | `string`   | User's street address                 |
| `city`       | `string`   | City where the user resides           |
| `state`      | `string`   | State where the user resides          |
| `zipcode`    | `string`   | User's postal code                    |
| `created_at` | `datetime` | Timestamp when the user was created   |

### **Post**

| **Field**    | **Type**   | **Description**                       |
| ------------ | ---------- | --------------------------------------|
| `id`         | `string`   | Unique identifier for the post (UUID) |
| `user_id`    | `string`   | ID of the user who created the post   |
| `title`      | `string`   | Title of the post                     |
| `content`    | `string`   | Content of the post                   |
| `created_at` | `datetime` | Timestamp when the post was created   |

---

### **Endpoints**

### Users

### Retrieve all users.

#### `GET /users?pageNumber=3&pageSize=2`

**Request Query Parameters:**

- `pageNumber` (optional)
- `pageSize` (optional)

**Response:**

```json
{
  "status": "success",
  "message": "Users listed successfully",
  "pagination": {
    "current_page": 3,
    "total_pages": 25,
    "total_size": 50
  },
  "data": [
    {
      "id": "12296bff-6b03-42f1-a934-cf1995413d8c",
      "firstname": "Michael",
      "lastname": "Johnson",
      "email": "Michael.Johnson.6@acme.corp",
      "street": "15 Pine Ln.",
      "city": "Chicago",
      "state": "IL",
      "zipcode": "60007",
      "createdAt": "2025-02-09T23:28:04.3599836+01:00"
    },
    {
      "id": "70fbdfcd-a513-4e92-9e0d-9c74dac01a72",
      "firstname": "Emily",
      "lastname": "Williams",
      "email": "Emily.Williams.7@acme.corp",
      "street": "20 Maple Dr.",
      "city": "Houston",
      "state": "TX",
      "zipcode": "77001",
      "createdAt": "2025-02-09T23:28:04.3599836+01:00"
    }
  ]
}
```

### Retrieve user by ID.

#### `GET /users/:userId`

**Request Path Variables:**

- `userId` (required)

**Response:**

```json
{
  "status": "success",
  "message": "User retrieved successfully",
  "data": {
    "id": "963de191-8278-40f0-a367-e2e45e724aad",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@example.com",
    "street": "123 Elm Street",
    "city": "New York",
    "state": "NY",
    "zipcode": "10001",
    "createdAt": "2025-02-09T17:15:06.6062919+01:00"
  }
}
```

### Retrieve the total count of users.

#### `GET /users/count`

**Request Path Variables:**

- `userId` (required)

**Response:**

```json
{
  "status": "success",
  "message": "User retrieved successfully",
  "data": {
    "id": "963de191-8278-40f0-a367-e2e45e724aad",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@example.com",
    "street": "123 Elm Street",
    "city": "New York",
    "state": "NY",
    "zipcode": "10001",
    "createdAt": "2025-02-09T17:15:06.6062919+01:00"
  }
}
```

### Posts

### Create a new post.

#### `POST /posts`

**Request Body:**

```json
{
  "userId": "963de191-8278-40f0-a367-e2e45e724aad", // required
  "title": "the title", // required
  "body": "a random body" // required
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Post created successfully",
  "data": {
    "id": "438c550c-33b8-4fd4-9a27-631c720f3d43",
    "userId": "963de191-8278-40f0-a367-e2e45e724aad",
    "title": "the title",
    "body": "a random body",
    "createdAt": "2025-02-09T22:26:24.0343903+01:00"
  }
}
```

### Retrieve all posts for a specific user.

#### `GET /posts?userId=18de9b2e-7ebc-4624-9bb6-4c1ba4ea11e2`

**Request Query Parameters:**

- `userId` (required)

**Response:**

```json
{
  "status": "success",
  "message": "Posts listed successfully",
  "data": [
    {
      "id": "4f83e4ad-8325-4f20-a87b-50c74a294ecf",
      "userId": "18de9b2e-7ebc-4624-9bb6-4c1ba4ea11e1",
      "title": "Post 3",
      "body": "Content of post 3",
      "createdAt": "2025-02-09T17:15:06.6162837+01:00"
    }
  ]
}
```

### Delete a post by ID.

#### `DELETE /posts/:id`

**Request path Parameters:**

- `id` (required)

**Response:**

```json
{
  "status": "success",
  "message": "Post deleted successfully"
}
```

---

### Errors

**General Error Response:**

```json
{
  "status": "error",
  "code": "USR-404001",
  "message": "User not found"
}
```

---

### **API Error Codes**

| **Name**            | **Code**     | **Message**                                        | **Description**                                       |
| ------------------- | ------------ | -------------------------------------------------- | ----------------------------------------------------- |
| `ErrInternalServer` | `APP-500`    | `Internal server error - Unable to handle request` | A server error occurred while processing the request. |
| `ErrInvalidInput`   | `APP-400`    | `Invalid input data`                               | The request body contains invalid or missing fields.  |
| `ErrUserNotFound`   | `USR-404001` | `User not found`                                   | The specified user could not be found.                |
| `ErrPostNotFound`   | `PST-404001` | `Post not found`                                   | The specified post could not be found.                |
| `ErrCreateUser`     | `USR-400101` | `Failed to create user`                            | An error occurred while trying to create a user.      |

---

### **Postman collection**

A Postman collection is provided in the root of the project as .postman_collection.json. You can use this collection to quickly explore and test the API endpoints.

To use the Postman collection:

1. Open Postman.
2. Click on "Import" and select the .postman_collection.json file from the project root.
3. Explore the endpoints included in the collection and run sample requests.

---
