package db

import (
	"time"

	"github.com/google/uuid"
	// "github.com/segmentio/ksuid"
)

type User struct {
	PK                string    `dynamodbav:"PK"`
	SK                string    `dynamodbav:"SK"`
	Username          string    `dynamodbav:"username"`
	Email             string    `dynamodbav:"email"`
	HashedPassword    string    `dynamodbav:"hashed_password"`
	PasswordChangedAt time.Time `dynamodbav:"password_changed_at"`
	FullName          string    `dynamodbav:"full_name"`
	Address           string    `dynamodbav:"address"`
	Phone             string    `dynamodbav:"phone"`
	CreatedAt         time.Time `dynamodbav:"created_at"`
}

type Order struct {
	PK            string    `dynamodbav:"PK"`
	SK            string    `dynamodbav:"SK"`
	GSI1PK        string    `dynamodbav:"GSI1PK"`
	GSI1SK        string    `dynamodbav:"GSI1SK"`
	OrderId       string    `json:"orderId" dynamodbav:"OrderId"`
	Status        string    `json:"status" dynamodbav:"Status"`
	Amount        float64   `json:"amount" dynamodbav:"Amount"`
	NumberOfItems int       `json:"numberOfItems" dynamodbav:"NumberOfItems"`
	CreatedAt     time.Time `json:"created_at" dynamodbav:"CreatedAt"`
}

type OrderItem struct {
	PK          string  `dynamodbav:"PK"`
	SK          string  `dynamodbav:"SK"`
	GSI1PK      string  `dynamodbav:"GSI1PK"`
	GSI1SK      string  `dynamodbav:"GSI1SK"`
	OrderId     string  `json:"orderId" dynamodbav:"OrderId"`
	ItemId      string  `json:"itemId" dynamodbav:"ItemId"`
	Description string  `json:"description" dynamodbav:"Description"`
	Price       float64 `json:"price" dynamodbav:"Price"`
	Quantity    int64   `json:"quantity" dynamodbav:"Quantity"`
}

type Session struct {
	PK           string    `dynamodbav:"PK"`
	SK           string    `dynamodbav:"SK"`
	ID           uuid.UUID `dyanodbav:"id"`
	Username     string    `dynamodbav:"username"`
	RefreshToken string    `dynamodbav:"refresh_token"`
	UserAgent    string    `dynamodbav:"user_agent"`
	ClientIP     string    `dynamodbav:"client_ip"`
	IsBlocked    bool      `dynamodbab:"is_blocked"`
	ExpiresAt    time.Time `dynamodbav:"expires_at"`
}
