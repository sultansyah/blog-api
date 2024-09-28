package domain

import "time"

type Comment struct {
	Id        int
	UserId    int
	ParentId  int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
