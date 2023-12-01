# 🐣 Egghead - Coin Management System 💰

[![Build & Deploy](https://github.com/ZoopSign/egghead/actions/workflows/production.ci-cd.yaml/badge.svg)](https://github.com/ZoopSign/egghead/actions/workflows/production.ci-cd.yaml)

[![Development CI / CD Pipeline](https://github.com/ZoopSign/egghead/actions/workflows/develop.ci-cd.yaml/badge.svg?branch=develop)](https://github.com/ZoopSign/egghead/actions/workflows/develop.ci-cd.yaml)


## Overview

This repository contains the implementation of the Coin Management System service written in Go. The service provides functionality for managing coin balances, performing credit and debit transactions, and maintaining transaction history for different products.

## 🛠️ Tech Stack

- **Go:** Programming language
- **PostgreSQL:** Database management system
- **Docker:** Containerization platform

## 🚀 Features

- **Product Management:**
  - **Get Products:** Retrieve the list of available products.
  - **Create Product:** Add a new product to the system.
  - **Get Single Products:** Get a single product by the `ProductID`.
  - **Update Products:** Update the product detail by the `productID`.
  - **Delete Product:** Delete the product from the system.

- **Coin Balance Management:**
  - **Get Balance:** Retrieve the current balance for a specific product for the user.
  - **Credit:** Add a specified amount to the balance of a specific product for the user.
  - **Debit:** Subtract a specified amount from the balance of a specific product for the user.

- **Transaction Management:**
  - **Get Transactions:** Retrieve the transaction history for a specific user.
  - **Get Single Transactions:** Get the transaction detail.

- **User Management:**
  - **Register User** Register the user for the first transaction.
  - **Archieve User** Archieve the user for the future use.

## 🛠️ Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:ZoopSign/egghead.git
   ```

2. Change to the project directory:

   ```bash
   cd egghead
   ```

3. Install dependencies:

   ```bash
   go get -u github.com/gorilla/mux
   go get -u github.com/lib/pq
   ```

4. Set up your PostgreSQL database and update the connection string in `main.go`:

   ```go
   // main.go
   connStr := "user=youruser dbname=yourdb sslmode=disable"
   ```

5. Build and run the service using Docker Compose:

    ```bash
    docker-compose up
    ```

## 🚀 Usage

### Running the Service

```bash
go run main.go
```

The service will be accessible at `http://localhost:8000`.

### 🛣️ Endpoints

#### Product Management

- **Get Products:**
  ```bash
  curl http://localhost:8000/internal/api/v1/products
  ```

- **Create Product:**
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"name":"New Product"}' http://localhost:8000/internal/api/v1/products
  ```

- **Get Products:**
  ```bash
  curl http://localhost:8000/internal/api/v1/products/{productID}
  ```

- **Update Product:**
  ```bash
  curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Product"}' http://localhost:8000/internal/api/v1/products/{productID}

  ```

- **Delete Product:**
  ```bash
  curl -X DELETE http://localhost:8000/internal/api/v1/products/{productID}

  ```


#### Coin Balance Management

- **Get Balance:**
  ```bash
  curl http://localhost:8000/private/api/v1/user/{userID}/balance
  ```

- **Credit:**
  ```bash
  curl -X POST http://localhost:8000/private/api/v1/user/{userID}/credit/
  ```

- **Debit:**
  ```bash
  curl -X POST http://localhost:8000/private/api/v1/user/{userID}/debit/
  ```

#### Transaction Management

- **Get Transactions:**
  ```bash
  curl http://localhost:8000/private/api/v1/user/{userID}/transactions
  ```

- **Get Transaction:**
  ```bash
  curl http://localhost:8000/private/api/v1/user/{userID}/transactions/{transactionID}
  ```

#### User Management

- **Archieve User:**
  ```bash
  curl http://localhost:8000private/api/v1/user/{userID}/archieve
  ```

- **Register User:**
  ```bash
  curl -X POST http://localhost:8000/private/api/v1/user/register
  ```

## 🧪 Testing

```bash
go test
```

Run the test suite to ensure the functionality of the service.

## Github Actions

```bash
# Install the act
brew install act

# dry run the jobs
act -n 

# run the actions
act

# verbose the jobs
act -v

# list all the actions
act -l

# secrets
act --secret-file my.secrets -s MY_SECRET=somevalue

# variables
act --var VARIABLE=somevalue

# env file
act --env-file my.env

# Run a job in a specific workflow (useful if you have duplicate job names)
act -j lint -W .github/workflows/checks.yml

# Collect artifacts to the /tmp/artifacts folder:
act --artifact-server-path /tmp/artifacts

# Run a specific job:
act -j test

# Run a specific event:
act <event_name>

#
act [<event>] [options]


```

## 📦 Dependencies

- [gorilla/mux](https://github.com/gorilla/mux): Router and dispatcher for building Go web servers.
- [lib/pq](https://github.com/lib/pq): PostgreSQL driver for Go's database/sql package.
- [migrations ](https://golang-migrate/migrate) for apply migrations
- [swaggerAPIdosc](https://github.com/swaggo/swag) for auto-generating Swagger API docs
- [security](https//github.com/securego/gosec) for checking Go security issues
- [best-practices](https://github.com/go-critic/go-critic) for checking Go the best practice issues
- [go-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

## 🤝 Contributing

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/new-feature`.
3. Make changes, commit, and push.
4. Submit a pull request.

