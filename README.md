# 🔗 URL Shortener API

A modern and efficient URL shortening service built with Go, following clean architecture principles. This service allows you to create shortened versions of long URLs for easier sharing and management.
This project was implemented based on the requirements from [roadmap.sh URL Shortener project](https://roadmap.sh/projects/url-shortening-service).

## ✨ Features

- 🎯 Create short URLs from long ones
- 🔄 Redirect from short to original URLs
- 📚 Swagger API documentation
- 🏗️ Clean Architecture design
- 🐘 PostgreSQL storage
- 📦 Redis caching
- 🔒 Error handling and validation
- 🔒 Rate limiting
- 🕹️ Testing

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher
- PostgreSQL
- Docker (optional)

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```
DB_NAME=your_db_name
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_SSL_MODE=disable

HTTP_PORT=8080
HTTP_HOST=http://localhost:8080

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=your_redis_password
REDIS_DB=0
```

### Installation

```
# Clone the repository
git clone https://github.com/idmaksim/url-shortener-api.git

# Install dependencies
go mod tidy

# Run the application
make start
```

## 📚 API Documentation

After starting the application, visit `http://localhost:8080/swagger/` for the Swagger API documentation.

### API Endpoints

```
POST /url           # Create a short URL
GET /{shortURL}     # Redirect to original URL
```

## 🏗️ Project Structure

```
.
├── cmd/
│   └── app/              # Application entry point
├── internal/
│   ├── app/             # Application setup
│   ├── config/          # Configuration
│   ├── db/              # Database connection
│   ├── domain/          # Business logic and entities
│   │   ├── models/      # Domain models
│   │   ├── services/    # Business logic
│   │   └── repositories/# Repository interfaces
│   ├── delivery/        # HTTP handlers and requests
│   │   └── http/
│   └── infrastructure/  # External implementations
│       └── repositories/# Repository implementations
```

## 🛠️ Technical Details

- **Clean Architecture**: Clear separation of concerns
- **Database**: PostgreSQL with GORM
- **API Framework**: Echo
- **Testing**: Testify
- **Documentation**: Swagger/OpenAPI
- **Configuration**: Environment variables with godotenv
- **Error Handling**: Custom domain errors
- **URL Generation**: UUID-based short URLs
- **Caching**: Redis caching
- **Rate Limiting**: IP-based rate limiting

## 🔨 Development

```
# Run tests
make test

# Run linter
make lint

# Generate Swagger docs
make swagger

# Build the application
make build
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ✨ Acknowledgments

- Built with [Echo Framework](https://echo.labstack.com/)
- Database powered by [GORM](https://gorm.io/)
- API documentation with [Swagger](https://swagger.io/)
