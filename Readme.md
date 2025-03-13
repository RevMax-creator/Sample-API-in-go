# Sample API in Go

This repository provides a sample implementation of a RESTful API built with Go. It demonstrates a modular, clean architecture and best practices for building scalable web services. The project is ideal for developers who want to learn how to structure a Go API with proper routing, middleware, error handling, and database abstraction.

---

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Clean Architecture & Modular Design:**  
  The code is organized into distinct directories to separate API specifications, command-line entry, handlers, middleware, and tools. This makes the codebase easier to maintain and extend.

- **Routing & Middleware:**  
  Uses the [Chi](https://github.com/go-chi/chi) router for lightweight and flexible routing. Custom middleware handles authentication and error responses.

- **Database Abstraction:**  
  Implements a database interface along with a mock database (`mockdb`) for demonstration purposes. This allows for easy swapping of database implementations without altering the business logic.

- **JSON Handling & Error Responses:**  
  Includes examples for decoding JSON requests, encoding JSON responses, and consistent error handling using custom error handler functions.

- **External Package Integration:**  
  Utilizes popular Go packages such as `logrus` for advanced logging and `gorilla/schema` for URL query parameter decoding.

- **Generics & Modern Go Features:**  
  Demonstrates the usage of generics (where applicable) and modern Go features to keep the code concise and type-safe.

---

## Project Structure

```
API/
├── api
│   └── api.go                # API specifications: request/response structs and helper functions
├── cmd
│   └── api
│       └── main.go           # Application entry point: sets up the router and starts the server
├── internal
│   ├── handlers              # API endpoint handlers
│   │   ├── api.go            # Common handler functions and endpoint route definitions
│   │   └── get_coin_balance.go  # Handler for the GET /account/coins endpoint
│   ├── middleware            # Custom middleware
│   │   └── authorization.go  # Authorization middleware for endpoint protection
│   └── tools                 # Database interface and mock database implementation
│       ├── database.go       # Database interface definitions and constructor
│       └── mockdb.go         # Mock database implementation for testing/demo purposes
├── go.mod                    # Module definition
└── go.sum                    # Dependency checksums
```

---

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/RevMax-creator/sample-api-in-go.git
   cd sample-api-in-go
   ```

2. **Install Dependencies:**

   The repository uses Go modules. Ensure you have Go 1.18 or later installed, then run:

   ```bash
   go mod tidy
   ```

---

## Usage

1. **Build & Run:**

   You can run the API server using the following command:

   ```bash
   go run cmd/api/main.go
   ```

2. **Server Details:**

   The server starts on `localhost:8080`. You should see logs indicating that the service has started.

3. **Testing the API:**

   Use tools like [Postman](https://www.postman.com/) or `curl` to test the endpoints. For example, to get the coin balance for a user named "Alex":

   ```bash
   curl -X GET "http://localhost:8080/account/coins?username=alex" -H "Authorization: 123ABC"
   ```

   A successful response will include a JSON object with the coin balance and a status code.

---

## Endpoints

### GET `/account/coins`

- **Description:**  
  Returns the coin balance for a given user account.

- **Parameters:**
  - Query parameter `username`: The username of the account.
  - Header `Authorization`: The token required for authentication.

- **Response:**

  ```json
  {
    "code": 200,
    "balance": 100
  }
  ```

- **Error Handling:**  
  If the username or token is missing, or if authentication fails, the API returns an appropriate error message using standardized error responses.

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch with your feature or bugfix.
3. Commit your changes.
4. Submit a pull request describing your changes.

---

## License

This project is licensed under the Creative Commons Modified License. See the [LICENSE](https://github.com/RevMax-creator/Sample-API-in-go?tab=License-1-ov-file) file for details.

---

Feel free to explore, modify, and extend this sample API to suit your needs. Happy coding!