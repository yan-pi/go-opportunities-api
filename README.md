# Go Opportunities API Documentation

## Table of Contents

- [Go Opportunities API Documentation](#go-opportunities-api-documentation)
  - [Table of Contents](#table-of-contents)
  - [Opening Object Structure](#opening-object-structure)
    - [`Opening` Object](#opening-object)
      - [Object Fields](#object-fields)
  - [Running the Project in Development Environment](#running-the-project-in-development-environment)
    - [Prerequisites](#prerequisites)
    - [Execution Steps](#execution-steps)
  - [API Endpoints](#api-endpoints)
  - [Conclusion](#conclusion)

---

## Opening Object Structure

### `Opening` Object

```go
type Opening struct {
    gorm.Model
    Role        string
    Description string
    Company     string
    Location    string
    Remote      bool
    Link        string
    Salary      int64
}
```

#### Object Fields

| Field       | Description                                                                 |
|-------------|-----------------------------------------------------------------------------|
| **ID**      | Unique identifier for the job opening.                                       |
| **Role**    | Job role title.                                                               |
| **Description** | Detailed description of the job.                                           |
| **Company** | Company offering the job.                                                    |
| **Location**| Location of the company.                                                     |
| **Remote**  | Indicates if the job is remote (true) or on-site (false).                    |
| **Link**    | Link for more information about the job or how to apply.                     |
| **Salary**  | Salary offered for the job, in cents.                                         |

---

## Running the Project in Development Environment

### Prerequisites

- Go (version 1.16 or higher) installed
- Git to clone the repository

### Execution Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/yan-pi/go-opportunities-api.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-opportunities-api
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

After executing the above commands, the API will be running at `http://localhost:8080`.

---

## API Endpoints
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/list-openings", handler.ListOpeningHandler)

## Conclusion

This project provides a Go API to manage job openings, using the Gin framework for routing and GORM for database interaction. The `Opening` object structure and available endpoints have been specified above to facilitate understanding and usage of the API.