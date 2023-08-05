package app

import (
	"atom-project/internal/models"
	"context"
)

type Repository interface {
	SetUserSettings(ctx context.Context, token string, mode int, complexity int) error
	GetUserLessons(ctx context.Context, token string) ([]models.ListLesson, error)
}
