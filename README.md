# HEX CMS service

This project is a Content Management System (CMS) implemented in Go, using the Gin web framework. It includes features for managing users, products, and invoices, with scalable architecture and robust API endpoints.

## Project structure

The project is structured as follows:

```bash
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
│   ├── models              # Data models (e.g., User, Product, Invoice)
│   ├── repositories        # Data access layer (repositories)
│   ├── routes              # API route definitions
│   ├── services            # Business logic services
│   └── utils               # Utility functions and helpers
└── tests                   # Unit tests and mock implementations
    ├── mocks               # Mocks for testing purposes
    └── services            # Service tests
```

## Installation

To run this project locally, follow these steps:

1. Clone the repository:

```bash
Copiar código
git clone <repository_url>
cd <project_directory>
Set up environment variables:
```

2. Create a .env file in the root directory with the following environment variables:

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
DH_HOST
DB_PORT
DOMAIN
JWT_SECRET
```

3. Install dependencies:

```bash
go mod tidy
```

4. Install third party packages
This project relies on the following third-party packages:
[Link text Here](https://link-url-here.org)

- gofumpt:

  - ***Description***: A formatter for Go code that enforces additional style rules.
  - ***Installation***: ```go install mvdan.cc/gofumpt@latest```
  - ***Documentation***: [gofumpt documentation](https://github.com/mvdan/gofumpt)


- air:

  - ***Description***: A live reload utility for Go applications.
  - ***Installation***: ```go install github.com/air-verse/air@latest```
  - ***Documentation***: [air documentation](https://github.com/air-verse/air)

- mockery:

  - ***Description***: A mock generation tool for Go interfaces.
  - ***Installation***: ```go install github.com/vektra/mockery/v2@v2.43.2```
  - ***Documentation***: [Mockery documentation](https://vektra.github.io/mockery/latest/)


**Ensure these packages are installed and set up correctly to work with the project.**

5. Create your database instance

```bash
  docker-compose up -D
```

6. Build and run the application:

```bash
air
```

7. Access the API:

The API will be accessible at http://localhost:8080.

## Testing
To run tests for this project:
``` bash
go test ./...
```
Unit tests are located in the tests directory, including mocks for repositories and external interfaces.

To generate required mocks for unit tests simply run mockery on your terminal

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
  - `style`: Changes that do not affect the meaning of the code
  - `refactor`: Code refactorings
  - `test`: Adding or updating tests
  - `chore`: Maintenance tasks, build changes, etc.

- **`<scope>`** (optional): Specifies the scope of the change. It can be anything relevant to the project.

- **`<subject>`**: A succinct description of the change. Use the imperative tense (e.g., "add", "fix", "update") and keep it concise (50 characters or less is recommended).

Example:
`feat(auth): add JWT token expiration handling`

## Guidelines

- **Be Clear and Concise**: Describe the change clearly and succinctly.
- **Use Imperative Mood**: Start the `<subject>` with a verb in the imperative mood.
- **Separate Header and Body**: Use the body for additional context if needed.
- **Follow Project Conventions**: Adhere to any specific commit message guidelines established by the team.

By following these guidelines, we can ensure that our commit messages are informative, consistent, and contribute to a clear and meaningful Git history for our project.