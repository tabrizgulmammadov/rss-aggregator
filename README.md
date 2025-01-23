# RSS Aggregator

## Overview

This RSS Aggregator is a Go-based web application that allows users to:
- Create and manage RSS feed subscriptions
- Automatically scrape and store articles from followed feeds
- Retrieve and view collected posts

## Features

### Core Functionality
- User registration and authentication
- RSS feed creation and tracking
- Automatic background scraping of RSS feeds
- Post collection and storage
- API endpoints for managing feeds and posts

### Technical Highlights
- Built with Go
- Uses PostgreSQL for data storage
- Implements concurrent feed scraping
- RESTful API design with Chi router
- SQLC for type-safe database queries

## Prerequisites

- Go 1.20+
- PostgreSQL
- Docker (optional)

## Project Structure

```
├── internal/
│   └── database/         # SQLC generated database methods
├── sql/
│   ├── queries/          # SQL query definitions
│   └── schema/           # Database schema
├── main.go               # Application entry point
├── handler_*.go          # API route handlers
├── middleware_auth.go    # Authentication middleware
├── models.go             # Data models and conversions
└── sqlc.yaml             # SQLC configuration
```

## Setup and Installation

### Environment Configuration

1. Clone the repository
2. Create a `.env` file with the following variables:
   ```
   PORT=8080
   DB_URL=postgresql://username:password@localhost:5432/rss_aggregator
   ```

### Database Setup

1. Create PostgreSQL database
2. Run database migrations
3. Configure connection in `.env`

### Running the Application

```bash
# Install dependencies
go mod tidy

# Run the application
go run .
```

## API Endpoints

### Users
- `POST /v1/users`: Create a new user
- `GET /v1/users`: Get current user details

### Feeds
- `POST /v1/feeds`: Create a new RSS feed
- `GET /v1/feeds`: List all feeds

### Feed Follows
- `POST /v1/feed-follows`: Follow a feed
- `GET /v1/feed-follows`: List feed follows
- `DELETE /v1/feed-follows/{feedFollowID}`: Unfollow a feed

### Posts
- `GET /v1/posts`: Retrieve user's collected posts

## Authentication

The application uses API key-based authentication:
- Each user receives a unique API key
- Include API key in request headers for authenticated routes

## Scraping Mechanism

- Background goroutine runs every minute
- Concurrently fetches updates from multiple feeds
- Stores new posts in the database
- Handles duplicate prevention

## Error Handling

- Comprehensive error responses
- Logging for server-side errors
- JSON-formatted error messages

## Performance Considerations

- Concurrent feed scraping
- Efficient database queries
- Timeout mechanisms for external requests

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request