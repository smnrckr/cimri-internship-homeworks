# Cimri Internship Homeworks

This repository contains various Go projects I worked on to better understand the language and its core concepts during my internship. Each project is located in a separate directory and focuses on different aspects of Go, including concurrency, data structures, and layered architecture.

## Projects

### 1. Guessing Game (`guess-game/`)
A simple command-line guessing game to get familiar with Go syntax and basic programming structures.

- The program generates a random number.
- The user tries to guess the number.

### 2. HTTP Response Status Checker (`http-status-checker/`)
A program that fetches response status codes from 10 different websites using Go's `http.Get` method. 

- Implements goroutines for concurrent requests.
- Uses a semaphore to limit the number of concurrent requests.
- Includes mutex for safe data access.

### 3. Binary Tree with Controlled Goroutines (`binary-tree/`)
A project demonstrating the use of binary trees with a concurrency optimization approach.

- Spawns goroutines only when necessary (leaf nodes).
- Efficiently processes nodes while maintaining controlled parallel execution.

### 4. Layered Architecture with Mock Database (`layered-architecture/`)
A project designed to practice and understand Go's layered architecture.

- Implements repository, service, and handler layers.
- Uses a mock database for data storage and retrieval.
- Follows clean architecture principles for better maintainability.

## Running the Projects
Each project can be run independently. Navigate to the respective folder and execute:

```bash
go run main.go
```

For projects requiring dependencies, ensure you have Go installed and run:

```bash
go mod tidy
go run main.go
```

## License
This repository is open-source and available for educational purposes.

---

Bunu doğrudan README olarak ekleyebilirsin. İstersen bazı detayları daha da genişletebilirim. Ne dersin? 😊
