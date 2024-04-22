package models

import "time"

// type User struct {
// 	ID             int64     `json:"id"`
// 	Email          string    `json:"email"`
// 	HashedPassword string    `json:"hashed_password"`
// 	CreatedAt      time.Time `json:"created_at"`
// 	UpdatedAt      time.Time `json:"updated_at"`
// }

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
