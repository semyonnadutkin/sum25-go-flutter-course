package models

import (
	"errors"
	"time"
) // used in "Message" struct, field: "Timestamp"

// Message represents a chat message
type Message struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// CreateMessageRequest represents the request to create a new message
type CreateMessageRequest struct {
	Username string `json:"username" validation:"required"`
	Content  string `json:"content" validation:"required"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	Content string `json:"content" validate:"required"`
}

// HTTPStatusResponse represents the response for HTTP status code endpoint
type HTTPStatusResponse struct {
	StatusCode  int    `json:"status_code"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewMessage creates a new message with the current timestamp
func NewMessage(id int, username, content string) *Message {
	return &Message{
		ID:        id,
		Username:  username,
		Content:   content,
		Timestamp: time.Now(), // "current timestamp"
	}
}

// Validate checks if the create message request is valid
func (r *CreateMessageRequest) Validate() error {
	// Check the "Username" field
	if len(r.Username) == 0 {
		// I believe these error messages suck
		// But somehow they are the "Go-to" choice for the course
		return errors.New("empty name error")
	}

	// Check the "Content" field
	if len(r.Content) == 0 {
		return errors.New("empty content error")
	}

	return nil
}

// Validate checks if the update message request is valid
func (r *UpdateMessageRequest) Validate() error {
	// Check the "Content" field for being empty
	if len(r.Content) == 0 {
		return errors.New("empty content error")
	}

	return nil
}
