# Go Echo Hexagonal Architecture

This is a boilerplate project for a Go application using the Echo web framework and following the hexagonal architecture pattern.

## Getting Started

### Prerequisites

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/go-echo-hexagonal.git
   cd go-echo-hexagonal
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up the environment:**

   Create an `app.env` file in the root of the project and add the following environment variables:

   ```
   DB_DRIVER=postgres
   DB_SOURCE="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
   SERVER_PORT=":8080"
   ```

4. **Run the database:**

   ```bash
   docker-compose up -d
   ```

5. **Run the application:**

   ```bash
   go run cmd/main.go
   ```

## Usage

- **Create a user:**

  ```bash
  curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"email":"user@example.com","password":"password"}'
  ```

- **Get a user:**

  ```bash
  curl http://localhost:8080/users/1
  ```
