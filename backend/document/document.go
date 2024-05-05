package document

import (
	"context"
	"time"

	"encore.dev/beta/auth"
	"github.com/google/uuid"
)

type Document struct {
	ID           string            `json:"id"`
	AuthorID     string            `json:"author_id"`
	Content      string            `json:"content"`
	Version      int               `json:"version"`
	History      []DocumentVersion `json:"history"`
	CreatedAt    time.Time         `json:"created_at"`
	LastEditedAt time.Time         `json:"last_edited_at"`
}

type DocumentVersion struct {
	Version int    `json:"version"`
	Content string `json:"content"`
}

type CreateDocumentParams struct{}

//encore:api auth method=POST path=/document
func CreateDocument(ctx context.Context) (*Document, error) {

	id := uuid.New()

	creatorID, _ := auth.UserID()

	doc := &Document{
		ID:           id.String(),
		AuthorID:     string(creatorID),
		CreatedAt:    time.Now(),
		LastEditedAt: time.Now(),
	}

	return doc, nil
}
