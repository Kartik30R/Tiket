package models

import (
	"context"
	"time"

 )

type Event struct {
	ID       uint  `json:"id" gorm:"primarykey"`
	Name     string	`json:"name"`
	Location string `json:"location"`
	Date     time.Time `json:"date"`
	CreateAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}



type EventRepository interface{
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, eventId uint) (*Event, error)
	CreateOne(ctx context.Context, event *Event) (*Event, error)
	DeleteOne(ctx context.Context, eventId uint) error
	UpdateOne(ctx context.Context, eventId uint, updateEvent map[string]any) (*Event, error)
}