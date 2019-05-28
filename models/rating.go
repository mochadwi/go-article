package models

import (
	"time"
)

//go:generate goqueryset -in rating.go

// Rating struct represent rating model. Next line (gen:qs) is needed to autogenerate RatingQuerySet.
// gen:qs
type Rating struct {
	ID           int64      `json:"userId" validate:"required"`
	LessonID     int64      `json:"lessonId" validate:"required"`
	RatingNumber int64      `json:"rating" validate:"required"`
	UpdatedAt    time.Time  `json:"updated_at"`
	CreatedAt    time.Time  `json:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}
