# GoldRush Integration Service

## Overview

GoldRush Integration is a Go-based microservice that provides a unified API for blockchain data access and user management. It serves as an integration layer between client applications and blockchain data providers, offering features such as transaction information retrieval, wallet balance tracking, and user authentication.

## Features

- **User Authentication**: Register and login functionality with JWT-based authentication
- **Blockchain Data Access**: Query transaction information across multiple chains
- **Wallet Monitoring**: Track recent transactions and token balances for addresses
- **Historical Data**: Access historical portfolio values and transaction history
- **Observability**: Integrated with OpenTelemetry for tracing, metrics, and logging

## Architecture

The service follows a clean architecture pattern with the following components:

- **Handler Layer**: HTTP request handling and routing
- **Service Layer**: Business logic implementation
- **Repository Layer**: Data persistence and retrieval
- **Client Layer**: Integration with external blockchain data providers

### Tech Stack

- **Go 1.23.6**: Core programming language
- **Chi Router**: HTTP routing
- **Redis**: Token storage and caching
- **JWT**: Authentication mechanism
- **gRPC**: Communication with SSO service
- **Docker**: Containerization
- **OpenTelemetry**: Distributed tracing and metrics
- **Prometheus**: Metrics collection
- **Grafana/Promtail**: Log aggregation and visualization

## Installation

### Prerequisites

- Go 1.23.6 or higher
- Docker and Docker Compose
- Redis

### Local Development

1. Clone the repository:

```bash
git clone https://github.com/vladislavprovich/goldrush-integration.git
cd goldrush-integration
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file based on environment variables used in the configuration

4. Run the application:

```bash
go run cmd/main.go
```

### Docker Deployment

1. Build and start the containers:

```bash
docker-compose up -d
```

The service will be available at `http://localhost:8090`.

## API Endpoints

### Authentication

- `POST /api/v1/register` - Register a new user
- `POST /api/v1/login` - Login and receive JWT token

### Blockchain Data (Requires Authentication)

- `GET /api/v1/transaction` - Get transaction information
- `GET /api/v1/balances` - Get token balances for an address
- `GET /api/v1/recent-transactions` - Get recent transactions for an address
- `GET /api/v1/historical-transactions` - Get historical portfolio transactions

## Configuration

The application uses environment variables for configuration. Key configuration parameters include:

- Server address and timeouts
- Redis connection details
- JWT secret key
- External service endpoints
- Telemetry settings

Configuration is managed through the `cleanenv` package and can be provided via environment variables or a configuration file.

## Development

### Project Structure

```
├── cmd/                # Application entry points
│   ├── config/         # Configuration loading
│   └── main.go         # Main application
├── config/             # Configuration files
├── internal/           # Internal packages
│   ├── handler/        # HTTP handlers
│   ├── middleware/     # HTTP middleware
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
├── pkg/                # Shared packages
│   ├── client/         # External API clients
│   ├── slogwriter/     # Logging utilities
│   ├── telemetry/      # Observability tools
│   └── tokenjwtparsing/# JWT utilities
└── logs/               # Application logs
```

### Testing

Run tests with:

```bash
go test ./...
```

## Observability

The service is instrumented with OpenTelemetry for distributed tracing and metrics collection. Logs are collected using Promtail and can be visualized in Grafana.