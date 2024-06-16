# MjolnirCore

**Description**: MjolnirCore is the central repository and core framework of our Content Management System (CMS), implemented in Go with the Gin web framework. Named after Thor's mighty hammer from Norse mythology, MjolnirCore embodies robustness, power, and reliability in managing content and data within our application.

### Key Features:

- **Scalable Architecture**: Built with scalability in mind, MjolnirCore provides a solid foundation to handle growing demands and data volumes.
- **Robust API Endpoints**: Includes a comprehensive set of API endpoints for managing users, products, and invoices, ensuring efficient data operations.
- **Secure Management**: Offers secure storage and management of essential CMS components, maintaining data integrity and accessibility.
- **Usage**: MjolnirCore serves as the backbone for our CMS, facilitating efficient content management and data operations with its powerful features and scalable architecture.

***Dependencies***: Built on Go programming language and Gin web framework, leveraging their strengths in performance, concurrency, and web application development.

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

### 1. Clone the repository:

```bash
Copiar código
git clone <repository_url>
cd <project_directory>
Set up environment variables:
```

### 2. Create a .env file in the root directory with the following environment variables:

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

### 3. Install dependencies:

```bash
go mod tidy
```

### 4. Install pre-commit

We use pre-commit to manage our pre-commit hooks. Follow the installation steps below:

Installation: Install pre-commit using Homebrew (macOS) or pip (cross-platform):

```bash
# Using Homebrew
brew install pre-commit

# Using pip
pip install pre-commit
```

Documentation: Refer to the pre-commit [documentation](https://pre-commit.com/index.html#install) for more details.

### 5. Install third party packages

This project relies on the following third-party packages:

- gofumpt:

  - **_Description_**: A formatter for Go code that enforces additional style rules.
  - **_Installation_**: `go install mvdan.cc/gofumpt@latest`
  - **_Documentation_**: [gofumpt documentation](https://github.com/mvdan/gofumpt)

- air:

  - **_Description_**: A live reload utility for Go applications.
  - **_Installation_**: `go install github.com/air-verse/air@latest`
  - **_Documentation_**: [air documentation](https://github.com/air-verse/air)

- mockery:

  - **_Description_**: A mock generation tool for Go interfaces.
  - **_Installation_**: `go install github.com/vektra/mockery/v2@v2.43.2`
  - **_Documentation_**: [Mockery documentation](https://vektra.github.io/mockery/latest/)

- pre-commit-hooks:
  - **_Description_**: A collection of various pre-commit hooks maintained by the pre-commit team.
  - **_Installation_**: Configuration is handled via .pre-commit-config.yaml.
  - **_Documentation_**: [pre-commit-hooks documentation](https://github.com/pre-commit/pre-commit-hooks)

**Ensure these packages are installed and set up correctly to work with the project.**

### 6. Create your database instance

```bash
  docker-compose up -D
```

### 7. Build and run the application:

```bash
air
```

### 8. Access the API:

The API will be accessible at http://localhost:8080.

## Testing

To run tests for this project:

```bash
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
