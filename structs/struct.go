package structs

import "time"

type Book struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url"`
	ReleaseYear int        `json:"release_year"`
	Price       int        `json:"price"`
	TotalPage   int        `json:"total_page"`
	Thickness   string     `json:"thickness"`
	CategoryID  int        `json:"category_id"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	ModifiedAt  *time.Time `json:"modified_at"` // Menggunakan pointer
	ModifiedBy  *string    `json:"modified_by"` // Menggunakan pointer
}

type Category struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"` // Gunakan pointer untuk menangani NULL
	ModifiedBy *string    `json:"modified_by"` // Sama untuk modified_by jika bisa NULL
}

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
