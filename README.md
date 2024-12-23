Book Management API
A RESTful API service for managing books and their categories. This project is built using Go with Gin framework and PostgreSQL database.
Features

JWT Authentication
Book Management (CRUD operations)
Category Management (CRUD operations)
Automatic book thickness calculation
Input validation
PostgreSQL database integration
Audit trails (created_at, created_by, modified_at, modified_by)

Prerequisites

Go 1.21 or higher
PostgreSQL 12 or higher
Git

Installation & Setup

Clone the repository

bashCopygit clone https://github.com/yourusername/book-management-api.git
cd book-management-api

Install dependencies

bashCopygo mod tidy

Set up environment variables
Create a .env file in the root directory with the following variables:

CopyDB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=your-db-name
JWT_SECRET=your-jwt-secret

Run database migrations

bashCopygo run migrations/migrate.go

Start the server

bashCopygo run main.go
The server will start on http://localhost:8080
API Endpoints
Authentication
Login
CopyPOST /api/users/login
Request body:
jsonCopy{
    "username": "your_username",
    "password": "your_password"
}
Response:
jsonCopy{
    "token": "your.jwt.token"
}
Categories
All category endpoints require JWT authentication. Include the token in the Authorization header:
CopyAuthorization: Bearer your.jwt.token
Get All Categories
CopyGET /api/categories
Response:
jsonCopy[
    {
        "id": 1,
        "name": "Fiction",
        "created_at": "2024-01-01T00:00:00Z",
        "created_by": "admin",
        "modified_at": "2024-01-01T00:00:00Z",
        "modified_by": "admin"
    }
]
Create Category
CopyPOST /api/categories
Request body:
jsonCopy{
    "name": "Fiction"
}
Get Category by ID
CopyGET /api/categories/:id
Delete Category
CopyDELETE /api/categories/:id
Get Books by Category
CopyGET /api/categories/:id/books
Books
All book endpoints require JWT authentication.
Get All Books
CopyGET /api/books
Response:
jsonCopy[
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
Create Book
CopyPOST /api/books
Request body:
jsonCopy{
    "title": "Sample Book",
    "description": "A great book",
    "image_url": "http://example.com/image.jpg",
    "release_year": 2020,
    "price": 29900,
    "total_page": 250,
    "category_id": 1
}
Note: thickness is automatically calculated based on total_page:

If total_page > 100: "tebal"
If total_page ≤ 100: "tipis"

Get Book by ID
CopyGET /api/books/:id
Delete Book
CopyDELETE /api/books/:id
Validation Rules
Books

Release year must be between 1980 and 2024
Total page must be a positive number
Category ID must reference an existing category
Title is required
Price must be a positive number

Categories

Name is required
Name must be unique

Error Responses
The API returns appropriate HTTP status codes:

200: Success
201: Created
400: Bad Request (validation errors)
401: Unauthorized (invalid or missing token)
404: Not Found
500: Internal Server Error

Error response format:
jsonCopy{
    "error": "Error message here"
}
Deployment
This project can be deployed to Railway:

Push your code to GitHub
Connect your GitHub repository to Railway
Set up the environment variables in Railway dashboard
Railway will automatically build and deploy your application

Development Notes

The project uses a layered architecture (Handler -> Service -> Repository)
JWT is used for authentication
PostgreSQL is used as the database
All timestamps are in UTC
Audit fields are automatically populated

License
MIT License