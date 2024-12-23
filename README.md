# Book Management API

A RESTful API service for managing books and their categories. This project is built using Go with Gin framework and PostgreSQL database.

## Features

- JWT Authentication
- Book Management (CRUD operations)
- Category Management (CRUD operations)
- Automatic book thickness calculation
- Input validation
- PostgreSQL database integration
- Audit trails (created_at, created_by, modified_at, modified_by)

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

## Installation & Setup

1. Clone the repository
```bash
git clone https://github.com/yourusername/book-management-api.git
cd book-management-api
```

2. Install dependencies
```bash
go mod tidy
```

3. Set up environment variables
Create a `.env` file in the root directory with the following variables:
```
DB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name
JWT_SECRET=your-jwt-secret
```

4. Run database migrations
```bash
go run migrations/migrate.go
```

5. Start the server
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

#### Login
```
POST /api/users/login
```
Request body:
```json
{
    "username": "your_username",
    "password": "your_password"
}
```
Response:
```json
{
    "token": "your.jwt.token"
}
```

### Categories

All category endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer your.jwt.token
```

#### Get All Categories
```
GET /api/categories
```
Response:
```json
[
    {
        "id": 1,
        "name": "Fiction",
        "created_at": "2024-01-01T00:00:00Z",
        "created_by": "admin",
        "modified_at": "2024-01-01T00:00:00Z",
        "modified_by": "admin"
    }
]
```

#### Create Category
```
POST /api/categories
```
Request body:
```json
{
    "name": "Fiction"
}
```

#### Get Category by ID
```
GET /api/categories/:id
```

#### Delete Category
```
DELETE /api/categories/:id
```

#### Get Books by Category
```
GET /api/categories/:id/books
```

### Books

All book endpoints require JWT authentication.

#### Get All Books
```
GET /api/books
```
Response:
```json
[
    {
        "id": 1,
        "title": "Sample Book",
        "description": "A great book",
        "image_url": "http://example.com/image.jpg",
        "release_year": 2020,
        "price": 29900,
        "total_page": 250,
        "thickness": "tebal",
        "category_id": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "created_by": "admin",
        "modified_at": "2024-01-01T00:00:00Z",
        "modified_by": "admin"
    }
]
```

#### Create Book
```
POST /api/books
```
Request body:
```json
{
    "title": "Sample Book",
    "description": "A great book",
    "image_url": "http://example.com/image.jpg",
    "release_year": 2020,
    "price": 29900,
    "total_page": 250,
    "category_id": 1
}
```

Note: `thickness` is automatically calculated based on `total_page`:
- If total_page > 100: "tebal"
- If total_page ≤ 100: "tipis"

#### Get Book by ID
```
GET /api/books/:id
```

#### Delete Book
```
DELETE /api/books/:id
```

## Validation Rules

### Books
- Release year must be between 1980 and 2024
- Total page must be a positive number
- Category ID must reference an existing category
- Title is required
- Price must be a positive number

### Categories
- Name is required
- Name must be unique

## Error Responses

The API returns appropriate HTTP status codes:

- 200: Success
- 201: Created
- 400: Bad Request (validation errors)
- 401: Unauthorized (invalid or missing token)
- 404: Not Found
- 500: Internal Server Error

Error response format:
```json
{
    "error": "Error message here"
}
```

## Deployment

This project can be deployed to Railway:

1. Push your code to GitHub
2. Connect your GitHub repository to Railway
3. Set up the environment variables in Railway dashboard
4. Railway will automatically build and deploy your application

## Development Notes

- The project uses a layered architecture (Handler -> Service -> Repository)
- JWT is used for authentication
- PostgreSQL is used as the database
- All timestamps are in UTC
- Audit fields are automatically populated

## License

[MIT License](LICENSE)