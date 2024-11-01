Here's the updated README with the revised API documentation:

---

# MjolnirCore

**Description**: MjolnirCore is the central repository and core framework of our Content Management System (CMS), implemented in Go with the Gin web framework. Named after Thor's mighty hammer from Norse mythology, MjolnirCore embodies robustness, power, and reliability in managing content and data within our application.

### Key Features

- **Scalable Architecture**: Built with scalability in mind, MjolnirCore provides a solid foundation to handle growing demands and data volumes.
- **Robust API Endpoints**: Includes a comprehensive set of API endpoints for managing users, products, companies, and teams, ensuring efficient data operations.
- **Secure Management**: Offers secure storage and management of essential CMS components, maintaining data integrity and accessibility.

**Dependencies**: Built on Go programming language and Gin web framework, leveraging their strengths in performance, concurrency, and web application development.

## Project Structure

The project is structured as follows:

```plaintext
.
├── cmd
│   └── main.go             # Entry point of the application
├── docker-compose.yml      # Docker Compose file for development environment
├── Dockerfile              # Dockerfile for containerizing the application
├── go.mod                  # Go modules file
├── go.sum                  # Go modules checksum file
├── internal_deprecated     # Deprecated internal directory (may not be in use)
├── pkg                     # Core packages and modules
│   ├── config              # Configuration package
│   ├── controllers         # API controllers
│   ├── interfaces          # Interfaces for external services (e.g., bcrypt, jwt)
│   ├── models              # Data models (e.g., User, Product, Company, Team)
│   ├── repositories        # Data access layer (repositories)
│   ├── routes              # API route definitions
│   ├── services            # Business logic services
│   └── utils               # Utility functions and helpers
└── tests                   # Unit tests and mock implementations
    ├── mocks               # Mocks for testing purposes
    └── services            # Service tests
```

---

## API Documentation

MjolnirCore provides various endpoints to manage Users, Products, Companies, and Teams. Here’s an overview of the endpoints:

### Users API

| Method | Endpoint                      | Description                       |
| ------ | ----------------------------- | --------------------------------- |
| POST   | `/api/users`                  | Creates a new user                |
| POST   | `/api/users/login`            | Authenticates user and issues JWT |
| GET    | `/api/users/invite/:inviteId` | Fetches user by invitation ID     |
| PUT    | `/api/users/:id`              | Updates user details              |

### Companies API

| Method | Endpoint         | Description                 |
| ------ | ---------------- | --------------------------- |
| PUT    | `/api/companies` | Updates company information |

### Products API

| Method | Endpoint            | Description               |
| ------ | ------------------- | ------------------------- |
| POST   | `/api/products`     | Creates a new product     |
| GET    | `/api/products`     | Retrieves all products    |
| GET    | `/api/products/:id` | Retrieves a product by ID |

### Teams API

| Method | Endpoint                                          | Description                                            |
| ------ | ------------------------------------------------- | ------------------------------------------------------ |
| GET    | `/api/companies/:companyId/teams`                 | Lists all teams under a specified company              |
| GET    | `/api/companies/:companyId/teams/:teamId/members` | Lists all members of a specified team within a company |

---

## Installation

To run this project locally, follow these steps:

### 1. Clone the repository:

```bash
git clone <repository_url>
cd <project_directory>
```

### 2. Set up environment variables:

Create a `.env` file in the root directory with the following variables:

```
ICLOUD_AUTH_EMAIL
ICLOUD_SENDER_EMAIL
ICLOUD_PASSWORD
SMTP_HOST
SMTP_PORT
PORT
DB_USER
DB_PASSWORD
DB_NAME
DB_HOST
DB_PORT
DOMAIN
JWT_SECRET
```

### 3. Install dependencies:

```bash
go mod tidy
```

### 4. Install pre-commit

We use pre-commit to manage pre-commit hooks. Install it with Homebrew (macOS) or pip:

```bash
# Using Homebrew
brew install pre-commit

# Using pip
pip install pre-commit
```

Refer to the [pre-commit documentation](https://pre-commit.com/index.html#install) for more details.

### 5. Install third-party packages

This project relies on the following third-party packages:

- **gofumpt**: A formatter for Go code.
  ```bash
  go install mvdan.cc/gofumpt@latest
  ```
- **air**: A live reload utility for Go applications.
  ```bash
  go install github.com/air-verse/air@latest
  ```
- **mockery**: A mock generation tool for Go interfaces.
  ```bash
  go install github.com/vektra/mockery/v2@v2.43.2
  ```

### 6. Create your database instance

```bash
docker-compose up -D
```

### 7. Build and run the application:

```bash
air
```

### 8. Access the API:

The API will be accessible at [http://localhost:8080](http://localhost:8080).

## Testing

To run tests for this project:

```bash
go test ./...
```

Unit tests are located in the `tests` directory, including mocks for repositories and external interfaces.

To generate required mocks for unit tests:

```bash
mockery
```

---

# Commit Message Guidelines

## Format

A correct commit message should follow the conventional commits specification, consisting of a header and an optional body:

### Header

The header includes three main parts:

- **`<type>`**: Describes the kind of change being made. Common types include:

  - `feat`: A new feature
  - `fix`: A bug fix
  - `docs`: Documentation changes
  - `style`: Code style changes
  - `refactor`: Code refactoring
  - `test`: Adding or updating tests
  - `chore`: Maintenance tasks

- **`<scope>`** (optional): Specifies the scope of the change.

- **`<subject>`**: A concise description (50 characters or less is recommended).

Example:

```
feat(auth): add JWT token expiration handling
```

By following these guidelines, we ensure that commit messages are informative and consistent.

---

This should work correctly as Markdown when displayed in a Markdown renderer or text editor that supports it!
