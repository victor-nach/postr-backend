# **User Management API**

This is a **Go-based** API that uses **SQLite** as the database. The project is designed for simplicity and ease of deployment, with **manual migrations**, **automatic migrations on startup**, and **optional database seeding**.

---

## **Table of Contents**
1. [Installation and Setup](#installation-and-setup)
2. [Makefile Commands](#makefile-commands)
3. [Database Migrations and Seeding](#database-migrations-and-seeding)
4. [API Documentation](#api-documentation)
   - Endpoints
   - Errors
   - Entities

---

## **Installation and Setup**

### Prerequisites:
- **Go** (version 1.18+)
- **Git**

### Steps:
1. Clone the repository:
   ```sh
   git clone git@github.com:victor-nach/postr-backend.git
   cd your-repo-name
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build and run the application:
   ```sh
   make build-run
   ```

The application will start on `http://localhost:8080`.

---

## **Makefile Commands**

The following `Makefile` commands are available to simplify development and management:

| Command          | Description                                 |
|------------------|---------------------------------------------|
| `make build-run` | Build and run the application.              |
| `make migrate`   | Run the migrations (without seeding).       |
| `make migrate-seed` | Run the migrations and seed the database. |
| `make run`       | Start the application using `go run`.       |

### Explanation:
- **`make build-run`**: Builds the Go application and starts it from the `bin` directory.  
- **`make migrate`**: Applies all pending migrations to the SQLite database without seeding data.  
- **`make migrate-seed`**: Applies migrations and seeds the database with initial data (users and posts).  
- **`make run`**: Starts the app in development mode using `go run`.

---

## **Database Migrations and Seeding**

### Automatic Migrations
The application will apply the latest migrations **automatically on startup**. This ensures that your database schema is always up-to-date.

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

### Endpoints
#### `GET /users`
**Description:** Retrieve all users.  
**Response:**
```json
[
  {
    "id": "uuid",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@example.com",
    "street": "123 Elm Street",
    "city": "New York",
    "state": "NY",
    "zipcode": "10001"
  }
]
```

#### `POST /users`
**Description:** Create a new user.  
**Request Body:**
```json
{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john@example.com",
  "street": "123 Elm Street",
  "city": "New York",
  "state": "NY",
  "zipcode": "10001"
}
```

**Response:**
```json
{
  "message": "User created successfully"
}
```

---

### Errors

| **Name**           | **Code** | **Message**              | **Description**                                      |
|---------------------|----------|--------------------------|------------------------------------------------------|
| `ValidationError`   | 400      | `Invalid request data.`  | The request body contains invalid or missing fields. |
| `NotFoundError`     | 404      | `Resource not found.`    | The requested resource does not exist.               |
| `InternalServerError` | 500    | `Something went wrong.`  | A server error occurred.                             |

---

## **Entities**

### User
| **Field**   | **Type** | **Description**                       |
|-------------|-----------|---------------------------------------|
| `id`        | `string`  | Unique identifier for the user (UUID). |
| `firstname` | `string`  | User's first name.                    |
| `lastname`  | `string`  | User's last name.                     |
| `email`     | `string`  | User's email (must be unique).        |
| `street`    | `string`  | User's street address.                |
| `city`      | `string`  | City where the user resides.          |
| `state`     | `string`  | State where the user resides.         |
| `zipcode`   | `string`  | User's postal code.                   |
| `created_at` | `datetime` | Timestamp when the user was created. |

### Post
| **Field**   | **Type** | **Description**                       |
|-------------|-----------|---------------------------------------|
| `id`        | `string`  | Unique identifier for the post (UUID). |
| `user_id`   | `string`  | ID of the user who created the post.  |
| `title`     | `string`  | Title of the post.                    |
| `content`   | `string`  | Content of the post.                  |
| `created_at` | `datetime` | Timestamp when the post was created. |

---

## **Development and Contribution**

### Running Tests
You can write tests and run them using:
```sh
go test ./...
```

### Code Style
- Follow Go best practices.
- Use `golangci-lint` for linting.

### Contributing
Feel free to submit issues or pull requests.

---

# **User Management API**

This is a **Go-based** API that uses **SQLite** as the database. The project is designed for simplicity and ease of deployment, with **manual migrations**, **automatic migrations on startup**, and **optional database seeding**.

---

## **Table of Contents**
1. [Installation and Setup](#installation-and-setup)
2. [Makefile Commands](#makefile-commands)
3. [Database Migrations and Seeding](#database-migrations-and-seeding)
4. [API Documentation](#api-documentation)
   - Endpoints
   - Errors
   - Entities
5. [Development and Contribution](#development-and-contribution)

---

## **Installation and Setup**

### Prerequisites:
- **Go** (version 1.18+)
- **Git**

### Steps:
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/your-repo-name.git
   cd your-repo-name
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build and run the application:
   ```sh
   make build-run
   ```

The application will start on `http://localhost:8080`.

---

## **Makefile Commands**

The following `Makefile` commands are available to simplify development and management:

| Command          | Description                                 |
|------------------|---------------------------------------------|
| `make build-run` | Build and run the application.              |
| `make migrate`   | Run the migrations (without seeding).       |
| `make migrate-seed` | Run the migrations and seed the database. |
| `make run`       | Start the application using `go run`.       |

### Explanation:
- **`make build-run`**: Builds the Go application and starts it from the `bin` directory.  
- **`make migrate`**: Applies all pending migrations to the SQLite database without seeding data.  
- **`make migrate-seed`**: Applies migrations and seeds the database with initial data (users and posts).  
- **`make run`**: Starts the app in development mode using `go run`.

---

## **Database Migrations and Seeding**

### Automatic Migrations
The application will apply the latest migrations **automatically on startup**. This ensures that your database schema is always up-to-date.

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

### Endpoints
#### `GET /users`
**Description:** Retrieve all users.  
**Response:**
```json
[
  {
    "id": "uuid",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@example.com",
    "street": "123 Elm Street",
    "city": "New York",
    "state": "NY",
    "zipcode": "10001"
  }
]
```

#### `POST /users`
**Description:** Create a new user.  
**Request Body:**
```json
{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john@example.com",
  "street": "123 Elm Street",
  "city": "New York",
  "state": "NY",
  "zipcode": "10001"
}
```

**Response:**
```json
{
  "message": "User created successfully"
}
```

---

### Errors

| **Name**           | **Code** | **Message**              | **Description**                                      |
|---------------------|----------|--------------------------|------------------------------------------------------|
| `ValidationError`   | 400      | `Invalid request data.`  | The request body contains invalid or missing fields. |
| `NotFoundError`     | 404      | `Resource not found.`    | The requested resource does not exist.               |
| `InternalServerError` | 500    | `Something went wrong.`  | A server error occurred.                             |

---

## **Entities**

### User
| **Field**   | **Type** | **Description**                       |
|-------------|-----------|---------------------------------------|
| `id`        | `string`  | Unique identifier for the user (UUID). |
| `firstname` | `string`  | User's first name.                    |
| `lastname`  | `string`  | User's last name.                     |
| `email`     | `string`  | User's email (must be unique).        |
| `street`    | `string`  | User's street address.                |
| `city`      | `string`  | City where the user resides.          |
| `state`     | `string`  | State where the user resides.         |
| `zipcode`   | `string`  | User's postal code.                   |
| `created_at` | `datetime` | Timestamp when the user was created. |

### Post
| **Field**   | **Type** | **Description**                       |
|-------------|-----------|---------------------------------------|
| `id`        | `string`  | Unique identifier for the post (UUID). |
| `user_id`   | `string`  | ID of the user who created the post.  |
| `title`     | `string`  | Title of the post.                    |
| `content`   | `string`  | Content of the post.                  |
| `created_at` | `datetime` | Timestamp when the post was created. |

---

## **Development and Contribution**

### Running Tests
You can write tests and run them using:
```sh
go test ./...
```

### Code Style
- Follow Go best practices.
- Use `golangci-lint` for linting.

---

