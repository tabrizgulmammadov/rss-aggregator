# RSS Aggregator

An RSS Aggregator application built using Go, designed to fetch, parse, and serve RSS feeds with support for database operations, authentication, and scalable service components.

---

## Project Structure

```
- cmd
    - rss-aggregator
        main.go
- config
    config.go
- internal
    - auth
        auth.go                # Handles API key extraction and validation from requests
    - database
        # Auto-generated files by sqlc for database interactions
    - handler
        *_handler.go          # Handles requests and performs database operations
    - middleware
        auth.go               # Middleware for authentication
        config.go             # Middleware for application configuration
    - models
        # Model classes to structure responses returned to the client
    - service
        rss_service.go        # Core business logic for RSS feed management
        scraper_service.go    # Logic for scraping RSS feeds from external sources
    - utils
        json_response.go      # Utility for standardized JSON responses
- sql
    - queries
        # SQL query files for interacting with database tables
    - schema
        # SQL schema files for creating tables and other database structures
- vendor
.env                         # Environment variables
.env.example                 # Example environment file
.env.local                   # Local environment variables
.gitignore                   # Git ignore file
go.mod                       # Go module definition
go.sum                       # Dependency checksums
LICENSE                      # Project license
README.md                    # Project documentation
sql.yaml                     # SQLC configuration file
```

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (v1.18 or higher)
- [PostgreSQL](https://www.postgresql.org/) or any compatible database
- `sqlc` for generating database interaction code

### Installation

1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd rss-aggregator
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Configure the environment variables:
   - Create a `.env` file based on the `.env.example` file and fill in the necessary values.

4.. Generate database interaction files with `sqlc`:
   ```sh
   sqlc generate
   ```

### Running the Application

1. Start the server:
   ```sh
   go run cmd/rss-aggregator/main.go
   ```

2. The application should now be running on the configured port (default: `8080`).

---

## Features

- **Authentication**: Secure API key-based authentication for requests.
- **RSS Management**: Fetch, parse, and serve RSS feeds.
- **Database Integration**: Store and manage RSS data in a relational database using `sqlc`.
- **Scalable Design**: Organized project structure to support modular development.

---

## Directory Details

### `cmd/rss-aggregator`
Contains the entry point of the application (`main.go`). This is where the application starts.

### `config`
Defines application configuration settings, such as environment variable parsing and application constants.

### `internal`

#### `auth`
Handles API key extraction and validation from incoming requests.

#### `database`
Contains `sqlc`-generated files for database interaction.

#### `handler`
Defines HTTP handlers for managing requests and performing operations on the database.

#### `middleware`
Defines middleware for authentication and configuration loading.

#### `models`
Contains model classes for structuring responses returned to the client.

#### `service`
- `rss_service.go`: Contains core business logic for RSS feed management.
- `scraper_service.go`: Implements scraping logic to fetch RSS feeds from external sources.

#### `utils`
Contains utility functions, such as standardized JSON response formatting.

### `sql`

#### `queries`
Contains SQL query files for database interactions.

#### `schema`
Contains SQL schema files for creating tables and other database structures.

### Other Files

- `.env`, `.env.example`, `.env.local`: Environment variable files.
- `.gitignore`: Specifies intentionally untracked files to ignore.
- `go.mod`, `go.sum`: Go module and dependency files.
- `LICENSE`: Specifies the project license.
- `README.md`: Project documentation.
- `sql.yaml`: Configuration for `sqlc`.

---

## License

This project is licensed under the terms of the [MIT License](LICENSE).

---

## Contribution

Feel free to contribute to the project by opening issues or submitting pull requests. Make sure to follow the contribution guidelines.

---

## Contact

For any inquiries, please contact me at gulmammadovtabriz@gmail.com.
