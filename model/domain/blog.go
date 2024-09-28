package domain

import "time"

type Blog struct {
	Id            int
	UserId        int
	CategoryId    int
	Title         string
	Content       string
	Likes         int
	CommentsCount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
